package oauth

import (
	"github.com/aritradevelops/authinfinity/server/internal/core"
	"github.com/aritradevelops/authinfinity/server/internal/db"
)

type OauthRepository struct {
	core.Repository[*Oauth]
}

func Repository() *OauthRepository {
	var oauthModel = Model()
	return &OauthRepository{
		Repository: core.NewRepository[*Oauth](oauthModel, db.Instance()),
	}
}
