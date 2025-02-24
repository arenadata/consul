// Code generated by mockery v2.15.0. DO NOT EDIT.

package catalog

import (
	proxycfg "github.com/shulutkov/yellow-pages/agent/proxycfg"
	mock "github.com/stretchr/testify/mock"

	structs "github.com/shulutkov/yellow-pages/agent/structs"
)

// MockConfigManager is an autogenerated mock type for the ConfigManager type
type MockConfigManager struct {
	mock.Mock
}

// Deregister provides a mock function with given fields: proxyID, source
func (_m *MockConfigManager) Deregister(proxyID proxycfg.ProxyID, source proxycfg.ProxySource) {
	_m.Called(proxyID, source)
}

// Register provides a mock function with given fields: proxyID, service, source, token, overwrite
func (_m *MockConfigManager) Register(proxyID proxycfg.ProxyID, service *structs.NodeService, source proxycfg.ProxySource, token string, overwrite bool) error {
	ret := _m.Called(proxyID, service, source, token, overwrite)

	var r0 error
	if rf, ok := ret.Get(0).(func(proxycfg.ProxyID, *structs.NodeService, proxycfg.ProxySource, string, bool) error); ok {
		r0 = rf(proxyID, service, source, token, overwrite)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Watch provides a mock function with given fields: req
func (_m *MockConfigManager) Watch(req proxycfg.ProxyID) (<-chan *proxycfg.ConfigSnapshot, proxycfg.CancelFunc) {
	ret := _m.Called(req)

	var r0 <-chan *proxycfg.ConfigSnapshot
	if rf, ok := ret.Get(0).(func(proxycfg.ProxyID) <-chan *proxycfg.ConfigSnapshot); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan *proxycfg.ConfigSnapshot)
		}
	}

	var r1 proxycfg.CancelFunc
	if rf, ok := ret.Get(1).(func(proxycfg.ProxyID) proxycfg.CancelFunc); ok {
		r1 = rf(req)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(proxycfg.CancelFunc)
		}
	}

	return r0, r1
}

type mockConstructorTestingTNewMockConfigManager interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockConfigManager creates a new instance of MockConfigManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockConfigManager(t mockConstructorTestingTNewMockConfigManager) *MockConfigManager {
	mock := &MockConfigManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
