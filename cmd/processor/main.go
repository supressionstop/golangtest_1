package main

import (
	"log"
	"softpro6/config"
)

func main() {
	// config
	_, err := config.NewConfig()
	if err != nil {
		log.Fatalf("config error: %s", err)
	}

	// run
}
