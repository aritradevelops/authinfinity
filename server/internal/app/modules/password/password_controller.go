package password

import "github.com/aritradevelops/authinfinity/server/internal/pkg/core"

type PasswordController struct {
	core.Controller[*Password]
}

var passwordController  *PasswordController

