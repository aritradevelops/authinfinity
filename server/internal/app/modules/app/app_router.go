package app 

import (
	"fmt"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
	"github.com/gofiber/fiber/v2"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/db"
	"github.com/aritradevelops/authinfinity/server/internal/authn"
)

func setup() {
	appModel = core.NewModel("apps", []string{"name"})

	appRepository = &AppRepository{
		Repository: core.NewRepository[*App](appModel, db.Instance()),
	}

	appService = &AppService{
		Service: core.NewService(core.Repository[*App](appRepository)),
	}

	appController = &AppController{
		Controller: core.NewController(core.Service[*App](appService)),
	}

}

func RegisterRoutes(router fiber.Router) {
	setup()
	fmt.Println("Module: App is registered successfully")
	appRouter := router.Group("/apps")
	appRouter.Use(core.SetModule("App"), authn.AuthMiddleware())
	appRouter.Get("/list", appController.List)
	appRouter.Post("/create", appController.Create)
	appRouter.Get("/view/:id", appController.View)
	appRouter.Put("/update/:id", appController.Update)
	appRouter.Delete("/delete/:id", appController.Delete)
}