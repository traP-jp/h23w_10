package repository

import "github.com/traP-jp/h23w_10/pkg/domain"

type AnswerRepository interface {
	FindByID(id string) (*domain.Answer, error)
	FindByQuestionID(id string) ([]domain.Answer, error)
	Create(answer *domain.Answer) (*domain.Answer, error)
	Update(answer *domain.Answer) (*domain.Answer, error)
}
