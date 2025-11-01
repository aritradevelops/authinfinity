package resetpasswordrequest 

import (
	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
)

type ResetPasswordRequestRepository struct {
	core.Repository[*ResetPasswordRequest]
}

var resetPasswordRequestRepository *ResetPasswordRequestRepository 

func Repository() *ResetPasswordRequestRepository {
	return resetPasswordRequestRepository 
}
	
