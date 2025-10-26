package migrate

import (
	"fmt"
	"io"
	"os"
	"strings"

	"ariga.io/atlas-provider-gorm/gormschema"
	"github.com/aritradevelops/authinfinity/server/internal/app/modules/account"
	"github.com/aritradevelops/authinfinity/server/internal/app/modules/app"
	"github.com/aritradevelops/authinfinity/server/internal/app/modules/oauth"
	"github.com/aritradevelops/authinfinity/server/internal/app/modules/session"
	"github.com/aritradevelops/authinfinity/server/internal/app/modules/user"
	"github.com/spf13/cobra"
)

func newMigrateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate generate",
		Short: "Generate migrations from files",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			sb := &strings.Builder{}
			models := []interface{}{
				&user.User{},
				&account.Account{},
				&app.App{},
				&session.Session{},
				&oauth.Oauth{},
			}
			stmts, err := gormschema.New("postgres").Load(models...)
			if err != nil {
				fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
				os.Exit(1)
			}
			sb.WriteString(stmts)
			sb.WriteString(";\n")
			io.WriteString(os.Stdout, sb.String())
			return nil
		},
	}
	return cmd
}
func RegisterCommand(parent *cobra.Command) {
	cmd := newMigrateCommand()
	parent.AddCommand(cmd)
}
