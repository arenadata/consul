// Code generated by mockery v2.15.0. DO NOT EDIT.

package autopilotevents

import (
	stream "github.com/arenadata/consul/agent/consul/stream"
	mock "github.com/stretchr/testify/mock"
)

// MockPublisher is an autogenerated mock type for the Publisher type
type MockPublisher struct {
	mock.Mock
}

// Publish provides a mock function with given fields: _a0
func (_m *MockPublisher) Publish(_a0 []stream.Event) {
	_m.Called(_a0)
}

type mockConstructorTestingTNewMockPublisher interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockPublisher creates a new instance of MockPublisher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockPublisher(t mockConstructorTestingTNewMockPublisher) *MockPublisher {
	mock := &MockPublisher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
