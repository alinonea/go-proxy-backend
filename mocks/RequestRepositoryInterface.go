// Code generated by mockery v2.34.2. DO NOT EDIT.

package mocks

import (
	domain "github.com/alinonea/main/domain"
	mock "github.com/stretchr/testify/mock"
)

// RequestRepositoryInterface is an autogenerated mock type for the RequestRepositoryInterface type
type RequestRepositoryInterface struct {
	mock.Mock
}

// SaveRequest provides a mock function with given fields: request
func (_m *RequestRepositoryInterface) SaveRequest(request domain.Request) error {
	ret := _m.Called(request)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.Request) error); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewRequestRepositoryInterface creates a new instance of RequestRepositoryInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRequestRepositoryInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *RequestRepositoryInterface {
	mock := &RequestRepositoryInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}