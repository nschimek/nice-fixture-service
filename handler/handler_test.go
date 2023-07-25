package handler

import (
	"net/http"
	"testing"

	"github.com/nschimek/nice-fixture-service/handler/mocks"
	"github.com/nschimek/nice-fixture-service/model"
	"github.com/nschimek/nice-fixture-service/model/rest_error"
	"github.com/nschimek/nice-fixture-service/service"
	svc_mocks "github.com/nschimek/nice-fixture-service/service/mocks"
	"github.com/stretchr/testify/suite"
)

type handlerTestSuite struct {
	suite.Suite
	result *model.Season
}

func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(handlerTestSuite))
}

func (s *handlerTestSuite) SetupTest() {
	s.result = &model.Season{Season: 2022}
}

func (s *handlerTestSuite) TestSetupRouter() {
	res := CreateRouter(&service.ServiceRegistry{League: &svc_mocks.League{}, Season: &svc_mocks.Season{}})

	s.Equal("/", res.RouterGroup.BasePath())
	s.NotNil(res)
}

func (s *handlerTestSuite) TestJsonResult() {
	m := &mocks.MockResponse{}
	jsonResult[model.Season](m.JSON, s.result, nil)

	s.Equal(http.StatusOK, m.Code)
	s.Equal(s.result, m.Obj)
}

func (s *handlerTestSuite) TestJsonResultError() {
	m := &mocks.MockResponse{}
	err := rest_error.New(rest_error.BadRequest)
	jsonResult[model.Season](m.JSON, s.result, err)

	s.Equal(err.Code, m.Code)
	s.Equal(err, m.Obj)
}

func (s *handlerTestSuite) TestJsonResultNil() {
	m := &mocks.MockResponse{}
	jsonResult[model.Season](m.JSON, nil, nil)
	exp := rest_error.NewNotFound()

	s.Equal(exp.Code, m.Code)
	s.Equal(exp, m.Obj)
}