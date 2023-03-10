package usecase

import (
	"context"
	"softpro6/internal/valueobject"
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
		Provider() string
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

type (
	IsAppReadyUseCase interface {
		Execute(ctx context.Context) (Readiness, error)
	}

	CheckedStorage interface {
		IsReady() (bool, error)
	}

	CheckedLine interface {
		IsSynced() (bool, error)
	}
)

type (
	FirstSyncSubscribeUseCase interface {
		Execute(ctx context.Context, publishers []FirstSyncPublisher) (FirstSyncSubscription, error)
	}

	FirstSyncPublisher interface {
		IAmSynced() <-chan struct{}
	}

	FirstSyncSubscription interface {
		IsSynced() <-chan struct{}
	}
)

type (
	GetRecentSportsUseCase interface {
		Execute(ctx context.Context, sports ...valueobject.Sport) ([]Sport, error)
	}
)

// Entities and repos
type (
	Sport interface {
		Name() string
		Rate() valueobject.Rate
		CreatedAt() time.Time
	}
	Readiness interface {
		IsReady() bool
		Reasons() []error
	}

	SportRepository interface {
		GetRecent(ctx context.Context) (Sport, error)
		IsSynced(ctx context.Context, after time.Time) (bool, error)
		Store(ctx context.Context, sport Sport) error
	}
)

// Policies

type (
	LineToSportPolicy interface {
		Export(Line) (Sport, error)
	}
)
