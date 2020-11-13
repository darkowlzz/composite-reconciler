// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/darkowlzz/composite-reconciler/controller/v1 (interfaces: Controller)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	v1 "github.com/darkowlzz/composite-reconciler/event/v1"
	gomock "github.com/golang/mock/gomock"
	v10 "github.com/openshift/custom-resource-status/conditions/v1"
	v11 "k8s.io/apimachinery/pkg/apis/meta/v1"
	reflect "reflect"
	reconcile "sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// MockController is a mock of Controller interface
type MockController struct {
	ctrl     *gomock.Controller
	recorder *MockControllerMockRecorder
}

// MockControllerMockRecorder is the mock recorder for MockController
type MockControllerMockRecorder struct {
	mock *MockController
}

// NewMockController creates a new mock instance
func NewMockController(ctrl *gomock.Controller) *MockController {
	mock := &MockController{ctrl: ctrl}
	mock.recorder = &MockControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockController) EXPECT() *MockControllerMockRecorder {
	return m.recorder
}

// AddFinalizer mocks base method
func (m *MockController) AddFinalizer(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFinalizer", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddFinalizer indicates an expected call of AddFinalizer
func (mr *MockControllerMockRecorder) AddFinalizer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFinalizer", reflect.TypeOf((*MockController)(nil).AddFinalizer), arg0)
}

// Cleanup mocks base method
func (m *MockController) Cleanup() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cleanup")
	ret0, _ := ret[0].(error)
	return ret0
}

// Cleanup indicates an expected call of Cleanup
func (mr *MockControllerMockRecorder) Cleanup() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cleanup", reflect.TypeOf((*MockController)(nil).Cleanup))
}

// Default mocks base method
func (m *MockController) Default() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Default")
}

// Default indicates an expected call of Default
func (mr *MockControllerMockRecorder) Default() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Default", reflect.TypeOf((*MockController)(nil).Default))
}

// FetchInstance mocks base method
func (m *MockController) FetchInstance() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchInstance")
	ret0, _ := ret[0].(error)
	return ret0
}

// FetchInstance indicates an expected call of FetchInstance
func (mr *MockControllerMockRecorder) FetchInstance() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchInstance", reflect.TypeOf((*MockController)(nil).FetchInstance))
}

// FetchStatus mocks base method
func (m *MockController) FetchStatus() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchStatus")
	ret0, _ := ret[0].(error)
	return ret0
}

// FetchStatus indicates an expected call of FetchStatus
func (mr *MockControllerMockRecorder) FetchStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchStatus", reflect.TypeOf((*MockController)(nil).FetchStatus))
}

// GetObjectMetadata mocks base method
func (m *MockController) GetObjectMetadata() v11.ObjectMeta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetObjectMetadata")
	ret0, _ := ret[0].(v11.ObjectMeta)
	return ret0
}

// GetObjectMetadata indicates an expected call of GetObjectMetadata
func (mr *MockControllerMockRecorder) GetObjectMetadata() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObjectMetadata", reflect.TypeOf((*MockController)(nil).GetObjectMetadata))
}

// InitReconcile mocks base method
func (m *MockController) InitReconcile(arg0 context.Context, arg1 reconcile.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "InitReconcile", arg0, arg1)
}

// InitReconcile indicates an expected call of InitReconcile
func (mr *MockControllerMockRecorder) InitReconcile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitReconcile", reflect.TypeOf((*MockController)(nil).InitReconcile), arg0, arg1)
}

// Initialize mocks base method
func (m *MockController) Initialize(arg0 v10.Condition) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Initialize", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Initialize indicates an expected call of Initialize
func (mr *MockControllerMockRecorder) Initialize(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialize", reflect.TypeOf((*MockController)(nil).Initialize), arg0)
}

// IsUninitialized mocks base method
func (m *MockController) IsUninitialized() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUninitialized")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsUninitialized indicates an expected call of IsUninitialized
func (mr *MockControllerMockRecorder) IsUninitialized() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUninitialized", reflect.TypeOf((*MockController)(nil).IsUninitialized))
}

// Operate mocks base method
func (m *MockController) Operate() (reconcile.Result, v1.ReconcilerEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Operate")
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(v1.ReconcilerEvent)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Operate indicates an expected call of Operate
func (mr *MockControllerMockRecorder) Operate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Operate", reflect.TypeOf((*MockController)(nil).Operate))
}

// SaveClone mocks base method
func (m *MockController) SaveClone() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SaveClone")
}

// SaveClone indicates an expected call of SaveClone
func (mr *MockControllerMockRecorder) SaveClone() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveClone", reflect.TypeOf((*MockController)(nil).SaveClone))
}

// UpdateConditions mocks base method
func (m *MockController) UpdateConditions(arg0 []v10.Condition) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateConditions", arg0)
}

// UpdateConditions indicates an expected call of UpdateConditions
func (mr *MockControllerMockRecorder) UpdateConditions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateConditions", reflect.TypeOf((*MockController)(nil).UpdateConditions), arg0)
}

// UpdateStatus mocks base method
func (m *MockController) UpdateStatus() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus")
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStatus indicates an expected call of UpdateStatus
func (mr *MockControllerMockRecorder) UpdateStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockController)(nil).UpdateStatus))
}

// Validate mocks base method
func (m *MockController) Validate() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate")
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate
func (mr *MockControllerMockRecorder) Validate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockController)(nil).Validate))
}
