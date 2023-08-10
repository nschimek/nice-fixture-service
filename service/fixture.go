package service

import (
	"github.com/nschimek/nice-fixture-service/model"
	"github.com/nschimek/nice-fixture-service/model/rest_error"
	"github.com/nschimek/nice-fixture-service/repository"
)

//go:generate mockery --name League --filename league_mock.go
type Fixture interface {
	GetAll() ([]model.Fixture, *rest_error.Error)
}

type fixture struct {
	repo repository.Fixture
}

func NewFixture(repo repository.Fixture) *fixture {
	return &fixture{repo: repo}
}

func (s *fixture) GetAll() ([]model.Fixture, *rest_error.Error) {
	r, err := s.repo.GetAll()
	return r, rest_error.NewInternal(err)
}