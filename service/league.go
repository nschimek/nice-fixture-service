package service

import (
	"github.com/nschimek/nice-fixture-service/model"
	"github.com/nschimek/nice-fixture-service/repository"
)

type League interface {
	GetByParams(params model.LeagueParams) ([]model.League, error)
}

type league struct {
	repo repository.League
}

func NewLeague(repo repository.League) *league {
	return &league{repo: repo}
}

func (s *league) GetByParams(params model.LeagueParams) ([]model.League, error) {
	// if there's a season specified, we have to join to league_season - so it's a different repo method
	if params.Season > 0 {
		return s.repo.GetAllBySeason(&model.LeagueSeason{Season: params.Season})
	}
	return s.repo.GetAll()
}