package validator

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/aritradevelops/authinfinity/server/internal/middlewares/translator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

type ValidationError struct {
	Message string `json:"message,omitempty"`
	Field   string `json:"field,omitempty"`
	Value   any    `json:"value,omitempty"`
}

func Validate(data any, c *fiber.Ctx) []ValidationError {
	// TODO: consider it moving into an init function
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})
	errs := validate.Struct(data)
	if errs != nil {
		errors := []ValidationError{}
		for _, err := range errs.(validator.ValidationErrors) {
			errors = append(errors, ValidationError{
				Message: translator.Localize(c, fmt.Sprintf("validation.%s", err.Tag()), map[string]any{"Field": Labelize(err.Field()), "Value": err.Value(), "Param": err.Param()}),
				Field:   err.Field(),
				Value:   err.Value(),
			})
		}
		return errors
	}
	return nil
}

func Labelize(field string) string {
	formatted := []string{}
	for _, part := range strings.Split(field, "_") {
		formatted = append(formatted, strings.Title(part))
	}
	fmt.Println("field", field, "label", strings.Join(formatted, " "))
	return strings.Join(formatted, " ")
}
