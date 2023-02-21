package usecase

import (
	"context"
)

type StoreSport struct {
	// Dependencies
	repo SportRepository
}

func NewStoreSport(sportRepo SportRepository) *StoreSport {
	return &StoreSport{repo: sportRepo}
}

func (uc *StoreSport) Execute(ctx context.Context, sport Sport) error {
	return uc.repo.Store(ctx, sport)
}
