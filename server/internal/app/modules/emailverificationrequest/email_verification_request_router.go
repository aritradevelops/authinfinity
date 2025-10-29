package emailverificationrequest 

import (
	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router) {
	controller := Controller()
	emailVerificationRequestRouter := router.Group("/email-verification-requests")
	emailVerificationRequestRouter.Use(core.SetModule("EmailVerificationRequest"))
	emailVerificationRequestRouter.Get("/list", controller.List)
	emailVerificationRequestRouter.Post("/create", controller.Create)
	emailVerificationRequestRouter.Get("/view/:id", controller.View)
	emailVerificationRequestRouter.Put("/update/:id", controller.Update)
	emailVerificationRequestRouter.Delete("/delete/:id", controller.Delete)
}