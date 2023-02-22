package config

import (
	"github.com/spf13/viper"
	"time"
)

type (
	Config struct {
		App       `mapstructure:"app"`
		Log       `mapstructure:"log"`
		DB        `mapstructure:"db"`
		Workers   []Worker            `mapstructure:"workers"`
		Providers map[string]Provider `mapstructure:"providers"`
	}

	App struct {
		Name        string `mapstructure:"name"`
		Environment string `mapstructure:"environment"`
	}

	Log struct {
		Level string `mapstructure:"level"`
	}

	DB struct {
		URL string `mapstructure:"url"`
	}

	Worker struct {
		Sport        string        `mapstructure:"sport"`
		PollInterval time.Duration `mapstructure:"poll_interval"`
		Provider     string        `mapstructure:"provider"`
	}

	Provider struct {
		BaseUrl     string        `mapstructure:"base_url"`
		HttpTimeout time.Duration `mapstructure:"http_timeout"`
	}
)

func NewConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	err = viper.Unmarshal(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
