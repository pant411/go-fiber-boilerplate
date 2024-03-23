package main

import (
	"log"

	"go-fiber-boilerplate/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("failed to start the application: %v", err)
	}
}
