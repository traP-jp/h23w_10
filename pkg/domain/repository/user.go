package repository

import "github.com/traP-jp/h23w_10/pkg/domain"

type UserRepository interface {
	FindUserByID(userID string) (*domain.User, error)
}
