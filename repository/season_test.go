package repository

import (
	"errors"
	"testing"

	"github.com/nschimek/nice-fixture-service/core"
	"github.com/nschimek/nice-fixture-service/core/mocks"
	"github.com/nschimek/nice-fixture-service/model"
	"github.com/stretchr/testify/suite"
)

type seasonTestSuite struct {
	suite.Suite
	mockDatabase *mocks.Database
	repo Season
}

func TestSeasonRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(seasonTestSuite))
}

func (s *seasonTestSuite) SetupTest() {
	s.mockDatabase = &mocks.Database{}
	s.repo = NewSeason(s.mockDatabase)
}

func (s *seasonTestSuite) TestGetAll() {
	var entities []model.Season

	s.mockDatabase.EXPECT().GroupBy(&entities, 
		&model.LeagueSeason{}, seasonGroupBySelectSql, "season").Return(core.DatabaseResult{RowsAffected: 3})

	res, err := s.repo.GetAll()

	s.mockDatabase.AssertExpectations(s.T())
	s.Equal(entities, res)
	s.Nil(err)
}

func (s *seasonTestSuite) TestGetAllError() {
	var entities []model.Season

	s.mockDatabase.EXPECT().GroupBy(&entities, 
		&model.LeagueSeason{}, seasonGroupBySelectSql, "season").Return(core.DatabaseResult{Error: errors.New("test")})

	res, err := s.repo.GetAll()

	s.mockDatabase.AssertExpectations(s.T())
	s.Nil(res)
	s.ErrorContains(err, "test")
}