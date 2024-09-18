// Code generated by MockGen. DO NOT EDIT.
// Source: issuelinksservice_inf.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=issuelinksservice_inf.go -destination=mocks/issuelinksservice_inf.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gitlab "github.com/xanzy/go-gitlab"
	gomock "go.uber.org/mock/gomock"
)

// MockIssueLinksService is a mock of IssueLinksService interface.
type MockIssueLinksService struct {
	ctrl     *gomock.Controller
	recorder *MockIssueLinksServiceMockRecorder
}

// MockIssueLinksServiceMockRecorder is the mock recorder for MockIssueLinksService.
type MockIssueLinksServiceMockRecorder struct {
	mock *MockIssueLinksService
}

// NewMockIssueLinksService creates a new mock instance.
func NewMockIssueLinksService(ctrl *gomock.Controller) *MockIssueLinksService {
	mock := &MockIssueLinksService{ctrl: ctrl}
	mock.recorder = &MockIssueLinksServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIssueLinksService) EXPECT() *MockIssueLinksServiceMockRecorder {
	return m.recorder
}

// CreateIssueLink mocks base method.
func (m *MockIssueLinksService) CreateIssueLink(pid any, issue int, opt *gitlab.CreateIssueLinkOptions, options ...gitlab.RequestOptionFunc) (*gitlab.IssueLink, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, issue, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateIssueLink", varargs...)
	ret0, _ := ret[0].(*gitlab.IssueLink)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateIssueLink indicates an expected call of CreateIssueLink.
func (mr *MockIssueLinksServiceMockRecorder) CreateIssueLink(pid, issue, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, issue, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateIssueLink", reflect.TypeOf((*MockIssueLinksService)(nil).CreateIssueLink), varargs...)
}

// DeleteIssueLink mocks base method.
func (m *MockIssueLinksService) DeleteIssueLink(pid any, issue, issueLink int, options ...gitlab.RequestOptionFunc) (*gitlab.IssueLink, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, issue, issueLink}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteIssueLink", varargs...)
	ret0, _ := ret[0].(*gitlab.IssueLink)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// DeleteIssueLink indicates an expected call of DeleteIssueLink.
func (mr *MockIssueLinksServiceMockRecorder) DeleteIssueLink(pid, issue, issueLink any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, issue, issueLink}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteIssueLink", reflect.TypeOf((*MockIssueLinksService)(nil).DeleteIssueLink), varargs...)
}

// GetIssueLink mocks base method.
func (m *MockIssueLinksService) GetIssueLink(pid any, issue, issueLink int, options ...gitlab.RequestOptionFunc) (*gitlab.IssueLink, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, issue, issueLink}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetIssueLink", varargs...)
	ret0, _ := ret[0].(*gitlab.IssueLink)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetIssueLink indicates an expected call of GetIssueLink.
func (mr *MockIssueLinksServiceMockRecorder) GetIssueLink(pid, issue, issueLink any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, issue, issueLink}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIssueLink", reflect.TypeOf((*MockIssueLinksService)(nil).GetIssueLink), varargs...)
}

// ListIssueRelations mocks base method.
func (m *MockIssueLinksService) ListIssueRelations(pid any, issue int, options ...gitlab.RequestOptionFunc) ([]*gitlab.IssueRelation, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, issue}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListIssueRelations", varargs...)
	ret0, _ := ret[0].([]*gitlab.IssueRelation)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListIssueRelations indicates an expected call of ListIssueRelations.
func (mr *MockIssueLinksServiceMockRecorder) ListIssueRelations(pid, issue any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, issue}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListIssueRelations", reflect.TypeOf((*MockIssueLinksService)(nil).ListIssueRelations), varargs...)
}
