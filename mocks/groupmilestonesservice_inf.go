// Code generated by MockGen. DO NOT EDIT.
// Source: groupmilestonesservice_inf.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=groupmilestonesservice_inf.go -destination=mocks/groupmilestonesservice_inf.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gitlab "gitlab.com/gitlab-org/api/client-go"
	gomock "go.uber.org/mock/gomock"
)

// MockGroupMilestonesService is a mock of GroupMilestonesService interface.
type MockGroupMilestonesService struct {
	ctrl     *gomock.Controller
	recorder *MockGroupMilestonesServiceMockRecorder
	isgomock struct{}
}

// MockGroupMilestonesServiceMockRecorder is the mock recorder for MockGroupMilestonesService.
type MockGroupMilestonesServiceMockRecorder struct {
	mock *MockGroupMilestonesService
}

// NewMockGroupMilestonesService creates a new mock instance.
func NewMockGroupMilestonesService(ctrl *gomock.Controller) *MockGroupMilestonesService {
	mock := &MockGroupMilestonesService{ctrl: ctrl}
	mock.recorder = &MockGroupMilestonesServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGroupMilestonesService) EXPECT() *MockGroupMilestonesServiceMockRecorder {
	return m.recorder
}

// CreateGroupMilestone mocks base method.
func (m *MockGroupMilestonesService) CreateGroupMilestone(gid any, opt *gitlab.CreateGroupMilestoneOptions, options ...gitlab.RequestOptionFunc) (*gitlab.GroupMilestone, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{gid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateGroupMilestone", varargs...)
	ret0, _ := ret[0].(*gitlab.GroupMilestone)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateGroupMilestone indicates an expected call of CreateGroupMilestone.
func (mr *MockGroupMilestonesServiceMockRecorder) CreateGroupMilestone(gid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{gid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGroupMilestone", reflect.TypeOf((*MockGroupMilestonesService)(nil).CreateGroupMilestone), varargs...)
}

// DeleteGroupMilestone mocks base method.
func (m *MockGroupMilestonesService) DeleteGroupMilestone(pid any, milestone int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, milestone}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteGroupMilestone", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteGroupMilestone indicates an expected call of DeleteGroupMilestone.
func (mr *MockGroupMilestonesServiceMockRecorder) DeleteGroupMilestone(pid, milestone any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, milestone}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGroupMilestone", reflect.TypeOf((*MockGroupMilestonesService)(nil).DeleteGroupMilestone), varargs...)
}

// GetGroupMilestone mocks base method.
func (m *MockGroupMilestonesService) GetGroupMilestone(gid any, milestone int, options ...gitlab.RequestOptionFunc) (*gitlab.GroupMilestone, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{gid, milestone}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetGroupMilestone", varargs...)
	ret0, _ := ret[0].(*gitlab.GroupMilestone)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetGroupMilestone indicates an expected call of GetGroupMilestone.
func (mr *MockGroupMilestonesServiceMockRecorder) GetGroupMilestone(gid, milestone any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{gid, milestone}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupMilestone", reflect.TypeOf((*MockGroupMilestonesService)(nil).GetGroupMilestone), varargs...)
}

// GetGroupMilestoneBurndownChartEvents mocks base method.
func (m *MockGroupMilestonesService) GetGroupMilestoneBurndownChartEvents(gid any, milestone int, opt *gitlab.GetGroupMilestoneBurndownChartEventsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.BurndownChartEvent, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{gid, milestone, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetGroupMilestoneBurndownChartEvents", varargs...)
	ret0, _ := ret[0].([]*gitlab.BurndownChartEvent)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetGroupMilestoneBurndownChartEvents indicates an expected call of GetGroupMilestoneBurndownChartEvents.
func (mr *MockGroupMilestonesServiceMockRecorder) GetGroupMilestoneBurndownChartEvents(gid, milestone, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{gid, milestone, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupMilestoneBurndownChartEvents", reflect.TypeOf((*MockGroupMilestonesService)(nil).GetGroupMilestoneBurndownChartEvents), varargs...)
}

// GetGroupMilestoneIssues mocks base method.
func (m *MockGroupMilestonesService) GetGroupMilestoneIssues(gid any, milestone int, opt *gitlab.GetGroupMilestoneIssuesOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.Issue, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{gid, milestone, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetGroupMilestoneIssues", varargs...)
	ret0, _ := ret[0].([]*gitlab.Issue)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetGroupMilestoneIssues indicates an expected call of GetGroupMilestoneIssues.
func (mr *MockGroupMilestonesServiceMockRecorder) GetGroupMilestoneIssues(gid, milestone, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{gid, milestone, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupMilestoneIssues", reflect.TypeOf((*MockGroupMilestonesService)(nil).GetGroupMilestoneIssues), varargs...)
}

// GetGroupMilestoneMergeRequests mocks base method.
func (m *MockGroupMilestonesService) GetGroupMilestoneMergeRequests(gid any, milestone int, opt *gitlab.GetGroupMilestoneMergeRequestsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.BasicMergeRequest, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{gid, milestone, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetGroupMilestoneMergeRequests", varargs...)
	ret0, _ := ret[0].([]*gitlab.BasicMergeRequest)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetGroupMilestoneMergeRequests indicates an expected call of GetGroupMilestoneMergeRequests.
func (mr *MockGroupMilestonesServiceMockRecorder) GetGroupMilestoneMergeRequests(gid, milestone, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{gid, milestone, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupMilestoneMergeRequests", reflect.TypeOf((*MockGroupMilestonesService)(nil).GetGroupMilestoneMergeRequests), varargs...)
}

// ListGroupMilestones mocks base method.
func (m *MockGroupMilestonesService) ListGroupMilestones(gid any, opt *gitlab.ListGroupMilestonesOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.GroupMilestone, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{gid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListGroupMilestones", varargs...)
	ret0, _ := ret[0].([]*gitlab.GroupMilestone)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListGroupMilestones indicates an expected call of ListGroupMilestones.
func (mr *MockGroupMilestonesServiceMockRecorder) ListGroupMilestones(gid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{gid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGroupMilestones", reflect.TypeOf((*MockGroupMilestonesService)(nil).ListGroupMilestones), varargs...)
}

// UpdateGroupMilestone mocks base method.
func (m *MockGroupMilestonesService) UpdateGroupMilestone(gid any, milestone int, opt *gitlab.UpdateGroupMilestoneOptions, options ...gitlab.RequestOptionFunc) (*gitlab.GroupMilestone, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{gid, milestone, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateGroupMilestone", varargs...)
	ret0, _ := ret[0].(*gitlab.GroupMilestone)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UpdateGroupMilestone indicates an expected call of UpdateGroupMilestone.
func (mr *MockGroupMilestonesServiceMockRecorder) UpdateGroupMilestone(gid, milestone, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{gid, milestone, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGroupMilestone", reflect.TypeOf((*MockGroupMilestonesService)(nil).UpdateGroupMilestone), varargs...)
}
