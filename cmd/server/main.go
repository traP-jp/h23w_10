package main

import (
	"encoding/hex"
	"errors"
	"net"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	mysql_migrate "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/traP-jp/h23w_10/handler"
	"github.com/traP-jp/h23w_10/pkg/infra/repository"
)

func main() {
	db := ConnectDB()
	defer db.Close()
	driver, err := mysql_migrate.WithInstance(db.DB, &mysql_migrate.Config{})
	if err != nil {
		panic(err)
	}
	m, err := migrate.NewWithDatabaseInstance("file://db/migrations", "mysql", driver)
	if err != nil {
		panic(err)
	}
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		panic(err)
	}

	sessionSecret, err := hex.DecodeString(getEnvOrDefault("SESSION_SECRET", "12345678"))

	h := handler.NewHandler(repository.NewQuestionRepository(db), repository.NewAnswerRepository(db), repository.NewUserRepository(db))

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(session.Middleware(sessions.NewCookieStore(sessionSecret)))

	e.GET("/health", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	e.GET("/questions", h.GetQuestions)
	e.POST("/questions", h.PostQuestion)
	e.GET("/questions/:id", h.GetQuestionByID)
	e.POST("/questions/:id/answers", h.PostAnswer)
	e.GET("/tags", h.GetTags)
	e.POST("/tags", h.PostTag)
	e.GET("/users/:id", h.GetUserByID)

	e.Logger.Fatal(e.Start(":8080"))
}

func ConnectDB() *sqlx.DB {
	// wait for mysql
	time.Sleep(10 * time.Second)
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	config := mysql.Config{
		User:                 getEnvOrDefault("MYSQL_USER", "root"),
		Passwd:               getEnvOrDefault("MYSQL_PASSWORD", "h23w10"),
		Net:                  "tcp",
		Addr:                 net.JoinHostPort(getEnvOrDefault("MYSQL_HOST", "db"), getEnvOrDefault("MYSQL_PORT", "3306")),
		DBName:               getEnvOrDefault("MYSQL_DATABASE", "h23w10"),
		Loc:                  jst,
		AllowNativePasswords: true,
		MultiStatements:      true,
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
