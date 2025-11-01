package emailverificationrequest 

import (
	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
)

type EmailVerificationRequestRepository struct {
	core.Repository[*EmailVerificationRequest]
}

var emailVerificationRequestRepository *EmailVerificationRequestRepository 

func Repository() *EmailVerificationRequestRepository {
	return emailVerificationRequestRepository 
}
	
