package main

import (
	"log"
	"softpro6/config"
	"softpro6/internal/app/processor"
)

func main() {
	// config
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("config error: %s", err)
	}

	// run
	processor.Run(cfg)
}
