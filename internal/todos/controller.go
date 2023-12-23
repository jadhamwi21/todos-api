package todos

import (
	"fmt"
	"todos-api/internal/validation"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type TodosController struct {
	Repo *TodosRepository
}

func NewTodosController(repo *TodosRepository) TodosController {
	return TodosController{repo}
}

func (controller TodosController) getTodosHandler(ctx *fiber.Ctx) error {
	repo := controller.Repo
	username := ctx.Locals("username").(string)

	todos, err := repo.GetTodos(username)
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
	repo := controller.Repo
	todo := new(NewTodo)
	if err := ctx.BodyParser(todo); err != nil {
		return err
	}

	if err := validator.New().Struct(todo); err != nil {
		validationError := err.(validator.ValidationErrors)
		return &validation.InvalidError{Errors: validationError}
	}

	username := ctx.Locals("username").(string)
	if err := repo.CreateTodo(todo, username); err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{"code": fiber.StatusOK, "data": todo})
}

type TodoUpdate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (controller TodosController) updateTodoHandler(ctx *fiber.Ctx) error {
	repo := controller.Repo
	todo := &TodoUpdate{}
	if err := ctx.BodyParser(todo); err != nil {
		fmt.Printf("error = %v", err)
		return err
	}

	if err := validator.New().Struct(todo); err != nil {
		validationError := err.(validator.ValidationErrors)
		return &validation.InvalidError{Errors: validationError}
	}

	name := ctx.Params("name")
	username := ctx.Locals("username").(string)

	if err := repo.UpdateTodo(username, name, todo); err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{"code": fiber.StatusOK, "data": todo})
}
func (controller TodosController) deleteTodoHandler(ctx *fiber.Ctx) error {
	repo := controller.Repo

	name := ctx.Params("name")
	username := ctx.Locals("username").(string)
	id, err := repo.DeleteTodo(username, name)
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{"code": fiber.StatusOK, "message": fmt.Sprintf("deleted todo %v", id)})
}
