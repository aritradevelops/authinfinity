package emailverificationrequest 

import (
	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/db"
)

type EmailVerificationRequestRepository struct {
	core.Repository[*EmailVerificationRequest]
}


func Repository() *EmailVerificationRequestRepository {
	var emailVerificationRequestModel = Model()
	return &EmailVerificationRequestRepository{
		Repository: core.NewRepository[*EmailVerificationRequest](emailVerificationRequestModel, db.Instance()),
	}
}
