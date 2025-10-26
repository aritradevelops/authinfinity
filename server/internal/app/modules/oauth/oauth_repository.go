package oauth

import (
	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/db"
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
