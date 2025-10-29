package auth

import (
	"fmt"
	"net/http"

	"github.com/aritradevelops/authinfinity/server/internal/app/modules/user"
	"github.com/aritradevelops/authinfinity/server/internal/middlewares/translator"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/response"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	core.Controller[*user.User]
	service *AuthService
}

func Controller() *AuthController {
	var authService = Service()
	return &AuthController{
		Controller: core.NewController(core.Service[*user.User](authService)),
		service:    authService,
	}
}

func (bc *AuthController) Register(c *fiber.Ctx) error {
	err := bc.service.Register(c)
	if err != nil {
		fmt.Printf("Error : %+v", err)
		return err
	}

	c.Status(http.StatusCreated)
	return c.JSON(
		response.NewServerResponse(
			translator.Localize(c, "user.register", map[string]string{"Entity": c.Locals("module").(string)}),
			fiber.Map{},
		),
	)
}
