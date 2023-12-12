package repository

import (
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/traP-jp/h23w_10/pkg/domain"
	"github.com/traP-jp/h23w_10/pkg/domain/repository"
)

// エラーが出ないことを確認するためだけのテスト

func TestFindAllQuestions(t *testing.T) {
	db := NewDB(t)
	defer db.Close()

	repo := NewQuestionRepository(db)
	questions, err := repo.Find(&repository.FindQuestionsCondition{
		Limit:  10,
		Offset: 0,
		Statuses: []domain.QuestionStatus{
			domain.QuestionStatusClosed,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", questions)
}

func TestFindQuestionsByTagID(t *testing.T) {
	db := NewDB(t)
	defer db.Close()

	repo := NewQuestionRepository(db)
	questions, err := repo.FindByTagID("bc6c1c8d-9898-11ee-906b-0242ac120002", &repository.FindQuestionsCondition{
		Limit:  10,
		Offset: 0,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", questions)
}

func TestFindQuestionByID(t *testing.T) {
	db := NewDB(t)
	defer db.Close()

	repo := NewQuestionRepository(db)
	question, err := repo.FindByID("d6e88dd1-9892-11ee-906b-0242ac120002")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", question)
}

func TestCreateQuestion(t *testing.T) {
	db := NewDB(t)
	defer db.Close()

	repo := NewQuestionRepository(db)
	question := &domain.Question{
		ID:        domain.NewUUID(),
		UserID:    domain.NewUUID(),
		Title:     "test",
		Content:   "test",
		CreatedAt: time.Now(),
		Tags: []domain.Tag{
			{
				ID: "bc6c1c8d-9898-11ee-906b-0242ac120002",
			},
		},
		Status: domain.QuestionStatusOpen,
	}
	_, err := repo.Create(question)
	if err != nil {
		t.Fatal(err)
	}
}

// 流れを確認するためのテスト
func TestQuestionRepository(t *testing.T) {
	db := NewDB(t)
	defer db.Close()

	repo := NewQuestionRepository(db)

	var (
		questionID = domain.NewUUID()
		tagID      = domain.NewUUID()
	)

	// Create Tag
	t.Run("Create Tag", func(t *testing.T) {
		tag := &domain.Tag{
			ID:   tagID,
			Name: "test",
		}
		_, err := repo.CreateTag(tag)
		if err != nil {
			t.Fatal(err)
		}
	})

	// Create Question
	t.Run("Create Question", func(t *testing.T) {
		question := &domain.Question{
			ID:        questionID,
			UserID:    domain.NewUUID(),
			Title:     "test title",
			Content:   "test content",
			CreatedAt: time.Now(),
			Tags: []domain.Tag{
				{ID: tagID},
			},
			Status: domain.QuestionStatusOpen,
		}
		_, err := repo.Create(question)
		if err != nil {
			t.Fatal(err)
		}
	})

	// Find Question
	t.Run("Find Question", func(t *testing.T) {
		questions, err := repo.Find(&repository.FindQuestionsCondition{
			Limit:    10,
			Offset:   0,
			Statuses: []domain.QuestionStatus{domain.QuestionStatusOpen, domain.QuestionStatusClosed},
		})
		if err != nil {
			t.Fatal(err)
		}
		if len(questions) != 1 {
			t.Errorf("len(questions) = %d, want %d", len(questions), 1)
		}
		q := questions[0]
		if q.Title == "" || q.Content == "" || q.CreatedAt.IsZero() || len(q.Tags) != 1 || q.Status != domain.QuestionStatusOpen {
			t.Errorf("invalid question: %+v", q)
		}
	})

	// Find Question By ID
	t.Run("Find Question By ID", func(t *testing.T) {
		q, err := repo.FindByID(questionID)
		if err != nil {
			t.Fatal(err)
		}
		if q.Title == "" || q.Content == "" || q.CreatedAt.IsZero() || len(q.Tags) != 1 || q.Status != domain.QuestionStatusOpen {
			t.Errorf("invalid question: %+v", q)
		}
	})

	// Find Question By Tag ID
	t.Run("Find Question By Tag ID", func(t *testing.T) {
		questions, err := repo.FindByTagID(tagID, &repository.FindQuestionsCondition{
			Limit:    10,
			Offset:   0,
			Statuses: []domain.QuestionStatus{domain.QuestionStatusOpen, domain.QuestionStatusClosed},
		})
		if err != nil {
			t.Fatal(err)
		}
		if len(questions) != 1 {
			t.Errorf("len(questions) = %d, want %d", len(questions), 1)
		}
	})
	t.Run("Find Question By Tag ID (invalid)", func(t *testing.T) {
		questions, err := repo.FindByTagID("invalid", &repository.FindQuestionsCondition{
			Limit:    10,
			Offset:   0,
			Statuses: []domain.QuestionStatus{domain.QuestionStatusOpen, domain.QuestionStatusClosed},
		})
		if err != nil {
			t.Fatal(err)
		}
		if len(questions) != 0 {
			t.Errorf("len(questions) = %d, want %d", len(questions), 0)
		}
	})
}
