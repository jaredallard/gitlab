// Code generated by MockGen. DO NOT EDIT.
// Source: projectmirrorservice_inf.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=projectmirrorservice_inf.go -destination=mocks/projectmirrorservice_inf.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gitlab "gitlab.com/gitlab-org/api/client-go"
	gomock "go.uber.org/mock/gomock"
)

// MockProjectMirrorService is a mock of ProjectMirrorService interface.
type MockProjectMirrorService struct {
	ctrl     *gomock.Controller
	recorder *MockProjectMirrorServiceMockRecorder
	isgomock struct{}
}

// MockProjectMirrorServiceMockRecorder is the mock recorder for MockProjectMirrorService.
type MockProjectMirrorServiceMockRecorder struct {
	mock *MockProjectMirrorService
}

// NewMockProjectMirrorService creates a new mock instance.
func NewMockProjectMirrorService(ctrl *gomock.Controller) *MockProjectMirrorService {
	mock := &MockProjectMirrorService{ctrl: ctrl}
	mock.recorder = &MockProjectMirrorServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectMirrorService) EXPECT() *MockProjectMirrorServiceMockRecorder {
	return m.recorder
}

// AddProjectMirror mocks base method.
func (m *MockProjectMirrorService) AddProjectMirror(pid any, opt *gitlab.AddProjectMirrorOptions, options ...gitlab.RequestOptionFunc) (*gitlab.ProjectMirror, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddProjectMirror", varargs...)
	ret0, _ := ret[0].(*gitlab.ProjectMirror)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AddProjectMirror indicates an expected call of AddProjectMirror.
func (mr *MockProjectMirrorServiceMockRecorder) AddProjectMirror(pid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProjectMirror", reflect.TypeOf((*MockProjectMirrorService)(nil).AddProjectMirror), varargs...)
}

// DeleteProjectMirror mocks base method.
func (m *MockProjectMirrorService) DeleteProjectMirror(pid any, mirror int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, mirror}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteProjectMirror", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteProjectMirror indicates an expected call of DeleteProjectMirror.
func (mr *MockProjectMirrorServiceMockRecorder) DeleteProjectMirror(pid, mirror any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, mirror}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProjectMirror", reflect.TypeOf((*MockProjectMirrorService)(nil).DeleteProjectMirror), varargs...)
}

// EditProjectMirror mocks base method.
func (m *MockProjectMirrorService) EditProjectMirror(pid any, mirror int, opt *gitlab.EditProjectMirrorOptions, options ...gitlab.RequestOptionFunc) (*gitlab.ProjectMirror, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, mirror, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EditProjectMirror", varargs...)
	ret0, _ := ret[0].(*gitlab.ProjectMirror)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// EditProjectMirror indicates an expected call of EditProjectMirror.
func (mr *MockProjectMirrorServiceMockRecorder) EditProjectMirror(pid, mirror, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, mirror, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditProjectMirror", reflect.TypeOf((*MockProjectMirrorService)(nil).EditProjectMirror), varargs...)
}

// GetProjectMirror mocks base method.
func (m *MockProjectMirrorService) GetProjectMirror(pid any, mirror int, options ...gitlab.RequestOptionFunc) (*gitlab.ProjectMirror, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, mirror}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetProjectMirror", varargs...)
	ret0, _ := ret[0].(*gitlab.ProjectMirror)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetProjectMirror indicates an expected call of GetProjectMirror.
func (mr *MockProjectMirrorServiceMockRecorder) GetProjectMirror(pid, mirror any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, mirror}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjectMirror", reflect.TypeOf((*MockProjectMirrorService)(nil).GetProjectMirror), varargs...)
}

// ListProjectMirror mocks base method.
func (m *MockProjectMirrorService) ListProjectMirror(pid any, opt *gitlab.ListProjectMirrorOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.ProjectMirror, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListProjectMirror", varargs...)
	ret0, _ := ret[0].([]*gitlab.ProjectMirror)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListProjectMirror indicates an expected call of ListProjectMirror.
func (mr *MockProjectMirrorServiceMockRecorder) ListProjectMirror(pid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjectMirror", reflect.TypeOf((*MockProjectMirrorService)(nil).ListProjectMirror), varargs...)
}
