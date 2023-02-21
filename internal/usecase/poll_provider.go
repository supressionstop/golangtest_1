package usecase

import (
	"context"
)

type PollProvider struct {
	// Dependencies
	getLine     GetLineUseCase
	insertSport StoreSportUseCase

	// Arguments
	sport string

	// Policies
	lineToSport LineToSportPolicy
}

func NewPollProvider(getLine GetLineUseCase, insertSport StoreSportUseCase, sport string, port LineToSportPolicy) *PollProvider {
	return &PollProvider{
		getLine:     getLine,
		insertSport: insertSport,
		sport:       sport,
		lineToSport: port,
	}
}

func (uc *PollProvider) Execute(ctx context.Context) error {
	line, err := uc.getLine.Execute(ctx, uc.sport)
	if err != nil {
		return err
	}

	entity, err := uc.lineToSport.Export(line)
	if err != nil {
		return err
	}

	return uc.insertSport.Execute(ctx, entity)
}

// Policies
