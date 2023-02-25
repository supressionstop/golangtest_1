package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/url"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

type (
	Config struct {
		App        `mapstructure:"app"`
		Log        `mapstructure:"log"`
		DB         `mapstructure:"db"`
		Workers    []Worker            `mapstructure:"workers"`
		Providers  map[string]Provider `mapstructure:"providers"`
		HttpServer `mapstructure:"http_server"`
	}

	App struct {
		Name        string `mapstructure:"name"`
		Environment string `mapstructure:"environment"`
	}

	Log struct {
		Level string `mapstructure:"level"`
	}

	DB struct {
		URL           string `mapstructure:"url"`
		MigrationsUrl string `mapstructure:"migrations_url"`
	}

	Worker struct {
		Sport        string        `mapstructure:"sport"`
		PollInterval time.Duration `mapstructure:"poll_interval"`
		Provider     string        `mapstructure:"provider" example:"some_provider"`
	}

	Provider struct {
		BaseUrl     string        `mapstructure:"base_url" example:"http://localhost:8080"`
		HttpTimeout time.Duration `mapstructure:"http_timeout" example:"5s"`
	}

	HttpServer struct {
		Address string `mapstructure:"address" example:":80"`
	}
)

func NewConfig(environment string) (*Config, error) {
	// 1st priority - env
	// E.g:
	// (struct) Config.DB.URL == DB_URL (env)
	viper.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))
	viper.AutomaticEnv()
	// 2nd priority - json
	if environment == "" {
		viper.SetConfigName("config")
	} else {
		viper.SetConfigName(fmt.Sprintf("config_%s", strings.ToLower(environment)))
	}
	viper.SetConfigType("json")

	_, b, _, _ := runtime.Caller(0)
	configFolderPath := filepath.Dir(b)
	viper.AddConfigPath(configFolderPath)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	err = viper.Unmarshal(cfg)
	if err != nil {
		return nil, err
	}

	appRootPath := filepath.Join(b, "../..")
	setPathsFromRoot(appRootPath, cfg)

	return cfg, nil
}

func setPathsFromRoot(projectRoot string, config *Config) {
	// migrationsUrl
	mUrl, err := url.Parse(config.DB.MigrationsUrl)
	if err != nil {
		log.Fatalf("NewConfig - setPathsFromRoot - MigrationsUrl: %s", err)
	}
	newUrl := url.URL{Path: projectRoot}
	newUrl.Scheme = mUrl.Scheme
	newUrl = *newUrl.JoinPath(mUrl.Path)
	config.DB.MigrationsUrl = newUrl.String()
}
