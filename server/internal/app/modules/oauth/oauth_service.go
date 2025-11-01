package oauth 

import (
	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
)

type OauthService struct {
	core.Service[*Oauth]
}

var oauthService *OauthService

func Service() *OauthService {
	return oauthService
}
