package repository

import (
	"github.com/nschimek/nice-fixture-service/core"
	"github.com/nschimek/nice-fixture-service/model"
)

const teamLeagueSeasons = "TeamLeagueSeasons"

//go:generate mockery --name Team --filename team_mock.go
type Team interface {
	GetByIdRepository[model.Team, int]
	GetAllByLeagueSeason(tls *model.TeamLeagueSeason) ([]model.Team, error)
}

type team struct {
	getByIdRepository[model.Team, int]
	db core.Database	
}

func NewTeam(db core.Database) *team {
	return &team{
		getByIdRepository: getByIdRepository[model.Team, int]{repository: newRepo(db)},
		db: db,
	}
}

func (r *team) GetAllByLeagueSeason(tls *model.TeamLeagueSeason) ([]model.Team, error) {
	var teams []model.Team
	if err := r.db.Preload(&teams, nil, teamLeagueSeasons, tls).Error; err != nil {
		return nil, err
	}
	return teams, nil
}

