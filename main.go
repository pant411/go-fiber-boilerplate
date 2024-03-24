package main

import (
	"log"

	"go-fiber-boilerplate/app"
)

func main() {
	if err := app.Execute(); err != nil {
		log.Fatalf("failed to start the application: %v", err)
	}
}
