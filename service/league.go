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
	var r []model.League
	var err error

	// if there are no search params, there is no need to join to season
	if params == (model.LeagueParams{}) {
		r, err = s.repo.GetAll()
	} else {
		r, err = s.repo.GetAllBySeason(&model.LeagueSeason{
			Season: params.Season, 
			Current: params.Current,
		})
	}

	return r, rest_error.NewInternal(err)
}

func (s *league) GetById(id int) (*model.League, *rest_error.Error) {
	r, err := s.repo.GetById(id)
	return r, rest_error.NewInternal(err)
}