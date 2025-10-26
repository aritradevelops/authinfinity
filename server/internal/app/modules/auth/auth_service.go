package auth

import (
	"github.com/aritradevelops/authinfinity/server/internal/app/modules/user"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
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
