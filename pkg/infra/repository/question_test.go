package repository

import (
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/traP-jp/h23w_10/pkg/domain"
)

// エラーが出ないことを確認するためだけのテスト

func NewDB(t *testing.T) *sqlx.DB {
	dsn := "root:h23w10@tcp(localhost:3306)/h23w10?parseTime=true&loc=Asia%2FTokyo"
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		t.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}
	return db
}

func TestFindAllQuestions(t *testing.T) {
	db := NewDB(t)
	defer db.Close()

	repo := NewQuestionRepository(db)
	questions, err := repo.Find(10, 0)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", questions)
}

func TestFindQuestionByID(t *testing.T) {
	db := NewDB(t)
	defer db.Close()

	repo := NewQuestionRepository(db)
	question, err := repo.FindByID("d6e88dd1-9892-11ee-906b-0242ac120002")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", question)
}

func TestCreateQuestion(t *testing.T) {
	db := NewDB(t)
	defer db.Close()

	repo := NewQuestionRepository(db)
	question := &domain.Question{
		ID:        domain.NewUUID(),
		UserID:    domain.NewUUID(),
		Title:     "test",
		Content:   "test",
		CreatedAt: time.Now(),
		Tags: []domain.Tag{
			{
				ID: "bc6c1c8d-9898-11ee-906b-0242ac120002",
			},
		},
		Status: domain.QuestionStatusOpen,
	}
	_, err := repo.Create(question)
	if err != nil {
		t.Fatal(err)
	}
}
