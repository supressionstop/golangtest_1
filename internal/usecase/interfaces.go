package usecase

import (
	"context"
	"time"
)

type (
	GetLineUseCase interface {
		Execute(ctx context.Context, sport string) (Line, error)
	}

	GetLineProvider interface {
		GetLine(ctx context.Context, sportName string) (Line, error)
	}

	Line interface {
		Sport() string
		Rate() string
	}
)

type PollProviderUseCase interface {
	Execute(ctx context.Context) error
}

type StoreSportUseCase interface {
	Execute(ctx context.Context, sport Sport) error
}

// Entities and repos
type (
	Sport interface {
		Name() string
		Rate() float64 // todo rate
		CreatedAt() time.Time
	}

	SportRepository interface {
		GetRecent(ctx context.Context) (Sport, error)
		IsSynced(ctx context.Context, after time.Time) (bool, error)
		Store(ctx context.Context, sport Sport) error // todo rate
	}
)

// Policies -

type (
	LineToSportPolicy interface {
		Export(Line) (Sport, error)
	}
)
