package main

import (
	"io"
	"os"

	"ariga.io/atlas-provider-gorm/gormschema"
	"github.com/aritradevelops/authinfinity/server/internal/app/modules/account"
	"github.com/aritradevelops/authinfinity/server/internal/app/modules/app"
	"github.com/aritradevelops/authinfinity/server/internal/app/modules/emailverificationrequest"
	"github.com/aritradevelops/authinfinity/server/internal/app/modules/oauth"
	"github.com/aritradevelops/authinfinity/server/internal/app/modules/password"
	"github.com/aritradevelops/authinfinity/server/internal/app/modules/resetpasswordrequest"
	"github.com/aritradevelops/authinfinity/server/internal/app/modules/session"
	"github.com/aritradevelops/authinfinity/server/internal/app/modules/user"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/logger"
)

// Models array - same as in migrate/generate.go
var models = []any{&account.Account{}, &app.App{}, &emailverificationrequest.EmailVerificationRequest{}, &oauth.Oauth{}, &password.Password{}, &resetpasswordrequest.ResetPasswordRequest{}, &session.Session{}, &user.User{}}

func main() {
	stmts, err := gormschema.New("postgres").Load(models...)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to load gorm schema")
		os.Exit(1)
	}
	io.WriteString(os.Stdout, stmts)
}
