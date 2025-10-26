package app

import (
	"github.com/aritradevelops/authinfinity/server/internal/core"
	"github.com/aritradevelops/authinfinity/server/internal/db"
)

type AppRepository struct {
	core.Repository[*App]
}

func Repository() *AppRepository {
	var appModel = Model()
	return &AppRepository{
		Repository: core.NewRepository[*App](appModel, db.Instance()),
	}
}
