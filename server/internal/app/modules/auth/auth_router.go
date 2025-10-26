package auth

import (
	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router) {
	// controller := Controller()
	authRouter := router.Group("/auth")
	authRouter.Use(core.SetModule("Auth"))
	// authRouter.Get("/list", controller.List)
	// authRouter.Post("/create", controller.Create)
	// authRouter.Get("/view/:id", controller.View)
	// authRouter.Put("/update/:id", controller.Update)
	// authRouter.Delete("/delete/:id", controller.Delete)
}
