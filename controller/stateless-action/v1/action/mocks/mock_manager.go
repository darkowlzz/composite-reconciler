// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/darkowlzz/operator-toolkit/controller/stateless-action/v1/action (interfaces: Manager)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockManager is a mock of Manager interface.
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager.
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// Check mocks base method.
func (m *MockManager) Check(arg0 context.Context, arg1 interface{}) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Check", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Check indicates an expected call of Check.
func (mr *MockManagerMockRecorder) Check(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Check", reflect.TypeOf((*MockManager)(nil).Check), arg0, arg1)
}

// Defer mocks base method.
func (m *MockManager) Defer(arg0 context.Context, arg1 interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Defer", arg0, arg1)
}

// Defer indicates an expected call of Defer.
func (mr *MockManagerMockRecorder) Defer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Defer", reflect.TypeOf((*MockManager)(nil).Defer), arg0, arg1)
}

// GetObjects mocks base method.
func (m *MockManager) GetObjects() ([]interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetObjects")
	ret0, _ := ret[0].([]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObjects indicates an expected call of GetObjects.
func (mr *MockManagerMockRecorder) GetObjects() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObjects", reflect.TypeOf((*MockManager)(nil).GetObjects))
}

// Run mocks base method.
func (m *MockManager) Run(arg0 context.Context, arg1 interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Run", arg0, arg1)
}

// Run indicates an expected call of Run.
func (mr *MockManagerMockRecorder) Run(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockManager)(nil).Run), arg0, arg1)
}