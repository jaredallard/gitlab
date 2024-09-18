// Code generated by MockGen. DO NOT EDIT.
// Source: snippetsservice_inf.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=snippetsservice_inf.go -destination=mocks/snippetsservice_inf.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gitlab "github.com/xanzy/go-gitlab"
	gomock "go.uber.org/mock/gomock"
)

// MockSnippetsService is a mock of SnippetsService interface.
type MockSnippetsService struct {
	ctrl     *gomock.Controller
	recorder *MockSnippetsServiceMockRecorder
}

// MockSnippetsServiceMockRecorder is the mock recorder for MockSnippetsService.
type MockSnippetsServiceMockRecorder struct {
	mock *MockSnippetsService
}

// NewMockSnippetsService creates a new mock instance.
func NewMockSnippetsService(ctrl *gomock.Controller) *MockSnippetsService {
	mock := &MockSnippetsService{ctrl: ctrl}
	mock.recorder = &MockSnippetsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSnippetsService) EXPECT() *MockSnippetsServiceMockRecorder {
	return m.recorder
}

// CreateSnippet mocks base method.
func (m *MockSnippetsService) CreateSnippet(opt *gitlab.CreateSnippetOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Snippet, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateSnippet", varargs...)
	ret0, _ := ret[0].(*gitlab.Snippet)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateSnippet indicates an expected call of CreateSnippet.
func (mr *MockSnippetsServiceMockRecorder) CreateSnippet(opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSnippet", reflect.TypeOf((*MockSnippetsService)(nil).CreateSnippet), varargs...)
}

// DeleteSnippet mocks base method.
func (m *MockSnippetsService) DeleteSnippet(snippet int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{snippet}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteSnippet", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteSnippet indicates an expected call of DeleteSnippet.
func (mr *MockSnippetsServiceMockRecorder) DeleteSnippet(snippet any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{snippet}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSnippet", reflect.TypeOf((*MockSnippetsService)(nil).DeleteSnippet), varargs...)
}

// ExploreSnippets mocks base method.
func (m *MockSnippetsService) ExploreSnippets(opt *gitlab.ExploreSnippetsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.Snippet, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ExploreSnippets", varargs...)
	ret0, _ := ret[0].([]*gitlab.Snippet)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ExploreSnippets indicates an expected call of ExploreSnippets.
func (mr *MockSnippetsServiceMockRecorder) ExploreSnippets(opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExploreSnippets", reflect.TypeOf((*MockSnippetsService)(nil).ExploreSnippets), varargs...)
}

// GetSnippet mocks base method.
func (m *MockSnippetsService) GetSnippet(snippet int, options ...gitlab.RequestOptionFunc) (*gitlab.Snippet, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{snippet}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSnippet", varargs...)
	ret0, _ := ret[0].(*gitlab.Snippet)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetSnippet indicates an expected call of GetSnippet.
func (mr *MockSnippetsServiceMockRecorder) GetSnippet(snippet any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{snippet}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSnippet", reflect.TypeOf((*MockSnippetsService)(nil).GetSnippet), varargs...)
}

// ListAllSnippets mocks base method.
func (m *MockSnippetsService) ListAllSnippets(opt *gitlab.ListAllSnippetsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.Snippet, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAllSnippets", varargs...)
	ret0, _ := ret[0].([]*gitlab.Snippet)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListAllSnippets indicates an expected call of ListAllSnippets.
func (mr *MockSnippetsServiceMockRecorder) ListAllSnippets(opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllSnippets", reflect.TypeOf((*MockSnippetsService)(nil).ListAllSnippets), varargs...)
}

// ListSnippets mocks base method.
func (m *MockSnippetsService) ListSnippets(opt *gitlab.ListSnippetsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.Snippet, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListSnippets", varargs...)
	ret0, _ := ret[0].([]*gitlab.Snippet)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListSnippets indicates an expected call of ListSnippets.
func (mr *MockSnippetsServiceMockRecorder) ListSnippets(opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSnippets", reflect.TypeOf((*MockSnippetsService)(nil).ListSnippets), varargs...)
}

// SnippetContent mocks base method.
func (m *MockSnippetsService) SnippetContent(snippet int, options ...gitlab.RequestOptionFunc) ([]byte, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{snippet}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SnippetContent", varargs...)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SnippetContent indicates an expected call of SnippetContent.
func (mr *MockSnippetsServiceMockRecorder) SnippetContent(snippet any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{snippet}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SnippetContent", reflect.TypeOf((*MockSnippetsService)(nil).SnippetContent), varargs...)
}

// SnippetFileContent mocks base method.
func (m *MockSnippetsService) SnippetFileContent(snippet int, ref, filename string, options ...gitlab.RequestOptionFunc) ([]byte, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{snippet, ref, filename}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SnippetFileContent", varargs...)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SnippetFileContent indicates an expected call of SnippetFileContent.
func (mr *MockSnippetsServiceMockRecorder) SnippetFileContent(snippet, ref, filename any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{snippet, ref, filename}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SnippetFileContent", reflect.TypeOf((*MockSnippetsService)(nil).SnippetFileContent), varargs...)
}

// UpdateSnippet mocks base method.
func (m *MockSnippetsService) UpdateSnippet(snippet int, opt *gitlab.UpdateSnippetOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Snippet, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{snippet, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateSnippet", varargs...)
	ret0, _ := ret[0].(*gitlab.Snippet)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UpdateSnippet indicates an expected call of UpdateSnippet.
func (mr *MockSnippetsServiceMockRecorder) UpdateSnippet(snippet, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{snippet, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSnippet", reflect.TypeOf((*MockSnippetsService)(nil).UpdateSnippet), varargs...)
}
