package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"github.com/traP-jp/h23w_10/pkg/domain"
	"github.com/traP-jp/h23w_10/pkg/domain/repository"
)

type User struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	IconURL     string `json:"icon_url,omitempty"`
	UserType    string `json:"user_type,omitempty"`
}

type QuestionWithUser struct {
	ID        string                `json:"id,omitempty"`
	User      User                  `json:"user,omitempty"`
	Title     string                `json:"title,omitempty"`
	Content   string                `json:"content,omitempty"`
	CreatedAt time.Time             `json:"created_at,omitempty"`
	Tags      []domain.Tag          `json:"tags,omitempty"`
	Status    domain.QuestionStatus `json:"status,omitempty"`
}

type AnswerWithUser struct {
	ID         string    `json:"id,omitempty"`
	User       User      `json:"user,omitempty"`
	QuestionID string    `json:"question_id,omitempty"`
	Content    string    `json:"content,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
}

type GetQuestionsResponse struct {
	Total     int                `json:"total,omitempty"`
	Questions []QuestionWithUser `json:"questions,omitempty"`
}

type GetQuestionByIDResponse struct {
	ID        string                `json:"id,omitempty"`
	User      User                  `json:"user,omitempty"`
	Title     string                `json:"title,omitempty"`
	Content   string                `json:"content,omitempty"`
	CreatedAt time.Time             `json:"created_at,omitempty"`
	Tags      []domain.Tag          `json:"tags,omitempty"`
	Answers   []AnswerWithUser      `json:"answers,omitempty"`
	Status    domain.QuestionStatus `json:"status,omitempty"`
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
	BotPost bool         `json:"bot_post,omitempty"`
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

type PutQuestionRequest struct {
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

	result := make([]QuestionWithUser, 0, len(questions))
	for _, q := range questions {
		user, err := h.urepo.FindUserByID(q.UserID)
		if errors.Is(err, repository.ErrNotFound) {
			total--
			continue
		} else if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		result = append(result, QuestionWithUser{
			ID: q.ID,
			User: User{
				ID:          user.ID,
				Name:        user.Name,
				DisplayName: user.DisplayName,
				IconURL:     user.IconURL.String(),
				UserType:    string(user.UserType),
			},
			Title:     q.Title,
			Content:   q.Content,
			CreatedAt: q.CreatedAt,
			Tags:      q.Tags,
			Status:    q.Status,
		})
	}

	response := GetQuestionsResponse{
		Total:     total,
		Questions: result,
	}
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) GetQuestionByID(c echo.Context) error {
	id := c.Param("id")
	question, err := h.qrepo.FindByID(id)
	if errors.Is(err, repository.ErrNotFound) {
		log.Println("question not found")
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	user, err := h.urepo.FindUserByID(question.UserID)
	if errors.Is(err, repository.ErrNotFound) {
		log.Println("user not found")
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	answers, err := h.arepo.FindByQuestionID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	answerUsers := make(map[string]domain.User)
	for _, answer := range answers {
		if _, ok := answerUsers[answer.UserID]; !ok {
			user, err := h.urepo.FindUserByID(answer.UserID)
			if errors.Is(err, repository.ErrNotFound) {
				continue
			} else if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			answerUsers[answer.UserID] = *user
		}
	}

	response := GetQuestionByIDResponse{
		ID: question.ID,
		User: User{
			ID:          user.ID,
			Name:        user.Name,
			DisplayName: user.DisplayName,
			IconURL:     user.IconURL.String(),
			UserType:    string(user.UserType),
		},
		Title:     question.Title,
		Content:   question.Content,
		CreatedAt: question.CreatedAt,
		Tags:      question.Tags,
		Answers: lo.Map(answers, func(a domain.Answer, _ int) AnswerWithUser {
			u := answerUsers[a.UserID]
			return AnswerWithUser{
				ID: a.ID,
				User: User{
					ID:          u.ID,
					Name:        u.Name,
					DisplayName: u.DisplayName,
					IconURL:     u.IconURL.String(),
					UserType:    string(u.UserType),
				},
				QuestionID: a.QuestionID,
				Content:    a.Content,
				CreatedAt:  a.CreatedAt,
			}
		}),
		Status: question.Status,
	}

	return c.JSON(http.StatusOK, response)
}

func (h *Handler) PostQuestion(c echo.Context) error {
	var request PostQuestionRequest
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if uid := c.Get("userID"); uid != request.UserID {
		log.Printf("uid: %s, request.UserID: %s", uid, request.UserID)
		return echo.NewHTTPError(http.StatusForbidden, "not allowed to create")
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
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
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

	go func() {
		if request.BotPost {
			err := h.trapSvc.PostQuestionInfo(
				request.Title,
				request.Content,
				fmt.Sprintf("https://staqoverflow.trap.show/questions/%s", result.ID),
			)
			if err != nil {
				log.Println(err)
			}
		}
	}()

	return c.JSON(http.StatusCreated, response)
}

func (h *Handler) PutQuestion(c echo.Context) error {
	var req PutQuestionRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	question, err := h.qrepo.FindByID(req.ID)
	if errors.Is(err, repository.ErrNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if uid := c.Get("userID"); uid != question.UserID {
		return echo.NewHTTPError(http.StatusForbidden, "not allowed to update")
	}

	question = &domain.Question{
		ID:      req.ID,
		Title:   req.Title,
		Content: req.Content,
		Tags:    req.Tags,
		Status:  domain.QuestionStatus(req.Status),
	}
	_, err = h.qrepo.Update(question)
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
