package todos

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
}

func NewTodosRepo(db *gorm.DB) *TodosRepository {
	return &TodosRepository{db}
}

type TodosRepository struct {
	DB *gorm.DB
}

type ApiTodo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (repo TodosRepository) GetTodos() ([]ApiTodo, error) {
	todos := []ApiTodo{}
	res := repo.DB.Model(&Todo{}).Find(&todos)

	if res.Error != nil {

		return nil, res.Error
	}

	return todos, nil
}
func (repo TodosRepository) CreateTodo(todo *NewTodo) error {
	todoRecord := &Todo{Name: todo.Name, Description: todo.Description}
	res := repo.DB.Create(todoRecord)

	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (repo TodosRepository) UpdateTodo(name string, todoUpdate *TodoUpdate) error {
	var todo Todo
	res := repo.DB.Where(&Todo{Name: name}).First(&todo)

	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return &fiber.Error{Code: fiber.StatusNotFound, Message: "Todo not found"}
		}
		return res.Error
	}

	todoUpdated := &Todo{Name: todoUpdate.Name, Description: todoUpdate.Description}

	res = repo.DB.Model(&todo).Updates(todoUpdated)

	if res.Error != nil {
		return res.Error
	}
	return nil
}
func (repo TodosRepository) DeleteTodo(name string) (uint, error) {
	var todo Todo
	var id uint
	res := repo.DB.Where(&Todo{Name: name}).First(&todo)

	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return id, &fiber.Error{Code: fiber.StatusNotFound, Message: "Todo not found"}
		}
		return id, res.Error
	}

	id = todo.ID
	res = repo.DB.Delete(&todo)

	if res.Error != nil {
		return id, res.Error
	}
	return todo.ID, nil
}
