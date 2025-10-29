package app

import (
	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
)

type AppService struct {
	core.Service[*App]
}

func Service() *AppService {
	var appRepository = Repository()
	return &AppService{
		Service: core.NewService(core.Repository[*App](appRepository)),
	}
}
