package main

import (
	"fmt"
	"todos-api/config"
	"todos-api/internal/database"

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
	app := fiber.New()

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Healthy!")
	})

	app.Listen(fmt.Sprintf(":%v", viper.Get("PORT")))
}
