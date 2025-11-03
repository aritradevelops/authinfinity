package oauth 

import (
	"github.com/gofiber/fiber/v2"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/db"
	"github.com/aritradevelops/authinfinity/server/internal/authn"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/logger"
)

func setup() {
	oauthModel = core.NewModel("oauths", []string{"name"})

	oauthRepository = &OauthRepository{
		Repository: core.NewRepository[*Oauth](oauthModel, db.Instance()),
	}

	oauthService = &OauthService{
		Service: core.NewService(core.Repository[*Oauth](oauthRepository)),
	}

	oauthController = &OauthController{
		Controller: core.NewController(core.Service[*Oauth](oauthService)),
	}

}

func RegisterRoutes(router fiber.Router) {
	setup()
	logger.Info().Str("module", "Oauth").Msg("Module registered s")
	oauthRouter := router.Group("/oauths")
	oauthRouter.Use(core.SetModule("Oauth"), authn.AuthMiddleware())
	oauthRouter.Get("/list", oauthController.List)
	oauthRouter.Post("/create", oauthController.Create)
	oauthRouter.Get("/view/:id", oauthController.View)
	oauthRouter.Put("/update/:id", oauthController.Update)
	oauthRouter.Delete("/delete/:id", oauthController.Delete)
}