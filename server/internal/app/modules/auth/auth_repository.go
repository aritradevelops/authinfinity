package auth

import (
	"github.com/aritradevelops/authinfinity/server/internal/app/modules/user"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/db"
)

type AuthRepository struct {
	core.Repository[*user.User]
}

func Repository() *AuthRepository {
	var authModel = Model()
	return &AuthRepository{
		Repository: core.NewRepository[*user.User](authModel, db.Instance()),
	}
}
