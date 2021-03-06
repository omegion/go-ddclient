// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/omegion/go-ddclient/internal/provider (interfaces: CloudflareAPIInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	cloudflare "github.com/cloudflare/cloudflare-go"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCloudflareAPIInterface is a mock of CloudflareAPIInterface interface
type MockCloudflareAPIInterface struct {
	ctrl     *gomock.Controller
	recorder *MockCloudflareAPIInterfaceMockRecorder
}

// MockCloudflareAPIInterfaceMockRecorder is the mock recorder for MockCloudflareAPIInterface
type MockCloudflareAPIInterfaceMockRecorder struct {
	mock *MockCloudflareAPIInterface
}

// NewMockCloudflareAPIInterface creates a new mock instance
func NewMockCloudflareAPIInterface(ctrl *gomock.Controller) *MockCloudflareAPIInterface {
	mock := &MockCloudflareAPIInterface{ctrl: ctrl}
	mock.recorder = &MockCloudflareAPIInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCloudflareAPIInterface) EXPECT() *MockCloudflareAPIInterfaceMockRecorder {
	return m.recorder
}

// CreateDNSRecord mocks base method
func (m *MockCloudflareAPIInterface) CreateDNSRecord(arg0 context.Context, arg1 string, arg2 cloudflare.DNSRecord) (*cloudflare.DNSRecordResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDNSRecord", arg0, arg1, arg2)
	ret0, _ := ret[0].(*cloudflare.DNSRecordResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDNSRecord indicates an expected call of CreateDNSRecord
func (mr *MockCloudflareAPIInterfaceMockRecorder) CreateDNSRecord(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDNSRecord", reflect.TypeOf((*MockCloudflareAPIInterface)(nil).CreateDNSRecord), arg0, arg1, arg2)
}

// DNSRecords mocks base method
func (m *MockCloudflareAPIInterface) DNSRecords(arg0 context.Context, arg1 string, arg2 cloudflare.DNSRecord) ([]cloudflare.DNSRecord, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DNSRecords", arg0, arg1, arg2)
	ret0, _ := ret[0].([]cloudflare.DNSRecord)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DNSRecords indicates an expected call of DNSRecords
func (mr *MockCloudflareAPIInterfaceMockRecorder) DNSRecords(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DNSRecords", reflect.TypeOf((*MockCloudflareAPIInterface)(nil).DNSRecords), arg0, arg1, arg2)
}

// UpdateDNSRecord mocks base method
func (m *MockCloudflareAPIInterface) UpdateDNSRecord(arg0 context.Context, arg1, arg2 string, arg3 cloudflare.DNSRecord) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDNSRecord", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateDNSRecord indicates an expected call of UpdateDNSRecord
func (mr *MockCloudflareAPIInterfaceMockRecorder) UpdateDNSRecord(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDNSRecord", reflect.TypeOf((*MockCloudflareAPIInterface)(nil).UpdateDNSRecord), arg0, arg1, arg2, arg3)
}

// ZoneIDByName mocks base method
func (m *MockCloudflareAPIInterface) ZoneIDByName(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZoneIDByName", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ZoneIDByName indicates an expected call of ZoneIDByName
func (mr *MockCloudflareAPIInterfaceMockRecorder) ZoneIDByName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZoneIDByName", reflect.TypeOf((*MockCloudflareAPIInterface)(nil).ZoneIDByName), arg0)
}
