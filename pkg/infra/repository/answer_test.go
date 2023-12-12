package repository

import "testing"

func TestFindAnswersByQuestionID(t *testing.T) {
	db := NewDB(t)
	defer db.Close()

	repo := NewAnswerRepository(db)
	answers, err := repo.FindByQuestionID("117bbc88-9b0f-4687-a9ce-8745ccbf58d2")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", answers)
}
