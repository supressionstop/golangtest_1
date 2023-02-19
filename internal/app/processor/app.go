package processor

import (
	"softpro6/config"
	"softpro6/pkg/logger"
	"softpro6/pkg/postgres"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// storage
	p, err := postgres.New(cfg.DB.URL)
	if err != nil {
		l.Fatal("app - Run - postgres.New", err)
	}
	defer p.Close()

	l.Info("Started.", cfg.App.Name)
}
