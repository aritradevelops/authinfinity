package account

import (
	"github.com/aritradevelops/authinfinity/server/internal/core"
	"github.com/aritradevelops/authinfinity/server/internal/db"
)

type AccountRepository struct {
	core.Repository[*Account]
}

func Repository() *AccountRepository {
	var accountModel = Model()
	return &AccountRepository{
		Repository: core.NewRepository[*Account](accountModel, db.Instance()),
	}
}
