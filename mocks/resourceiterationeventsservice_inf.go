// Code generated by MockGen. DO NOT EDIT.
// Source: resourceiterationeventsservice_inf.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=resourceiterationeventsservice_inf.go -destination=mocks/resourceiterationeventsservice_inf.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gitlab "gitlab.com/gitlab-org/api/client-go"
	gomock "go.uber.org/mock/gomock"
)

// MockResourceIterationEventsService is a mock of ResourceIterationEventsService interface.
type MockResourceIterationEventsService struct {
	ctrl     *gomock.Controller
	recorder *MockResourceIterationEventsServiceMockRecorder
	isgomock struct{}
}

// MockResourceIterationEventsServiceMockRecorder is the mock recorder for MockResourceIterationEventsService.
type MockResourceIterationEventsServiceMockRecorder struct {
	mock *MockResourceIterationEventsService
}

// NewMockResourceIterationEventsService creates a new mock instance.
func NewMockResourceIterationEventsService(ctrl *gomock.Controller) *MockResourceIterationEventsService {
	mock := &MockResourceIterationEventsService{ctrl: ctrl}
	mock.recorder = &MockResourceIterationEventsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResourceIterationEventsService) EXPECT() *MockResourceIterationEventsServiceMockRecorder {
	return m.recorder
}

// GetIssueIterationEvent mocks base method.
func (m *MockResourceIterationEventsService) GetIssueIterationEvent(pid any, issue, event int, options ...gitlab.RequestOptionFunc) (*gitlab.IterationEvent, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, issue, event}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetIssueIterationEvent", varargs...)
	ret0, _ := ret[0].(*gitlab.IterationEvent)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetIssueIterationEvent indicates an expected call of GetIssueIterationEvent.
func (mr *MockResourceIterationEventsServiceMockRecorder) GetIssueIterationEvent(pid, issue, event any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, issue, event}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIssueIterationEvent", reflect.TypeOf((*MockResourceIterationEventsService)(nil).GetIssueIterationEvent), varargs...)
}

// ListIssueIterationEvents mocks base method.
func (m *MockResourceIterationEventsService) ListIssueIterationEvents(pid any, issue int, opt *gitlab.ListIterationEventsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.IterationEvent, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, issue, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListIssueIterationEvents", varargs...)
	ret0, _ := ret[0].([]*gitlab.IterationEvent)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListIssueIterationEvents indicates an expected call of ListIssueIterationEvents.
func (mr *MockResourceIterationEventsServiceMockRecorder) ListIssueIterationEvents(pid, issue, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, issue, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListIssueIterationEvents", reflect.TypeOf((*MockResourceIterationEventsService)(nil).ListIssueIterationEvents), varargs...)
}
