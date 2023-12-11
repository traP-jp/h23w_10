package domain

import "time"

type Answer struct {
	ID         string
	UserID     string
	QuestionID string
	Content    string
	CreatedAt  time.Time
}
