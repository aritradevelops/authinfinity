package app

import "github.com/aritradevelops/authinfinity/server/internal/pkg/core"

type AppController struct {
	core.Controller[*App]
}

func Controller() *AppController {
	var appService = Service()
	return &AppController{
		Controller: core.NewController(core.Service[*App](appService)),
	}
}
