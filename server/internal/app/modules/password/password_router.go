package password 

import (
	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router) {
	controller := Controller()
	passwordRouter := router.Group("/passwords")
	passwordRouter.Use(core.SetModule("Password"))
	passwordRouter.Get("/list", controller.List)
	passwordRouter.Post("/create", controller.Create)
	passwordRouter.Get("/view/:id", controller.View)
	passwordRouter.Put("/update/:id", controller.Update)
	passwordRouter.Delete("/delete/:id", controller.Delete)
}