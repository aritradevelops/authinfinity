package auth

import (
	"github.com/aritradevelops/authinfinity/server/internal/core"
	"github.com/aritradevelops/authinfinity/server/internal/modules/user"
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
