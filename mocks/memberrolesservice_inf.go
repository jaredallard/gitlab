// Code generated by MockGen. DO NOT EDIT.
// Source: memberrolesservice_inf.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=memberrolesservice_inf.go -destination=mocks/memberrolesservice_inf.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gitlab "gitlab.com/gitlab-org/api/client-go"
	gomock "go.uber.org/mock/gomock"
)

// MockMemberRolesService is a mock of MemberRolesService interface.
type MockMemberRolesService struct {
	ctrl     *gomock.Controller
	recorder *MockMemberRolesServiceMockRecorder
	isgomock struct{}
}

// MockMemberRolesServiceMockRecorder is the mock recorder for MockMemberRolesService.
type MockMemberRolesServiceMockRecorder struct {
	mock *MockMemberRolesService
}

// NewMockMemberRolesService creates a new mock instance.
func NewMockMemberRolesService(ctrl *gomock.Controller) *MockMemberRolesService {
	mock := &MockMemberRolesService{ctrl: ctrl}
	mock.recorder = &MockMemberRolesServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMemberRolesService) EXPECT() *MockMemberRolesServiceMockRecorder {
	return m.recorder
}

// CreateMemberRole mocks base method.
func (m *MockMemberRolesService) CreateMemberRole(gid any, opt *gitlab.CreateMemberRoleOptions, options ...gitlab.RequestOptionFunc) (*gitlab.MemberRole, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{gid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateMemberRole", varargs...)
	ret0, _ := ret[0].(*gitlab.MemberRole)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateMemberRole indicates an expected call of CreateMemberRole.
func (mr *MockMemberRolesServiceMockRecorder) CreateMemberRole(gid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{gid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMemberRole", reflect.TypeOf((*MockMemberRolesService)(nil).CreateMemberRole), varargs...)
}

// DeleteMemberRole mocks base method.
func (m *MockMemberRolesService) DeleteMemberRole(gid any, memberRole int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{gid, memberRole}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteMemberRole", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteMemberRole indicates an expected call of DeleteMemberRole.
func (mr *MockMemberRolesServiceMockRecorder) DeleteMemberRole(gid, memberRole any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{gid, memberRole}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMemberRole", reflect.TypeOf((*MockMemberRolesService)(nil).DeleteMemberRole), varargs...)
}

// ListMemberRoles mocks base method.
func (m *MockMemberRolesService) ListMemberRoles(gid any, options ...gitlab.RequestOptionFunc) ([]*gitlab.MemberRole, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{gid}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListMemberRoles", varargs...)
	ret0, _ := ret[0].([]*gitlab.MemberRole)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListMemberRoles indicates an expected call of ListMemberRoles.
func (mr *MockMemberRolesServiceMockRecorder) ListMemberRoles(gid any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{gid}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMemberRoles", reflect.TypeOf((*MockMemberRolesService)(nil).ListMemberRoles), varargs...)
}
