package oauth

import "github.com/aritradevelops/authinfinity/server/internal/pkg/core"

type OauthController struct {
	core.Controller[*Oauth]
}

var oauthController  *OauthController

