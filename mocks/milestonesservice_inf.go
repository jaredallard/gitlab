// Code generated by MockGen. DO NOT EDIT.
// Source: milestonesservice_inf.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=milestonesservice_inf.go -destination=mocks/milestonesservice_inf.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gitlab "github.com/xanzy/go-gitlab"
	gomock "go.uber.org/mock/gomock"
)

// MockMilestonesService is a mock of MilestonesService interface.
type MockMilestonesService struct {
	ctrl     *gomock.Controller
	recorder *MockMilestonesServiceMockRecorder
}

// MockMilestonesServiceMockRecorder is the mock recorder for MockMilestonesService.
type MockMilestonesServiceMockRecorder struct {
	mock *MockMilestonesService
}

// NewMockMilestonesService creates a new mock instance.
func NewMockMilestonesService(ctrl *gomock.Controller) *MockMilestonesService {
	mock := &MockMilestonesService{ctrl: ctrl}
	mock.recorder = &MockMilestonesServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMilestonesService) EXPECT() *MockMilestonesServiceMockRecorder {
	return m.recorder
}

// CreateMilestone mocks base method.
func (m *MockMilestonesService) CreateMilestone(pid any, opt *gitlab.CreateMilestoneOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Milestone, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateMilestone", varargs...)
	ret0, _ := ret[0].(*gitlab.Milestone)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateMilestone indicates an expected call of CreateMilestone.
func (mr *MockMilestonesServiceMockRecorder) CreateMilestone(pid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMilestone", reflect.TypeOf((*MockMilestonesService)(nil).CreateMilestone), varargs...)
}

// DeleteMilestone mocks base method.
func (m *MockMilestonesService) DeleteMilestone(pid any, milestone int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, milestone}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteMilestone", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteMilestone indicates an expected call of DeleteMilestone.
func (mr *MockMilestonesServiceMockRecorder) DeleteMilestone(pid, milestone any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, milestone}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMilestone", reflect.TypeOf((*MockMilestonesService)(nil).DeleteMilestone), varargs...)
}

// GetMilestone mocks base method.
func (m *MockMilestonesService) GetMilestone(pid any, milestone int, options ...gitlab.RequestOptionFunc) (*gitlab.Milestone, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, milestone}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMilestone", varargs...)
	ret0, _ := ret[0].(*gitlab.Milestone)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetMilestone indicates an expected call of GetMilestone.
func (mr *MockMilestonesServiceMockRecorder) GetMilestone(pid, milestone any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, milestone}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMilestone", reflect.TypeOf((*MockMilestonesService)(nil).GetMilestone), varargs...)
}

// GetMilestoneIssues mocks base method.
func (m *MockMilestonesService) GetMilestoneIssues(pid any, milestone int, opt *gitlab.GetMilestoneIssuesOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.Issue, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, milestone, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMilestoneIssues", varargs...)
	ret0, _ := ret[0].([]*gitlab.Issue)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetMilestoneIssues indicates an expected call of GetMilestoneIssues.
func (mr *MockMilestonesServiceMockRecorder) GetMilestoneIssues(pid, milestone, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, milestone, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMilestoneIssues", reflect.TypeOf((*MockMilestonesService)(nil).GetMilestoneIssues), varargs...)
}

// GetMilestoneMergeRequests mocks base method.
func (m *MockMilestonesService) GetMilestoneMergeRequests(pid any, milestone int, opt *gitlab.GetMilestoneMergeRequestsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.MergeRequest, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, milestone, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMilestoneMergeRequests", varargs...)
	ret0, _ := ret[0].([]*gitlab.MergeRequest)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetMilestoneMergeRequests indicates an expected call of GetMilestoneMergeRequests.
func (mr *MockMilestonesServiceMockRecorder) GetMilestoneMergeRequests(pid, milestone, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, milestone, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMilestoneMergeRequests", reflect.TypeOf((*MockMilestonesService)(nil).GetMilestoneMergeRequests), varargs...)
}

// ListMilestones mocks base method.
func (m *MockMilestonesService) ListMilestones(pid any, opt *gitlab.ListMilestonesOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.Milestone, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListMilestones", varargs...)
	ret0, _ := ret[0].([]*gitlab.Milestone)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListMilestones indicates an expected call of ListMilestones.
func (mr *MockMilestonesServiceMockRecorder) ListMilestones(pid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMilestones", reflect.TypeOf((*MockMilestonesService)(nil).ListMilestones), varargs...)
}

// UpdateMilestone mocks base method.
func (m *MockMilestonesService) UpdateMilestone(pid any, milestone int, opt *gitlab.UpdateMilestoneOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Milestone, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, milestone, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateMilestone", varargs...)
	ret0, _ := ret[0].(*gitlab.Milestone)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UpdateMilestone indicates an expected call of UpdateMilestone.
func (mr *MockMilestonesServiceMockRecorder) UpdateMilestone(pid, milestone, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, milestone, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMilestone", reflect.TypeOf((*MockMilestonesService)(nil).UpdateMilestone), varargs...)
}