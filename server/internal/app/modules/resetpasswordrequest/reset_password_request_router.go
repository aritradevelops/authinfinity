package resetpasswordrequest 

import (
	"fmt"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
	"github.com/gofiber/fiber/v2"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/db"
	"github.com/aritradevelops/authinfinity/server/internal/authn"
)

func setup() {
	resetPasswordRequestModel = core.NewModel("reset_password_requests", []string{"name"})

	resetPasswordRequestRepository = &ResetPasswordRequestRepository{
		Repository: core.NewRepository[*ResetPasswordRequest](resetPasswordRequestModel, db.Instance()),
	}

	resetPasswordRequestService = &ResetPasswordRequestService{
		Service: core.NewService(core.Repository[*ResetPasswordRequest](resetPasswordRequestRepository)),
	}

	resetPasswordRequestController = &ResetPasswordRequestController{
		Controller: core.NewController(core.Service[*ResetPasswordRequest](resetPasswordRequestService)),
	}

}

func RegisterRoutes(router fiber.Router) {
	setup()
	fmt.Println("Module: ResetPasswordRequest is registered successfully")
	resetPasswordRequestRouter := router.Group("/reset-password-requests")
	resetPasswordRequestRouter.Use(core.SetModule("ResetPasswordRequest"), authn.AuthMiddleware())
	resetPasswordRequestRouter.Get("/list", resetPasswordRequestController.List)
	resetPasswordRequestRouter.Post("/create", resetPasswordRequestController.Create)
	resetPasswordRequestRouter.Get("/view/:id", resetPasswordRequestController.View)
	resetPasswordRequestRouter.Put("/update/:id", resetPasswordRequestController.Update)
	resetPasswordRequestRouter.Delete("/delete/:id", resetPasswordRequestController.Delete)
}