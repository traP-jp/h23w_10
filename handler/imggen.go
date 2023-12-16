package handler

import (
	"errors"
	"image"
	"image/png"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type PostImageRequest struct {
	IconURLs []string `json:"icon_url,omitempty"`
}

func (h *Handler) PostImage(c echo.Context) error {
	var request PostImageRequest
	if err := c.Bind(&request); err != nil {
		return err
	}
	icons := make([]image.Image, len(request.IconURLs))
	for i, iconURL := range request.IconURLs {
		var err error
		icons[i], err = openImage(iconURL)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	res, err := h.imggenSvc.GenerateImage(icons)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	fileName := uuid.New().String()

	file, err := os.Create(fileName + ".png")
	if err != nil {
		return err
	}
	defer file.Close()

	png.Encode(file, res)
	return c.File(fileName + ".png")
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
