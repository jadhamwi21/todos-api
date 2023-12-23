package auth

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AddAuthRoutes(app *fiber.App, db *gorm.DB) {
	authRepo := NewAuthRepo(db)
	authController := NewAuthController(authRepo)
	router := app.Group("/auth")

	router.Post("/signup", authController.signupHandler)
	router.Post("/login", authController.loginHandler)
	router.Post("/logout", AuthMiddleware, authController.loginHandler)
}
