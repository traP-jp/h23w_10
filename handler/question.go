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

type GetQuestionsResponse struct {
	Total     int               `json:"total,omitempty"`
	Questions []domain.Question `json:"questions,omitempty"`
}

type GetTagsResponse struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type PostQuestionRequest struct {
	UserID  string       `json:"user_id,omitempty"`
	Title   string       `json:"title,omitempty"`
	Content string       `json:"content,omitempty"`
	Tags    []domain.Tag `json:"tags,omitempty"`
	Status  string       `json:"status,omitempty"`
}

type PostQuestionResponse struct {
	ID        string       `json:"id,omitempty"`
	UserID    string       `json:"user_id,omitempty"`
	Title     string       `json:"title,omitempty"`
	Content   string       `json:"content,omitempty"`
	CreatedAt string       `json:"created_at,omitempty"`
	Tags      []domain.Tag `json:"tags,omitempty"`
	Status    string       `json:"status,omitempty"`
}

type UpdateQuestionRequest struct {
	ID      string       `json:"id,omitempty"`
	Title   string       `json:"title,omitempty"`
	Content string       `json:"content,omitempty"`
	Tags    []domain.Tag `json:"tags,omitempty"`
	Status  string       `json:"status,omitempty"`
}

type PostTagRequest struct {
	Name string `json:"name,omitempty"`
}

type PostTagResponse struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
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
	var statuses []domain.QuestionStatus
	if s := c.QueryParam("status"); s != "" {
		statuses = lo.Map(strings.Split(s, ","), func(i string, _ int) domain.QuestionStatus {
			return domain.QuestionStatus(i)
		})
	} else {
		statuses = domain.AvailableQuestionStatus()
	}
	condition := &repository.FindQuestionsCondition{
		Limit:    limit,
		Offset:   offset,
		Statuses: statuses,
	}

	tag := c.QueryParam("tag")
	var questions []domain.Question
	var total int
	var err error
	if tag != "" {
		questions, total, err = h.qrepo.FindByTagID(tag, condition)
	} else {
		questions, total, err = h.qrepo.Find(condition)
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	response := GetQuestionsResponse{
		Total:     total,
		Questions: questions,
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

func (h *Handler) UpdateQuestion(c echo.Context) error {
	var req UpdateQuestionRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	question := &domain.Question{
		ID:      req.ID,
		Title:   req.Title,
		Content: req.Content,
		Tags:    req.Tags,
		Status:  domain.QuestionStatus(req.Status),
	}
	_, err := h.qrepo.Update(question)
	if errors.Is(err, repository.ErrTagNotFound) {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *Handler) GetTags(c echo.Context) error {
	tags, err := h.qrepo.FindTags()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := make([]GetTagsResponse, len(tags))
	for i, t := range tags {
		response[i] = GetTagsResponse{
			ID:   t.ID,
			Name: t.Name,
		}
	}
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
		Status:    domain.QuestionStatus(request.Status),
	}
	result, err := h.qrepo.Create(question)
	if err != nil {
		return err
	}

	response := PostQuestionResponse{
		ID:        result.ID,
		UserID:    result.UserID,
		Title:     result.Title,
		Content:   result.Content,
		CreatedAt: result.CreatedAt.String(),
		Tags:      result.Tags,
		Status:    string(result.Status),
	}

	return c.JSON(http.StatusCreated, response)
}

func (h *Handler) PostTag(c echo.Context) error {
	var request PostTagRequest
	if err := c.Bind(&request); err != nil {
		return err
	}

	tag := &domain.Tag{
		ID:   domain.NewUUID(),
		Name: request.Name,
	}
	result, err := h.qrepo.CreateTag(tag)
	if errors.Is(err, repository.ErrTagAlreadyExists) {
		return echo.NewHTTPError(http.StatusBadRequest, repository.ErrTagAlreadyExists)
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := PostTagResponse{
		ID:   result.ID,
		Name: result.Name,
	}

	return c.JSON(http.StatusCreated, response)
}
