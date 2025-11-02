package main

import (
	"io"
	"os"

	"ariga.io/atlas-provider-gorm/gormschema"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/logger"
)

// Models array - same as in migrate/generate.go
var models = []any{}

func main() {
	stmts, err := gormschema.New("postgres").Load(models...)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to load gorm schema")
		os.Exit(1)
	}
	io.WriteString(os.Stdout, stmts)
}
