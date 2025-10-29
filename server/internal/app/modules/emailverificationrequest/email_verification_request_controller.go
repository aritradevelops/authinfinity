package emailverificationrequest

import "github.com/aritradevelops/authinfinity/server/internal/pkg/core"

type EmailVerificationRequestController struct {
	core.Controller[*EmailVerificationRequest]
}


func Controller() *EmailVerificationRequestController {
	var emailVerificationRequestService = Service()
	return &EmailVerificationRequestController{
		Controller: core.NewController(core.Service[*EmailVerificationRequest](emailVerificationRequestService)),
	}
}
