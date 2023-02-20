package sport

import (
	"context"
	"time"
)

type Sport interface {
	Name() string
	Rate() float64 // todo rate
	CreatedAt() time.Time
}

type Repository interface {
	//GetRecent(ctx context.Context) (Sport, error)
	//IsSynced(ctx context.Context, after time.Time) (bool, error)

	Store(ctx context.Context, sport Sport) error // todo rate
}
