package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/nschimek/nice-fixture-service/core"
	"github.com/nschimek/nice-fixture-service/model"
	"github.com/nschimek/nice-fixture-service/model/rest_error"
	"github.com/nschimek/nice-fixture-service/service/mocks"
	"github.com/stretchr/testify/suite"
)

type teamHandlerTestSuite struct {
	suite.Suite
	mockService *mocks.Team
	router *gin.Engine
	teams []model.Team
}

func TestTeamHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(teamHandlerTestSuite))
}

func (s *teamHandlerTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	s.mockService = &mocks.Team{}
	s.router = gin.Default()
	s.teams = []model.Team{
		{Id: 1, TeamLeagueSeasons: []model.TeamLeagueSeason{{LeagueId: 1, Season: 2022}}},
		{Id: 2, TeamLeagueSeasons: []model.TeamLeagueSeason{{LeagueId: 1, Season: 2022}}},
		{Id: 3, TeamLeagueSeasons: []model.TeamLeagueSeason{{LeagueId: 2, Season: 2022}}},
	}
	setupTeam(s.router.Group(core.ApiBasePath(1)), s.mockService)
}

func (s *teamHandlerTestSuite) TestGetByParams() {
	s.mockService.EXPECT().GetByParams(model.TeamParams{League: 1, Season: 2022}).Return(s.teams[0:2], nil)
	
	rr := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/teams?league=1&season=2022", nil)
	s.router.ServeHTTP(rr, req)

	exp, _ := json.Marshal(s.teams[0:2])

	s.Equal(http.StatusOK, rr.Code)
	s.Equal(exp, rr.Body.Bytes())
	s.mockService.AssertExpectations(s.T())
}

func (s *teamHandlerTestSuite) TestGetByParamsBindingError() {
	rr := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/teams?league=-1&season=asdf", nil)
	s.router.ServeHTTP(rr, req)

	s.Equal(http.StatusBadRequest, rr.Code)
	s.mockService.AssertExpectations(s.T())
}

func (s *teamHandlerTestSuite) TestGetByParamsServiceError() {
	re := rest_error.NewWithDetail(rest_error.Internal, "test")
	s.mockService.EXPECT().GetByParams(model.TeamParams{}).Return(nil, re)

	rr := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/teams", nil)
	s.router.ServeHTTP(rr, req)

	exp, _ := json.Marshal(re)

	s.Equal(http.StatusInternalServerError, rr.Code)
	s.Equal(exp, rr.Body.Bytes())
	s.mockService.AssertExpectations(s.T())
}

func (s *teamHandlerTestSuite) TestGetById() {
	s.mockService.EXPECT().GetById(1).Return(&s.teams[0], nil)

	rr := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/teams/1", nil)
	s.router.ServeHTTP(rr, req)

	exp, _ := json.Marshal(s.teams[0])

	s.Equal(http.StatusOK, rr.Code)
	s.Equal(exp, rr.Body.Bytes())
	s.mockService.AssertExpectations(s.T())
}

func (s *teamHandlerTestSuite) TestGetByIdBindingError() {
	rr := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/teams/asdf", nil)
	s.router.ServeHTTP(rr, req)

	s.Equal(http.StatusBadRequest, rr.Code)
	s.mockService.AssertExpectations(s.T())
}

func (s *teamHandlerTestSuite) TestGetByIdNotFound() {
	s.mockService.EXPECT().GetById(99).Return(nil, nil)
	rr := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/teams/99", nil)
	s.router.ServeHTTP(rr, req)

	exp, _ := json.Marshal(rest_error.NewNotFound())

	s.Equal(http.StatusNotFound, rr.Code)
	s.Equal(exp, rr.Body.Bytes())
	s.mockService.AssertExpectations(s.T())
}