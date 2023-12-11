package handler

import "github.com/traP-jp/h23w_10/pkg/domain/repository"

type Handler struct {
	qrepo repository.QuestionRepository
	arepo repository.AnswerRepository
}
