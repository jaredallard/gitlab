// Code generated by MockGen. DO NOT EDIT.
// Source: branchesservice_inf.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=branchesservice_inf.go -destination=mocks/branchesservice_inf.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gitlab "github.com/xanzy/go-gitlab"
	gomock "go.uber.org/mock/gomock"
)

// MockBranchesService is a mock of BranchesService interface.
type MockBranchesService struct {
	ctrl     *gomock.Controller
	recorder *MockBranchesServiceMockRecorder
	isgomock struct{}
}

// MockBranchesServiceMockRecorder is the mock recorder for MockBranchesService.
type MockBranchesServiceMockRecorder struct {
	mock *MockBranchesService
}

// NewMockBranchesService creates a new mock instance.
func NewMockBranchesService(ctrl *gomock.Controller) *MockBranchesService {
	mock := &MockBranchesService{ctrl: ctrl}
	mock.recorder = &MockBranchesServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBranchesService) EXPECT() *MockBranchesServiceMockRecorder {
	return m.recorder
}

// CreateBranch mocks base method.
func (m *MockBranchesService) CreateBranch(pid any, opt *gitlab.CreateBranchOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Branch, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateBranch", varargs...)
	ret0, _ := ret[0].(*gitlab.Branch)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateBranch indicates an expected call of CreateBranch.
func (mr *MockBranchesServiceMockRecorder) CreateBranch(pid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBranch", reflect.TypeOf((*MockBranchesService)(nil).CreateBranch), varargs...)
}

// DeleteBranch mocks base method.
func (m *MockBranchesService) DeleteBranch(pid any, branch string, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, branch}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteBranch", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteBranch indicates an expected call of DeleteBranch.
func (mr *MockBranchesServiceMockRecorder) DeleteBranch(pid, branch any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, branch}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBranch", reflect.TypeOf((*MockBranchesService)(nil).DeleteBranch), varargs...)
}

// DeleteMergedBranches mocks base method.
func (m *MockBranchesService) DeleteMergedBranches(pid any, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteMergedBranches", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteMergedBranches indicates an expected call of DeleteMergedBranches.
func (mr *MockBranchesServiceMockRecorder) DeleteMergedBranches(pid any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMergedBranches", reflect.TypeOf((*MockBranchesService)(nil).DeleteMergedBranches), varargs...)
}

// GetBranch mocks base method.
func (m *MockBranchesService) GetBranch(pid any, branch string, options ...gitlab.RequestOptionFunc) (*gitlab.Branch, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, branch}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBranch", varargs...)
	ret0, _ := ret[0].(*gitlab.Branch)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetBranch indicates an expected call of GetBranch.
func (mr *MockBranchesServiceMockRecorder) GetBranch(pid, branch any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, branch}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBranch", reflect.TypeOf((*MockBranchesService)(nil).GetBranch), varargs...)
}

// ListBranches mocks base method.
func (m *MockBranchesService) ListBranches(pid any, opts *gitlab.ListBranchesOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.Branch, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opts}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListBranches", varargs...)
	ret0, _ := ret[0].([]*gitlab.Branch)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListBranches indicates an expected call of ListBranches.
func (mr *MockBranchesServiceMockRecorder) ListBranches(pid, opts any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opts}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBranches", reflect.TypeOf((*MockBranchesService)(nil).ListBranches), varargs...)
}

// ProtectBranch mocks base method.
func (m *MockBranchesService) ProtectBranch(pid any, branch string, opts *gitlab.ProtectBranchOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Branch, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, branch, opts}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ProtectBranch", varargs...)
	ret0, _ := ret[0].(*gitlab.Branch)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ProtectBranch indicates an expected call of ProtectBranch.
func (mr *MockBranchesServiceMockRecorder) ProtectBranch(pid, branch, opts any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, branch, opts}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProtectBranch", reflect.TypeOf((*MockBranchesService)(nil).ProtectBranch), varargs...)
}

// UnprotectBranch mocks base method.
func (m *MockBranchesService) UnprotectBranch(pid any, branch string, options ...gitlab.RequestOptionFunc) (*gitlab.Branch, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, branch}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UnprotectBranch", varargs...)
	ret0, _ := ret[0].(*gitlab.Branch)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UnprotectBranch indicates an expected call of UnprotectBranch.
func (mr *MockBranchesServiceMockRecorder) UnprotectBranch(pid, branch any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, branch}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnprotectBranch", reflect.TypeOf((*MockBranchesService)(nil).UnprotectBranch), varargs...)
}
