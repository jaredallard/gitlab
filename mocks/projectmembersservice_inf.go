// Code generated by MockGen. DO NOT EDIT.
// Source: projectmembersservice_inf.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=projectmembersservice_inf.go -destination=mocks/projectmembersservice_inf.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gitlab "github.com/xanzy/go-gitlab"
	gomock "go.uber.org/mock/gomock"
)

// MockProjectMembersService is a mock of ProjectMembersService interface.
type MockProjectMembersService struct {
	ctrl     *gomock.Controller
	recorder *MockProjectMembersServiceMockRecorder
	isgomock struct{}
}

// MockProjectMembersServiceMockRecorder is the mock recorder for MockProjectMembersService.
type MockProjectMembersServiceMockRecorder struct {
	mock *MockProjectMembersService
}

// NewMockProjectMembersService creates a new mock instance.
func NewMockProjectMembersService(ctrl *gomock.Controller) *MockProjectMembersService {
	mock := &MockProjectMembersService{ctrl: ctrl}
	mock.recorder = &MockProjectMembersServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectMembersService) EXPECT() *MockProjectMembersServiceMockRecorder {
	return m.recorder
}

// AddProjectMember mocks base method.
func (m *MockProjectMembersService) AddProjectMember(pid any, opt *gitlab.AddProjectMemberOptions, options ...gitlab.RequestOptionFunc) (*gitlab.ProjectMember, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddProjectMember", varargs...)
	ret0, _ := ret[0].(*gitlab.ProjectMember)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AddProjectMember indicates an expected call of AddProjectMember.
func (mr *MockProjectMembersServiceMockRecorder) AddProjectMember(pid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProjectMember", reflect.TypeOf((*MockProjectMembersService)(nil).AddProjectMember), varargs...)
}

// DeleteProjectMember mocks base method.
func (m *MockProjectMembersService) DeleteProjectMember(pid any, user int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, user}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteProjectMember", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteProjectMember indicates an expected call of DeleteProjectMember.
func (mr *MockProjectMembersServiceMockRecorder) DeleteProjectMember(pid, user any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, user}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProjectMember", reflect.TypeOf((*MockProjectMembersService)(nil).DeleteProjectMember), varargs...)
}

// EditProjectMember mocks base method.
func (m *MockProjectMembersService) EditProjectMember(pid any, user int, opt *gitlab.EditProjectMemberOptions, options ...gitlab.RequestOptionFunc) (*gitlab.ProjectMember, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, user, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EditProjectMember", varargs...)
	ret0, _ := ret[0].(*gitlab.ProjectMember)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// EditProjectMember indicates an expected call of EditProjectMember.
func (mr *MockProjectMembersServiceMockRecorder) EditProjectMember(pid, user, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, user, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditProjectMember", reflect.TypeOf((*MockProjectMembersService)(nil).EditProjectMember), varargs...)
}

// GetInheritedProjectMember mocks base method.
func (m *MockProjectMembersService) GetInheritedProjectMember(pid any, user int, options ...gitlab.RequestOptionFunc) (*gitlab.ProjectMember, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, user}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetInheritedProjectMember", varargs...)
	ret0, _ := ret[0].(*gitlab.ProjectMember)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetInheritedProjectMember indicates an expected call of GetInheritedProjectMember.
func (mr *MockProjectMembersServiceMockRecorder) GetInheritedProjectMember(pid, user any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, user}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInheritedProjectMember", reflect.TypeOf((*MockProjectMembersService)(nil).GetInheritedProjectMember), varargs...)
}

// GetProjectMember mocks base method.
func (m *MockProjectMembersService) GetProjectMember(pid any, user int, options ...gitlab.RequestOptionFunc) (*gitlab.ProjectMember, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, user}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetProjectMember", varargs...)
	ret0, _ := ret[0].(*gitlab.ProjectMember)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetProjectMember indicates an expected call of GetProjectMember.
func (mr *MockProjectMembersServiceMockRecorder) GetProjectMember(pid, user any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, user}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjectMember", reflect.TypeOf((*MockProjectMembersService)(nil).GetProjectMember), varargs...)
}

// ListAllProjectMembers mocks base method.
func (m *MockProjectMembersService) ListAllProjectMembers(pid any, opt *gitlab.ListProjectMembersOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.ProjectMember, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAllProjectMembers", varargs...)
	ret0, _ := ret[0].([]*gitlab.ProjectMember)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListAllProjectMembers indicates an expected call of ListAllProjectMembers.
func (mr *MockProjectMembersServiceMockRecorder) ListAllProjectMembers(pid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllProjectMembers", reflect.TypeOf((*MockProjectMembersService)(nil).ListAllProjectMembers), varargs...)
}

// ListProjectMembers mocks base method.
func (m *MockProjectMembersService) ListProjectMembers(pid any, opt *gitlab.ListProjectMembersOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.ProjectMember, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListProjectMembers", varargs...)
	ret0, _ := ret[0].([]*gitlab.ProjectMember)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListProjectMembers indicates an expected call of ListProjectMembers.
func (mr *MockProjectMembersServiceMockRecorder) ListProjectMembers(pid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjectMembers", reflect.TypeOf((*MockProjectMembersService)(nil).ListProjectMembers), varargs...)
}
