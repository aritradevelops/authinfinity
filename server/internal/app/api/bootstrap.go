package api

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aritradevelops/authinfinity/server/internal/pkg/config"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/db"
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

	srv := server.New(&conf, database)

	go func() {
		if err := srv.Start(); err != nil {
			log.Printf("failed to start the server due to %v", err)
			os.Exit(1)
		}
	}()

	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	<-quitCh
	log.Println("shutting down to server")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("server forced shutdown %v", err)
	}

	err = database.Disconnect()
	if err != nil {
		return err
	}
	return nil
}
