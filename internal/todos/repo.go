package todos

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
}

func NewTodosRepo(db *gorm.DB) *TodosRepository {
	return &TodosRepository{db: db}
}

type TodosRepository struct {
	db *gorm.DB
}

func (repo TodosRepository) GetTodos() ([]Todo, error) {
	todos := []Todo{}
	res := repo.db.Find(&todos)

	if res.Error != nil {
		return nil, res.Error
	}

	return todos, nil
}
