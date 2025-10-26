package main

import (
	"log"

	"github.com/aritradevelops/authinfinity/server/cmd/cli/internal/logx"
	"github.com/aritradevelops/authinfinity/server/cmd/cli/migrate"
	"github.com/aritradevelops/authinfinity/server/cmd/cli/module"
	"github.com/spf13/cobra"
)

func main() {
	lg := logx.New()
	lg.Info("Welcome to AuthInfinity CLI", "version", "v0.1")
	root := &cobra.Command{Use: "authinfinity-cli"}
	module.RegisterCommand(root)
	migrate.RegisterCommand(root)
	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}
