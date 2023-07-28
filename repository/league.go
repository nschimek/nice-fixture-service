package repository

import (
	"github.com/nschimek/nice-fixture-service/core"
	"github.com/nschimek/nice-fixture-service/model"
)

//go:generate mockery --name League --filename league_mock.go
type League interface {
	GetByIdRepository[model.League, int]
	GetAllBySeason(season *model.LeagueSeason) ([]model.League, error)
}

type league struct {
	getByIdRepository[model.League, int]
	db core.Database
}

func NewLeague(db core.Database) *league {
	return &league{
		getByIdRepository: getByIdRepository[model.League, int]{repository: newRepo(db)},
		db: db,
	}
}

func (r *league) GetAllBySeason(season *model.LeagueSeason) ([]model.League, error) {
	var leagues []model.League
	if err := r.db.Preload(&leagues, nil, "Seasons", season).Error; err != nil {
		return nil, err
	}
	return leagues, nil
}