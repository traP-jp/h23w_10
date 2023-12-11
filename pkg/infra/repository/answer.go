package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/traP-jp/h23w_10/pkg/domain"
)

type AnswerRepository struct {
	db *sqlx.DB
}

func (r *AnswerRepository) FindByQuestionID(questionID int) ([]domain.Answer, error) {
	var answers []domain.Answer
	err := r.db.Select(&answers, "SELECT * FROM answers WHERE question_id=?", questionID)
	if err != nil {
		return nil, err
	}
	return answers, nil
}
