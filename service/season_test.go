package service

import (
	"errors"
	"testing"

	"github.com/nschimek/nice-fixture-service/model"
	"github.com/nschimek/nice-fixture-service/model/rest_error"
	"github.com/nschimek/nice-fixture-service/repository/mocks"
	"github.com/stretchr/testify/suite"
)

type seasonServiceTestSuite struct {
	suite.Suite
	mockRepository *mocks.Season
	service Season
	seasons []model.Season
}

func TestSeasonServiceTestSuite(t *testing.T) {
	suite.Run(t, new(seasonServiceTestSuite))
}

func (s *seasonServiceTestSuite) SetupTest() {
	s.mockRepository = &mocks.Season{}
	s.service = NewSeason(s.mockRepository)
	s.seasons = []model.Season{
		{Season: 2023, Current: true},
		{Season: 2022, Current: false},
		{Season: 2021, Current: false},
	}
}

func (s *seasonServiceTestSuite) TestGetAll() {
	s.mockRepository.EXPECT().GetAll().Return(s.seasons, nil)

	res, err := s.service.GetAll()

	s.ElementsMatch(s.seasons, res)
	s.Nil(err)
}

func (s *seasonServiceTestSuite) TestGetAllError() {
	s.mockRepository.EXPECT().GetAll().Return(nil, errors.New("test"))
	
	res, err := s.service.GetAll()

	s.Nil(res)
	s.Equal(rest_error.NewInternal(errors.New("test")), err)
}