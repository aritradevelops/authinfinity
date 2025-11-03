package resetpasswordrequest

import "github.com/aritradevelops/authinfinity/server/internal/pkg/core"

type ResetPasswordRequestController struct {
	core.Controller[*ResetPasswordRequest]
}

var resetPasswordRequestController  *ResetPasswordRequestController

