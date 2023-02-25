package ticker

import (
	"context"
	"sync"
	"time"
)

const (
	_defaultInterval = 5 * time.Second
)

type Worker struct {
	useCase Executor

	id                string
	interval          time.Duration
	notify            chan error
	stop              chan struct{}
	isFirstTimeSynced bool
	onlyOneSync       sync.Once
}

type Executor interface {
	Execute(context.Context) error
}

func New(id string, useCase Executor, options ...Option) *Worker {
	w := &Worker{
		useCase: useCase,

		id:       id,
		interval: _defaultInterval,
		notify:   make(chan error, 1),
		stop:     make(chan struct{}, 1),
	}

	for _, option := range options {
		option(w)
	}

	return w
}

func (p *Worker) ID() string {
	return p.id
}

func (p *Worker) Start(ctx context.Context) {
	go p.start(ctx)
}

func (p *Worker) start(ctx context.Context) {
	ticker := time.NewTicker(p.interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			err := p.useCase.Execute(ctx)
			if err != nil {
				p.notify <- err
			}
			p.onlyOneSync.Do(func() {
				p.isFirstTimeSynced = true
			})
		case <-ctx.Done():
			return
		case <-p.stop:
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

func (p *Worker) IsSynced() (bool, error) {
	return p.isFirstTimeSynced, nil
}
