package handler

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
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

type PostQuestionRequest struct {
	UserID  string
	Title   string
	Content string
	Tags    []domain.Tag
	Status  domain.QuestionStatus
}

type PostQuestionResponce struct {
	ID        string
	UserID    string
	Title     string
	Content   string
	CreatedAt time.Time
	Tags      []domain.Tag
	Status    domain.QuestionStatus
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

	statuses := strings.Split(c.QueryParam("status"), ",")

	tag := c.QueryParam("tag")
	var questions []domain.Question
	var err error
	if tag != "" {
		questions, err = h.qrepo.FindByTagID(tag, &repository.FindQuestionsCondition{
			Limit:  limit,
			Offset: offset,
			Statuses: lo.Map(statuses, func(i string, _ int) domain.QuestionStatus {
				return domain.QuestionStatus(i)
			}),
		})
	} else {
		questions, err = h.qrepo.Find(&repository.FindQuestionsCondition{
			Limit:  limit,
			Offset: offset,
			Statuses: lo.Map(statuses, func(i string, _ int) domain.QuestionStatus {
				return domain.QuestionStatus(i)
			}),
		})
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

func (h *Handler) PostQuestion(c echo.Context) error {
	var request PostQuestionRequest
	if err := c.Bind(&request); err != nil {
		return err
	}

	question := &domain.Question{
		ID:        domain.NewUUID(),
		UserID:    request.UserID,
		Title:     request.Title,
		Content:   request.Content,
		CreatedAt: time.Now(),
		Tags:      request.Tags,
		Status:    request.Status,
	}
	result, err := h.qrepo.Create(question)
	if err != nil {
		return err
	}

	responce := PostQuestionResponce{
		ID:        result.ID,
		UserID:    result.UserID,
		Title:     result.Title,
		Content:   result.Content,
		CreatedAt: result.CreatedAt,
		Tags:      result.Tags,
		Status:    result.Status,
	}

	return c.JSON(http.StatusOK, responce)
}
