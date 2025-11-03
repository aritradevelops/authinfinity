package password 

import (
	"github.com/gofiber/fiber/v2"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/db"
	"github.com/aritradevelops/authinfinity/server/internal/authn"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/logger"
)

func setup() {
	passwordModel = core.NewModel("passwords", []string{"name"})

	passwordRepository = &PasswordRepository{
		Repository: core.NewRepository[*Password](passwordModel, db.Instance()),
	}

	passwordService = &PasswordService{
		Service: core.NewService(core.Repository[*Password](passwordRepository)),
	}

	passwordController = &PasswordController{
		Controller: core.NewController(core.Service[*Password](passwordService)),
	}

}

func RegisterRoutes(router fiber.Router) {
	setup()
	logger.Info().Str("module", "Password").Msg("Module registered s")
	passwordRouter := router.Group("/passwords")
	passwordRouter.Use(core.SetModule("Password"), authn.AuthMiddleware())
	passwordRouter.Get("/list", passwordController.List)
	passwordRouter.Post("/create", passwordController.Create)
	passwordRouter.Get("/view/:id", passwordController.View)
	passwordRouter.Put("/update/:id", passwordController.Update)
	passwordRouter.Delete("/delete/:id", passwordController.Delete)
}