package repository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/traP-jp/h23w_10/pkg/domain"
	"github.com/traP-jp/h23w_10/pkg/domain/repository"
)

type AnswerRepository struct {
	db *sqlx.DB
}

func NewAnswerRepository(db *sqlx.DB) *AnswerRepository {
	return &AnswerRepository{
		db: db,
	}
}

type Answer struct {
	ID         string    `db:"id"`
	QuestionID string    `db:"question_id"`
	UserID     string    `db:"user_id"`
	Content    string    `db:"content"`
	CreatedAt  time.Time `db:"created_at"`
}

func (r *AnswerRepository) FindByQuestionID(questionID string) ([]domain.Answer, error) {
	var answers []Answer
	err := r.db.Select(&answers, "SELECT * FROM answers WHERE question_id = ? ORDER BY created_at ASC", questionID)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, repository.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	result := make([]domain.Answer, len(answers))
	for i, answer := range answers {
		result[i] = fromAnswerModel(answer)
	}

	return result, nil
}

func fromAnswerModel(answer Answer) domain.Answer {
	return domain.Answer{
		ID:         answer.ID,
		QuestionID: answer.QuestionID,
		UserID:     answer.UserID,
		Content:    answer.Content,
		CreatedAt:  answer.CreatedAt,
	}
}
