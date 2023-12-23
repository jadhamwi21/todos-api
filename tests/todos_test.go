package tests

import (
	"testing"
	"todos-api/internal/database"
	"todos-api/internal/todos"
)

func TestGenerateTodosRepository(t *testing.T) {

	db, _ := database.SetupDatabase()

	got := todos.NewTodosRepo(db)
	want := &todos.TodosRepository{DB: db}
	if *got != *want {
		t.Errorf("invalid repository generated")
	}
}
func TestGenerateTodosController(t *testing.T) {
	db, _ := database.SetupDatabase()
	repo := todos.NewTodosRepo(db)
	want := todos.NewTodosController(repo)
	got := &todos.TodosController{Repo: repo}

	if *got != want {
		t.Errorf("controllers mismatch")
	}

}
