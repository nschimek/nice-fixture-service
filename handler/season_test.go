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

type seasonHandlerTestSuite struct {
	suite.Suite
	mockService *mocks.Season
	router *gin.Engine
	seasons []model.Season
}

func TestSeasonHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(seasonHandlerTestSuite))
}

func (s *seasonHandlerTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	s.mockService = &mocks.Season{}
	s.router = gin.Default()
	s.seasons = []model.Season{
		{Season: 2023, Current: true},
		{Season: 2022, Current: false},
		{Season: 2021, Current: false},
	}
	setupSeason(s.router.Group(core.ApiBasePath(1)), s.mockService)
}

func (s *seasonHandlerTestSuite) TestGetAll() {
	s.mockService.EXPECT().GetAll().Return(s.seasons, nil)

	rr := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/seasons", nil)
	s.router.ServeHTTP(rr, req)

	exp, _ := json.Marshal(s.seasons)

	s.Equal(http.StatusOK, rr.Code)
	s.Equal(exp, rr.Body.Bytes())
	s.mockService.AssertExpectations(s.T())
}

func (s *seasonHandlerTestSuite) TestGetAllError() {
	re := rest_error.NewWithDetail(rest_error.Internal, "test")
	s.mockService.EXPECT().GetAll().Return(nil, re)

	rr := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/seasons", nil)
	s.router.ServeHTTP(rr, req)

	exp, _ := json.Marshal(re)

	s.Equal(http.StatusInternalServerError, rr.Code)
	s.Equal(exp, rr.Body.Bytes())
	s.mockService.AssertExpectations(s.T())
}