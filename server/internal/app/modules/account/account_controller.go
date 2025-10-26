package account

import "github.com/aritradevelops/authinfinity/server/internal/pkg/core"

type AccountController struct {
	core.Controller[*Account]
}

func Controller() *AccountController {
	var accountService = Service()
	return &AccountController{
		Controller: core.NewController(core.Service[*Account](accountService)),
	}
}
