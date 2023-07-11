package repository

import (
	"github.com/nschimek/nice-fixture-service/core"
	"github.com/nschimek/nice-fixture-service/model"
)

type League interface {
	GetAll() ([]model.League, error)
	GetAllBySeason(season int) ([]model.League, error)
}

type league struct {
	db core.Database
}

func NewLeague(db core.Database) *league {
	return &league{db: db}
}

func (r *league) GetAll() ([]model.League, error) {
	var leagues []model.League
	if err := r.db.GetAll(&leagues).Error; err != nil {
		return nil, err
	}
	return leagues, nil
}

func (r *league) GetAllBySeason(season int) ([]model.League, error) {
	var leagues []model.League
	if err := r.db.InnerJoin("LeagueSeason", &model.LeagueSeason{Season: season}, leagues).Error; err != nil {
		return nil, err
	}
	return leagues, nil
}