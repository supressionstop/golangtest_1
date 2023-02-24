package integration

import (
	"log"
	"os"
	"softpro6/config"
	"softpro6/pkg/postgres"
	"testing"
)

func GetConfig(t *testing.T) *config.Config {
	t.Helper()
	appEnv := os.Getenv("APP_ENVIRONMENT")
	cfg, err := config.NewConfig(appEnv)
	if err != nil {
		log.Fatalf("getConfig - config.NewConfig: %s", err)
	}
	return cfg
}

func InitPg(t *testing.T) *postgres.Postgres {
	t.Helper()

	cfg := GetConfig(t)
	pg, err := postgres.New(cfg.DB.URL)
	if err != nil {
		log.Fatalf("initPg - postgres.New: %s", err)
	}

	return pg
}
