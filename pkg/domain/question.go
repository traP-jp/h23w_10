package domain

import "time"

type Question struct {
	ID        string         `json:"id,omitempty"`
	UserID    string         `json:"user_id,omitempty"`
	Title     string         `json:"title,omitempty"`
	Content   string         `json:"content,omitempty"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	Tags      []Tag          `json:"tags,omitempty"`
	Answers   []Answer       `json:"answers,omitempty"`
	Status    QuestionStatus `json:"status,omitempty"`
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

func AvailableQuestionStatus() []QuestionStatus {
	return []QuestionStatus{
		QuestionStatusOpen,
		QuestionStatusClosed,
	}
}
