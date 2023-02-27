package usecase

import (
	"context"
	"fmt"
	"softpro6/internal/valueobject"
)

type GetRecentSports struct {
	sportRepos map[valueobject.Sport]SportRepository
}

func NewGetRecentSport(repos map[valueobject.Sport]SportRepository) *GetRecentSports {
	return &GetRecentSports{sportRepos: repos}
}

func (uc *GetRecentSports) Execute(ctx context.Context, sports ...valueobject.Sport) ([]Sport, error) {
	var notFoundSports []valueobject.Sport
	var repos []SportRepository
	for i := range sports {
		repository, isRepoFound := uc.sportRepos[sports[i]]
		if !isRepoFound {
			notFoundSports = append(notFoundSports, sports[i])
			continue
		}
		repos = append(repos, repository)
	}

	if len(notFoundSports) > 0 {
		return nil, fmt.Errorf("getRecentSport - repositories not found for sports %v", notFoundSports)
	}

	var result []Sport
	for i := range repos {
		recentSport, err := repos[i].GetRecent(ctx)
		if err != nil {
			return nil, fmt.Errorf("getRecentSport - %T.GetRecent %w", repos[i], err)
		}
		result = append(result, recentSport)
	}

	return result, nil
}
