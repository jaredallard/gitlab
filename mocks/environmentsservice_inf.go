// Code generated by MockGen. DO NOT EDIT.
// Source: environmentsservice_inf.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=environmentsservice_inf.go -destination=mocks/environmentsservice_inf.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gitlab "github.com/xanzy/go-gitlab"
	gomock "go.uber.org/mock/gomock"
)

// MockEnvironmentsService is a mock of EnvironmentsService interface.
type MockEnvironmentsService struct {
	ctrl     *gomock.Controller
	recorder *MockEnvironmentsServiceMockRecorder
}

// MockEnvironmentsServiceMockRecorder is the mock recorder for MockEnvironmentsService.
type MockEnvironmentsServiceMockRecorder struct {
	mock *MockEnvironmentsService
}

// NewMockEnvironmentsService creates a new mock instance.
func NewMockEnvironmentsService(ctrl *gomock.Controller) *MockEnvironmentsService {
	mock := &MockEnvironmentsService{ctrl: ctrl}
	mock.recorder = &MockEnvironmentsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEnvironmentsService) EXPECT() *MockEnvironmentsServiceMockRecorder {
	return m.recorder
}

// CreateEnvironment mocks base method.
func (m *MockEnvironmentsService) CreateEnvironment(pid any, opt *gitlab.CreateEnvironmentOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Environment, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateEnvironment", varargs...)
	ret0, _ := ret[0].(*gitlab.Environment)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateEnvironment indicates an expected call of CreateEnvironment.
func (mr *MockEnvironmentsServiceMockRecorder) CreateEnvironment(pid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEnvironment", reflect.TypeOf((*MockEnvironmentsService)(nil).CreateEnvironment), varargs...)
}

// DeleteEnvironment mocks base method.
func (m *MockEnvironmentsService) DeleteEnvironment(pid any, environment int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, environment}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteEnvironment", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteEnvironment indicates an expected call of DeleteEnvironment.
func (mr *MockEnvironmentsServiceMockRecorder) DeleteEnvironment(pid, environment any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, environment}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEnvironment", reflect.TypeOf((*MockEnvironmentsService)(nil).DeleteEnvironment), varargs...)
}

// EditEnvironment mocks base method.
func (m *MockEnvironmentsService) EditEnvironment(pid any, environment int, opt *gitlab.EditEnvironmentOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Environment, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, environment, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EditEnvironment", varargs...)
	ret0, _ := ret[0].(*gitlab.Environment)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// EditEnvironment indicates an expected call of EditEnvironment.
func (mr *MockEnvironmentsServiceMockRecorder) EditEnvironment(pid, environment, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, environment, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditEnvironment", reflect.TypeOf((*MockEnvironmentsService)(nil).EditEnvironment), varargs...)
}

// GetEnvironment mocks base method.
func (m *MockEnvironmentsService) GetEnvironment(pid any, environment int, options ...gitlab.RequestOptionFunc) (*gitlab.Environment, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, environment}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetEnvironment", varargs...)
	ret0, _ := ret[0].(*gitlab.Environment)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetEnvironment indicates an expected call of GetEnvironment.
func (mr *MockEnvironmentsServiceMockRecorder) GetEnvironment(pid, environment any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, environment}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEnvironment", reflect.TypeOf((*MockEnvironmentsService)(nil).GetEnvironment), varargs...)
}

// ListEnvironments mocks base method.
func (m *MockEnvironmentsService) ListEnvironments(pid any, opts *gitlab.ListEnvironmentsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.Environment, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opts}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListEnvironments", varargs...)
	ret0, _ := ret[0].([]*gitlab.Environment)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListEnvironments indicates an expected call of ListEnvironments.
func (mr *MockEnvironmentsServiceMockRecorder) ListEnvironments(pid, opts any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opts}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEnvironments", reflect.TypeOf((*MockEnvironmentsService)(nil).ListEnvironments), varargs...)
}

// StopEnvironment mocks base method.
func (m *MockEnvironmentsService) StopEnvironment(pid any, environmentID int, opt *gitlab.StopEnvironmentOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Environment, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, environmentID, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StopEnvironment", varargs...)
	ret0, _ := ret[0].(*gitlab.Environment)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// StopEnvironment indicates an expected call of StopEnvironment.
func (mr *MockEnvironmentsServiceMockRecorder) StopEnvironment(pid, environmentID, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, environmentID, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopEnvironment", reflect.TypeOf((*MockEnvironmentsService)(nil).StopEnvironment), varargs...)
}