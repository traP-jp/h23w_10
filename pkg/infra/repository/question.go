package repository

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/traP-jp/h23w_10/pkg/domain"
)

type QuestionRepository struct {
	db *sqlx.DB
}

func NewQuestionRepository(db *sqlx.DB) *QuestionRepository {
	return &QuestionRepository{
		db: db,
	}
}

type Question struct {
	ID        string    `db:"id"`
	UserID    string    `db:"user_id"`
	Title     string    `db:"title"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
}

type Tag struct {
	ID   string `db:"id"`
	Name string `db:"name"`
}

type QuestionTag struct {
	QuestionID string `db:"question_id"`
	TagID      string `db:"tag_id"`
}

func (r *QuestionRepository) Find(limit, offset int) ([]domain.Question, error) {
	var questions []Question
	err := r.db.Select(&questions, "SELECT * FROM questions ORDER BY created_at DESC LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return nil, err
	}
	ids := make([]string, len(questions))
	for i, q := range questions {
		ids[i] = q.ID
	}
	var tags []Tag
	err = r.db.Select(&tags, "SELECT * FROM tags  WHERE id IN (?)", ids)
	if err != nil {
		return nil, err
	}

	return results, nil
}
