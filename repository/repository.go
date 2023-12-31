package repository

import "github.com/nschimek/nice-fixture-service/core"

type RepositoryRegistry struct {
	League League
	Season Season
	Team Team
	Fixture Fixture
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
		Fixture: NewFixture(db),
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

type GetAllRepository[T any] interface {
	GetAll() ([]T, error)
}

type getAllRepository[T any] struct {
	*repository
}

func (r getAllRepository[T]) GetAll() ([]T, error) {
	var dest []T
	if err := r.db.GetAll(&dest).Error; err != nil {
		return nil, err
	}
	return dest, nil
}