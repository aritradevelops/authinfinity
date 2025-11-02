package main

import (
	"os"

	"github.com/aritradevelops/authinfinity/server/internal/app/api"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/logger"
)

func main() {
	err := api.Bootstrap()
	if err != nil {
		logger.Error().Err(err).Msg("failed to start the apis")
		os.Exit(1)
	}
}
