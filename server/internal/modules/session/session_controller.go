package session

import "github.com/aritradevelops/authinfinity/server/internal/core"

type SessionController struct {
	core.Controller[*Session]
}

func Controller() *SessionController {
	var sessionService = Service()
	return &SessionController{
		Controller: core.NewController(core.Service[*Session](sessionService)),
	}
}
