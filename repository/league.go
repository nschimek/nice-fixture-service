package repository

import (
	"github.com/nschimek/nice-fixture-service/core"
	"github.com/nschimek/nice-fixture-service/model"
)

type League interface {
	GetByIdRepository[model.League, int]
	GetAllBySeason(season *model.LeagueSeason) ([]model.League, error)
}

type league struct {
	getByIdRepository[model.League, int]
	db core.Database
}

func NewLeague(db core.Database) *league {
	r := newRepo(db)
	return &league{
		getByIdRepository: getByIdRepository[model.League, int]{repository: r},
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