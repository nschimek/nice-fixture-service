// Code generated by mockery v2.28.0. DO NOT EDIT.

package mocks

import (
	model "github.com/nschimek/nice-fixture-service/model"
	mock "github.com/stretchr/testify/mock"
)

// Season is an autogenerated mock type for the Season type
type Season struct {
	mock.Mock
}

type Season_Expecter struct {
	mock *mock.Mock
}

func (_m *Season) EXPECT() *Season_Expecter {
	return &Season_Expecter{mock: &_m.Mock}
}

// GetAll provides a mock function with given fields:
func (_m *Season) GetAll() ([]model.Season, error) {
	ret := _m.Called()

	var r0 []model.Season
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]model.Season, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []model.Season); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Season)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Season_GetAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAll'
type Season_GetAll_Call struct {
	*mock.Call
}

// GetAll is a helper method to define mock.On call
func (_e *Season_Expecter) GetAll() *Season_GetAll_Call {
	return &Season_GetAll_Call{Call: _e.mock.On("GetAll")}
}

func (_c *Season_GetAll_Call) Run(run func()) *Season_GetAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Season_GetAll_Call) Return(_a0 []model.Season, _a1 error) *Season_GetAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Season_GetAll_Call) RunAndReturn(run func() ([]model.Season, error)) *Season_GetAll_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewSeason interface {
	mock.TestingT
	Cleanup(func())
}

// NewSeason creates a new instance of Season. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSeason(t mockConstructorTestingTNewSeason) *Season {
	mock := &Season{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
