package app 

import (
	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
)

type AppService struct {
	core.Service[*App]
}

var appService *AppService

func Service() *AppService {
	return appService
}
