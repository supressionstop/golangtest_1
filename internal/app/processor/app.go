package processor

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"softpro6/config"
	"softpro6/internal/usecase"
	"softpro6/internal/usecase/repo"
	"softpro6/pkg/logger"
	"softpro6/pkg/postgres"
	"softpro6/pkg/providers/kiddy"
	"syscall"
)

func Run(cfg *config.Config) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	ctx := context.Background()

	l := logger.New(cfg.Log.Level, cfg.App.Environment, cfg.App.Name)

	// storage
	pg, err := postgres.New(cfg.DB.URL)
	if err != nil {
		l.Fatal("app - Run - postgres.New", err)
	}
	defer pg.Close()

	// workers
	providers, err := initProviders(cfg.Providers)
	if err != nil {
		l.Fatal("app - Run - initProviders", err)
	}
	wp, errs := newWorkerPool(ctx, cfg.Workers, providers, l, pg)
	if len(errs) != 0 {
		l.Fatal("app - Run - newWorkerPool", errs)
	}
	wp.watchAndRestart(ctx)

	l.Info(fmt.Sprintf("%s started", cfg.App.Name))

	select {
	case s := <-interrupt:
		l.Info("got signal from os", zap.String("signal", s.String()))
	}

	l.Info("done")
}

func initProviders(providersCfg map[string]config.Provider) (map[string]usecase.GetLineProvider, error) {
	m := map[string]usecase.GetLineProvider{}
	for providerId, cfg := range providersCfg {
		provider, err := providerFactory(providerId, &http.Client{Timeout: cfg.HttpTimeout}, cfg.BaseUrl)
		if err != nil {
			return nil, err
		}
		m[providerId] = provider
	}
	return m, nil
}

func providerFactory(id string, httpClient *http.Client, baseUrl string) (usecase.GetLineProvider, error) {
	switch id {
	case "kiddy":
		return kiddy.NewKiddy(httpClient, baseUrl)
	default:
		return nil, fmt.Errorf("unknown provider: %q", id)
	}
}

func repositoryFactory(pg *postgres.Postgres, sport string) (usecase.SportRepository, error) {
	switch sport {
	case "baseball":
		return repo.NewBaseballRepo(pg), nil
	default:
		return nil, fmt.Errorf("unknown sport: %q", sport)
	}
}
