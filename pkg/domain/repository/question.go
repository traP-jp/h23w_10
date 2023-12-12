package repository

import "github.com/traP-jp/h23w_10/pkg/domain"

type QuestionRepository interface {
	Find(condition *FindQuestionsCondition) ([]domain.Question, error)
	// answerは含めないで返す
	FindByID(id string) (*domain.Question, error)
	FindByTagID(tagID string, condition *FindQuestionsCondition) ([]domain.Question, error)
	Create(question *domain.Question) (*domain.Question, error)
}

type FindQuestionsCondition struct {
	Limit    int
	Offset   int
	Statuses []domain.QuestionStatus
}
