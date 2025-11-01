package user

import "github.com/aritradevelops/authinfinity/server/internal/pkg/core"

type UserController struct {
	core.Controller[*User]
}

var userController  *UserController

