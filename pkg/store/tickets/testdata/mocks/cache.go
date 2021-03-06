// Code generated by MockGen. DO NOT EDIT.
// Source: cache.go

// Package ticketcache is a generated GoMock package.
package ticketcache

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	tickets "github.com/tagpro/zd-search-cli/pkg/store/tickets"
)

// MockCache is a mock of Cache interface.
type MockCache struct {
	ctrl     *gomock.Controller
	recorder *MockCacheMockRecorder
}

// MockCacheMockRecorder is the mock recorder for MockCache.
type MockCacheMockRecorder struct {
	mock *MockCache
}

// NewMockCache creates a new mock instance.
func NewMockCache(ctrl *gomock.Controller) *MockCache {
	mock := &MockCache{ctrl: ctrl}
	mock.recorder = &MockCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCache) EXPECT() *MockCacheMockRecorder {
	return m.recorder
}

// GetTickets mocks base method.
func (m *MockCache) GetTickets(key, input string) (tickets.Tickets, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTickets", key, input)
	ret0, _ := ret[0].(tickets.Tickets)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTickets indicates an expected call of GetTickets.
func (mr *MockCacheMockRecorder) GetTickets(key, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTickets", reflect.TypeOf((*MockCache)(nil).GetTickets), key, input)
}

// Optimise mocks base method.
func (m *MockCache) Optimise() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Optimise")
	ret0, _ := ret[0].(error)
	return ret0
}

// Optimise indicates an expected call of Optimise.
func (mr *MockCacheMockRecorder) Optimise() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Optimise", reflect.TypeOf((*MockCache)(nil).Optimise))
}
