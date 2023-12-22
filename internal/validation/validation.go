package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type InvalidError struct {
	Errors validator.ValidationErrors
}

func (err *InvalidError) Error() string {
	return err.Error()
}

func TagFormatter(field string, tag string) string {
	switch tag {
	case "required":
		return fmt.Sprintf("%v is required", field)
	}
	return tag
}

func (err *InvalidError) JSON() map[string]interface{} {
	errors := make(map[string]interface{})
	for _, currentErr := range err.Errors {
		errors[currentErr.Field()] = TagFormatter(currentErr.Field(), currentErr.Tag())
	}

	response := make(map[string]interface{})
	response["code"] = fiber.StatusBadRequest
	response["errors"] = errors

	return response
}
