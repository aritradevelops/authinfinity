package oauth

import (
	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
)

type OauthService struct {
	core.Service[*Oauth]
}

func Service() *OauthService {
	var oauthRepository = Repository()
	return &OauthService{
		Service: core.NewService(core.Repository[*Oauth](oauthRepository)),
	}
}
