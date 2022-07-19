package helper

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(message string, code int,
	status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	json_response := Response{
		Meta: meta,
		Data: data,
	}
	return json_response
}

func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}

func GenerateUUID() string {
	id := uuid.New()
	fmt.Println(id.String())

	return id.String()
}
