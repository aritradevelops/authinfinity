package session

import (
	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/db"
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
