package app

import "github.com/aritradevelops/authinfinity/server/internal/pkg/core"

type AppController struct {
	core.Controller[*App]
}

var appController  *AppController

