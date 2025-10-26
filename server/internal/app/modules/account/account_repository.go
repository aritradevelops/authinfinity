package account

import (
	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/db"
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
