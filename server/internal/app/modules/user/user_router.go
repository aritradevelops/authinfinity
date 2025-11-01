package user 

import (
	"fmt"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
	"github.com/gofiber/fiber/v2"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/db"
	"github.com/aritradevelops/authinfinity/server/internal/authn"
)

func setup() {
	userModel = core.NewModel("users", []string{"name"})

	userRepository = &UserRepository{
		Repository: core.NewRepository[*User](userModel, db.Instance()),
	}

	userService = &UserService{
		Service: core.NewService(core.Repository[*User](userRepository)),
	}

	userController = &UserController{
		Controller: core.NewController(core.Service[*User](userService)),
	}

}

func RegisterRoutes(router fiber.Router) {
	setup()
	fmt.Println("Module: User is registered successfully")
	userRouter := router.Group("/users")
	userRouter.Use(core.SetModule("User"), authn.AuthMiddleware())
	userRouter.Get("/list", userController.List)
	userRouter.Post("/create", userController.Create)
	userRouter.Get("/view/:id", userController.View)
	userRouter.Put("/update/:id", userController.Update)
	userRouter.Delete("/delete/:id", userController.Delete)
}