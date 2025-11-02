package main

import (
	"log"

	"github.com/aritradevelops/authinfinity/server/cmd/cli/migrate"
	"github.com/aritradevelops/authinfinity/server/cmd/cli/module"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/logger"
	"github.com/spf13/cobra"
)

func main() {
	logger.Info().Str("version", "v0.1").Msg("Welcome to AuthInfinity CLI")
	root := &cobra.Command{Use: "authinfinity-cli"}
	module.RegisterCommand(root)
	migrate.RegisterCommand(root)
	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}
