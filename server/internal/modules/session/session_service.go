package session

import (
	"github.com/aritradevelops/authinfinity/server/internal/core"
)

type SessionService struct {
	core.Service[*Session]
}

func Service() *SessionService {
	var sessionRepository = Repository()
	return &SessionService{
		Service: core.NewService(core.Repository[*Session](sessionRepository)),
	}
}
