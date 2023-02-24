package repository_test

import (
	"context"
	"github.com/stretchr/testify/require"
	"softpro6/internal/entity/sport"
	"softpro6/internal/usecase"
	"softpro6/internal/usecase/repo"
	"softpro6/internal/valueobject"
	"softpro6/pkg/postgres"
	"softpro6/test/integration"
	"testing"
	"time"
)

func TestSportRepositories(t *testing.T) {
	cfg := integration.GetConfig(t)
	repositories := createRepositories(t)
	postgres.Up(cfg.DB.MigrationsUrl, cfg.DB.URL)
	defer t.Cleanup(func() {
		postgres.Down(cfg.DB.MigrationsUrl, cfg.DB.URL)
	})

	for i := range repositories {
		r := repositories[i]

		t.Run(r.Name, func(t *testing.T) {
			t.Run("testStore", func(t *testing.T) {
				testStore(t, r)
			})
		})
	}
}

func testStore(t *testing.T, r Repository) {
	t.Helper()
	ctx := context.Background()

	testCases := []struct {
		Name  string
		Sport usecase.Sport
	}{
		{
			Name:  "positive_value",
			Sport: sport.NewBaseball(valueobject.NewRate("foo", 1.001), time.Now()),
		},
		{
			Name:  "negative_value",
			Sport: sport.NewBaseball(valueobject.NewRate("foo", -1.001), time.Now()),
		},
		{
			Name:  "zero_value",
			Sport: sport.NewBaseball(valueobject.NewRate("foo", 0), time.Now()),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			err := r.Repository.Store(ctx, tc.Sport)
			require.NoError(t, err)
		})
	}
}

func createRepositories(t *testing.T) []Repository {
	pg := integration.InitPg(t)
	return []Repository{
		{
			Name:       "BaseballRepository",
			Repository: repo.NewBaseballRepo(pg),
		},
	}
}

type Repository struct {
	Name       string
	Repository usecase.SportRepository
}
