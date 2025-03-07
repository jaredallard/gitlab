// Code generated by MockGen. DO NOT EDIT.
// Source: errortrackingservice_inf.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=errortrackingservice_inf.go -destination=mocks/errortrackingservice_inf.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gitlab "gitlab.com/gitlab-org/api/client-go"
	gomock "go.uber.org/mock/gomock"
)

// MockErrorTrackingService is a mock of ErrorTrackingService interface.
type MockErrorTrackingService struct {
	ctrl     *gomock.Controller
	recorder *MockErrorTrackingServiceMockRecorder
	isgomock struct{}
}

// MockErrorTrackingServiceMockRecorder is the mock recorder for MockErrorTrackingService.
type MockErrorTrackingServiceMockRecorder struct {
	mock *MockErrorTrackingService
}

// NewMockErrorTrackingService creates a new mock instance.
func NewMockErrorTrackingService(ctrl *gomock.Controller) *MockErrorTrackingService {
	mock := &MockErrorTrackingService{ctrl: ctrl}
	mock.recorder = &MockErrorTrackingServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockErrorTrackingService) EXPECT() *MockErrorTrackingServiceMockRecorder {
	return m.recorder
}

// CreateClientKey mocks base method.
func (m *MockErrorTrackingService) CreateClientKey(pid any, options ...gitlab.RequestOptionFunc) (*gitlab.ErrorTrackingClientKey, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateClientKey", varargs...)
	ret0, _ := ret[0].(*gitlab.ErrorTrackingClientKey)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateClientKey indicates an expected call of CreateClientKey.
func (mr *MockErrorTrackingServiceMockRecorder) CreateClientKey(pid any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateClientKey", reflect.TypeOf((*MockErrorTrackingService)(nil).CreateClientKey), varargs...)
}

// DeleteClientKey mocks base method.
func (m *MockErrorTrackingService) DeleteClientKey(pid any, keyID int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, keyID}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteClientKey", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteClientKey indicates an expected call of DeleteClientKey.
func (mr *MockErrorTrackingServiceMockRecorder) DeleteClientKey(pid, keyID any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, keyID}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteClientKey", reflect.TypeOf((*MockErrorTrackingService)(nil).DeleteClientKey), varargs...)
}

// EnableDisableErrorTracking mocks base method.
func (m *MockErrorTrackingService) EnableDisableErrorTracking(pid any, opt *gitlab.EnableDisableErrorTrackingOptions, options ...gitlab.RequestOptionFunc) (*gitlab.ErrorTrackingSettings, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EnableDisableErrorTracking", varargs...)
	ret0, _ := ret[0].(*gitlab.ErrorTrackingSettings)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// EnableDisableErrorTracking indicates an expected call of EnableDisableErrorTracking.
func (mr *MockErrorTrackingServiceMockRecorder) EnableDisableErrorTracking(pid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableDisableErrorTracking", reflect.TypeOf((*MockErrorTrackingService)(nil).EnableDisableErrorTracking), varargs...)
}

// GetErrorTrackingSettings mocks base method.
func (m *MockErrorTrackingService) GetErrorTrackingSettings(pid any, options ...gitlab.RequestOptionFunc) (*gitlab.ErrorTrackingSettings, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetErrorTrackingSettings", varargs...)
	ret0, _ := ret[0].(*gitlab.ErrorTrackingSettings)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetErrorTrackingSettings indicates an expected call of GetErrorTrackingSettings.
func (mr *MockErrorTrackingServiceMockRecorder) GetErrorTrackingSettings(pid any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetErrorTrackingSettings", reflect.TypeOf((*MockErrorTrackingService)(nil).GetErrorTrackingSettings), varargs...)
}

// ListClientKeys mocks base method.
func (m *MockErrorTrackingService) ListClientKeys(pid any, opt *gitlab.ListClientKeysOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.ErrorTrackingClientKey, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListClientKeys", varargs...)
	ret0, _ := ret[0].([]*gitlab.ErrorTrackingClientKey)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListClientKeys indicates an expected call of ListClientKeys.
func (mr *MockErrorTrackingServiceMockRecorder) ListClientKeys(pid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListClientKeys", reflect.TypeOf((*MockErrorTrackingService)(nil).ListClientKeys), varargs...)
}
