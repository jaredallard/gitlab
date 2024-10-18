// Code generated by MockGen. DO NOT EDIT.
// Source: grouplabelsservice_inf.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=grouplabelsservice_inf.go -destination=mocks/grouplabelsservice_inf.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gitlab "github.com/xanzy/go-gitlab"
	gomock "go.uber.org/mock/gomock"
)

// MockGroupLabelsService is a mock of GroupLabelsService interface.
type MockGroupLabelsService struct {
	ctrl     *gomock.Controller
	recorder *MockGroupLabelsServiceMockRecorder
	isgomock struct{}
}

// MockGroupLabelsServiceMockRecorder is the mock recorder for MockGroupLabelsService.
type MockGroupLabelsServiceMockRecorder struct {
	mock *MockGroupLabelsService
}

// NewMockGroupLabelsService creates a new mock instance.
func NewMockGroupLabelsService(ctrl *gomock.Controller) *MockGroupLabelsService {
	mock := &MockGroupLabelsService{ctrl: ctrl}
	mock.recorder = &MockGroupLabelsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGroupLabelsService) EXPECT() *MockGroupLabelsServiceMockRecorder {
	return m.recorder
}

// CreateGroupLabel mocks base method.
func (m *MockGroupLabelsService) CreateGroupLabel(gid any, opt *gitlab.CreateGroupLabelOptions, options ...gitlab.RequestOptionFunc) (*gitlab.GroupLabel, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{gid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateGroupLabel", varargs...)
	ret0, _ := ret[0].(*gitlab.GroupLabel)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateGroupLabel indicates an expected call of CreateGroupLabel.
func (mr *MockGroupLabelsServiceMockRecorder) CreateGroupLabel(gid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{gid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGroupLabel", reflect.TypeOf((*MockGroupLabelsService)(nil).CreateGroupLabel), varargs...)
}

// DeleteGroupLabel mocks base method.
func (m *MockGroupLabelsService) DeleteGroupLabel(gid, lid any, opt *gitlab.DeleteGroupLabelOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{gid, lid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteGroupLabel", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteGroupLabel indicates an expected call of DeleteGroupLabel.
func (mr *MockGroupLabelsServiceMockRecorder) DeleteGroupLabel(gid, lid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{gid, lid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGroupLabel", reflect.TypeOf((*MockGroupLabelsService)(nil).DeleteGroupLabel), varargs...)
}

// GetGroupLabel mocks base method.
func (m *MockGroupLabelsService) GetGroupLabel(gid, labelID any, options ...gitlab.RequestOptionFunc) (*gitlab.GroupLabel, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{gid, labelID}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetGroupLabel", varargs...)
	ret0, _ := ret[0].(*gitlab.GroupLabel)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetGroupLabel indicates an expected call of GetGroupLabel.
func (mr *MockGroupLabelsServiceMockRecorder) GetGroupLabel(gid, labelID any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{gid, labelID}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupLabel", reflect.TypeOf((*MockGroupLabelsService)(nil).GetGroupLabel), varargs...)
}

// ListGroupLabels mocks base method.
func (m *MockGroupLabelsService) ListGroupLabels(gid any, opt *gitlab.ListGroupLabelsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.GroupLabel, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{gid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListGroupLabels", varargs...)
	ret0, _ := ret[0].([]*gitlab.GroupLabel)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListGroupLabels indicates an expected call of ListGroupLabels.
func (mr *MockGroupLabelsServiceMockRecorder) ListGroupLabels(gid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{gid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGroupLabels", reflect.TypeOf((*MockGroupLabelsService)(nil).ListGroupLabels), varargs...)
}

// SubscribeToGroupLabel mocks base method.
func (m *MockGroupLabelsService) SubscribeToGroupLabel(gid, labelID any, options ...gitlab.RequestOptionFunc) (*gitlab.GroupLabel, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{gid, labelID}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SubscribeToGroupLabel", varargs...)
	ret0, _ := ret[0].(*gitlab.GroupLabel)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SubscribeToGroupLabel indicates an expected call of SubscribeToGroupLabel.
func (mr *MockGroupLabelsServiceMockRecorder) SubscribeToGroupLabel(gid, labelID any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{gid, labelID}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeToGroupLabel", reflect.TypeOf((*MockGroupLabelsService)(nil).SubscribeToGroupLabel), varargs...)
}

// UnsubscribeFromGroupLabel mocks base method.
func (m *MockGroupLabelsService) UnsubscribeFromGroupLabel(gid, labelID any, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{gid, labelID}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UnsubscribeFromGroupLabel", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnsubscribeFromGroupLabel indicates an expected call of UnsubscribeFromGroupLabel.
func (mr *MockGroupLabelsServiceMockRecorder) UnsubscribeFromGroupLabel(gid, labelID any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{gid, labelID}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsubscribeFromGroupLabel", reflect.TypeOf((*MockGroupLabelsService)(nil).UnsubscribeFromGroupLabel), varargs...)
}

// UpdateGroupLabel mocks base method.
func (m *MockGroupLabelsService) UpdateGroupLabel(gid any, opt *gitlab.UpdateGroupLabelOptions, options ...gitlab.RequestOptionFunc) (*gitlab.GroupLabel, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{gid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateGroupLabel", varargs...)
	ret0, _ := ret[0].(*gitlab.GroupLabel)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UpdateGroupLabel indicates an expected call of UpdateGroupLabel.
func (mr *MockGroupLabelsServiceMockRecorder) UpdateGroupLabel(gid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{gid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGroupLabel", reflect.TypeOf((*MockGroupLabelsService)(nil).UpdateGroupLabel), varargs...)
}
