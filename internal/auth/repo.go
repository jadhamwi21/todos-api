package auth

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type User struct {
	Username string `json:"username" gorm:"primaryKey"`
	Password string `json:"password"`
}

type AuthRepository struct {
	DB *gorm.DB
}

func NewAuthRepo(db *gorm.DB) *AuthRepository {
	return &AuthRepository{DB: db}
}

func (authRepo *AuthRepository) CreateNewUser(username string, password string) error {
	user := &User{Username: username}
	res := authRepo.DB.First(&user)
	if res.Error == nil {
		return &fiber.Error{Code: 409, Message: "username already exists"}
	}
	user = &User{Username: username, Password: password}
	res = authRepo.DB.Create(user)
	if res.Error != nil {

		return res.Error
	}
	return nil
}

func (authRepo *AuthRepository) AuthenticateUser(username string, password string) (string, error) {
	user := &User{Username: username}
	db := authRepo.DB.First(user)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			return "", &fiber.Error{Code: fiber.StatusNotFound, Message: "user not found"}
		}
		return "", db.Error
	}
	if user.Password != password {
		return "", &fiber.Error{Code: fiber.StatusForbidden, Message: "incorrect password"}
	}
	token, err := GenerateJwtToken(user.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}
