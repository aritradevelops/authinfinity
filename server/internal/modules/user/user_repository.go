package user

import (
	"github.com/aritradevelops/authinfinity/server/internal/core"
	"github.com/aritradevelops/authinfinity/server/internal/db"
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
