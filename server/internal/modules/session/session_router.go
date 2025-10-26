package session

import (
	"github.com/aritradevelops/authinfinity/server/internal/core"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router) {
	controller := Controller()
	sessionRouter := router.Group("/sessions")
	sessionRouter.Use(core.SetModule("Session"))
	sessionRouter.Get("/list", controller.List)
	sessionRouter.Post("/create", controller.Create)
	sessionRouter.Get("/view/:id", controller.View)
	sessionRouter.Put("/update/:id", controller.Update)
	sessionRouter.Delete("/delete/:id", controller.Delete)
}
