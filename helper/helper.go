package helper

import (
	"strings"

	"github.com/go-playground/validator/v10"
)
type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type ResponseError struct {
	Meta 	Meta        `json:"meta"`
	Errors 	interface{} `json:"errors"`
}
type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}
	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}

func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}

func APIResponseError(message string, code int, status string, err string) ResponseError {
	splittedError := strings.Split(err, "\n")

	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}
	jsonResponse := ResponseError{
		Meta: meta,
		Errors: splittedError,
		
	}

	return jsonResponse
}
