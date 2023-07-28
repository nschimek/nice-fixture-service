package repository

import "github.com/nschimek/nice-fixture-service/core"

type RepositoryRegistry struct {
	League League
	Season Season
	Team Team
}

type repository struct {
	db core.Database
}

func newRepo(db core.Database) *repository {
	return &repository{db: db}
}

func Setup(db core.Database) *RepositoryRegistry {
	core.Log.Debug("Setting up Repositories...")
	return &RepositoryRegistry{
		League: NewLeague(db),
		Season: NewSeason(db),
		Team: NewTeam(db),
	}
}

type GetByIdRepository[T any, I any] interface {
	GetById(id I) (*T, error)
}

type getByIdRepository[T any, I any] struct {
	*repository
}

func (r getByIdRepository[T, I]) GetById(id I) (*T, error) {
	var dest T
	if res := r.db.GetById(id, &dest); res.RowsAffected == 0 && res.Error == nil {
		return nil, nil
	} else if res.Error != nil {
		return nil, res.Error
	}
	return &dest, nil
}