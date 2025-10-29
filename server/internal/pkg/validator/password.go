package validator

import "github.com/gofiber/fiber/v2"

type PasswordConstraints struct {
	minLength int
	maxLength int
}

func ValidatePassword(c *fiber.Ctx, password string) *ValidationError {
	constraints := PasswordConstraints{
		minLength: 8,
		maxLength: 50,
	}
	if password == "" {
		return &ValidationError{
			Field:   "password",
			Message: "",
			Value:   password,
		}
	}
	if len(password) < constraints.minLength {
		return &ValidationError{
			Field:   "password",
			Message: "",
			Value:   password,
		}
	}
	return nil
}
