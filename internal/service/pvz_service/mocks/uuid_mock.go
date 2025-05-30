// Code generated by MockGen. DO NOT EDIT.
// Source: uuid.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUUIDInterface is a mock of UUIDInterface interface.
type MockUUIDInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUUIDInterfaceMockRecorder
}

// MockUUIDInterfaceMockRecorder is the mock recorder for MockUUIDInterface.
type MockUUIDInterfaceMockRecorder struct {
	mock *MockUUIDInterface
}

// NewMockUUIDInterface creates a new mock instance.
func NewMockUUIDInterface(ctrl *gomock.Controller) *MockUUIDInterface {
	mock := &MockUUIDInterface{ctrl: ctrl}
	mock.recorder = &MockUUIDInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUUIDInterface) EXPECT() *MockUUIDInterfaceMockRecorder {
	return m.recorder
}

// NewString mocks base method.
func (m *MockUUIDInterface) NewString() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewString")
	ret0, _ := ret[0].(string)
	return ret0
}

// NewString indicates an expected call of NewString.
func (mr *MockUUIDInterfaceMockRecorder) NewString() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewString", reflect.TypeOf((*MockUUIDInterface)(nil).NewString))
}
