package todos

import (
	"errors"
	"todos-api/internal/app_error"

	"github.com/gofiber/fiber/v2"
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

type ApiTodo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (repo TodosRepository) GetTodos() ([]ApiTodo, error) {
	todos := []ApiTodo{}
	res := repo.db.Model(&Todo{}).Find(&todos)

	if res.Error != nil {

		return nil, res.Error
	}

	return todos, nil
}
func (repo TodosRepository) CreateTodo(todo *Todo) error {
	res := repo.db.Create(todo)

	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (repo TodosRepository) UpdateTodo(name string, todoUpdate *Todo) error {
	var todo Todo
	res := repo.db.Where(&Todo{Name: name}).First(&todo)

	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return &app_error.ResponseError{Code: fiber.StatusNotFound, Message: "Todo not found"}
		}
		return res.Error
	}

	res = repo.db.Model(&todo).Save(todoUpdate)

	if res.Error != nil {
		return res.Error
	}
	return nil
}
func (repo TodosRepository) DeleteTodo(name string) error {
	var todo Todo
	res := repo.db.Where(&Todo{Name: name}).First(&todo)

	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return &app_error.ResponseError{Code: fiber.StatusNotFound, Message: "Todo not found"}
		}
		return res.Error
	}

	res = repo.db.Delete(&todo)

	if res.Error != nil {
		return res.Error
	}
	return nil
}
