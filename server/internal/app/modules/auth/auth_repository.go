package auth

import (
	"github.com/aritradevelops/authinfinity/server/internal/app/modules/user"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
)

type AuthRepository struct {
	core.Repository[*user.User]
}

var authRepository *AuthRepository

func Repository() *AuthRepository {
	return authRepository
}
