// Code generated by MockGen. DO NOT EDIT.
// Source: resourcelabeleventsservice_inf.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=resourcelabeleventsservice_inf.go -destination=mocks/resourcelabeleventsservice_inf.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gitlab "gitlab.com/gitlab-org/api/client-go"
	gomock "go.uber.org/mock/gomock"
)

// MockResourceLabelEventsService is a mock of ResourceLabelEventsService interface.
type MockResourceLabelEventsService struct {
	ctrl     *gomock.Controller
	recorder *MockResourceLabelEventsServiceMockRecorder
	isgomock struct{}
}

// MockResourceLabelEventsServiceMockRecorder is the mock recorder for MockResourceLabelEventsService.
type MockResourceLabelEventsServiceMockRecorder struct {
	mock *MockResourceLabelEventsService
}

// NewMockResourceLabelEventsService creates a new mock instance.
func NewMockResourceLabelEventsService(ctrl *gomock.Controller) *MockResourceLabelEventsService {
	mock := &MockResourceLabelEventsService{ctrl: ctrl}
	mock.recorder = &MockResourceLabelEventsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResourceLabelEventsService) EXPECT() *MockResourceLabelEventsServiceMockRecorder {
	return m.recorder
}

// GetGroupEpicLabelEvent mocks base method.
func (m *MockResourceLabelEventsService) GetGroupEpicLabelEvent(gid any, epic, event int, options ...gitlab.RequestOptionFunc) (*gitlab.LabelEvent, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{gid, epic, event}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetGroupEpicLabelEvent", varargs...)
	ret0, _ := ret[0].(*gitlab.LabelEvent)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetGroupEpicLabelEvent indicates an expected call of GetGroupEpicLabelEvent.
func (mr *MockResourceLabelEventsServiceMockRecorder) GetGroupEpicLabelEvent(gid, epic, event any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{gid, epic, event}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupEpicLabelEvent", reflect.TypeOf((*MockResourceLabelEventsService)(nil).GetGroupEpicLabelEvent), varargs...)
}

// GetIssueLabelEvent mocks base method.
func (m *MockResourceLabelEventsService) GetIssueLabelEvent(pid any, issue, event int, options ...gitlab.RequestOptionFunc) (*gitlab.LabelEvent, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, issue, event}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetIssueLabelEvent", varargs...)
	ret0, _ := ret[0].(*gitlab.LabelEvent)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetIssueLabelEvent indicates an expected call of GetIssueLabelEvent.
func (mr *MockResourceLabelEventsServiceMockRecorder) GetIssueLabelEvent(pid, issue, event any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, issue, event}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIssueLabelEvent", reflect.TypeOf((*MockResourceLabelEventsService)(nil).GetIssueLabelEvent), varargs...)
}

// GetMergeRequestLabelEvent mocks base method.
func (m *MockResourceLabelEventsService) GetMergeRequestLabelEvent(pid any, request, event int, options ...gitlab.RequestOptionFunc) (*gitlab.LabelEvent, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, request, event}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMergeRequestLabelEvent", varargs...)
	ret0, _ := ret[0].(*gitlab.LabelEvent)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetMergeRequestLabelEvent indicates an expected call of GetMergeRequestLabelEvent.
func (mr *MockResourceLabelEventsServiceMockRecorder) GetMergeRequestLabelEvent(pid, request, event any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, request, event}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMergeRequestLabelEvent", reflect.TypeOf((*MockResourceLabelEventsService)(nil).GetMergeRequestLabelEvent), varargs...)
}

// ListGroupEpicLabelEvents mocks base method.
func (m *MockResourceLabelEventsService) ListGroupEpicLabelEvents(gid any, epic int, opt *gitlab.ListLabelEventsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.LabelEvent, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{gid, epic, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListGroupEpicLabelEvents", varargs...)
	ret0, _ := ret[0].([]*gitlab.LabelEvent)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListGroupEpicLabelEvents indicates an expected call of ListGroupEpicLabelEvents.
func (mr *MockResourceLabelEventsServiceMockRecorder) ListGroupEpicLabelEvents(gid, epic, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{gid, epic, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGroupEpicLabelEvents", reflect.TypeOf((*MockResourceLabelEventsService)(nil).ListGroupEpicLabelEvents), varargs...)
}

// ListIssueLabelEvents mocks base method.
func (m *MockResourceLabelEventsService) ListIssueLabelEvents(pid any, issue int, opt *gitlab.ListLabelEventsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.LabelEvent, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, issue, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListIssueLabelEvents", varargs...)
	ret0, _ := ret[0].([]*gitlab.LabelEvent)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListIssueLabelEvents indicates an expected call of ListIssueLabelEvents.
func (mr *MockResourceLabelEventsServiceMockRecorder) ListIssueLabelEvents(pid, issue, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, issue, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListIssueLabelEvents", reflect.TypeOf((*MockResourceLabelEventsService)(nil).ListIssueLabelEvents), varargs...)
}

// ListMergeRequestsLabelEvents mocks base method.
func (m *MockResourceLabelEventsService) ListMergeRequestsLabelEvents(pid any, request int, opt *gitlab.ListLabelEventsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.LabelEvent, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, request, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListMergeRequestsLabelEvents", varargs...)
	ret0, _ := ret[0].([]*gitlab.LabelEvent)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListMergeRequestsLabelEvents indicates an expected call of ListMergeRequestsLabelEvents.
func (mr *MockResourceLabelEventsServiceMockRecorder) ListMergeRequestsLabelEvents(pid, request, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, request, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMergeRequestsLabelEvents", reflect.TypeOf((*MockResourceLabelEventsService)(nil).ListMergeRequestsLabelEvents), varargs...)
}
