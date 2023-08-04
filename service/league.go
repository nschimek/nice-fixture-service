package service

import (
	"github.com/nschimek/nice-fixture-service/model"
	"github.com/nschimek/nice-fixture-service/model/rest_error"
	"github.com/nschimek/nice-fixture-service/repository"
)

//go:generate mockery --name League --filename league_mock.go
type League interface {
	GetByParams(params model.LeagueParams) ([]model.League, *rest_error.Error)
	GetById(id int) (*model.League, *rest_error.Error)
}

type league struct {
	repo repository.League
}

func NewLeague(repo repository.League) *league {
	return &league{repo: repo}
}

func (s *league) GetByParams(params model.LeagueParams) ([]model.League, *rest_error.Error) {
	if r, err := s.repo.GetAllBySeason(&model.LeagueSeason{
			Season: params.Season, 
			Current: params.Current,
		}); err == nil {
		return OnlyPopulatedChildren[model.League, model.LeagueSeason](r, func(p model.League) []model.LeagueSeason {
			return p.Seasons
		}), nil
	} else {
		return nil, rest_error.NewInternal(err)
	}
}

func (s *league) GetById(id int) (*model.League, *rest_error.Error) {
	r, err := s.repo.GetById(id)
	return r, rest_error.NewInternal(err)
}