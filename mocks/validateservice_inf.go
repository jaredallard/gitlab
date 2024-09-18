// Code generated by MockGen. DO NOT EDIT.
// Source: validateservice_inf.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=validateservice_inf.go -destination=mocks/validateservice_inf.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gitlab "github.com/xanzy/go-gitlab"
	gomock "go.uber.org/mock/gomock"
)

// MockValidateService is a mock of ValidateService interface.
type MockValidateService struct {
	ctrl     *gomock.Controller
	recorder *MockValidateServiceMockRecorder
}

// MockValidateServiceMockRecorder is the mock recorder for MockValidateService.
type MockValidateServiceMockRecorder struct {
	mock *MockValidateService
}

// NewMockValidateService creates a new mock instance.
func NewMockValidateService(ctrl *gomock.Controller) *MockValidateService {
	mock := &MockValidateService{ctrl: ctrl}
	mock.recorder = &MockValidateServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockValidateService) EXPECT() *MockValidateServiceMockRecorder {
	return m.recorder
}

// Lint mocks base method.
func (m *MockValidateService) Lint(opts *gitlab.LintOptions, options ...gitlab.RequestOptionFunc) (*gitlab.LintResult, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{opts}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Lint", varargs...)
	ret0, _ := ret[0].(*gitlab.LintResult)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Lint indicates an expected call of Lint.
func (mr *MockValidateServiceMockRecorder) Lint(opts any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{opts}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lint", reflect.TypeOf((*MockValidateService)(nil).Lint), varargs...)
}

// ProjectLint mocks base method.
func (m *MockValidateService) ProjectLint(pid any, opt *gitlab.ProjectLintOptions, options ...gitlab.RequestOptionFunc) (*gitlab.ProjectLintResult, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ProjectLint", varargs...)
	ret0, _ := ret[0].(*gitlab.ProjectLintResult)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ProjectLint indicates an expected call of ProjectLint.
func (mr *MockValidateServiceMockRecorder) ProjectLint(pid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectLint", reflect.TypeOf((*MockValidateService)(nil).ProjectLint), varargs...)
}

// ProjectNamespaceLint mocks base method.
func (m *MockValidateService) ProjectNamespaceLint(pid any, opt *gitlab.ProjectNamespaceLintOptions, options ...gitlab.RequestOptionFunc) (*gitlab.ProjectLintResult, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ProjectNamespaceLint", varargs...)
	ret0, _ := ret[0].(*gitlab.ProjectLintResult)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ProjectNamespaceLint indicates an expected call of ProjectNamespaceLint.
func (mr *MockValidateServiceMockRecorder) ProjectNamespaceLint(pid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProjectNamespaceLint", reflect.TypeOf((*MockValidateService)(nil).ProjectNamespaceLint), varargs...)
}