package session 

import (
	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
)

type SessionRepository struct {
	core.Repository[*Session]
}

var sessionRepository *SessionRepository 

func Repository() *SessionRepository {
	return sessionRepository 
}
	
