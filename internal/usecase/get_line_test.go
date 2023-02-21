package usecase_test

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"softpro6/internal/usecase"
	"testing"
)

type test struct {
	name string
	mock func()
	res  interface{}
	err  error
}

func makeGetLine(t *testing.T) (usecase.GetLineUseCase, *MockGetLineProvider) {
	t.Helper()

	goMockCtl := gomock.NewController(t)
	defer goMockCtl.Finish()

	provider := NewMockGetLineProvider(goMockCtl)

	return usecase.NewGetLineUseCase(provider), provider
}

func TestGetLine_Execute(t *testing.T) {
	t.Parallel()
	getLine, provider := makeGetLine(t)

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		goMockCtl := gomock.NewController(t)
		defer goMockCtl.Finish()

		// arrange
		sportName := "baseball"
		line := NewMockLine(goMockCtl)
		line.EXPECT().Sport().Return(sportName).Times(2)
		line.EXPECT().Rate().Return("1.005").Times(2)
		provider.EXPECT().GetLine(context.Background(), sportName).Return(line, nil)

		// act
		res, err := getLine.Execute(context.Background(), sportName)

		// assert
		require.EqualValues(t, line.Sport(), res.Sport())
		require.EqualValues(t, line.Rate(), res.Rate())
		require.ErrorIs(t, nil, err)
	})
	t.Run("unknown sport", func(t *testing.T) {
		t.Parallel()

		// arrange
		customErr := errors.New("unknown sport")
		sportName := "fooBarBaz"
		provider.EXPECT().GetLine(context.Background(), sportName).Return(nil, customErr)

		// act
		res, err := getLine.Execute(context.Background(), sportName)

		// assert
		require.EqualValues(t, nil, res)
		require.ErrorIs(t, customErr, err)
	})
}
