package api

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aritradevelops/authinfinity/server/internal/pkg/config"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/db"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/logger"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/server"
)

func Bootstrap() error {
	conf, err := config.Load()
	if err != nil {
		return err
	}
	database := db.New(conf.Env.DBConnectionURI, db.DatabaseKind(conf.Env.DBKind))
	err = database.Connect()
	if err != nil {
		return err
	}
	logger.Info().Msg("Database connected successfully.")

	srv := server.New(&conf, database)

	go func() {
		if err := srv.Start(); err != nil {
			logger.Fatal().Err(err).Msg("failed to start the server")
			os.Exit(1)
		}
	}()

	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	<-quitCh
	logger.Info().Msg("Bye! shutting down gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("server forced shutdown %v", err)
	}

	err = database.Disconnect()
	if err != nil {
		return err
	}
	logger.Info().Msg("Database disconnected.")
	return nil
}
