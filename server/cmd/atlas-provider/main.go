package main

import (
	"fmt"
	"io"
	"os"

	"ariga.io/atlas-provider-gorm/gormschema"
	"github.com/aritradevelops/authinfinity/server/internal/app/modules/user"
	"github.com/aritradevelops/authinfinity/server/internal/app/modules/session"
	"github.com/aritradevelops/authinfinity/server/internal/app/modules/oauth"
	"github.com/aritradevelops/authinfinity/server/internal/app/modules/auth"
	"github.com/aritradevelops/authinfinity/server/internal/app/modules/app"
	"github.com/aritradevelops/authinfinity/server/internal/app/modules/password"
	"github.com/aritradevelops/authinfinity/server/internal/app/modules/account"
)

// Models array - same as in migrate/generate.go
var models = []any{&account.Account{}, &app.App{}, &auth.Auth{}, &oauth.Oauth{}, &password.Password{}, &session.Session{}, &user.User{},
}

func main() {
	stmts, err := gormschema.New("postgres").Load(models...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
		os.Exit(1)
	}
	io.WriteString(os.Stdout, stmts)
}
