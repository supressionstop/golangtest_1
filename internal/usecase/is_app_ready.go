package usecase

import (
	"context"
	"errors"
	"softpro6/internal/entity"
)

type IsAppReady struct {
	storage CheckedStorage
	lines   []CheckedLine
}

func NewIsAppReady(storage CheckedStorage, lines []CheckedLine) *IsAppReady {
	return &IsAppReady{
		storage: storage,
		lines:   lines,
	}
}

func (uc *IsAppReady) Execute(ctx context.Context) (Readiness, error) {
	var errs []error
	isReady, err := uc.storage.IsReady()
	if err != nil {
		errs = append(errs, err)
	}
	if !isReady {
		errs = append(errs, errors.New("storage is not ready"))
	}

	for _, line := range uc.lines {
		isSynced, err := line.IsSynced()
		if err != nil {
			errs = append(errs, err)
		}
		if !isSynced {
			errs = append(errs, errors.New("line is not ready"))
		}
	}

	return entity.NewReadiness(errs), nil
}
