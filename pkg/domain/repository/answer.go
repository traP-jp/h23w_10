package repository

import "github.com/traP-jp/h23w_10/pkg/domain"

type AnswerRepository interface {
	FindByQuestionID(id string) ([]domain.Answer, error)
	Create(answer domain.Answer) (*domain.Answer, error)
}
