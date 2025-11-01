package app 

import (
	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
)

type AppRepository struct {
	core.Repository[*App]
}

var appRepository *AppRepository 

func Repository() *AppRepository {
	return appRepository 
}
	
