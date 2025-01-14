// Code generated by MockGen. DO NOT EDIT.
// Source: runnersservice_inf.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=runnersservice_inf.go -destination=mocks/runnersservice_inf.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gitlab "gitlab.com/gitlab-org/api/client-go"
	gomock "go.uber.org/mock/gomock"
)

// MockRunnersService is a mock of RunnersService interface.
type MockRunnersService struct {
	ctrl     *gomock.Controller
	recorder *MockRunnersServiceMockRecorder
	isgomock struct{}
}

// MockRunnersServiceMockRecorder is the mock recorder for MockRunnersService.
type MockRunnersServiceMockRecorder struct {
	mock *MockRunnersService
}

// NewMockRunnersService creates a new mock instance.
func NewMockRunnersService(ctrl *gomock.Controller) *MockRunnersService {
	mock := &MockRunnersService{ctrl: ctrl}
	mock.recorder = &MockRunnersServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRunnersService) EXPECT() *MockRunnersServiceMockRecorder {
	return m.recorder
}

// DeleteRegisteredRunner mocks base method.
func (m *MockRunnersService) DeleteRegisteredRunner(opt *gitlab.DeleteRegisteredRunnerOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteRegisteredRunner", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteRegisteredRunner indicates an expected call of DeleteRegisteredRunner.
func (mr *MockRunnersServiceMockRecorder) DeleteRegisteredRunner(opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRegisteredRunner", reflect.TypeOf((*MockRunnersService)(nil).DeleteRegisteredRunner), varargs...)
}

// DeleteRegisteredRunnerByID mocks base method.
func (m *MockRunnersService) DeleteRegisteredRunnerByID(rid int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{rid}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteRegisteredRunnerByID", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteRegisteredRunnerByID indicates an expected call of DeleteRegisteredRunnerByID.
func (mr *MockRunnersServiceMockRecorder) DeleteRegisteredRunnerByID(rid any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{rid}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRegisteredRunnerByID", reflect.TypeOf((*MockRunnersService)(nil).DeleteRegisteredRunnerByID), varargs...)
}

// DisableProjectRunner mocks base method.
func (m *MockRunnersService) DisableProjectRunner(pid any, runner int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, runner}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DisableProjectRunner", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DisableProjectRunner indicates an expected call of DisableProjectRunner.
func (mr *MockRunnersServiceMockRecorder) DisableProjectRunner(pid, runner any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, runner}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableProjectRunner", reflect.TypeOf((*MockRunnersService)(nil).DisableProjectRunner), varargs...)
}

// EnableProjectRunner mocks base method.
func (m *MockRunnersService) EnableProjectRunner(pid any, opt *gitlab.EnableProjectRunnerOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Runner, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EnableProjectRunner", varargs...)
	ret0, _ := ret[0].(*gitlab.Runner)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// EnableProjectRunner indicates an expected call of EnableProjectRunner.
func (mr *MockRunnersServiceMockRecorder) EnableProjectRunner(pid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableProjectRunner", reflect.TypeOf((*MockRunnersService)(nil).EnableProjectRunner), varargs...)
}

// GetRunnerDetails mocks base method.
func (m *MockRunnersService) GetRunnerDetails(rid any, options ...gitlab.RequestOptionFunc) (*gitlab.RunnerDetails, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{rid}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRunnerDetails", varargs...)
	ret0, _ := ret[0].(*gitlab.RunnerDetails)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetRunnerDetails indicates an expected call of GetRunnerDetails.
func (mr *MockRunnersServiceMockRecorder) GetRunnerDetails(rid any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{rid}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRunnerDetails", reflect.TypeOf((*MockRunnersService)(nil).GetRunnerDetails), varargs...)
}

// ListAllRunners mocks base method.
func (m *MockRunnersService) ListAllRunners(opt *gitlab.ListRunnersOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.Runner, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAllRunners", varargs...)
	ret0, _ := ret[0].([]*gitlab.Runner)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListAllRunners indicates an expected call of ListAllRunners.
func (mr *MockRunnersServiceMockRecorder) ListAllRunners(opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllRunners", reflect.TypeOf((*MockRunnersService)(nil).ListAllRunners), varargs...)
}

// ListGroupsRunners mocks base method.
func (m *MockRunnersService) ListGroupsRunners(gid any, opt *gitlab.ListGroupsRunnersOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.Runner, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{gid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListGroupsRunners", varargs...)
	ret0, _ := ret[0].([]*gitlab.Runner)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListGroupsRunners indicates an expected call of ListGroupsRunners.
func (mr *MockRunnersServiceMockRecorder) ListGroupsRunners(gid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{gid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGroupsRunners", reflect.TypeOf((*MockRunnersService)(nil).ListGroupsRunners), varargs...)
}

// ListProjectRunners mocks base method.
func (m *MockRunnersService) ListProjectRunners(pid any, opt *gitlab.ListProjectRunnersOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.Runner, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListProjectRunners", varargs...)
	ret0, _ := ret[0].([]*gitlab.Runner)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListProjectRunners indicates an expected call of ListProjectRunners.
func (mr *MockRunnersServiceMockRecorder) ListProjectRunners(pid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjectRunners", reflect.TypeOf((*MockRunnersService)(nil).ListProjectRunners), varargs...)
}

// ListRunnerJobs mocks base method.
func (m *MockRunnersService) ListRunnerJobs(rid any, opt *gitlab.ListRunnerJobsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.Job, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{rid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRunnerJobs", varargs...)
	ret0, _ := ret[0].([]*gitlab.Job)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListRunnerJobs indicates an expected call of ListRunnerJobs.
func (mr *MockRunnersServiceMockRecorder) ListRunnerJobs(rid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{rid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRunnerJobs", reflect.TypeOf((*MockRunnersService)(nil).ListRunnerJobs), varargs...)
}

// ListRunners mocks base method.
func (m *MockRunnersService) ListRunners(opt *gitlab.ListRunnersOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.Runner, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRunners", varargs...)
	ret0, _ := ret[0].([]*gitlab.Runner)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListRunners indicates an expected call of ListRunners.
func (mr *MockRunnersServiceMockRecorder) ListRunners(opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRunners", reflect.TypeOf((*MockRunnersService)(nil).ListRunners), varargs...)
}

// RegisterNewRunner mocks base method.
func (m *MockRunnersService) RegisterNewRunner(opt *gitlab.RegisterNewRunnerOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Runner, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RegisterNewRunner", varargs...)
	ret0, _ := ret[0].(*gitlab.Runner)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RegisterNewRunner indicates an expected call of RegisterNewRunner.
func (mr *MockRunnersServiceMockRecorder) RegisterNewRunner(opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterNewRunner", reflect.TypeOf((*MockRunnersService)(nil).RegisterNewRunner), varargs...)
}

// RemoveRunner mocks base method.
func (m *MockRunnersService) RemoveRunner(rid any, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{rid}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RemoveRunner", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveRunner indicates an expected call of RemoveRunner.
func (mr *MockRunnersServiceMockRecorder) RemoveRunner(rid any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{rid}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveRunner", reflect.TypeOf((*MockRunnersService)(nil).RemoveRunner), varargs...)
}

// ResetGroupRunnerRegistrationToken mocks base method.
func (m *MockRunnersService) ResetGroupRunnerRegistrationToken(gid any, options ...gitlab.RequestOptionFunc) (*gitlab.RunnerRegistrationToken, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{gid}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ResetGroupRunnerRegistrationToken", varargs...)
	ret0, _ := ret[0].(*gitlab.RunnerRegistrationToken)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ResetGroupRunnerRegistrationToken indicates an expected call of ResetGroupRunnerRegistrationToken.
func (mr *MockRunnersServiceMockRecorder) ResetGroupRunnerRegistrationToken(gid any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{gid}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetGroupRunnerRegistrationToken", reflect.TypeOf((*MockRunnersService)(nil).ResetGroupRunnerRegistrationToken), varargs...)
}

// ResetInstanceRunnerRegistrationToken mocks base method.
func (m *MockRunnersService) ResetInstanceRunnerRegistrationToken(options ...gitlab.RequestOptionFunc) (*gitlab.RunnerRegistrationToken, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ResetInstanceRunnerRegistrationToken", varargs...)
	ret0, _ := ret[0].(*gitlab.RunnerRegistrationToken)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ResetInstanceRunnerRegistrationToken indicates an expected call of ResetInstanceRunnerRegistrationToken.
func (mr *MockRunnersServiceMockRecorder) ResetInstanceRunnerRegistrationToken(options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetInstanceRunnerRegistrationToken", reflect.TypeOf((*MockRunnersService)(nil).ResetInstanceRunnerRegistrationToken), options...)
}

// ResetProjectRunnerRegistrationToken mocks base method.
func (m *MockRunnersService) ResetProjectRunnerRegistrationToken(pid any, options ...gitlab.RequestOptionFunc) (*gitlab.RunnerRegistrationToken, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ResetProjectRunnerRegistrationToken", varargs...)
	ret0, _ := ret[0].(*gitlab.RunnerRegistrationToken)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ResetProjectRunnerRegistrationToken indicates an expected call of ResetProjectRunnerRegistrationToken.
func (mr *MockRunnersServiceMockRecorder) ResetProjectRunnerRegistrationToken(pid any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetProjectRunnerRegistrationToken", reflect.TypeOf((*MockRunnersService)(nil).ResetProjectRunnerRegistrationToken), varargs...)
}

// ResetRunnerAuthenticationToken mocks base method.
func (m *MockRunnersService) ResetRunnerAuthenticationToken(rid int, options ...gitlab.RequestOptionFunc) (*gitlab.RunnerAuthenticationToken, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{rid}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ResetRunnerAuthenticationToken", varargs...)
	ret0, _ := ret[0].(*gitlab.RunnerAuthenticationToken)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ResetRunnerAuthenticationToken indicates an expected call of ResetRunnerAuthenticationToken.
func (mr *MockRunnersServiceMockRecorder) ResetRunnerAuthenticationToken(rid any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{rid}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetRunnerAuthenticationToken", reflect.TypeOf((*MockRunnersService)(nil).ResetRunnerAuthenticationToken), varargs...)
}

// UpdateRunnerDetails mocks base method.
func (m *MockRunnersService) UpdateRunnerDetails(rid any, opt *gitlab.UpdateRunnerDetailsOptions, options ...gitlab.RequestOptionFunc) (*gitlab.RunnerDetails, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{rid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateRunnerDetails", varargs...)
	ret0, _ := ret[0].(*gitlab.RunnerDetails)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UpdateRunnerDetails indicates an expected call of UpdateRunnerDetails.
func (mr *MockRunnersServiceMockRecorder) UpdateRunnerDetails(rid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{rid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRunnerDetails", reflect.TypeOf((*MockRunnersService)(nil).UpdateRunnerDetails), varargs...)
}

// VerifyRegisteredRunner mocks base method.
func (m *MockRunnersService) VerifyRegisteredRunner(opt *gitlab.VerifyRegisteredRunnerOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "VerifyRegisteredRunner", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyRegisteredRunner indicates an expected call of VerifyRegisteredRunner.
func (mr *MockRunnersServiceMockRecorder) VerifyRegisteredRunner(opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyRegisteredRunner", reflect.TypeOf((*MockRunnersService)(nil).VerifyRegisteredRunner), varargs...)
}
