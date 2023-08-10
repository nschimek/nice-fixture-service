package repository

import (
	"github.com/nschimek/nice-fixture-service/core"
	"github.com/nschimek/nice-fixture-service/model"
)

//go:generate mockery --name Team --filename team_mock.go
type Team interface {
	GetByIdRepository[model.Team, int]
	GetAllRepository[model.Team]
	GetAllByLeagueSeason(tls *model.TeamLeagueSeason) ([]model.Team, error)
}

type team struct {
	getByIdRepository[model.Team, int]
	getAllRepository[model.Team]
	db core.Database	
}

func NewTeam(db core.Database) *team {
	return &team{
		getByIdRepository: getByIdRepository[model.Team, int]{repository: newRepo(db)},
		getAllRepository: getAllRepository[model.Team]{repository: newRepo(db)},
		db: db,
	}
}

func (r *team) GetAllByLeagueSeason(tls *model.TeamLeagueSeason) ([]model.Team, error) {
	var teams []model.Team
	if err := r.db.InnerJoin(&teams, "LeagueSeason", tls).Error; err != nil {
		return nil, err
	}
	return teams, nil
}

