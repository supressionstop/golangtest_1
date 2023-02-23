package main

import (
	"log"
	"os"
	"softpro6/config"
	"softpro6/internal/app/processor"
)

func main() {
	// config
	cfg, err := config.NewConfig(os.Getenv("APP_ENVIRONMENT"))
	if err != nil {
		log.Fatalf("config error: %s", err)
	}

	// run
	processor.Run(cfg)
}
