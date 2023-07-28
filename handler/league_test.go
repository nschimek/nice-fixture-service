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

type leagueHandlerTestSuite struct {
	suite.Suite
	mockService *mocks.League
	router *gin.Engine
	leagues []model.League
}

func TestLeagueHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(leagueHandlerTestSuite))
}

func (s *leagueHandlerTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	s.mockService = &mocks.League{}
	s.router = gin.Default()
	s.leagues = []model.League{
		{Id: 39, Name: "Premier League", Seasons: []model.LeagueSeason{{LeagueId: 39, Season: 2022, Current: true}}},
		{Id: 152, Name: "La Liga", Seasons: []model.LeagueSeason{{LeagueId: 39, Season: 2022, Current: true}}},
		{Id: 210, Name: "Serie A", Seasons: []model.LeagueSeason{{LeagueId: 39, Season: 2022, Current: true}}},
	}
	setupLeague(s.router.Group(core.ApiBasePath(1)), s.mockService)
}

func (s *leagueHandlerTestSuite) TestGetByParams() {
	s.mockService.EXPECT().GetByParams(model.LeagueParams{Season: 2022}).Return(s.leagues, nil)

	rr := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/leagues?season=2022", nil)
	s.router.ServeHTTP(rr, req)

	exp, _ := json.Marshal(s.leagues)

	s.Equal(http.StatusOK, rr.Code)
	s.Equal(exp, rr.Body.Bytes())
	s.mockService.AssertExpectations(s.T())
}

func (s *leagueHandlerTestSuite) TestGetByParamsBindingError() {
	rr := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/leagues?season=asdf", nil)
	s.router.ServeHTTP(rr, req)

	s.Equal(http.StatusBadRequest, rr.Code)
	s.mockService.AssertExpectations(s.T())
}

func (s *leagueHandlerTestSuite) TestGetByParamsServiceError() {
	re := rest_error.NewWithDetail(rest_error.Internal, "test")
	s.mockService.EXPECT().GetByParams(model.LeagueParams{}).Return(nil, re)

	rr := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/leagues", nil)
	s.router.ServeHTTP(rr, req)

	exp, _ := json.Marshal(re)

	s.Equal(http.StatusInternalServerError, rr.Code)
	s.Equal(exp, rr.Body.Bytes())
	s.mockService.AssertExpectations(s.T())
}

func (s *leagueHandlerTestSuite) TestGetById() {
	s.mockService.EXPECT().GetById(39).Return(&s.leagues[0], nil)

	rr := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/leagues/39", nil)
	s.router.ServeHTTP(rr, req)

	exp, _ := json.Marshal(s.leagues[0])

	s.Equal(http.StatusOK, rr.Code)
	s.Equal(exp, rr.Body.Bytes())
	s.mockService.AssertExpectations(s.T())
}

func (s *leagueHandlerTestSuite) TestGetByIdBindingError() {
	rr := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/leagues/asdf", nil)
	s.router.ServeHTTP(rr, req)

	s.Equal(http.StatusBadRequest, rr.Code)
	s.mockService.AssertExpectations(s.T())
}

func (s *leagueHandlerTestSuite) TestGetByIdNotFound() {
	s.mockService.EXPECT().GetById(39).Return(nil, nil)
	rr := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/leagues/39", nil)
	s.router.ServeHTTP(rr, req)

	exp, _ := json.Marshal(rest_error.New(rest_error.NotFound))

	s.Equal(http.StatusNotFound, rr.Code)
	s.Equal(exp, rr.Body.Bytes())
	s.mockService.AssertExpectations(s.T())
}



