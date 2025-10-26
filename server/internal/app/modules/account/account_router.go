package account

import (
	"github.com/aritradevelops/authinfinity/server/internal/auth"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router) {
	controller := Controller()
	accountRouter := router.Group("/accounts")
	accountRouter.Use(core.SetModule("Account"), auth.AuthMiddleware())
	accountRouter.Get("/list", controller.List)
	accountRouter.Post("/create", controller.Create)
	accountRouter.Get("/view/:id", controller.View)
	accountRouter.Put("/update/:id", controller.Update)
	accountRouter.Delete("/delete/:id", controller.Delete)
}
