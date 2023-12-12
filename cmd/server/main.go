package main

import (
	"net"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/traP-jp/h23w_10/handler"
	"github.com/traP-jp/h23w_10/pkg/infra/repository"
)

func main() {
	db := ConnectDB()
	defer db.Close()

	h := handler.NewHandler(repository.NewQuestionRepository(db), repository.NewAnswerRepository(db))

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.GET("/health", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	e.GET("/questions", h.GetQuestions)
	e.POST("/questions/:id/answers", h.PostAnswer)
	e.GET("/questions/:id", h.GetQuestionByID)

	e.Logger.Fatal(e.Start(":8080"))
}

func ConnectDB() *sqlx.DB {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	config := mysql.Config{
		User:                 getEnvOrDefault("MYSQL_USER", "root"),
		Passwd:               getEnvOrDefault("MYSQL_PASSWORD", "h23w10"),
		Net:                  "tcp",
		Addr:                 net.JoinHostPort(getEnvOrDefault("MYSQL_HOST", "localhost"), getEnvOrDefault("MYSQL_PORT", "3306")),
		DBName:               getEnvOrDefault("MYSQL_DATABASE", "h23w10"),
		Loc:                  jst,
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	db, err := sqlx.Open("mysql", config.FormatDSN())
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}
	return db
}

func getEnvOrDefault(key, def string) string {
	value := os.Getenv(key)
	if value == "" {
		return def
	}
	return value
}
