package oauth

import (
	"github.com/aritradevelops/authinfinity/server/internal/core"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router) {
	controller := Controller()
	oauthRouter := router.Group("/oauths")
	oauthRouter.Use(core.SetModule("Oauth"))
	oauthRouter.Get("/list", controller.List)
	oauthRouter.Post("/create", controller.Create)
	oauthRouter.Get("/view/:id", controller.View)
	oauthRouter.Put("/update/:id", controller.Update)
	oauthRouter.Delete("/delete/:id", controller.Delete)
}
