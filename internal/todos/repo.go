package todos

import (
	"errors"
	"todos-api/internal/auth"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Todo struct {
	Name         string `json:"name" gorm:"primaryKey"`
	Description  string `json:"description"`
	User         auth.User
	UserUsername string
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

func (repo TodosRepository) GetTodos(username string) ([]ApiTodo, error) {
	user := auth.User{Username: username}
	repo.DB.First(&user)
	todos := []ApiTodo{}
	res := repo.DB.Model(&Todo{}).Where(&Todo{User: user, UserUsername: username}).Find(&todos)

	if res.Error != nil {

		return nil, res.Error
	}

	return todos, nil
}
func (repo TodosRepository) CreateTodo(todo *NewTodo, username string) error {
	user := auth.User{Username: username}
	repo.DB.First(&user)
	todoRecord := &Todo{Name: todo.Name, Description: todo.Description, User: user}
	res := repo.DB.First(todoRecord)
	if res.Error == nil {
		return &fiber.Error{Code: fiber.StatusConflict, Message: "todo already exists"}
	}
	res = repo.DB.Create(todoRecord)

	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (repo TodosRepository) UpdateTodo(username string, name string, todoUpdate *TodoUpdate) error {
	user := auth.User{Username: username}
	repo.DB.First(&user)
	var todo Todo
	res := repo.DB.Where(&Todo{Name: name, User: user, UserUsername: user.Username}).First(&todo)

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
func (repo TodosRepository) DeleteTodo(username string, name string) (string, error) {
	user := auth.User{Username: username}
	repo.DB.First(&user)
	var todo Todo

	res := repo.DB.Where(&Todo{Name: name, User: user, UserUsername: user.Username}).First(&todo)

	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return "", &fiber.Error{Code: fiber.StatusNotFound, Message: "Todo not found"}
		}
		return "", res.Error
	}

	res = repo.DB.Delete(&todo)

	if res.Error != nil {
		return "", res.Error
	}
	return name, nil
}
