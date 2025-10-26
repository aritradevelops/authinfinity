package oauth

import "github.com/aritradevelops/authinfinity/server/internal/core"

type OauthController struct {
	core.Controller[*Oauth]
}

func Controller() *OauthController {
	var oauthService = Service()
	return &OauthController{
		Controller: core.NewController(core.Service[*Oauth](oauthService)),
	}
}
