package auth 

import (
	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
)

type AuthService struct {
	core.Service[*Auth]
}

var authService *AuthService

func Service() *AuthService {
	return authService
}
