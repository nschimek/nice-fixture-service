package service

import (
	"github.com/nschimek/nice-fixture-service/model"
	"github.com/nschimek/nice-fixture-service/repository"
)

type Season interface {
	GetAll() ([]model.Season, error)
}

type season struct {
	repo repository.Season
}

func NewSeason(repo repository.Season) *season {
	return &season{repo: repo}
}

func (s *season) GetAll() ([]model.Season, error) {
	return s.repo.GetAll()
}
