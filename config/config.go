package config

import (
	"os"
	"time"
)

type (
	Config struct {
		App
		Log
		DB
		Workers []Worker
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

	Worker struct {
		ID           string
		PollInterval string // todo to duration
		Provider     Provider
	}

	Provider struct {
		ID          string
		BaseUrl     string
		HttpTimeout time.Duration
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
