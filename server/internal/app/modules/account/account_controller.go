package account

import "github.com/aritradevelops/authinfinity/server/internal/pkg/core"

type AccountController struct {
	core.Controller[*Account]
}

var accountController  *AccountController

