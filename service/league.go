package service

import (
	"github.com/nschimek/nice-fixture-service/model"
	"github.com/nschimek/nice-fixture-service/repository"
)

type League interface {}

type league struct {
	repo repository.League
}

func NewLeague(repo repository.League) *league {
	return &league{repo: repo}
}

func (s *league) GetAll() ([]model.League, error) {
	return s.repo.GetAll()
}

func (s *league) GetAllBySeason(season int) ([]model.League, error) {
	return s.repo.GetAllBySeason(season)
}