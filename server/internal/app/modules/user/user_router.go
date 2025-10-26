package user

import (
	"github.com/aritradevelops/authinfinity/server/internal/auth"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router) {
	controller := Controller()
	userRouter := router.Group("/users")
	userRouter.Use(core.SetModule("User"), auth.AuthMiddleware())
	userRouter.Get("/list", controller.List)
	userRouter.Post("/create", controller.Create)
	userRouter.Get("/view/:id", controller.View)
	userRouter.Put("/update/:id", controller.Update)
	userRouter.Delete("/delete/:id", controller.Delete)
}
