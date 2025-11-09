package jwtutil

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/aritradevelops/authinfinity/server/internal/app/modules/account"
	"github.com/aritradevelops/authinfinity/server/internal/app/modules/app"
	"github.com/aritradevelops/authinfinity/server/internal/app/modules/user"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/config"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/crypto"
)

type Tokens struct {
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

type JWTPayload struct {
	Id        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	DP        string `json:"dp,omitempty"`
	Email     string `json:"email,omitempty"`
	AccountId string `json:"account_id,omitempty"`
	AppId     string `json:"app_id,omitempty"`
}

func Sign(user user.User, app app.App, account account.Account, audience string) (Tokens, error) {
	tokens := Tokens{}
	refreshToken, err := crypto.GenerateHash(config.Instance().RefreshTokenHashLength)
	if err != nil {
		return tokens, err
	}
	tokens.RefreshToken = refreshToken
	payload := JWTPayload{
		Id:        user.ID.String(),
		Name:      user.Name,
		Email:     user.Email,
		DP:        *user.Dp,
		AppId:     app.ID.String(),
		AccountId: user.AccountID.String(),
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return tokens, err
	}
	duration, err := time.ParseDuration(app.JwtLifetime)
	if err != nil {
		return tokens, err
	}

	expiration := time.Now().Add(duration)

	switch app.JwtAlgo {
	case "HS256":
		// TODO:

	default:
		return tokens, errors.New("Unsupported jwt signing algorithm.")
	}
	return tokens, nil
}
