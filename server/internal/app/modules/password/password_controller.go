package password

import "github.com/aritradevelops/authinfinity/server/internal/pkg/core"

type PasswordController struct {
	core.Controller[*Password]
}


func Controller() *PasswordController {
	var passwordService = Service()
	return &PasswordController{
		Controller: core.NewController(core.Service[*Password](passwordService)),
	}
}
