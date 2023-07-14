package repository

import (
	"github.com/nschimek/nice-fixture-service/core"
	"github.com/nschimek/nice-fixture-service/model"
)

type Season interface {
	GetAll() ([]model.Season, error)
}

type season struct {
	db core.Database
}

func NewSeason(db core.Database) *season {
	return &season{db: db}
}

func (r *season) GetAll() ([]model.Season, error) {
	var seasons []model.Season
	if err := r.db.GroupBy(&seasons, &model.LeagueSeason{}, "season, max(current) as current", "season").Error; err != nil {
		return nil, err
	}
	return seasons, nil
}
