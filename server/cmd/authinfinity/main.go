package main

import (
	"log"
	"os"

	"github.com/aritradevelops/authinfinity/server/internal/api"
)

func main() {
	err := api.Bootstrap()
	if err != nil {
		log.Printf("%v", err)
		os.Exit(1)
	}
}
