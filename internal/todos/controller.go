package todos

import (
	"fmt"
	"todos-api/internal/validation"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type TodosController struct {
	repo *TodosRepository
}

func NewTodosController(repo *TodosRepository) TodosController {
	return TodosController{repo}
}

func (controller TodosController) getTodosHandler(ctx *fiber.Ctx) error {
	repo := controller.repo

	todos, err := repo.GetTodos()
	if err != nil {
		return err
	}

	return ctx.JSON(todos)
}

type NewTodo struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

func (controller TodosController) createTodoHandler(ctx *fiber.Ctx) error {
	repo := controller.repo

	todo := new(NewTodo)
	if err := ctx.BodyParser(todo); err != nil {
		return err
	}

	if err := validator.New().Struct(todo); err != nil {
		validationError := err.(validator.ValidationErrors)
		return &validation.InvalidError{Errors: validationError}
	}

	if err := repo.CreateTodo(todo); err != nil {
		return err
	}

	return ctx.SendString("Todo Created")
}

type TodoUpdate struct {
	PrevName    string `json:"prevName" validate:"required"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (controller TodosController) updateTodoHandler(ctx *fiber.Ctx) error {
	repo := controller.repo
	todo := &TodoUpdate{}
	if err := ctx.BodyParser(todo); err != nil {
		fmt.Printf("error = %v", err)
		return err
	}
	name := ctx.Params("name")

	todo.PrevName = name

	if err := validator.New().Struct(todo); err != nil {
		validationError := err.(validator.ValidationErrors)
		return &validation.InvalidError{Errors: validationError}
	}

	err := repo.UpdateTodo(name, todo)
	if err != nil {
		return err
	}
	return ctx.SendString("Todo Updated")
}
func (controller TodosController) deleteTodoHandler(ctx *fiber.Ctx) error {
	repo := controller.repo

	name := ctx.Params("name")

	print(name)
	err := repo.DeleteTodo(name)
	if err != nil {
		return err
	}
	return ctx.SendString("Todo Deleted")
}
