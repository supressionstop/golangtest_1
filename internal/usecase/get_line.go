package usecase

import (
	"context"
)

type GetLine struct {
	// Dependencies
	provider GetLineProvider
}

func NewGetLineUseCase(provider GetLineProvider) GetLineUseCase {
	return &GetLine{provider: provider}
}

func (uc *GetLine) Execute(ctx context.Context, sport string) (Line, error) {
	return uc.provider.GetLine(ctx, sport)
}
