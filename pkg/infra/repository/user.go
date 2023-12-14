package repository

import (
	"database/sql"
	"errors"
	"net/url"

	"github.com/jmoiron/sqlx"
	"github.com/traP-jp/h23w_10/pkg/domain"
	"github.com/traP-jp/h23w_10/pkg/domain/repository"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

type User struct {
	ID          string `db:"id"`
	Name        string `db:"name"`
	DisplayName string `db:"display_name"`
	IconURL     string `db:"icon_url"`
	UserType    string `db:"user_type"`
}

func (r *UserRepository) FindUserByID(id string) (*domain.User, error) {
	var user User
	err := r.db.Get(&user, "SELECT * FROM users WHERE id = ?", id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, repository.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	url, err := url.Parse(user.IconURL)
	if err != nil {
		return nil, err
	}

	result := domain.User{
		ID:       user.ID,
		Name:     user.Name,
		IconURL:  *url,
		UserType: domain.UserType(user.UserType),
	}
	return &result, nil
}

func (r *UserRepository) Create(user *domain.User) (*domain.User, error) {
	_, err := r.db.Exec("INSERT INTO users (id, name, display_name, icon_url, user_type) VALUES (?, ?, ?, ?, ?)", user.ID, user.Name, user.DisplayName, user.IconURL.String(), user.UserType)
	if err != nil {
		return nil, err
	}

	return user, nil
}
