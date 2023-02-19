package processor

import (
	"softpro6/config"
	"softpro6/pkg/logger"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	l.Info("Started.", cfg.App.Name)
}
