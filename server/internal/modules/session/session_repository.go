package session

import (
	"github.com/aritradevelops/authinfinity/server/internal/core"
	"github.com/aritradevelops/authinfinity/server/internal/db"
)

type SessionRepository struct {
	core.Repository[*Session]
}

func Repository() *SessionRepository {
	var sessionModel = Model()
	return &SessionRepository{
		Repository: core.NewRepository[*Session](sessionModel, db.Instance()),
	}
}
