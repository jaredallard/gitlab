// Code generated by MockGen. DO NOT EDIT.
// Source: keysservice_inf.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=keysservice_inf.go -destination=mocks/keysservice_inf.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gitlab "github.com/xanzy/go-gitlab"
	gomock "go.uber.org/mock/gomock"
)

// MockKeysService is a mock of KeysService interface.
type MockKeysService struct {
	ctrl     *gomock.Controller
	recorder *MockKeysServiceMockRecorder
}

// MockKeysServiceMockRecorder is the mock recorder for MockKeysService.
type MockKeysServiceMockRecorder struct {
	mock *MockKeysService
}

// NewMockKeysService creates a new mock instance.
func NewMockKeysService(ctrl *gomock.Controller) *MockKeysService {
	mock := &MockKeysService{ctrl: ctrl}
	mock.recorder = &MockKeysServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKeysService) EXPECT() *MockKeysServiceMockRecorder {
	return m.recorder
}

// GetKeyByFingerprint mocks base method.
func (m *MockKeysService) GetKeyByFingerprint(opt *gitlab.GetKeyByFingerprintOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Key, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetKeyByFingerprint", varargs...)
	ret0, _ := ret[0].(*gitlab.Key)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetKeyByFingerprint indicates an expected call of GetKeyByFingerprint.
func (mr *MockKeysServiceMockRecorder) GetKeyByFingerprint(opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKeyByFingerprint", reflect.TypeOf((*MockKeysService)(nil).GetKeyByFingerprint), varargs...)
}

// GetKeyWithUser mocks base method.
func (m *MockKeysService) GetKeyWithUser(key int, options ...gitlab.RequestOptionFunc) (*gitlab.Key, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{key}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetKeyWithUser", varargs...)
	ret0, _ := ret[0].(*gitlab.Key)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetKeyWithUser indicates an expected call of GetKeyWithUser.
func (mr *MockKeysServiceMockRecorder) GetKeyWithUser(key any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{key}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKeyWithUser", reflect.TypeOf((*MockKeysService)(nil).GetKeyWithUser), varargs...)
}