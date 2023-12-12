package repository

import (
	"database/sql"
	"errors"
	"slices"
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
	StatusID  int       `db:"status_id"`
}

type Tag struct {
	ID   string `db:"id"`
	Name string `db:"name"`
}

type QuestionTag struct {
	QuestionID string `db:"question_id"`
	TagID      string `db:"tag_id"`
}

type QuestionStatus struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

func (r *QuestionRepository) Find(limit, offset int) ([]domain.Question, error) {
	// get status
	statuses, err := r.getAllStatuses()
	if err != nil {
		return nil, err
	}

	// get questions
	questions := make(map[string]*domain.Question)
	rows, err := r.db.Queryx("SELECT * FROM questions ORDER BY created_at DESC LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var question Question
	var ids []string
	for rows.Next() {
		err := rows.StructScan(&question)
		if err != nil {
			return nil, err
		}
		questions[question.ID] = fromQuestionModel(question, statuses[question.StatusID])
		ids = append(ids, question.ID)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	if len(ids) == 0 {
		return nil, nil
	}

	// get tags
	query, params, err := sqlx.In("SELECT tags.id, tags.name, qt.question_id FROM tags INNER JOIN question_tags qt ON tags.id = qt.tag_id WHERE qt.question_id IN (?)", ids)
	if err != nil {
		return nil, err
	}
	rows, err = r.db.Queryx(query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var (
		tagID      string
		tagName    string
		questionID string
	)
	for rows.Next() {
		err := rows.Scan(&tagID, &tagName, &questionID)
		if err != nil {
			return nil, err
		}
		questions[questionID].Tags = append(questions[questionID].Tags, domain.Tag{ID: tagID, Name: tagName})
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// map to slice
	result := make([]domain.Question, 0, len(questions))
	for _, question := range questions {
		result = append(result, *question)
	}
	// sort by created_at desc
	slices.SortFunc(result, func(a, b domain.Question) int {
		return b.CreatedAt.Compare(a.CreatedAt)
	})

	return result, nil
}

func (r *QuestionRepository) FindByID(id string) (*domain.Question, error) {
	// get status
	statuses, err := r.getAllStatuses()
	if err != nil {
		return nil, err
	}

	// get question
	var question Question
	err = r.db.Get(&question, "SELECT * FROM questions WHERE id = ?", id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	result := fromQuestionModel(question, statuses[question.StatusID])

	// get tags
	rows, err := r.db.Queryx("SELECT tags.id, tags.name FROM tags INNER JOIN question_tags qt ON tags.id = qt.tag_id WHERE qt.question_id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tag Tag
	for rows.Next() {
		err := rows.StructScan(&tag)
		if err != nil {
			return nil, err
		}
		result.Tags = append(result.Tags, domain.Tag{ID: tag.ID, Name: tag.Name})
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *QuestionRepository) Create(question *domain.Question) (*domain.Question, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// insert question
	statusID, err := r.getStatusIDByName(string(question.Status))
	if err != nil {
		return nil, err
	}
	questionModel := Question{
		ID:        question.ID,
		UserID:    question.UserID,
		Title:     question.Title,
		Content:   question.Content,
		CreatedAt: question.CreatedAt,
		StatusID:  statusID,
	}
	_, err = tx.NamedExec("INSERT INTO questions (id, user_id, title, content, created_at, status_id) VALUES (:id, :user_id, :title, :content, :created_at, :status_id)", questionModel)
	if err != nil {
		return nil, err
	}

	// insert tags
	for _, tag := range question.Tags {
		// check tag exists
		count := 0
		err := tx.Get(&count, "SELECT COUNT(*) FROM tags WHERE id = ?", tag.ID)
		if count == 0 {
			return nil, errors.New("tag not found")
		}

		// insert question_tag
		_, err = tx.Exec("INSERT INTO question_tags (question_id, tag_id) VALUES (?, ?)", question.ID, tag.ID)
		if err != nil {
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return question, nil
}

func (r *QuestionRepository) getAllStatuses() (map[int]string, error) {
	statuses := make(map[int]string)
	rows, err := r.db.Queryx("SELECT * FROM question_statuses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var status QuestionStatus
	for rows.Next() {
		err := rows.StructScan(&status)
		if err != nil {
			return nil, err
		}
		statuses[status.ID] = status.Name
	}
	return statuses, nil
}

func (r *QuestionRepository) getStatusIDByName(name string) (int, error) {
	var status QuestionStatus
	err := r.db.Get(&status, "SELECT * FROM question_statuses WHERE name = ?", name)
	if errors.Is(err, sql.ErrNoRows) {
		return 0, nil
	} else if err != nil {
		return 0, err
	}
	return status.ID, nil
}

func fromQuestionModel(question Question, status string) *domain.Question {
	return &domain.Question{
		ID:        question.ID,
		UserID:    question.UserID,
		Title:     question.Title,
		Content:   question.Content,
		CreatedAt: question.CreatedAt,
		Status:    domain.QuestionStatus(status),
	}
}