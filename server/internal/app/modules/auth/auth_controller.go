package auth

import "github.com/aritradevelops/authinfinity/server/internal/pkg/core"

type AuthController struct {
	core.Controller[*Auth]
}

var authController  *AuthController

