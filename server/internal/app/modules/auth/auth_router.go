package auth 

import (
	"fmt"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
	"github.com/gofiber/fiber/v2"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/db"
	"github.com/aritradevelops/authinfinity/server/internal/authn"
)

func setup() {
	authModel = core.NewModel("auths", []string{"name"})

	authRepository = &AuthRepository{
		Repository: core.NewRepository[*Auth](authModel, db.Instance()),
	}

	authService = &AuthService{
		Service: core.NewService(core.Repository[*Auth](authRepository)),
	}

	authController = &AuthController{
		Controller: core.NewController(core.Service[*Auth](authService)),
	}

}

func RegisterRoutes(router fiber.Router) {
	setup()
	fmt.Println("Module: Auth is registered successfully")
	authRouter := router.Group("/auths")
	authRouter.Use(core.SetModule("Auth"), authn.AuthMiddleware())
	authRouter.Get("/list", authController.List)
	authRouter.Post("/create", authController.Create)
	authRouter.Get("/view/:id", authController.View)
	authRouter.Put("/update/:id", authController.Update)
	authRouter.Delete("/delete/:id", authController.Delete)
}