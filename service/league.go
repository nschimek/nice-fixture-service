package service

import (
	"github.com/nschimek/nice-fixture-service/model"
	"github.com/nschimek/nice-fixture-service/repository"
)

type League interface {
	GetByParams(params model.LeagueParams) ([]model.League, error)
	GetById(id int) (*model.League, error)
}

type league struct {
	repo repository.League
}

func NewLeague(repo repository.League) *league {
	return &league{repo: repo}
}

func (s *league) GetByParams(params model.LeagueParams) ([]model.League, error) {
	r, err := s.repo.GetAllBySeason(&model.LeagueSeason{Season: params.Season, Current: params.Current})

	// remove entries with no seasons
	if err == nil {
		o := []model.League{}
		for _, l := range r {
			if len(l.Seasons) > 0 {
				o = append(o, l)
			}
		}
		return o, nil
	}

	return nil, err
}

func (s *league) GetById(id int) (*model.League, error) {
	return s.repo.GetById(id)
}