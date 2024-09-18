// Code generated by MockGen. DO NOT EDIT.
// Source: resourceweighteventsservice_inf.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=resourceweighteventsservice_inf.go -destination=mocks/resourceweighteventsservice_inf.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gitlab "github.com/xanzy/go-gitlab"
	gomock "go.uber.org/mock/gomock"
)

// MockResourceWeightEventsService is a mock of ResourceWeightEventsService interface.
type MockResourceWeightEventsService struct {
	ctrl     *gomock.Controller
	recorder *MockResourceWeightEventsServiceMockRecorder
}

// MockResourceWeightEventsServiceMockRecorder is the mock recorder for MockResourceWeightEventsService.
type MockResourceWeightEventsServiceMockRecorder struct {
	mock *MockResourceWeightEventsService
}

// NewMockResourceWeightEventsService creates a new mock instance.
func NewMockResourceWeightEventsService(ctrl *gomock.Controller) *MockResourceWeightEventsService {
	mock := &MockResourceWeightEventsService{ctrl: ctrl}
	mock.recorder = &MockResourceWeightEventsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResourceWeightEventsService) EXPECT() *MockResourceWeightEventsServiceMockRecorder {
	return m.recorder
}

// ListIssueWeightEvents mocks base method.
func (m *MockResourceWeightEventsService) ListIssueWeightEvents(pid any, issue int, opt *gitlab.ListWeightEventsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.WeightEvent, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, issue, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListIssueWeightEvents", varargs...)
	ret0, _ := ret[0].([]*gitlab.WeightEvent)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListIssueWeightEvents indicates an expected call of ListIssueWeightEvents.
func (mr *MockResourceWeightEventsServiceMockRecorder) ListIssueWeightEvents(pid, issue, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, issue, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListIssueWeightEvents", reflect.TypeOf((*MockResourceWeightEventsService)(nil).ListIssueWeightEvents), varargs...)
}