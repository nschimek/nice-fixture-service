package service

import (
	"github.com/nschimek/nice-fixture-service/core"
	"github.com/nschimek/nice-fixture-service/repository"
)

type ServiceRegistry struct {
	League League
	Season Season
	Team Team
	Fixture Fixture
}

func Setup(repos *repository.RepositoryRegistry) *ServiceRegistry {
	core.Log.Debug("Setting up Services...")
	return &ServiceRegistry{
		League: NewLeague(repos.League),
		Season: NewSeason(repos.Season),
		Team: NewTeam(repos.Team),
		Fixture: NewFixture(repos.Fixture),
	}
}