package auth

import (
	"github.com/aritradevelops/authinfinity/server/internal/core"
	"github.com/aritradevelops/authinfinity/server/internal/db"
	"github.com/aritradevelops/authinfinity/server/internal/modules/user"
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
