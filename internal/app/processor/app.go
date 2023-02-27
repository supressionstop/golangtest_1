package processor

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net/http"
	"os"
	"os/signal"
	"softpro6/config"
	grpcv1 "softpro6/internal/controller/grpc/v1"
	"softpro6/internal/controller/grpc/v1/pb"
	httpv1 "softpro6/internal/controller/http/v1"
	"softpro6/internal/usecase"
	"softpro6/internal/usecase/repo"
	"softpro6/internal/valueobject"
	"softpro6/pkg/grpcsrv"
	"softpro6/pkg/httpsrv"
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
	pg, err := postgres.NewWithContext(ctx, cfg.DB.URL)
	if err != nil {
		l.Fatal("app - Run - postgres.New", err)
	}
	defer pg.Close()
	sportRepoMap := map[valueobject.Sport]usecase.SportRepository{
		valueobject.Baseball: repo.NewBaseballRepo(pg),
	}

	// workers
	providers, err := initProviders(cfg.Providers)
	if err != nil {
		l.Fatal("app - Run - initProviders", err)
	}
	wp, errs := newWorkerPool(ctx, cfg.Workers, providers, l, sportRepoMap)
	if len(errs) != 0 {
		l.Fatal("app - Run - newWorkerPool", errs)
	}
	wp.watchAndRestart(ctx)

	// http server
	checkingAwareWorkers, err := wp.CheckingAwareLines()
	if err != nil {
		l.Fatal("app - Run - CheckingAwareLines", err)
	}
	isAppReady := usecase.NewIsAppReady(pg, checkingAwareWorkers)
	chiRouter := chi.NewRouter()
	httpv1.NewRouter(chiRouter, isAppReady, l)
	httpServer := httpsrv.NewAndStartOnAddr(chiRouter, cfg.HttpServer.Address)
	l.Info(fmt.Sprintf("http server started on %s", cfg.HttpServer.Address))
	l.Info(fmt.Sprintf("swagger available on %s", cfg.HttpServer.Address+"/swagger/"))

	// grpc server
	publishingAwareWorkers, err := wp.PublishingAwareLines()
	if err != nil {
		l.Fatal("app - Run - PublishingAwareLines", err)
	}
	firstSyncSubscribe := usecase.NewFirstSyncSubscribe()
	subscription, err := firstSyncSubscribe.Execute(ctx, publishingAwareWorkers)
	if err != nil {
		l.Fatal("app - Run - PublishingAwareLines", err)
	}
	getRecentSport := usecase.NewGetRecentSport(sportRepoMap)
	grpcServer, err := grpcsrv.NewServer(cfg.GrpcServer.Address, l, func(server *grpc.Server) {
		oursGrpcServer := grpcv1.NewGrpcServer(getRecentSport, l)
		pb.RegisterProcessorServiceServer(server, oursGrpcServer)
	})
	if err != nil {
		l.Fatal("app - Run - grpcsrv.NewServerAndRun", err)
	}
	grpcServer.StartLater(subscription.IsSynced())

	l.Info(fmt.Sprintf("%s started", cfg.App.Name))

	select {
	case s := <-interrupt:
		l.Info("got signal from os", zap.String("signal", s.String()))
	case s := <-httpServer.Notify():
		l.Error("app - Run - httpServer.Notify", s)
	case s := <-grpcServer.Notify():
		l.Error("app - Run - grpcServer.Notify", s)
	}

	// shutdown
	l.Info("http server is shutting")
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
	l.Info("http server has been shut down")

	l.Info("grpc server is shutting")
	grpcServer.Shutdown()
	l.Info("grpc server has been shut down")

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
