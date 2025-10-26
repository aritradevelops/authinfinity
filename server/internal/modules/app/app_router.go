package app

import (
	"github.com/aritradevelops/authinfinity/server/internal/auth"
	"github.com/aritradevelops/authinfinity/server/internal/core"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router) {
	controller := Controller()
	appRouter := router.Group("/apps")
	appRouter.Use(core.SetModule("App"), auth.AuthMiddleware())
	appRouter.Get("/list", controller.List)
	appRouter.Post("/create", controller.Create)
	appRouter.Get("/view/:id", controller.View)
	appRouter.Put("/update/:id", controller.Update)
	appRouter.Delete("/delete/:id", controller.Delete)
}
