package emailverificationrequest

import "github.com/aritradevelops/authinfinity/server/internal/pkg/core"

type EmailVerificationRequestController struct {
	core.Controller[*EmailVerificationRequest]
}

var emailVerificationRequestController  *EmailVerificationRequestController

