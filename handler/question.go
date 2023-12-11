package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h23w_10/pkg/domain"
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

func (h Handler) GetQuestions(c echo.Context, limit, offset int) error {
	questions, err := h.qrepo.Find(limit, offset)
	if err != nil {
		return err
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
