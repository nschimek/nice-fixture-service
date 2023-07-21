package repository

import (
	"errors"
	"testing"

	"github.com/nschimek/nice-fixture-service/core"
	"github.com/nschimek/nice-fixture-service/core/mocks"
	"github.com/nschimek/nice-fixture-service/model"
	"github.com/stretchr/testify/suite"
)

type leagueTestSuite struct {
	suite.Suite
	mockDatabase *mocks.Database
	repo League
}

func TestLeagueRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(leagueTestSuite))
}

func (s *leagueTestSuite) SetupTest() {
	s.mockDatabase = &mocks.Database{}
	s.repo = NewLeague(s.mockDatabase)
}

func (s *leagueTestSuite) TestGetByIdFound() {
	var entity model.League
	id := 1

	s.mockDatabase.EXPECT().GetById(id, &entity).Return(core.DatabaseResult{RowsAffected: 1})

	res, err := s.repo.GetById(id)

	s.mockDatabase.AssertExpectations(s.T())
	s.Equal(&entity, res)
	s.Nil(err)
}

func (s *leagueTestSuite) TestGetByIdError() {
	var entity model.League
	id := 1

	s.mockDatabase.EXPECT().GetById(id, &entity).Return(core.DatabaseResult{Error: errors.New("test")})

	res, err := s.repo.GetById(id)

	s.mockDatabase.AssertExpectations(s.T())
	s.ErrorContains(err, "test")
	s.Nil(res)
}

func (s *leagueTestSuite) TestGetByIdNotFound() {
	var entity model.League
	id := 1

	s.mockDatabase.EXPECT().GetById(id, &entity).Return(core.DatabaseResult{})

	res, err := s.repo.GetById(id)

	s.mockDatabase.AssertExpectations(s.T())
	s.Nil(res)
	s.Nil(err)
}

func (s *leagueTestSuite) TestGetAllBySeason() {
	var entities []model.League
	season := &model.LeagueSeason{Season: 2022}

	s.mockDatabase.EXPECT().Preload(&entities, nil, "Seasons", season).Return(core.DatabaseResult{RowsAffected: 2})

	res, err := s.repo.GetAllBySeason(season)

	s.mockDatabase.AssertExpectations(s.T())
	s.Equal(entities, res)
	s.Nil(err)
}