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
	res := r.db.Gorm().
		Preload("Status").
		Preload("LeagueSeason").
		Preload("TeamHome").
		Preload("TeamAway").
		Joins("TeamHomeStats").
		Joins("left join team_stats hts on (hts.team_id = fixtures.team_home_id and hts.league_id = fixtures.league_id and hts.season = fixtures.season and hts.next_fixture_id = fixtures.id and fixtures.status_id = 'FT')").
		Joins("TeamAwayStats").
		Joins("left join team_stats ats on (ats.team_id = fixtures.team_away_id and ats.league_id = fixtures.league_id and ats.season = fixtures.season and ats.next_fixture_id = fixtures.id and fixtures.status_id = 'FT')").
		Preload("TeamHomeTLS.TeamStats").
		Preload("TeamAwayTLS.TeamStats").
		Where("team_home_id = ? or team_away_id = ?", 33, 33).
		Limit(10).
		Order("id asc").
		Find(&fixtures)

	if err := res.Error; err != nil {
		return nil, err
	}
	return fixtures, nil
}