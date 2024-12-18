// Code generated by MockGen. DO NOT EDIT.
// Source: deploykeysservice_inf.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=deploykeysservice_inf.go -destination=mocks/deploykeysservice_inf.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gitlab "gitlab.com/gitlab-org/api/client-go"
	gomock "go.uber.org/mock/gomock"
)

// MockDeployKeysService is a mock of DeployKeysService interface.
type MockDeployKeysService struct {
	ctrl     *gomock.Controller
	recorder *MockDeployKeysServiceMockRecorder
	isgomock struct{}
}

// MockDeployKeysServiceMockRecorder is the mock recorder for MockDeployKeysService.
type MockDeployKeysServiceMockRecorder struct {
	mock *MockDeployKeysService
}

// NewMockDeployKeysService creates a new mock instance.
func NewMockDeployKeysService(ctrl *gomock.Controller) *MockDeployKeysService {
	mock := &MockDeployKeysService{ctrl: ctrl}
	mock.recorder = &MockDeployKeysServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeployKeysService) EXPECT() *MockDeployKeysServiceMockRecorder {
	return m.recorder
}

// AddDeployKey mocks base method.
func (m *MockDeployKeysService) AddDeployKey(pid any, opt *gitlab.AddDeployKeyOptions, options ...gitlab.RequestOptionFunc) (*gitlab.ProjectDeployKey, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddDeployKey", varargs...)
	ret0, _ := ret[0].(*gitlab.ProjectDeployKey)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AddDeployKey indicates an expected call of AddDeployKey.
func (mr *MockDeployKeysServiceMockRecorder) AddDeployKey(pid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddDeployKey", reflect.TypeOf((*MockDeployKeysService)(nil).AddDeployKey), varargs...)
}

// DeleteDeployKey mocks base method.
func (m *MockDeployKeysService) DeleteDeployKey(pid any, deployKey int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, deployKey}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteDeployKey", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteDeployKey indicates an expected call of DeleteDeployKey.
func (mr *MockDeployKeysServiceMockRecorder) DeleteDeployKey(pid, deployKey any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, deployKey}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDeployKey", reflect.TypeOf((*MockDeployKeysService)(nil).DeleteDeployKey), varargs...)
}

// EnableDeployKey mocks base method.
func (m *MockDeployKeysService) EnableDeployKey(pid any, deployKey int, options ...gitlab.RequestOptionFunc) (*gitlab.ProjectDeployKey, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, deployKey}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EnableDeployKey", varargs...)
	ret0, _ := ret[0].(*gitlab.ProjectDeployKey)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// EnableDeployKey indicates an expected call of EnableDeployKey.
func (mr *MockDeployKeysServiceMockRecorder) EnableDeployKey(pid, deployKey any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, deployKey}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableDeployKey", reflect.TypeOf((*MockDeployKeysService)(nil).EnableDeployKey), varargs...)
}

// GetDeployKey mocks base method.
func (m *MockDeployKeysService) GetDeployKey(pid any, deployKey int, options ...gitlab.RequestOptionFunc) (*gitlab.ProjectDeployKey, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, deployKey}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDeployKey", varargs...)
	ret0, _ := ret[0].(*gitlab.ProjectDeployKey)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetDeployKey indicates an expected call of GetDeployKey.
func (mr *MockDeployKeysServiceMockRecorder) GetDeployKey(pid, deployKey any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, deployKey}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeployKey", reflect.TypeOf((*MockDeployKeysService)(nil).GetDeployKey), varargs...)
}

// ListAllDeployKeys mocks base method.
func (m *MockDeployKeysService) ListAllDeployKeys(opt *gitlab.ListInstanceDeployKeysOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.InstanceDeployKey, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAllDeployKeys", varargs...)
	ret0, _ := ret[0].([]*gitlab.InstanceDeployKey)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListAllDeployKeys indicates an expected call of ListAllDeployKeys.
func (mr *MockDeployKeysServiceMockRecorder) ListAllDeployKeys(opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllDeployKeys", reflect.TypeOf((*MockDeployKeysService)(nil).ListAllDeployKeys), varargs...)
}

// ListProjectDeployKeys mocks base method.
func (m *MockDeployKeysService) ListProjectDeployKeys(pid any, opt *gitlab.ListProjectDeployKeysOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.ProjectDeployKey, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListProjectDeployKeys", varargs...)
	ret0, _ := ret[0].([]*gitlab.ProjectDeployKey)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListProjectDeployKeys indicates an expected call of ListProjectDeployKeys.
func (mr *MockDeployKeysServiceMockRecorder) ListProjectDeployKeys(pid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjectDeployKeys", reflect.TypeOf((*MockDeployKeysService)(nil).ListProjectDeployKeys), varargs...)
}

// UpdateDeployKey mocks base method.
func (m *MockDeployKeysService) UpdateDeployKey(pid any, deployKey int, opt *gitlab.UpdateDeployKeyOptions, options ...gitlab.RequestOptionFunc) (*gitlab.ProjectDeployKey, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, deployKey, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateDeployKey", varargs...)
	ret0, _ := ret[0].(*gitlab.ProjectDeployKey)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UpdateDeployKey indicates an expected call of UpdateDeployKey.
func (mr *MockDeployKeysServiceMockRecorder) UpdateDeployKey(pid, deployKey, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, deployKey, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDeployKey", reflect.TypeOf((*MockDeployKeysService)(nil).UpdateDeployKey), varargs...)
}
