package user

import "github.com/aritradevelops/authinfinity/server/internal/core"

type UserController struct {
	core.Controller[*User]
}

func Controller() *UserController {
	var userService = Service()
	return &UserController{
		Controller: core.NewController(core.Service[*User](userService)),
	}
}
