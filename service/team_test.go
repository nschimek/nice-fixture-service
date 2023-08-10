package service

import (
	"errors"
	"testing"

	"github.com/nschimek/nice-fixture-service/model"
	"github.com/nschimek/nice-fixture-service/model/rest_error"
	"github.com/nschimek/nice-fixture-service/repository/mocks"
	"github.com/stretchr/testify/suite"
)

type teamTestSuite struct {
	suite.Suite
	mockRepository *mocks.Team
	service        Team
	teams          []model.Team
}

func TestTeamRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(teamTestSuite))
}

func (s *teamTestSuite) SetupTest() {
	s.mockRepository = &mocks.Team{}
	s.service = NewTeam(s.mockRepository)
	s.teams = []model.Team{
		{Id: 1, Name: "Team 1", LeagueSeason: &model.TeamLeagueSeason{LeagueId: 1, Season: 2022}},
		{Id: 2, Name: "Team 2", LeagueSeason: &model.TeamLeagueSeason{LeagueId: 2, Season: 2022}},
		{Id: 3, Name: "Team 3", LeagueSeason: &model.TeamLeagueSeason{LeagueId: 3, Season: 2022}},
	}
}

func (s *teamTestSuite) TestGetByParams() {
	p := model.TeamParams{Season: 2002}
	s.mockRepository.EXPECT().GetAllByLeagueSeason(&model.TeamLeagueSeason{Season: p.Season}).Return(s.teams, nil)

	res, err := s.service.GetByParams(p)

	s.ElementsMatch(res, s.teams)
	s.Nil(err)
}

func (s *teamTestSuite) TestGetByParamsError() {
	te := rest_error.NewInternal(errors.New("test"))
	p := model.TeamParams{Season: 2022}
	s.mockRepository.EXPECT().GetAllByLeagueSeason(&model.TeamLeagueSeason{Season: p.Season}).Return(nil, te)

	res, err := s.service.GetByParams(p)

	s.Nil(res)
	s.Equal(te, err)
}

func (s *teamTestSuite) TestGetByParamsAll() {
	p := model.TeamParams{}
	s.mockRepository.EXPECT().GetAll().Return(s.teams, nil)

	res, err := s.service.GetByParams(p)

	s.ElementsMatch(res, s.teams)
	s.Nil(err)
}

func (s *teamTestSuite) TestGetByParamsAllError() {
	te := rest_error.NewInternal(errors.New("test"))
	p := model.TeamParams{}

	s.mockRepository.EXPECT().GetAll().Return(nil, te)

	res, err := s.service.GetByParams(p)

	s.Nil(res)
	s.Equal(te, err)
}

func (s *teamTestSuite) TestGetById() {
	id := 1
	s.mockRepository.EXPECT().GetById(id).Return(&s.teams[0], nil)

	res, err := s.service.GetById(id)

	s.Equal(&s.teams[0], res)
	s.Nil(err)
}

func (s *teamTestSuite) TestGetByIdError() {
	te := rest_error.NewInternal(errors.New("test"))
	id := 1
	s.mockRepository.EXPECT().GetById(id).Return(nil, te)	

	res, err := s.service.GetById(id)

	s.Nil(res)	
	s.Equal(te, err)
}