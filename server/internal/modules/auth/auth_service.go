package auth

import (
	"github.com/aritradevelops/authinfinity/server/internal/core"
	"github.com/aritradevelops/authinfinity/server/internal/modules/user"
)

type AuthService struct {
	core.Service[*user.User]
}

func Service() *AuthService {
	var authRepository = Repository()
	return &AuthService{
		Service: core.NewService(core.Repository[*user.User](authRepository)),
	}
}
