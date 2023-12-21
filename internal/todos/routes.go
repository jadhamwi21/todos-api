package todos

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AddTodosRoute(app *fiber.App, db *gorm.DB) {

	todosRepo := NewTodosRepo(db)
	todosController := NewTodosController(todosRepo)
	router := app.Group("/todos")

	router.Get("/", todosController.getTodosHandler)
	router.Post("/", todosController.createTodoHandler)
	router.Put("/:name", todosController.updateTodoHandler)
	router.Delete("/:name", todosController.deleteTodoHandler)
}
