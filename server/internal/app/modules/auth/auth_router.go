package auth

import (
	"github.com/aritradevelops/authinfinity/server/internal/app/modules/user"
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
	logger.Info().Str("module", "Auth").Msg("Module registered.")
	authRouter := router.Group("/auth")
	authRouter.Use(core.SetModule("Auth"))
	authRouter.Post("/register", authController.Register)
}
