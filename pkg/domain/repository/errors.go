package repository

import "errors"

var (
	ErrNotFound = errors.New("not found")

	ErrTagAlreadyExists = errors.New("tag already exists")
)
