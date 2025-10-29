package core

import (
	"fmt"
	"net/http"

	"github.com/aritradevelops/authinfinity/server/internal/middlewares/translator"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

type ErrorInfo struct {
	ValidationErrors []validator.ValidationError `json:"validation_errors,omitempty"`
}

// implements error
type HttpError struct {
	Message    string
	StatusCode int
	Info       ErrorInfo
}

func (e HttpError) Error() string {
	return e.Message
}

func NewInternalServerError(c *fiber.Ctx) HttpError {
	return HttpError{
		Message:    translator.Localize(c, fmt.Sprintf("errors.%d", http.StatusInternalServerError)),
		StatusCode: http.StatusInternalServerError,
	}
}

func NewNotFoundError(c *fiber.Ctx) HttpError {
	return HttpError{
		Message:    translator.Localize(c, fmt.Sprintf("errors.%d", http.StatusNotFound)),
		StatusCode: http.StatusNotFound,
	}
}

func NewBadRequestError(c *fiber.Ctx) HttpError {
	return HttpError{
		Message:    translator.Localize(c, fmt.Sprintf("errors.%d", http.StatusBadRequest)),
		StatusCode: http.StatusBadRequest,
	}
}

func NewRequestValidationError(c *fiber.Ctx, errors []validator.ValidationError) HttpError {
	return HttpError{
		Message:    translator.Localize(c, fmt.Sprintf("errors.%d", http.StatusUnprocessableEntity)),
		StatusCode: http.StatusUnprocessableEntity,
		Info: ErrorInfo{
			ValidationErrors: errors,
		},
	}
}

func NewConflictError(c *fiber.Ctx) HttpError {
	return HttpError{
		Message:    translator.Localize(c, fmt.Sprintf("errors.%d", http.StatusConflict)),
		StatusCode: http.StatusConflict,
	}
}

func NewDuplicateKeyError(c *fiber.Ctx, key string) HttpError {
	return HttpError{
		Message: translator.Localize(c, "validation.unique", map[string]string{
			"Field": key,
		}),
		StatusCode: http.StatusConflict,
	}
}
