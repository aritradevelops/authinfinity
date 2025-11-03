package session

import "github.com/aritradevelops/authinfinity/server/internal/pkg/core"

type SessionController struct {
	core.Controller[*Session]
}

var sessionController  *SessionController

