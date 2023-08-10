package repository

import (
	"github.com/nschimek/nice-fixture-service/core"
	"github.com/nschimek/nice-fixture-service/model"
)

//go:generate mockery --name League --filename league_mock.go
type League interface {
	GetByIdRepository[model.League, int]
	GetAllRepository[model.League]
	GetAllBySeason(season *model.LeagueSeason) ([]model.League, error)
}

type league struct {
	getByIdRepository[model.League, int]
	getAllRepository[model.League]
	db core.Database
}

func NewLeague(db core.Database) *league {
	return &league{
		getByIdRepository: getByIdRepository[model.League, int]{repository: newRepo(db)},
		getAllRepository: getAllRepository[model.League]{repository: newRepo(db)},
		db: db,
	}
}

func (r *league) GetAllBySeason(season *model.LeagueSeason) ([]model.League, error) {
	var leagues []model.League
	if err := r.db.InnerJoin(&leagues, "Season", season).Error; err != nil {
		return nil, err
	}
	return leagues, nil
}