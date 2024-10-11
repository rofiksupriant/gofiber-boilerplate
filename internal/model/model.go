package model

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type ApiResponse[T any] struct {
	Code    int           `json:"code"`
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Data    T             `json:"data"`
	Paging  *PageMetaData `json:"paging,omitempty"`
	Errors  []string      `json:"errors,omitempty"`
}

type PageMetaData struct {
	Page      int    `json:"page"`
	Size      int    `json:"size"`
	TotalItem uint64 `json:"total_item"`
	TotalPage uint64 `json:"total_page"`
}

func (r *ApiResponse[T]) Success(message string, data T) *ApiResponse[T] {
	r.Code = 200
	r.Status = "OK"
	r.Message = message
	r.Data = data
	return r
}

func (r *ApiResponse[T]) SuccessWithPaging(message string, data T, pageMetaData *PageMetaData) *ApiResponse[T] {
	r.Code = 200
	r.Status = "OK"
	r.Message = message
	r.Data = data
	r.Paging = pageMetaData
	return r
}

func (r *ApiResponse[T]) BadRequest(error error) *ApiResponse[T] {
	var errors []string

	if ValidationErrors, ok := error.(validator.ValidationErrors); ok {
		for _, er := range ValidationErrors {
			errors = append(errors, GenerateCustomMessage(er))
		}
	} else {
		errors = append(errors, error.Error())
	}

	r.Code = 400
	r.Status = "BAD_REQUEST"
	r.Message = errors[0]
	r.Errors = errors
	return r
}

func (r *ApiResponse[T]) NotFound(message string) *ApiResponse[T] {
	r.Code = 404
	r.Status = "NOT_FOUND"
	r.Message = message
	return r
}

func (r *ApiResponse[T]) ServerError(message string) *ApiResponse[T] {
	r.Code = 500
	r.Status = "SERVER_ERROR"
	r.Message = message
	return r
}

func GenerateCustomMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("Field '%s' is required", err.Field())
	case "email":
		return fmt.Sprintf("Field '%s' must be a valid email", err.Field())
	case "min":
		return fmt.Sprintf("Field '%s' minimum value: %s", err.Field(), err.Param())
	case "max":
		return fmt.Sprintf("Field '%s' max value: %s", err.Field(), err.Param())
	default:
		return fmt.Sprintf("Field '%s' is invalid", err.Field())
	}
}
