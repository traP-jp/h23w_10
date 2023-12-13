package repository

import (
	"net/url"

	"github.com/jmoiron/sqlx"
	"github.com/traP-jp/h23w_10/pkg/domain"
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
	ID       string `db:"id"`
	Name     string `db:"name"`
	IconURL  string `db:"icon_url"`
	UserType string `db:"user_type"`
}

func (r *UserRepository) FindUserByID(id string) (*domain.User, error) {
	var user User
	err := r.db.Get(&user, "SELECT * FROM users WHERE id = ?", id)
	if err != nil {
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
