package service

import (
	"errors"
	"testing"

	"github.com/nschimek/nice-fixture-service/model"
	"github.com/nschimek/nice-fixture-service/model/rest_error"
	"github.com/nschimek/nice-fixture-service/repository/mocks"
	"github.com/stretchr/testify/suite"
)

type leagueServiceTestSuite struct {
	suite.Suite
	mockRepository *mocks.League
	service League
	leagues []model.League
}

func TestLeagueServiceTestSuite(t *testing.T) {
	suite.Run(t, new(leagueServiceTestSuite))
}

func (s *leagueServiceTestSuite) SetupTest() {
	s.mockRepository = &mocks.League{}
	s.service = NewLeague(s.mockRepository)
	s.leagues = []model.League{
		{Id: 39, Name: "Premier League", Seasons: []model.LeagueSeason{{LeagueId: 39, Season: 2022, Current: true}}},
		{Id: 40, Name: "La Liga"},
		{Id: 41, Name: "Serie A", Seasons: []model.LeagueSeason{{LeagueId: 41, Season: 2022, Current: true}}},
	}
}

func (s *leagueServiceTestSuite) TestGetByParams() {
	p := model.LeagueParams{Season: 2022, Current: true}
	s.mockRepository.EXPECT().GetAllBySeason(&model.LeagueSeason{Season: p.Season, Current: p.Current}).Return(s.leagues, nil)

	res, err := s.service.GetByParams(p)

	s.Contains(res, s.leagues[0])
	s.NotContains(res, s.leagues[1]) // should not contain this element becuase it has no seasons
	s.Contains(res, s.leagues[2])
	s.Nil(err)
}

func (s *leagueServiceTestSuite) TestGetByParamsError() {
	te := rest_error.NewInternal(errors.New("test"))
	p := model.LeagueParams{}
	s.mockRepository.EXPECT().GetAllBySeason(&model.LeagueSeason{}).Return(nil, te)

	res, err := s.service.GetByParams(p)

	s.Nil(res)
	s.Equal(te, err)
}

func (s *leagueServiceTestSuite) TestGetById() {
	id := 39
	s.mockRepository.EXPECT().GetById(id).Return(&s.leagues[0], nil)

	res, err := s.service.GetById(id)

	s.Equal(&s.leagues[0], res)
	s.Nil(err)
}

func (s *leagueServiceTestSuite) TestGetByIdError() {
	te := rest_error.NewInternal(errors.New("test"))
	id := 39
	s.mockRepository.EXPECT().GetById(id).Return(nil, te)

	res, err := s.service.GetById(id)

	s.Nil(res)
	s.Equal(te, err)
}

