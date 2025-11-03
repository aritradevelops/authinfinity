package emailverificationrequest 

import (
	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
)

type EmailVerificationRequestService struct {
	core.Service[*EmailVerificationRequest]
}

var emailVerificationRequestService *EmailVerificationRequestService

func Service() *EmailVerificationRequestService {
	return emailVerificationRequestService
}
