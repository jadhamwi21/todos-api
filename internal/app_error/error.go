package app_error

import (
	"fmt"
)

type ResponseError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (err *ResponseError) Error() string {

	return fmt.Sprintf("status : %v\nmessage : %v\n", err.Code, err.Message)
}

func (err *ResponseError) Response() ResponseError {
	return *err
}
