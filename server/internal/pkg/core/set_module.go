package core

import "github.com/gofiber/fiber/v2"

func SetModule(module string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals("module", module)
		return c.Next()
	}
}
