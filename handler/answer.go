package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type GetAnswersByQuestionIDResponse struct {
	ID         string    `json:"id"`
	UserID     string    `json:"user_id"`
	QuestionID string    `json:"title"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
}

func (h Handler) GetAnswersByQuestionID(c echo.Context, questionID string) error {
	answers, err := h.arepo.FindByQuestionID(questionID)
	if err != nil {
		return err
	}
	response := make([]GetAnswersByQuestionIDResponse, len(answers))
	for i, a := range answers {
		response[i] = GetAnswersByQuestionIDResponse{
			ID:         a.ID,
			UserID:     a.UserID,
			QuestionID: a.QuestionID,
			Content:    a.Content,
			CreatedAt:  a.CreatedAt,
		}
	}
	return c.JSON(http.StatusOK, response)
}
