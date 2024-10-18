// Code generated by MockGen. DO NOT EDIT.
// Source: groupmembersservice_inf.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=groupmembersservice_inf.go -destination=mocks/groupmembersservice_inf.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gitlab "github.com/xanzy/go-gitlab"
	gomock "go.uber.org/mock/gomock"
)

// MockGroupMembersService is a mock of GroupMembersService interface.
type MockGroupMembersService struct {
	ctrl     *gomock.Controller
	recorder *MockGroupMembersServiceMockRecorder
	isgomock struct{}
}

// MockGroupMembersServiceMockRecorder is the mock recorder for MockGroupMembersService.
type MockGroupMembersServiceMockRecorder struct {
	mock *MockGroupMembersService
}

// NewMockGroupMembersService creates a new mock instance.
func NewMockGroupMembersService(ctrl *gomock.Controller) *MockGroupMembersService {
	mock := &MockGroupMembersService{ctrl: ctrl}
	mock.recorder = &MockGroupMembersServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGroupMembersService) EXPECT() *MockGroupMembersServiceMockRecorder {
	return m.recorder
}

// AddGroupMember mocks base method.
func (m *MockGroupMembersService) AddGroupMember(gid any, opt *gitlab.AddGroupMemberOptions, options ...gitlab.RequestOptionFunc) (*gitlab.GroupMember, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{gid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddGroupMember", varargs...)
	ret0, _ := ret[0].(*gitlab.GroupMember)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AddGroupMember indicates an expected call of AddGroupMember.
func (mr *MockGroupMembersServiceMockRecorder) AddGroupMember(gid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{gid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddGroupMember", reflect.TypeOf((*MockGroupMembersService)(nil).AddGroupMember), varargs...)
}

// DeleteShareWithGroup mocks base method.
func (m *MockGroupMembersService) DeleteShareWithGroup(gid any, groupID int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{gid, groupID}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteShareWithGroup", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteShareWithGroup indicates an expected call of DeleteShareWithGroup.
func (mr *MockGroupMembersServiceMockRecorder) DeleteShareWithGroup(gid, groupID any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{gid, groupID}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteShareWithGroup", reflect.TypeOf((*MockGroupMembersService)(nil).DeleteShareWithGroup), varargs...)
}

// EditGroupMember mocks base method.
func (m *MockGroupMembersService) EditGroupMember(gid any, user int, opt *gitlab.EditGroupMemberOptions, options ...gitlab.RequestOptionFunc) (*gitlab.GroupMember, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{gid, user, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EditGroupMember", varargs...)
	ret0, _ := ret[0].(*gitlab.GroupMember)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// EditGroupMember indicates an expected call of EditGroupMember.
func (mr *MockGroupMembersServiceMockRecorder) EditGroupMember(gid, user, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{gid, user, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditGroupMember", reflect.TypeOf((*MockGroupMembersService)(nil).EditGroupMember), varargs...)
}

// GetGroupMember mocks base method.
func (m *MockGroupMembersService) GetGroupMember(gid any, user int, options ...gitlab.RequestOptionFunc) (*gitlab.GroupMember, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{gid, user}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetGroupMember", varargs...)
	ret0, _ := ret[0].(*gitlab.GroupMember)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetGroupMember indicates an expected call of GetGroupMember.
func (mr *MockGroupMembersServiceMockRecorder) GetGroupMember(gid, user any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{gid, user}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupMember", reflect.TypeOf((*MockGroupMembersService)(nil).GetGroupMember), varargs...)
}

// GetInheritedGroupMember mocks base method.
func (m *MockGroupMembersService) GetInheritedGroupMember(gid any, user int, options ...gitlab.RequestOptionFunc) (*gitlab.GroupMember, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{gid, user}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetInheritedGroupMember", varargs...)
	ret0, _ := ret[0].(*gitlab.GroupMember)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetInheritedGroupMember indicates an expected call of GetInheritedGroupMember.
func (mr *MockGroupMembersServiceMockRecorder) GetInheritedGroupMember(gid, user any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{gid, user}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInheritedGroupMember", reflect.TypeOf((*MockGroupMembersService)(nil).GetInheritedGroupMember), varargs...)
}

// RemoveGroupMember mocks base method.
func (m *MockGroupMembersService) RemoveGroupMember(gid any, user int, opt *gitlab.RemoveGroupMemberOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{gid, user, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RemoveGroupMember", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveGroupMember indicates an expected call of RemoveGroupMember.
func (mr *MockGroupMembersServiceMockRecorder) RemoveGroupMember(gid, user, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{gid, user, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveGroupMember", reflect.TypeOf((*MockGroupMembersService)(nil).RemoveGroupMember), varargs...)
}

// ShareWithGroup mocks base method.
func (m *MockGroupMembersService) ShareWithGroup(gid any, opt *gitlab.ShareWithGroupOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Group, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{gid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ShareWithGroup", varargs...)
	ret0, _ := ret[0].(*gitlab.Group)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ShareWithGroup indicates an expected call of ShareWithGroup.
func (mr *MockGroupMembersServiceMockRecorder) ShareWithGroup(gid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{gid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShareWithGroup", reflect.TypeOf((*MockGroupMembersService)(nil).ShareWithGroup), varargs...)
}
