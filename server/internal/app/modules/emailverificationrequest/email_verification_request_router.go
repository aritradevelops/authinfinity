package emailverificationrequest 

import (
	"fmt"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
	"github.com/gofiber/fiber/v2"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/db"
	"github.com/aritradevelops/authinfinity/server/internal/authn"
)

func setup() {
	emailVerificationRequestModel = core.NewModel("email_verification_requests", []string{"name"})

	emailVerificationRequestRepository = &EmailVerificationRequestRepository{
		Repository: core.NewRepository[*EmailVerificationRequest](emailVerificationRequestModel, db.Instance()),
	}

	emailVerificationRequestService = &EmailVerificationRequestService{
		Service: core.NewService(core.Repository[*EmailVerificationRequest](emailVerificationRequestRepository)),
	}

	emailVerificationRequestController = &EmailVerificationRequestController{
		Controller: core.NewController(core.Service[*EmailVerificationRequest](emailVerificationRequestService)),
	}

}

func RegisterRoutes(router fiber.Router) {
	setup()
	fmt.Println("Module: EmailVerificationRequest is registered successfully")
	emailVerificationRequestRouter := router.Group("/email-verification-requests")
	emailVerificationRequestRouter.Use(core.SetModule("EmailVerificationRequest"), authn.AuthMiddleware())
	emailVerificationRequestRouter.Get("/list", emailVerificationRequestController.List)
	emailVerificationRequestRouter.Post("/create", emailVerificationRequestController.Create)
	emailVerificationRequestRouter.Get("/view/:id", emailVerificationRequestController.View)
	emailVerificationRequestRouter.Put("/update/:id", emailVerificationRequestController.Update)
	emailVerificationRequestRouter.Delete("/delete/:id", emailVerificationRequestController.Delete)
}