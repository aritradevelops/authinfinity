package user

import (
	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/db"
)

type UserRepository struct {
	core.Repository[*User]
}

func Repository() *UserRepository {
	var userModel = Model()
	return &UserRepository{
		Repository: core.NewRepository[*User](userModel, db.Instance()),
	}
}
