package app

import (
	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
	"github.com/google/uuid"
)

type AppService struct {
	core.Service[*App]
}

var appService *AppService

func Service() *AppService {
	return appService
}

func (s *AppService) GetSysAdminApp() (*App, error) {
	app := &App{}
	err := Repository().View(core.Filter{
		"account_id": uuid.Nil,
		"deleted_by": nil,
	}, &app)
	if err != nil {
		return nil, err
	}
	return app, nil
}
