package usecase_test

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"softpro6/internal/usecase"
	"testing"
)

func makeIsAppReady(t *testing.T) (*usecase.IsAppReady, *MockCheckedStorage, []*MockCheckedLine) {
	t.Helper()
	goMockCtl := gomock.NewController(t)
	defer goMockCtl.Finish()

	storage := NewMockCheckedStorage(goMockCtl)
	a, b, c := NewMockCheckedLine(goMockCtl), NewMockCheckedLine(goMockCtl), NewMockCheckedLine(goMockCtl)

	uc := usecase.NewIsAppReady(storage, []usecase.CheckedLine{a, b, c})

	return uc, storage, []*MockCheckedLine{a, b, c}
}

func TestIsAppReady_Execute(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		goMockCtl := gomock.NewController(t)
		defer goMockCtl.Finish()

		// arrange
		var e []error
		readinessMock := NewMockReadiness(goMockCtl)
		readinessMock.EXPECT().IsReady().Return(true)
		readinessMock.EXPECT().Reasons().Return(e)
		isAppReady, storage, lines := makeIsAppReady(t)
		storage.EXPECT().IsReady().Return(true, nil)
		for i := range lines {
			lines[i].EXPECT().IsSynced().Return(true, nil)
		}

		// act
		readiness, err := isAppReady.Execute(context.Background())

		// assert
		require.NoError(t, err)
		require.EqualValues(t, readinessMock.IsReady(), readiness.IsReady())
		require.EqualValues(t, readinessMock.Reasons(), readiness.Reasons())
	})
	t.Run("not synced lines", func(t *testing.T) {
		t.Parallel()
		goMockCtl := gomock.NewController(t)
		defer goMockCtl.Finish()

		// arrange
		var e []error
		readinessMock := NewMockReadiness(goMockCtl)
		readinessMock.EXPECT().IsReady().Return(false)
		isAppReady, storage, lines := makeIsAppReady(t)
		storage.EXPECT().IsReady().Return(true, nil)
		for i := range lines {
			if i%2 == 0 {
				lines[i].EXPECT().IsSynced().Return(false, nil)
				e = append(e, errors.New("line is not ready"))
				continue
			}
			lines[i].EXPECT().IsSynced().Return(true, nil)
		}
		readinessMock.EXPECT().Reasons().Return(e)

		// act
		readiness, err := isAppReady.Execute(context.Background())

		// assert
		require.NoError(t, err)
		require.EqualValues(t, readinessMock.IsReady(), readiness.IsReady())
		require.EqualValues(t, readinessMock.Reasons(), readiness.Reasons())
	})
	t.Run("not synced storage", func(t *testing.T) {
		t.Parallel()
		goMockCtl := gomock.NewController(t)
		defer goMockCtl.Finish()

		// arrange
		var e []error
		readinessMock := NewMockReadiness(goMockCtl)
		readinessMock.EXPECT().IsReady().Return(false)
		isAppReady, storage, lines := makeIsAppReady(t)
		storage.EXPECT().IsReady().Return(false, nil)
		e = append(e, errors.New("storage is not ready"))
		for i := range lines {
			lines[i].EXPECT().IsSynced().Return(true, nil)
		}
		readinessMock.EXPECT().Reasons().Return(e)

		// act
		readiness, err := isAppReady.Execute(context.Background())

		// assert
		require.NoError(t, err)
		require.EqualValues(t, readinessMock.IsReady(), readiness.IsReady())
		require.EqualValues(t, readinessMock.Reasons(), readiness.Reasons())
	})
}
