package repository

import "errors"

var (
	ErrNotFound = errors.New("not found")

	ErrTagNotFound      = errors.New("tag not found")
	ErrTagAlreadyExists = errors.New("tag already exists")
)
