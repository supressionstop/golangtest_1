package ticker

import (
	"context"
	"softpro6/pkg/logger"
	"time"
)

const (
	_defaultName     = "ticker"
	_defaultInterval = 5 * time.Second
)

type Worker struct {
	logger  logger.Interface
	useCase Executor

	name     string
	interval time.Duration
	notify   chan error
	stop     chan struct{}
}

type Executor interface {
	Execute(context.Context) error
}

func New(useCase Executor, logger logger.Interface, options ...Option) *Worker {
	w := &Worker{
		logger:  logger,
		useCase: useCase,

		name:     _defaultName,
		interval: _defaultInterval,
		notify:   make(chan error, 1),
		stop:     make(chan struct{}, 1),
	}

	for _, option := range options {
		option(w)
	}

	return w
}

func (p *Worker) Start(ctx context.Context) {
	go p.start(ctx)
	p.logger.Debug("worker started")
}

func (p *Worker) start(ctx context.Context) {
	ticker := time.NewTicker(p.interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			err := p.useCase.Execute(ctx)
			if err != nil {
				p.logger.Debug("worker got err", err)
				p.notify <- err
			}
		case <-ctx.Done():
			p.logger.Debug("worker got ctx.Done")
			return
		case <-p.stop:
			p.logger.Debug("worker got stop signal")
			return
		}
	}
}

func (p *Worker) Stop() {
	p.stop <- struct{}{}
}

func (p *Worker) Restart(ctx context.Context) {
	p.Stop()
	p.Start(ctx)
}

func (p *Worker) Notify() <-chan error {
	return p.notify
}
