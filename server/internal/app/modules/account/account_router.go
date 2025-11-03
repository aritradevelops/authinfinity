package account 

import (
	"github.com/gofiber/fiber/v2"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/db"
	"github.com/aritradevelops/authinfinity/server/internal/authn"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/logger"
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
	logger.Info().Str("module", "Account").Msg("Module registered s")
	accountRouter := router.Group("/accounts")
	accountRouter.Use(core.SetModule("Account"), authn.AuthMiddleware())
	accountRouter.Get("/list", accountController.List)
	accountRouter.Post("/create", accountController.Create)
	accountRouter.Get("/view/:id", accountController.View)
	accountRouter.Put("/update/:id", accountController.Update)
	accountRouter.Delete("/delete/:id", accountController.Delete)
}