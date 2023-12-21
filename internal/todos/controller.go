package todos

import (
	"fmt"

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
func (controller TodosController) createTodoHandler(ctx *fiber.Ctx) error {
	repo := controller.repo
	todo := &Todo{}
	if err := ctx.BodyParser(todo); err != nil {
		fmt.Printf("error = %v", err)
		return err
	}
	err := repo.CreateTodo(todo)
	if err != nil {
		return err
	}
	return ctx.SendString("Todo Created")
}
func (controller TodosController) updateTodoHandler(ctx *fiber.Ctx) error {
	repo := controller.repo
	todo := &Todo{}
	if err := ctx.BodyParser(todo); err != nil {
		fmt.Printf("error = %v", err)
		return err
	}
	name := ctx.Params("name")

	print(name)
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
