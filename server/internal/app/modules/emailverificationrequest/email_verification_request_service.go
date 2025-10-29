package emailverificationrequest 

import (
	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
)

type EmailVerificationRequestService struct {
	core.Service[*EmailVerificationRequest]
}


func Service() *EmailVerificationRequestService {
	var emailVerificationRequestRepository = Repository()
	return &EmailVerificationRequestService{
		Service: core.NewService(core.Repository[*EmailVerificationRequest](emailVerificationRequestRepository)),
	}
}
