package service

import (
	"github.com/nschimek/nice-fixture-service/model"
	"github.com/nschimek/nice-fixture-service/model/rest_error"
	"github.com/nschimek/nice-fixture-service/repository"
)

type Season interface {
	GetAll() ([]model.Season, *rest_error.Error)
}

type season struct {
	repo repository.Season
}

func NewSeason(repo repository.Season) *season {
	return &season{repo: repo}
}

func (s *season) GetAll() ([]model.Season, *rest_error.Error) {
	r, err := s.repo.GetAll()
	return r, rest_error.NewInternal(err)
}
