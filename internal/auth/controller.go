package auth

import (
	"todos-api/internal/validation"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	Repo *AuthRepository
}

func NewAuthController(repo *AuthRepository) AuthController {
	return AuthController{repo}
}

type UserBody struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (controller AuthController) signupHandler(ctx *fiber.Ctx) error {

	repo := controller.Repo

	user := new(UserBody)
	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	if err := validator.New().Struct(user); err != nil {
		validationError := err.(validator.ValidationErrors)
		return &validation.InvalidError{Errors: validationError}
	}

	err := repo.CreateNewUser(user.Username, user.Password)
	if err != nil {
		return err
	}
	return ctx.JSON(fiber.Map{"code": 200, "message": "signed up successfully"})
}
func (controller AuthController) loginHandler(ctx *fiber.Ctx) error {

	repo := controller.Repo
	user := new(UserBody)
	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	if err := validator.New().Struct(user); err != nil {
		validationError := err.(validator.ValidationErrors)
		return &validation.InvalidError{Errors: validationError}
	}

	res, err := repo.AuthenticateUser(user.Username, user.Password)
	if err != nil {
		return err
	}
	ctx.Cookie(&fiber.Cookie{Name: "jwt", Value: res})
	return ctx.JSON(fiber.Map{"code": fiber.StatusOK, "message": "logged in"})
}

func (controller AuthController) logoutHandler(ctx *fiber.Ctx) error {
	ctx.ClearCookie("jwt")
	return ctx.JSON(fiber.Map{"code": fiber.StatusOK, "message": "logged out"})
}
