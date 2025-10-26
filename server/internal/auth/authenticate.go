package auth

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

const authKey = "user"

type AuthUser struct {
	ID        string
	AccountID string
}

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals(authKey, &AuthUser{
			ID:        "9c35b66c-bf36-4466-b793-4ded5128b8de",
			AccountID: "c4c3a079-15a0-4d73-93fb-7d26269a5206",
		})
		return c.Next()
	}
}

func GetAuthUser(c *fiber.Ctx) (*AuthUser, error) {
	authUser, ok := c.Locals(authKey).(*AuthUser)
	if !ok {
		return nil, fmt.Errorf("non protected route")
	}
	return authUser, nil
}
