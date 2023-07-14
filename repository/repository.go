package repository

import "github.com/nschimek/nice-fixture-service/core"

type RepositoryRegistry struct {
	League League
}

func Setup(db core.Database) *RepositoryRegistry {
	core.Log.Debug("Setting up Repositories...")
	return &RepositoryRegistry{
		League: NewLeague(db),
	}
}