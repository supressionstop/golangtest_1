package config

import "os"

type (
	Config struct {
		App
		Log
	}

	App struct {
		Name string
	}

	Log struct {
		Level string
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	// todo cobra
	cfg.App.Name = os.Getenv("APP_NAME")
	cfg.Log.Level = os.Getenv("LOG_LEVEL")

	return cfg, nil
}
