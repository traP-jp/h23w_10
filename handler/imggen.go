package handler

import (
	"context"
	"errors"
	"image"
	"image/png"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h23w_10/pkg/domain"
	"github.com/traP-jp/h23w_10/pkg/domain/repository"
	"golang.org/x/sync/errgroup"
)

type PostImageRequest struct {
	UserID string `json:"user_id,omitempty"`
}

func (h *Handler) PostImage(c echo.Context) error {
	var request PostImageRequest
	if err := c.Bind(&request); err != nil {
		return err
	}

	if _, err := os.Stat(filepath.Join(imagesDir, request.UserID+".png")); err == nil {
		return c.File(filepath.Join(imagesDir, request.UserID+".png"))
	}

	user, err := h.urepo.FindUserByID(request.UserID)
	if errors.Is(err, repository.ErrNotFound) {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	questions, _, err := h.qrepo.FindByUserID(request.UserID, &repository.FindQuestionsCondition{
		Limit:    46,
		Offset:   0,
		Statuses: domain.AvailableQuestionStatus(),
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	iconURLs := make(map[string]string)
	for _, question := range questions {
		a, err := h.arepo.FindByQuestionID(question.ID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		for _, answer := range a {
			if _, ok := iconURLs[answer.UserID]; ok {
				continue
			}
			user, err := h.urepo.FindUserByID(answer.UserID)
			if errors.Is(err, repository.ErrNotFound) {
				continue
			} else if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			if answer.UserID != request.UserID {
				iconURLs[user.IconURL.String()] = user.IconURL.String()
			}
		}
	}
	// チャンネルにアイコン画像を送っていく
	eg, ctx := errgroup.WithContext(context.Background())
	icons := make(chan image.Image)
	eg.Go(func() error {
		img, err := openImage(user.IconURL.String())
		if err != nil {
			return err
		}
		icons <- img
		for _, iconURL := range iconURLs {
			img, err := openImage(iconURL)
			if err != nil {
				return err
			}
			icons <- img
		}
		close(icons)
		return nil
	})
	// 画像を生成する
	res, err := h.imggenSvc.GenerateImage(ctx, 1+len(iconURLs), icons)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := eg.Wait(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	filename := filepath.Join(imagesDir, user.ID+".png")
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	png.Encode(file, res)

	return c.File(filename)
}

func openImage(iconURL string) (image.Image, error) {
	response, err := http.Get(iconURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, errors.New("not StatusOK")
	}

	img, _, err := image.Decode(response.Body)
	if err != nil {
		return nil, err
	}
	return img, nil
}
