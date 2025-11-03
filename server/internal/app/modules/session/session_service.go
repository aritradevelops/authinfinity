package session 

import (
	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
)

type SessionService struct {
	core.Service[*Session]
}

var sessionService *SessionService

func Service() *SessionService {
	return sessionService
}
