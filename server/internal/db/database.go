package db

import "fmt"

type Database interface {
	Connect() error
	Disconnect() error
	Health() error
}

func NotInitializedErr(what string) error {
	return fmt.Errorf("%s is not initialized, have you forgot to call Connect() ?", what)
}
