package repository

import (
	"errors"
	"testing"

	"github.com/nschimek/nice-fixture-service/core"
	"github.com/nschimek/nice-fixture-service/core/mocks"
	"github.com/nschimek/nice-fixture-service/model"
	"github.com/stretchr/testify/suite"
)

type teamTestSuite struct {
	suite.Suite
	mockDatabase *mocks.Database
	repo Team
}

func TestTeamRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(teamTestSuite))
}

func (s *teamTestSuite) SetupTest() {
	s.mockDatabase = &mocks.Database{}
	s.repo = NewTeam(s.mockDatabase)
}

func (s *teamTestSuite) TestGetById() {
	var entity model.Team
	id := 1
	
	s.mockDatabase.EXPECT().GetById(id, &entity).Return(core.DatabaseResult{RowsAffected: 1})
	
	res, err := s.repo.GetById(id)

	s.mockDatabase.AssertExpectations(s.T())
	s.Equal(&entity, res)
	s.Nil(err)
}

func (s *teamTestSuite) TestGetByIdError() {
	var entity model.Team
	id := 1

	s.mockDatabase.EXPECT().GetById(id, &entity).Return(core.DatabaseResult{Error: errors.New("test")})

	res, err := s.repo.GetById(id)

	s.mockDatabase.AssertExpectations(s.T())
	s.Nil(res)
	s.ErrorContains(err, "test")
}

func (s *teamTestSuite) TestGetAllByTLS() {
	var entities []model.Team
	tls := &model.TeamLeagueSeason{Season: 1, LeagueId: 1}

	s.mockDatabase.EXPECT().Preload(&entities, nil, teamLeagueSeasons, tls).Return(core.DatabaseResult{RowsAffected: 3})
	
	res, err := s.repo.GetAllByLeagueSeason(tls)
	
	s.mockDatabase.AssertExpectations(s.T())
	s.Equal(entities, res)
	s.Nil(err)
}

func (s *teamTestSuite) TestGetAllByTLSError() {
	var entities []model.Team
	tls := &model.TeamLeagueSeason{Season: 1, LeagueId: 1}

	s.mockDatabase.EXPECT().Preload(&entities, nil, teamLeagueSeasons, tls).Return(core.DatabaseResult{Error: errors.New("test")})

	res, err := s.repo.GetAllByLeagueSeason(tls)

	s.mockDatabase.AssertExpectations(s.T())
	s.Nil(res)
	s.ErrorContains(err, "test")
}




