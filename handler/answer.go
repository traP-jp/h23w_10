package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h23w_10/pkg/domain"
)

type GetAnswersByQuestionIDResponse struct {
	ID         string `json:"id,omitempty"`
	UserID     string `json:"user_id,omitempty"`
	QuestionID string `json:"title,omitempty"`
	Content    string `json:"content,omitempty"`
	CreatedAt  string `json:"created_at,omitempty"`
}

type PostAnswerRequest struct {
	UserID     string `json:"user_id,omitempty"`
	QuestionID string `json:"question_id,omitempty"`
	Content    string `json:"content,omitempty"`
}

type PostAnswerResponse struct {
	ID         string `json:"id,omitempty"`
	UserID     string `json:"user_id,omitempty"`
	QuestionID string `json:"title,omitempty"`
	Content    string `json:"content,omitempty"`
	CreatedAt  string `json:"created_at,omitempty"`
}

func (h *Handler) GetAnswersByQuestionID(c echo.Context) error {
	questionID := c.Param("id")
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
			CreatedAt:  a.CreatedAt.String(),
		}
	}
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) PostAnswer(c echo.Context) error {
	var request PostAnswerRequest
	if err := c.Bind(&request); err != nil {
		return err
	}

	answer := &domain.Answer{
		ID:         domain.NewUUID(),
		UserID:     request.UserID,
		QuestionID: request.QuestionID,
		Content:    request.Content,
		CreatedAt:  time.Now(),
	}
	result, err := h.arepo.Create(answer)
	if err != nil {
		return err
	}

	response := PostAnswerResponse{
		ID:         result.ID,
		UserID:     result.UserID,
		QuestionID: result.QuestionID,
		Content:    result.Content,
		CreatedAt:  result.CreatedAt.String(),
	}

	return c.JSON(http.StatusCreated, response)
}
