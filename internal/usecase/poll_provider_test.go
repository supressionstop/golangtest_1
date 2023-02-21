package usecase_test

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"softpro6/internal/usecase"
	"testing"
)

func makePollProvider(t *testing.T, sportName string) (*usecase.PollProvider, *MockGetLineUseCase, *MockStoreSportUseCase, *MockLineToSportPolicy) {
	t.Helper()
	goMockCtl := gomock.NewController(t)
	defer goMockCtl.Finish()

	getLineMock := NewMockGetLineUseCase(goMockCtl)
	storeSportMock := NewMockStoreSportUseCase(goMockCtl)
	lineSportPolicyMock := NewMockLineToSportPolicy(goMockCtl)

	uc := usecase.NewPollProvider(getLineMock, storeSportMock, sportName, lineSportPolicyMock)

	return uc, getLineMock, storeSportMock, lineSportPolicyMock
}

func TestPollProvider_Execute(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		goMockCtl := gomock.NewController(t)
		defer goMockCtl.Finish()

		// arrange
		sportName := "baseball"
		line := NewMockLine(goMockCtl)
		sport := NewMockSport(goMockCtl)
		pollProvider, getLineUseCase, storeSportUseCase, lineToSportPolicy := makePollProvider(t, sportName)
		getLineUseCase.EXPECT().Execute(context.Background(), sportName).Return(line, nil)
		lineToSportPolicy.EXPECT().Export(line).Return(sport, nil)
		storeSportUseCase.EXPECT().Execute(context.Background(), NewMockSport(goMockCtl))

		// act
		err := pollProvider.Execute(context.Background())

		// assert
		require.NoError(t, err)
	})
}
