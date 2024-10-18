// Code generated by MockGen. DO NOT EDIT.
// Source: resourcestateeventsservice_inf.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=resourcestateeventsservice_inf.go -destination=mocks/resourcestateeventsservice_inf.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gitlab "github.com/xanzy/go-gitlab"
	gomock "go.uber.org/mock/gomock"
)

// MockResourceStateEventsService is a mock of ResourceStateEventsService interface.
type MockResourceStateEventsService struct {
	ctrl     *gomock.Controller
	recorder *MockResourceStateEventsServiceMockRecorder
	isgomock struct{}
}

// MockResourceStateEventsServiceMockRecorder is the mock recorder for MockResourceStateEventsService.
type MockResourceStateEventsServiceMockRecorder struct {
	mock *MockResourceStateEventsService
}

// NewMockResourceStateEventsService creates a new mock instance.
func NewMockResourceStateEventsService(ctrl *gomock.Controller) *MockResourceStateEventsService {
	mock := &MockResourceStateEventsService{ctrl: ctrl}
	mock.recorder = &MockResourceStateEventsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResourceStateEventsService) EXPECT() *MockResourceStateEventsServiceMockRecorder {
	return m.recorder
}

// GetIssueStateEvent mocks base method.
func (m *MockResourceStateEventsService) GetIssueStateEvent(pid any, issue, event int, options ...gitlab.RequestOptionFunc) (*gitlab.StateEvent, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, issue, event}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetIssueStateEvent", varargs...)
	ret0, _ := ret[0].(*gitlab.StateEvent)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetIssueStateEvent indicates an expected call of GetIssueStateEvent.
func (mr *MockResourceStateEventsServiceMockRecorder) GetIssueStateEvent(pid, issue, event any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, issue, event}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIssueStateEvent", reflect.TypeOf((*MockResourceStateEventsService)(nil).GetIssueStateEvent), varargs...)
}

// GetMergeRequestStateEvent mocks base method.
func (m *MockResourceStateEventsService) GetMergeRequestStateEvent(pid any, request, event int, options ...gitlab.RequestOptionFunc) (*gitlab.StateEvent, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, request, event}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMergeRequestStateEvent", varargs...)
	ret0, _ := ret[0].(*gitlab.StateEvent)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetMergeRequestStateEvent indicates an expected call of GetMergeRequestStateEvent.
func (mr *MockResourceStateEventsServiceMockRecorder) GetMergeRequestStateEvent(pid, request, event any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, request, event}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMergeRequestStateEvent", reflect.TypeOf((*MockResourceStateEventsService)(nil).GetMergeRequestStateEvent), varargs...)
}

// ListIssueStateEvents mocks base method.
func (m *MockResourceStateEventsService) ListIssueStateEvents(pid any, issue int, opt *gitlab.ListStateEventsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.StateEvent, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, issue, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListIssueStateEvents", varargs...)
	ret0, _ := ret[0].([]*gitlab.StateEvent)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListIssueStateEvents indicates an expected call of ListIssueStateEvents.
func (mr *MockResourceStateEventsServiceMockRecorder) ListIssueStateEvents(pid, issue, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, issue, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListIssueStateEvents", reflect.TypeOf((*MockResourceStateEventsService)(nil).ListIssueStateEvents), varargs...)
}

// ListMergeStateEvents mocks base method.
func (m *MockResourceStateEventsService) ListMergeStateEvents(pid any, request int, opt *gitlab.ListStateEventsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.StateEvent, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, request, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListMergeStateEvents", varargs...)
	ret0, _ := ret[0].([]*gitlab.StateEvent)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListMergeStateEvents indicates an expected call of ListMergeStateEvents.
func (mr *MockResourceStateEventsServiceMockRecorder) ListMergeStateEvents(pid, request, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, request, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMergeStateEvents", reflect.TypeOf((*MockResourceStateEventsService)(nil).ListMergeStateEvents), varargs...)
}
