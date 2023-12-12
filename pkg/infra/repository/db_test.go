package repository

import (
	"testing"

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
