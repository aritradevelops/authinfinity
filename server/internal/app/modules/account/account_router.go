package account 

import (
	"fmt"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
	"github.com/gofiber/fiber/v2"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/db"
	"github.com/aritradevelops/authinfinity/server/internal/authn"
)

func setup() {
	accountModel = core.NewModel("accounts", []string{"name"})

	accountRepository = &AccountRepository{
		Repository: core.NewRepository[*Account](accountModel, db.Instance()),
	}

	accountService = &AccountService{
		Service: core.NewService(core.Repository[*Account](accountRepository)),
	}

	accountController = &AccountController{
		Controller: core.NewController(core.Service[*Account](accountService)),
	}

}

func RegisterRoutes(router fiber.Router) {
	setup()
	fmt.Println("Module: Account is registered successfully")
	accountRouter := router.Group("/accounts")
	accountRouter.Use(core.SetModule("Account"), authn.AuthMiddleware())
	accountRouter.Get("/list", accountController.List)
	accountRouter.Post("/create", accountController.Create)
	accountRouter.Get("/view/:id", accountController.View)
	accountRouter.Put("/update/:id", accountController.Update)
	accountRouter.Delete("/delete/:id", accountController.Delete)
}