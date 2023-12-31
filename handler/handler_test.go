package handler

import (
	"errors"
	"net/http"
	"testing"

	"github.com/go-playground/validator/v10"
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

	s.NotNil(res)
	s.Equal("/", res.RouterGroup.BasePath())
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

func (s *handlerTestSuite) TestBind() {
	bindFunc := func(obj any) error {
		return nil
	}
	m := &mocks.MockResponse{}

	// the params do not matter for this test because we are mocking the bind function
	res := bind(m.JSON, bindFunc, &model.LeagueParams{})

	s.True(res)
	s.Nil(m.Obj)
}

func (s *handlerTestSuite) TestBindValidationErrors() {
	bindFunc := func(obj any) error {
		return &validator.ValidationErrors{
			mocks.NewMockFieldError("gte=2008,lte=9999", "season", "2007"),
			mocks.NewMockFieldError("required", "current", ""),
		}
	}
	m := &mocks.MockResponse{}

	// the params do not matter for this test because we are mocking the bind function
	res := bind(m.JSON, bindFunc, &model.LeagueParams{})

	exp := rest_error.NewBadRequest([]string{
		"season [2007]: tag gte=2008,lte=9999 failed validation",
		"current []: tag required failed validation",
	}...)

	s.False(res)
	s.Equal(http.StatusBadRequest, m.Code)
	s.Equal(exp, m.Obj)
}

func (s *handlerTestSuite) TestBindOtherError() {
	err := errors.New("test")
	bindFunc := func(obj any) error {
		return err
	}
	m := &mocks.MockResponse{}

	// the params do not matter for this test because we are mocking the bind function
	res := bind(m.JSON, bindFunc, &model.LeagueParams{})

	s.False(res)
	s.Equal(http.StatusBadRequest, m.Code)
	s.Equal(rest_error.NewBadRequest(err.Error()), m.Obj)
}