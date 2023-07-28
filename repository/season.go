package repository

import (
	"github.com/nschimek/nice-fixture-service/core"
	"github.com/nschimek/nice-fixture-service/model"
)

const seasonGroupBySelectSql = "season, max(current) as current"

//go:generate mockery --name Season --filename season_mock.go
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
	if err := r.db.GroupBy(&seasons, &model.LeagueSeason{}, seasonGroupBySelectSql, "season").Error; err != nil {
		return nil, err
	}
	return seasons, nil
}
