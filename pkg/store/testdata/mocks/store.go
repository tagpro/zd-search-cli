// Code generated by MockGen. DO NOT EDIT.
// Source: store.go

// Package store is a generated GoMock package.
package store

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	store "github.com/tagpro/zd-search-cli/pkg/store"
	organistations "github.com/tagpro/zd-search-cli/pkg/store/organistations"
	tickets "github.com/tagpro/zd-search-cli/pkg/store/tickets"
	users "github.com/tagpro/zd-search-cli/pkg/store/users"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// GetKeys mocks base method.
func (m *MockStore) GetKeys() store.Keys {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKeys")
	ret0, _ := ret[0].(store.Keys)
	return ret0
}

// GetKeys indicates an expected call of GetKeys.
func (mr *MockStoreMockRecorder) GetKeys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKeys", reflect.TypeOf((*MockStore)(nil).GetKeys))
}

// GetOrganisations mocks base method.
func (m *MockStore) GetOrganisations(key, input string) (organistations.Organisations, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrganisations", key, input)
	ret0, _ := ret[0].(organistations.Organisations)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrganisations indicates an expected call of GetOrganisations.
func (mr *MockStoreMockRecorder) GetOrganisations(key, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrganisations", reflect.TypeOf((*MockStore)(nil).GetOrganisations), key, input)
}

// GetTickets mocks base method.
func (m *MockStore) GetTickets(key, input string) (tickets.Tickets, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTickets", key, input)
	ret0, _ := ret[0].(tickets.Tickets)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTickets indicates an expected call of GetTickets.
func (mr *MockStoreMockRecorder) GetTickets(key, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTickets", reflect.TypeOf((*MockStore)(nil).GetTickets), key, input)
}

// GetUsers mocks base method.
func (m *MockStore) GetUsers(key, input string) (users.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers", key, input)
	ret0, _ := ret[0].(users.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockStoreMockRecorder) GetUsers(key, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockStore)(nil).GetUsers), key, input)
}
