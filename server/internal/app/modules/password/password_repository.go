package password 

import (
	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/db"
)

type PasswordRepository struct {
	core.Repository[*Password]
}


func Repository() *PasswordRepository {
	var passwordModel = Model()
	return &PasswordRepository{
		Repository: core.NewRepository[*Password](passwordModel, db.Instance()),
	}
}
