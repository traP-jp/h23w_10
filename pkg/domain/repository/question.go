package repository

import "github.com/traP-jp/h23w_10/pkg/domain"

type QuestionRepository interface {
	Find(limit, offset int) ([]domain.Question, error)
	// answerは含めないで返す
	FindByID(id string) (*domain.Question, error)
	FindByTagID(tagID string, limit, offset int) ([]domain.Question, error)
	Create(question *domain.Question) (*domain.Question, error)
	CreateTag(tag *domain.Tag) (*domain.Tag, error)
}
