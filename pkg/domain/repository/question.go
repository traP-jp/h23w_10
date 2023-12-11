package repository

import "github.com/traP-jp/h23w_10/pkg/domain"

type QuestionRepository interface {
	FindAll() ([]domain.Question, error)
	FindByLimitAndOffset(limit int, offset int) ([]domain.Question, error)
	FindByID(id string) (*domain.Question, error)
	Create(question *domain.Question) (*domain.Question, error)
}
