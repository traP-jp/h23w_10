package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h23w_10/pkg/domain"
	"github.com/traP-jp/h23w_10/pkg/domain/repository"
)

type GetQuestionsResponce struct {
	ID        string       `json:"id"`
	UserID    string       `json:"user_id"`
	Title     string       `json:"title"`
	Content   string       `json:"content"`
	CreatedAt string       `json:"created_at"`
	Tags      []domain.Tag `json:"tags"`
	Status    string       `json:"status"`
}

func (h *Handler) GetQuestions(c echo.Context) error {
	limit := 10
	offset := 0
	if l := c.QueryParam("limit"); l != "" {
		limit, _ = strconv.Atoi(l)
	}
	if o := c.QueryParam("offset"); o != "" {
		offset, _ = strconv.Atoi(o)
	}

	tag := c.QueryParam("tag")
	var questions []domain.Question
	var err error
	if tag != "" {
		questions, err = h.qrepo.FindByTagID(tag, limit, offset)
	} else {
		questions, err = h.qrepo.Find(limit, offset)
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	response := make([]GetQuestionsResponce, len(questions))
	for i, q := range questions {
		response[i] = GetQuestionsResponce{
			ID:        q.ID,
			UserID:    q.UserID,
			Title:     q.Title,
			Content:   q.Content,
			CreatedAt: q.CreatedAt.String(),
			Tags:      q.Tags,
			Status:    string(q.Status),
		}
	}
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) GetQuestionByID(c echo.Context) error {
	id := c.Param("id")
	response, err := h.qrepo.FindByID(id)
	if errors.Is(err, repository.ErrNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	answers, err := h.arepo.FindByQuestionID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	response.Answers = answers
	return c.JSON(http.StatusOK, response)
}
