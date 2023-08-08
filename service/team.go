package service

import (
	"github.com/nschimek/nice-fixture-service/model"
	"github.com/nschimek/nice-fixture-service/model/rest_error"
	"github.com/nschimek/nice-fixture-service/repository"
)

//go:generate mockery --name Team --filename team_mock.go
type Team interface {
	GetByParams(params model.TeamParams) ([]model.Team, *rest_error.Error)
	GetById(id int) (*model.Team, *rest_error.Error)
}

type team struct {
	repo repository.Team
}

func NewTeam(repo repository.Team) *team {
	return &team{repo: repo}
}

func (s *team) GetByParams(params model.TeamParams) ([]model.Team, *rest_error.Error) {
	if r, err := s.repo.GetAllByLeagueSeason(&model.TeamLeagueSeason{
			Season: params.Season, 
			LeagueId: params.League,
		}); err == nil {
		return OnlyPopulatedChildren[model.Team, model.TeamLeagueSeason](r, func(p model.Team) []model.TeamLeagueSeason {
			return p.TeamLeagueSeasons
		}), nil
	} else {
		return nil, rest_error.NewInternal(err)
	}
}

func (s *team) GetById(id int) (*model.Team, *rest_error.Error) {
	r, err := s.repo.GetById(id)
	return r, rest_error.NewInternal(err)
}