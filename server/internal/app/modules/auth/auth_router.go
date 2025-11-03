package auth

import (
	"github.com/aritradevelops/authinfinity/server/internal/app/modules/user"
	"github.com/aritradevelops/authinfinity/server/internal/authn"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/db"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

func setup() {

	authRepository = &AuthRepository{
		Repository: core.NewRepository[*user.User](user.Model(), db.Instance()),
	}

	authService = &AuthService{
		Service: core.NewService(core.Repository[*user.User](authRepository)),
	}

	authController = &AuthController{
		Controller: core.NewController(core.Service[*user.User](authService)),
	}

}

func RegisterRoutes(router fiber.Router) {
	setup()
	logger.Info().Str("module", "Auth").Msg("Module registered s")
	authRouter := router.Group("/auths")
	authRouter.Use(core.SetModule("Auth"), authn.AuthMiddleware())
	authRouter.Get("/list", authController.List)
	authRouter.Post("/create", authController.Create)
	authRouter.Get("/view/:id", authController.View)
	authRouter.Put("/update/:id", authController.Update)
	authRouter.Delete("/delete/:id", authController.Delete)
}
