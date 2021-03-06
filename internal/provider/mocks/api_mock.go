// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/omegion/go-ddclient/internal/provider (interfaces: API)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	provider "github.com/omegion/go-ddclient/internal/provider"
	reflect "reflect"
)

// MockAPI is a mock of API interface
type MockAPI struct {
	ctrl     *gomock.Controller
	recorder *MockAPIMockRecorder
}

// MockAPIMockRecorder is the mock recorder for MockAPI
type MockAPIMockRecorder struct {
	mock *MockAPI
}

// NewMockAPI creates a new mock instance
func NewMockAPI(ctrl *gomock.Controller) *MockAPI {
	mock := &MockAPI{ctrl: ctrl}
	mock.recorder = &MockAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAPI) EXPECT() *MockAPIMockRecorder {
	return m.recorder
}

// SetRecord mocks base method
func (m *MockAPI) SetRecord(arg0 context.Context, arg1 provider.DNSRecord) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetRecord", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetRecord indicates an expected call of SetRecord
func (mr *MockAPIMockRecorder) SetRecord(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRecord", reflect.TypeOf((*MockAPI)(nil).SetRecord), arg0, arg1)
}
