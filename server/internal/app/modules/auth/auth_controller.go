package auth

import (
	"github.com/aritradevelops/authinfinity/server/internal/app/modules/user"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
)

type AuthController struct {
	core.Controller[*user.User]
}

func Controller() *AuthController {
	var authService = Service()
	return &AuthController{
		Controller: core.NewController(core.Service[*user.User](authService)),
	}
}
