// Code generated by MockGen. DO NOT EDIT.
// Source: repositorysubmodulesservice_inf.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=repositorysubmodulesservice_inf.go -destination=mocks/repositorysubmodulesservice_inf.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gitlab "github.com/xanzy/go-gitlab"
	gomock "go.uber.org/mock/gomock"
)

// MockRepositorySubmodulesService is a mock of RepositorySubmodulesService interface.
type MockRepositorySubmodulesService struct {
	ctrl     *gomock.Controller
	recorder *MockRepositorySubmodulesServiceMockRecorder
	isgomock struct{}
}

// MockRepositorySubmodulesServiceMockRecorder is the mock recorder for MockRepositorySubmodulesService.
type MockRepositorySubmodulesServiceMockRecorder struct {
	mock *MockRepositorySubmodulesService
}

// NewMockRepositorySubmodulesService creates a new mock instance.
func NewMockRepositorySubmodulesService(ctrl *gomock.Controller) *MockRepositorySubmodulesService {
	mock := &MockRepositorySubmodulesService{ctrl: ctrl}
	mock.recorder = &MockRepositorySubmodulesServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositorySubmodulesService) EXPECT() *MockRepositorySubmodulesServiceMockRecorder {
	return m.recorder
}

// UpdateSubmodule mocks base method.
func (m *MockRepositorySubmodulesService) UpdateSubmodule(pid any, submodule string, opt *gitlab.UpdateSubmoduleOptions, options ...gitlab.RequestOptionFunc) (*gitlab.SubmoduleCommit, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, submodule, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateSubmodule", varargs...)
	ret0, _ := ret[0].(*gitlab.SubmoduleCommit)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UpdateSubmodule indicates an expected call of UpdateSubmodule.
func (mr *MockRepositorySubmodulesServiceMockRecorder) UpdateSubmodule(pid, submodule, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, submodule, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSubmodule", reflect.TypeOf((*MockRepositorySubmodulesService)(nil).UpdateSubmodule), varargs...)
}
