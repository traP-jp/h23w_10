package repository

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

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
