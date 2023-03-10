package processor

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"softpro6/config"
	"softpro6/internal/usecase"
	"softpro6/internal/usecase/policy"
	"softpro6/internal/valueobject"
	"softpro6/pkg/logger"
	"softpro6/pkg/worker"
	"softpro6/pkg/worker/ticker"
)

type workerPool struct {
	l logger.Interface

	workers      map[string]worker.Worker
	controlFanIn chan controlMsg
}

type controlMsg struct {
	workerId string
	err      error
}

func newWorkerPool(ctx context.Context, workersConfig []config.Worker, providers map[string]usecase.GetLineProvider, l logger.Interface, sportRepos map[valueobject.Sport]usecase.SportRepository) (*workerPool, []error) {
	wp := &workerPool{
		l:            l,
		workers:      make(map[string]worker.Worker, len(workersConfig)),
		controlFanIn: make(chan controlMsg, len(workersConfig)),
	}

	var errs []error
	for _, workerCfg := range workersConfig {
		w, err := RunWorker(ctx, workerCfg, providers, sportRepos)
		if err != nil {
			errs = append(errs, fmt.Errorf("app - Run - RunWorker - Sport %s: %w", workerCfg.Sport, err))
			continue
		}
		wp.workers[w.ID()] = w
		l.Info("worker started", zap.String("worker_id", w.ID()))
	}
	if len(errs) != 0 {
		return nil, errs
	}

	for _, w := range wp.workers {
		go func(w worker.Worker) {
			for err := range w.Notify() {
				wp.controlFanIn <- controlMsg{
					workerId: w.ID(),
					err:      err,
				}
			}
		}(w)
	}

	return wp, nil
}

func (wp workerPool) watchAndRestart(ctx context.Context) {
	go func() {
		for {
			select {
			case msg := <-wp.controlFanIn:
				wp.l.Error(fmt.Sprintf("worker error: %s", msg.err), zap.String("worker_id", msg.workerId))
				wp.workers[msg.workerId].Stop()
				wp.workers[msg.workerId].Start(ctx)
			case <-ctx.Done():
				return
			}
		}
	}()
}

func (wp workerPool) CheckingAwareLines() ([]usecase.CheckedLine, error) {
	var result []usecase.CheckedLine
	for id, w := range wp.workers {
		switch workerType := w.(type) {
		case usecase.CheckedLine:
			result = append(result, workerType)
			continue
		default:
			return nil, fmt.Errorf("workerPool - CheckingAwareLines - not all workers are CheckedLine, bad: ID %s Type %T", id, workerType)
		}
	}
	return result, nil
}

func (wp workerPool) PublishingAwareLines() ([]usecase.FirstSyncPublisher, error) {
	var result []usecase.FirstSyncPublisher
	for id, w := range wp.workers {
		switch workerType := w.(type) {
		case usecase.FirstSyncPublisher:
			result = append(result, workerType)
			continue
		default:
			return nil, fmt.Errorf("workerPool - PublishingAwareLines - not all workers are FirstSyncPublisher, bad: ID %s Type %T", id, workerType)
		}
	}
	return result, nil
}

func RunWorker(ctx context.Context, cfg config.Worker, providers map[string]usecase.GetLineProvider, repoMap map[valueobject.Sport]usecase.SportRepository) (worker.Worker, error) {
	// UseCase
	provider, isProviderFound := providers[cfg.Provider]
	if !isProviderFound {
		return nil, fmt.Errorf("unknown worker provider: %s", cfg.Provider)
	}
	getLine := usecase.NewGetLineUseCase(provider)

	// Repository
	sportRepo, isRepoFound := repoMap[valueobject.NewSport(cfg.Sport)]
	if !isRepoFound {
		return nil, fmt.Errorf("repo for sport %s not found", cfg.Sport)
	}

	// UseCase
	insertSport := usecase.NewStoreSport(sportRepo)
	uc := usecase.NewPollProvider(getLine, insertSport, cfg.Sport, new(policy.LineToSport))

	workerId := fmt.Sprintf("%s_%s", cfg.Sport, uuid.New().String())
	w := ticker.New(workerId, uc, ticker.Interval(cfg.PollInterval))
	w.Start(ctx)

	return w, nil
}
