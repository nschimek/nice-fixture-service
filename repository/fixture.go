package repository

import (
	"github.com/nschimek/nice-fixture-service/core"
	"github.com/nschimek/nice-fixture-service/model"
)

type Fixture interface {
	GetAll() ([]model.Fixture, error)
}

type fixture struct {
	db core.Database
}

func NewFixture(db core.Database) *fixture {
	return &fixture{db: db}
}

func (r *fixture) GetAll() ([]model.Fixture, error) {
	var fixtures []model.Fixture	
	if err := r.db.Gorm().Preload("LeagueSeason").Preload("Status").Preload("Teams.Home").Preload("Teams.Away").Find(&fixtures).Error; err != nil {
		return nil, err
	}
	return fixtures, nil
}