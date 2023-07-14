package repository

import (
	"github.com/nschimek/nice-fixture-service/core"
	"github.com/nschimek/nice-fixture-service/model"
)

type League interface {
	GetAll() ([]model.League, error)
	GetAllBySeason(season *model.LeagueSeason) ([]model.League, error)
}

type league struct {
	db core.Database
}

func NewLeague(db core.Database) *league {
	return &league{db: db}
}

func (r *league) GetAll() ([]model.League, error) {
	var leagues []model.League
	// if err := r.db.GetAll(&leagues).Error; err != nil {
	// 	return nil, err
	// }
	r.db.Gorm().Preload("Seasons").Find(&leagues)
	return leagues, nil
}

func (r *league) GetAllBySeason(season *model.LeagueSeason) ([]model.League, error) {
	var leagues []model.League
	// if err := r.db.InnerJoin(leagues, "LeagueSeason", nil).Error; err != nil {
	// 	return nil, err
	// }
	// r.db.Gorm().Joins("JOIN league_seasons ON league_seasons.league_id = leagues.id AND league_seasons.season = ?", season.Season).Find(&leagues)

	r.db.Gorm().Preload("Seasons", "season = ?", season.Season).Find(&leagues)

	return leagues, nil
}