package main

import (
	"fmt"
	"todos-api/config"
	"todos-api/internal/app_error"
	"todos-api/internal/database"
	"todos-api/internal/todos"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func main() {
	config.SetupConfig()
	db := database.SetupDatabase()
	setupServer(db)
}

func setupServer(db *gorm.DB) {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			if responseErr, ok := err.(*app_error.ResponseError); ok {
				ctx.Status(responseErr.Code)
				return ctx.JSON(responseErr.Response())
			} else {
				ctx.Status(fiber.StatusInternalServerError)
				return ctx.SendString(err.Error())
			}
		},
	})

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Healthy!")
	})

	todos.AddTodosRoute(app, db)

	app.Listen(fmt.Sprintf(":%v", viper.Get("PORT")))
}
