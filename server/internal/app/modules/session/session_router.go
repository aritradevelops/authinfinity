package session 

import (
	"fmt"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
	"github.com/gofiber/fiber/v2"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/db"
	"github.com/aritradevelops/authinfinity/server/internal/authn"
)

func setup() {
	sessionModel = core.NewModel("sessions", []string{"name"})

	sessionRepository = &SessionRepository{
		Repository: core.NewRepository[*Session](sessionModel, db.Instance()),
	}

	sessionService = &SessionService{
		Service: core.NewService(core.Repository[*Session](sessionRepository)),
	}

	sessionController = &SessionController{
		Controller: core.NewController(core.Service[*Session](sessionService)),
	}

}

func RegisterRoutes(router fiber.Router) {
	setup()
	fmt.Println("Module: Session is registered successfully")
	sessionRouter := router.Group("/sessions")
	sessionRouter.Use(core.SetModule("Session"), authn.AuthMiddleware())
	sessionRouter.Get("/list", sessionController.List)
	sessionRouter.Post("/create", sessionController.Create)
	sessionRouter.Get("/view/:id", sessionController.View)
	sessionRouter.Put("/update/:id", sessionController.Update)
	sessionRouter.Delete("/delete/:id", sessionController.Delete)
}