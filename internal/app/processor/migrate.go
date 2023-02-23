//go:build migrate

package processor

import (
	"log"
	"os"
	"softpro6/config"
	"softpro6/pkg/postgres"
)

func init() {
	cfg, err := config.NewConfig(os.Getenv("APP_ENVIRONMENT"))
	if err != nil {
		log.Panicln("migrate - config.NewConfig: %s", err)
	}
	postgres.Migrate(cfg.DB.MigrationsUrl, cfg.DB.URL)
}
