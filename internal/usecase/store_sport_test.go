package usecase_test

import (
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"softpro6/internal/usecase"
	"testing"
)

func makeStoreSport(t *testing.T) (usecase.StoreSportUseCase, *MockSportRepository) {
	t.Helper()

	goMockCtl := gomock.NewController(t)
	defer goMockCtl.Finish()

	repo := NewMockSportRepository(goMockCtl)

	return usecase.NewStoreSport(repo), repo
}

func TestStoreSport_Execute(t *testing.T) {
	t.Parallel()
	storeSport, sportRepository := makeStoreSport(t)

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		goMockCtl := gomock.NewController(t)
		defer goMockCtl.Finish()

		// arrange
		concreteSport := NewMockSport(goMockCtl)
		sportRepository.EXPECT().Store(context.Background(), concreteSport).Return(nil)

		// act
		err := storeSport.Execute(context.Background(), concreteSport)

		// assert
		require.NoError(t, err)
	})
	t.Run("unknown sport", func(t *testing.T) {
		t.Parallel()
		goMockCtl := gomock.NewController(t)
		defer goMockCtl.Finish()

		// arrange
		sportName := "fooBarBaz"
		unknownSportErr := fmt.Errorf("unknown sport provided %q", sportName)
		unknownSport := NewMockSport(goMockCtl)
		sportRepository.EXPECT().Store(context.Background(), unknownSport).Return(unknownSportErr)

		// act
		err := storeSport.Execute(context.Background(), unknownSport)

		// assert
		require.ErrorIs(t, unknownSportErr, err)
	})
}
