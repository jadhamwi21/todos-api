package tests

import (
	"testing"
	"todos-api/internal/todos"

	"gorm.io/gorm"
)

func TestGenerateTodosRepository(t *testing.T) {

	db := &gorm.DB{}
	got := todos.NewTodosRepo(db)
	want := &todos.TodosRepository{DB: db}
	if *got != *want {
		t.Errorf("invalid repository generated")
	}
}
