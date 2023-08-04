package service

import (
	"testing"

	"github.com/nschimek/nice-fixture-service/model"
	"github.com/stretchr/testify/assert"
)

func TestOnlyPopulatedChildren(t *testing.T) {
	d := []model.Team{
		{Id: 1, TeamLeagueSeasons: []model.TeamLeagueSeason{{LeagueId: 1, Season: 2022}}},
		{Id: 2, TeamLeagueSeasons: []model.TeamLeagueSeason{{LeagueId: 2, Season: 2022}}},
		{Id: 3, TeamLeagueSeasons: []model.TeamLeagueSeason{}},
	}

	res := OnlyPopulatedChildren[model.Team, model.TeamLeagueSeason](d, func(p model.Team) []model.TeamLeagueSeason {
		return p.TeamLeagueSeasons
	})

	assert.Contains(t, res, d[0])
	assert.Contains(t, res, d[1])
	assert.NotContains(t, res, d[2])
}

