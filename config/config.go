package config

import "os"

type (
	Config struct {
		App
		Log
		DB
	}

	App struct {
		Name string
	}

	Log struct {
		Level string
	}

	DB struct {
		URL string
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	// todo cobra
	cfg.App.Name = os.Getenv("APP_NAME")
	cfg.Log.Level = os.Getenv("LOG_LEVEL")
	cfg.DB.URL = os.Getenv("DB_URL")

	return cfg, nil
}
