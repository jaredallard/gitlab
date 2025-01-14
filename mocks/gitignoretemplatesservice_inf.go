// Code generated by MockGen. DO NOT EDIT.
// Source: gitignoretemplatesservice_inf.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=gitignoretemplatesservice_inf.go -destination=mocks/gitignoretemplatesservice_inf.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gitlab "gitlab.com/gitlab-org/api/client-go"
	gomock "go.uber.org/mock/gomock"
)

// MockGitIgnoreTemplatesService is a mock of GitIgnoreTemplatesService interface.
type MockGitIgnoreTemplatesService struct {
	ctrl     *gomock.Controller
	recorder *MockGitIgnoreTemplatesServiceMockRecorder
	isgomock struct{}
}

// MockGitIgnoreTemplatesServiceMockRecorder is the mock recorder for MockGitIgnoreTemplatesService.
type MockGitIgnoreTemplatesServiceMockRecorder struct {
	mock *MockGitIgnoreTemplatesService
}

// NewMockGitIgnoreTemplatesService creates a new mock instance.
func NewMockGitIgnoreTemplatesService(ctrl *gomock.Controller) *MockGitIgnoreTemplatesService {
	mock := &MockGitIgnoreTemplatesService{ctrl: ctrl}
	mock.recorder = &MockGitIgnoreTemplatesServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGitIgnoreTemplatesService) EXPECT() *MockGitIgnoreTemplatesServiceMockRecorder {
	return m.recorder
}

// GetTemplate mocks base method.
func (m *MockGitIgnoreTemplatesService) GetTemplate(key string, options ...gitlab.RequestOptionFunc) (*gitlab.GitIgnoreTemplate, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{key}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTemplate", varargs...)
	ret0, _ := ret[0].(*gitlab.GitIgnoreTemplate)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTemplate indicates an expected call of GetTemplate.
func (mr *MockGitIgnoreTemplatesServiceMockRecorder) GetTemplate(key any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{key}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTemplate", reflect.TypeOf((*MockGitIgnoreTemplatesService)(nil).GetTemplate), varargs...)
}

// ListTemplates mocks base method.
func (m *MockGitIgnoreTemplatesService) ListTemplates(opt *gitlab.ListTemplatesOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.GitIgnoreTemplateListItem, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTemplates", varargs...)
	ret0, _ := ret[0].([]*gitlab.GitIgnoreTemplateListItem)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListTemplates indicates an expected call of ListTemplates.
func (mr *MockGitIgnoreTemplatesServiceMockRecorder) ListTemplates(opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTemplates", reflect.TypeOf((*MockGitIgnoreTemplatesService)(nil).ListTemplates), varargs...)
}
