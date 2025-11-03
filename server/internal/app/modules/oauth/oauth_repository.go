package oauth 

import (
	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
)

type OauthRepository struct {
	core.Repository[*Oauth]
}

var oauthRepository *OauthRepository 

func Repository() *OauthRepository {
	return oauthRepository 
}
	
