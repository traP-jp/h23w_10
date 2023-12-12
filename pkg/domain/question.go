package domain

import "time"

type Question struct {
	ID        string
	UserID    string
	Title     string
	Content   string
	CreatedAt time.Time
	Tags      []Tag
	Answers   []Answer
	Status    QuestionStatus
}

type Tag struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type QuestionStatus string

const (
	QuestionStatusOpen   QuestionStatus = "open"
	QuestionStatusClosed QuestionStatus = "closed"
)
