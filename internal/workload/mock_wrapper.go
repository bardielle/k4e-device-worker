// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/jakub-dzon/k4e-device-worker/internal/workload (interfaces: WorkloadWrapper)

// Package workload is a generated GoMock package.
package workload

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	api "github.com/jakub-dzon/k4e-device-worker/internal/workload/api"
	v1 "k8s.io/api/core/v1"
)

// MockWorkloadWrapper is a mock of WorkloadWrapper interface.
type MockWorkloadWrapper struct {
	ctrl     *gomock.Controller
	recorder *MockWorkloadWrapperMockRecorder
}

// MockWorkloadWrapperMockRecorder is the mock recorder for MockWorkloadWrapper.
type MockWorkloadWrapperMockRecorder struct {
	mock *MockWorkloadWrapper
}

// NewMockWorkloadWrapper creates a new mock instance.
func NewMockWorkloadWrapper(ctrl *gomock.Controller) *MockWorkloadWrapper {
	mock := &MockWorkloadWrapper{ctrl: ctrl}
	mock.recorder = &MockWorkloadWrapperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWorkloadWrapper) EXPECT() *MockWorkloadWrapperMockRecorder {
	return m.recorder
}

// Init mocks base method.
func (m *MockWorkloadWrapper) Init() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init")
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockWorkloadWrapperMockRecorder) Init() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockWorkloadWrapper)(nil).Init))
}

// List mocks base method.
func (m *MockWorkloadWrapper) List() ([]api.WorkloadInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]api.WorkloadInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockWorkloadWrapperMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockWorkloadWrapper)(nil).List))
}

// PersistConfiguration mocks base method.
func (m *MockWorkloadWrapper) PersistConfiguration() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PersistConfiguration")
	ret0, _ := ret[0].(error)
	return ret0
}

// PersistConfiguration indicates an expected call of PersistConfiguration.
func (mr *MockWorkloadWrapperMockRecorder) PersistConfiguration() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PersistConfiguration", reflect.TypeOf((*MockWorkloadWrapper)(nil).PersistConfiguration))
}

// RemoveTable mocks base method.
func (m *MockWorkloadWrapper) RemoveTable() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveTable")
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveTable indicates an expected call of RemoveTable.
func (mr *MockWorkloadWrapperMockRecorder) RemoveTable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTable", reflect.TypeOf((*MockWorkloadWrapper)(nil).RemoveTable))
}

// RemoveTable mocks base method.
func (m *MockWorkloadWrapper) RemoveMappingFile() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveMappingFile")
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveTable indicates an expected call of RemoveTable.
func (mr *MockWorkloadWrapperMockRecorder) RemoveMappingFile() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveMappingFile", reflect.TypeOf((*MockWorkloadWrapper)(nil).RemoveMappingFile))
}

// RegisterObserver mocks base method.
func (m *MockWorkloadWrapper) RegisterObserver(arg0 Observer) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterObserver", arg0)
}

// RegisterObserver indicates an expected call of RegisterObserver.
func (mr *MockWorkloadWrapperMockRecorder) RegisterObserver(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterObserver", reflect.TypeOf((*MockWorkloadWrapper)(nil).RegisterObserver), arg0)
}

// Remove mocks base method.
func (m *MockWorkloadWrapper) Remove(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockWorkloadWrapperMockRecorder) Remove(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockWorkloadWrapper)(nil).Remove), arg0)
}

// Run mocks base method.
func (m *MockWorkloadWrapper) Run(arg0 *v1.Pod, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockWorkloadWrapperMockRecorder) Run(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockWorkloadWrapper)(nil).Run), arg0, arg1)
}

// Start mocks base method.
func (m *MockWorkloadWrapper) Start(arg0 *v1.Pod) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockWorkloadWrapperMockRecorder) Start(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockWorkloadWrapper)(nil).Start), arg0)
}
