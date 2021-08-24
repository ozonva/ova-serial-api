// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repo/repo.go

// Package mock_repo is a generated GoMock package.
package mock_repo

import (
	model "ova-serial-api/internal/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepo is a mock of Repo interface.
type MockRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRepoMockRecorder
}

// MockRepoMockRecorder is the mock recorder for MockRepo.
type MockRepoMockRecorder struct {
	mock *MockRepo
}

// NewMockRepo creates a new mock instance.
func NewMockRepo(ctrl *gomock.Controller) *MockRepo {
	mock := &MockRepo{ctrl: ctrl}
	mock.recorder = &MockRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepo) EXPECT() *MockRepoMockRecorder {
	return m.recorder
}

// AddEntities mocks base method.
func (m *MockRepo) AddEntities(entities []model.Serial) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddEntities", entities)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEntities indicates an expected call of AddEntities.
func (mr *MockRepoMockRecorder) AddEntities(entities interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEntities", reflect.TypeOf((*MockRepo)(nil).AddEntities), entities)
}

// GetEntity mocks base method.
func (m *MockRepo) GetEntity(entityId int64) (model.Serial, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEntity", entityId)
	ret0, _ := ret[0].(model.Serial)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEntity indicates an expected call of GetEntity.
func (mr *MockRepoMockRecorder) GetEntity(entityId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEntity", reflect.TypeOf((*MockRepo)(nil).GetEntity), entityId)
}

// ListEntities mocks base method.
func (m *MockRepo) ListEntities(limit, offset uint64) ([]model.Serial, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEntities", limit, offset)
	ret0, _ := ret[0].([]model.Serial)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEntities indicates an expected call of ListEntities.
func (mr *MockRepoMockRecorder) ListEntities(limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEntities", reflect.TypeOf((*MockRepo)(nil).ListEntities), limit, offset)
}
