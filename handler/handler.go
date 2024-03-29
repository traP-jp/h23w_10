package handler

import (
	"encoding/gob"

	"github.com/traP-jp/h23w_10/pkg/domain/repository"
	"github.com/traP-jp/h23w_10/pkg/infra/trap"
	"github.com/traP-jp/h23w_10/pkg/usecase/imggen"
	"golang.org/x/oauth2"
)

type Handler struct {
	trapSvc *trap.TrapService

	qrepo repository.QuestionRepository
	arepo repository.AnswerRepository
	urepo repository.UserRepository

	oauth2Config oauth2.Config

	imggenSvc *imggen.ImggenService
}

func NewHandler(
	trapSvc *trap.TrapService,
	qrepo repository.QuestionRepository,
	arepo repository.AnswerRepository,
	urepo repository.UserRepository,
	oauth2Conf oauth2.Config,
	imggenSvc *imggen.ImggenService,
) *Handler {
	gob.Register(&oauth2.Token{})
	return &Handler{
		trapSvc: trapSvc,

		qrepo: qrepo,
		arepo: arepo,
		urepo: urepo,

		oauth2Config: oauth2Conf,

		imggenSvc: imggenSvc,
	}
}

const (
	sessionName = "h23w_10"
	imagesDir   = "images"
)
