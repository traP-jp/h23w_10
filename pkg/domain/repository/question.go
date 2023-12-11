package repository

import "github.com/traP-jp/h23w_10/pkg/domain"

type QuestionRepository interface {
	Find(limit, offset int) ([]domain.Question, error)
	FindByID(id string) (*domain.Question, error)
	Create(question *domain.Question) (*domain.Question, error)
}
