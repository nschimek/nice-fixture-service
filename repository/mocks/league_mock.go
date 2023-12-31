// Code generated by mockery v2.28.0. DO NOT EDIT.

package mocks

import (
	model "github.com/nschimek/nice-fixture-service/model"
	mock "github.com/stretchr/testify/mock"
)

// League is an autogenerated mock type for the League type
type League struct {
	mock.Mock
}

type League_Expecter struct {
	mock *mock.Mock
}

func (_m *League) EXPECT() *League_Expecter {
	return &League_Expecter{mock: &_m.Mock}
}

// GetAll provides a mock function with given fields:
func (_m *League) GetAll() ([]model.League, error) {
	ret := _m.Called()

	var r0 []model.League
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]model.League, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []model.League); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.League)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// League_GetAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAll'
type League_GetAll_Call struct {
	*mock.Call
}

// GetAll is a helper method to define mock.On call
func (_e *League_Expecter) GetAll() *League_GetAll_Call {
	return &League_GetAll_Call{Call: _e.mock.On("GetAll")}
}

func (_c *League_GetAll_Call) Run(run func()) *League_GetAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *League_GetAll_Call) Return(_a0 []model.League, _a1 error) *League_GetAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *League_GetAll_Call) RunAndReturn(run func() ([]model.League, error)) *League_GetAll_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllBySeason provides a mock function with given fields: season
func (_m *League) GetAllBySeason(season *model.LeagueSeason) ([]model.League, error) {
	ret := _m.Called(season)

	var r0 []model.League
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.LeagueSeason) ([]model.League, error)); ok {
		return rf(season)
	}
	if rf, ok := ret.Get(0).(func(*model.LeagueSeason) []model.League); ok {
		r0 = rf(season)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.League)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.LeagueSeason) error); ok {
		r1 = rf(season)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// League_GetAllBySeason_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllBySeason'
type League_GetAllBySeason_Call struct {
	*mock.Call
}

// GetAllBySeason is a helper method to define mock.On call
//   - season *model.LeagueSeason
func (_e *League_Expecter) GetAllBySeason(season interface{}) *League_GetAllBySeason_Call {
	return &League_GetAllBySeason_Call{Call: _e.mock.On("GetAllBySeason", season)}
}

func (_c *League_GetAllBySeason_Call) Run(run func(season *model.LeagueSeason)) *League_GetAllBySeason_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.LeagueSeason))
	})
	return _c
}

func (_c *League_GetAllBySeason_Call) Return(_a0 []model.League, _a1 error) *League_GetAllBySeason_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *League_GetAllBySeason_Call) RunAndReturn(run func(*model.LeagueSeason) ([]model.League, error)) *League_GetAllBySeason_Call {
	_c.Call.Return(run)
	return _c
}

// GetById provides a mock function with given fields: id
func (_m *League) GetById(id int) (*model.League, error) {
	ret := _m.Called(id)

	var r0 *model.League
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*model.League, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) *model.League); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.League)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// League_GetById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetById'
type League_GetById_Call struct {
	*mock.Call
}

// GetById is a helper method to define mock.On call
//   - id int
func (_e *League_Expecter) GetById(id interface{}) *League_GetById_Call {
	return &League_GetById_Call{Call: _e.mock.On("GetById", id)}
}

func (_c *League_GetById_Call) Run(run func(id int)) *League_GetById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *League_GetById_Call) Return(_a0 *model.League, _a1 error) *League_GetById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *League_GetById_Call) RunAndReturn(run func(int) (*model.League, error)) *League_GetById_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewLeague interface {
	mock.TestingT
	Cleanup(func())
}

// NewLeague creates a new instance of League. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLeague(t mockConstructorTestingTNewLeague) *League {
	mock := &League{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
