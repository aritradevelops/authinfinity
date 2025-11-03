package resetpasswordrequest 

import (
	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
)

type ResetPasswordRequestService struct {
	core.Service[*ResetPasswordRequest]
}

var resetPasswordRequestService *ResetPasswordRequestService

func Service() *ResetPasswordRequestService {
	return resetPasswordRequestService
}
