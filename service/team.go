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
	var teams []model.Team
	var err error

	if params == (model.TeamParams{}) {
		teams, err = s.repo.GetAll()
	} else {
		teams, err = s.repo.GetAllByLeagueSeason(&model.TeamLeagueSeason{
			Season: params.Season, 
			LeagueId: params.League,
		})
	}
	
	return teams, rest_error.NewInternal(err)
}

func (s *team) GetById(id int) (*model.Team, *rest_error.Error) {
	r, err := s.repo.GetById(id)
	return r, rest_error.NewInternal(err)
}