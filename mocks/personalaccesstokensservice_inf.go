// Code generated by MockGen. DO NOT EDIT.
// Source: personalaccesstokensservice_inf.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=personalaccesstokensservice_inf.go -destination=mocks/personalaccesstokensservice_inf.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gitlab "gitlab.com/gitlab-org/api/client-go"
	gomock "go.uber.org/mock/gomock"
)

// MockPersonalAccessTokensService is a mock of PersonalAccessTokensService interface.
type MockPersonalAccessTokensService struct {
	ctrl     *gomock.Controller
	recorder *MockPersonalAccessTokensServiceMockRecorder
	isgomock struct{}
}

// MockPersonalAccessTokensServiceMockRecorder is the mock recorder for MockPersonalAccessTokensService.
type MockPersonalAccessTokensServiceMockRecorder struct {
	mock *MockPersonalAccessTokensService
}

// NewMockPersonalAccessTokensService creates a new mock instance.
func NewMockPersonalAccessTokensService(ctrl *gomock.Controller) *MockPersonalAccessTokensService {
	mock := &MockPersonalAccessTokensService{ctrl: ctrl}
	mock.recorder = &MockPersonalAccessTokensServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPersonalAccessTokensService) EXPECT() *MockPersonalAccessTokensServiceMockRecorder {
	return m.recorder
}

// GetSinglePersonalAccessToken mocks base method.
func (m *MockPersonalAccessTokensService) GetSinglePersonalAccessToken(options ...gitlab.RequestOptionFunc) (*gitlab.PersonalAccessToken, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSinglePersonalAccessToken", varargs...)
	ret0, _ := ret[0].(*gitlab.PersonalAccessToken)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetSinglePersonalAccessToken indicates an expected call of GetSinglePersonalAccessToken.
func (mr *MockPersonalAccessTokensServiceMockRecorder) GetSinglePersonalAccessToken(options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSinglePersonalAccessToken", reflect.TypeOf((*MockPersonalAccessTokensService)(nil).GetSinglePersonalAccessToken), options...)
}

// GetSinglePersonalAccessTokenByID mocks base method.
func (m *MockPersonalAccessTokensService) GetSinglePersonalAccessTokenByID(token int, options ...gitlab.RequestOptionFunc) (*gitlab.PersonalAccessToken, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{token}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSinglePersonalAccessTokenByID", varargs...)
	ret0, _ := ret[0].(*gitlab.PersonalAccessToken)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetSinglePersonalAccessTokenByID indicates an expected call of GetSinglePersonalAccessTokenByID.
func (mr *MockPersonalAccessTokensServiceMockRecorder) GetSinglePersonalAccessTokenByID(token any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{token}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSinglePersonalAccessTokenByID", reflect.TypeOf((*MockPersonalAccessTokensService)(nil).GetSinglePersonalAccessTokenByID), varargs...)
}

// ListPersonalAccessTokens mocks base method.
func (m *MockPersonalAccessTokensService) ListPersonalAccessTokens(opt *gitlab.ListPersonalAccessTokensOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.PersonalAccessToken, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListPersonalAccessTokens", varargs...)
	ret0, _ := ret[0].([]*gitlab.PersonalAccessToken)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListPersonalAccessTokens indicates an expected call of ListPersonalAccessTokens.
func (mr *MockPersonalAccessTokensServiceMockRecorder) ListPersonalAccessTokens(opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPersonalAccessTokens", reflect.TypeOf((*MockPersonalAccessTokensService)(nil).ListPersonalAccessTokens), varargs...)
}

// RevokePersonalAccessToken mocks base method.
func (m *MockPersonalAccessTokensService) RevokePersonalAccessToken(token int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{token}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RevokePersonalAccessToken", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RevokePersonalAccessToken indicates an expected call of RevokePersonalAccessToken.
func (mr *MockPersonalAccessTokensServiceMockRecorder) RevokePersonalAccessToken(token any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{token}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokePersonalAccessToken", reflect.TypeOf((*MockPersonalAccessTokensService)(nil).RevokePersonalAccessToken), varargs...)
}

// RevokePersonalAccessTokenByID mocks base method.
func (m *MockPersonalAccessTokensService) RevokePersonalAccessTokenByID(token int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{token}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RevokePersonalAccessTokenByID", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RevokePersonalAccessTokenByID indicates an expected call of RevokePersonalAccessTokenByID.
func (mr *MockPersonalAccessTokensServiceMockRecorder) RevokePersonalAccessTokenByID(token any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{token}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokePersonalAccessTokenByID", reflect.TypeOf((*MockPersonalAccessTokensService)(nil).RevokePersonalAccessTokenByID), varargs...)
}

// RevokePersonalAccessTokenSelf mocks base method.
func (m *MockPersonalAccessTokensService) RevokePersonalAccessTokenSelf(options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RevokePersonalAccessTokenSelf", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RevokePersonalAccessTokenSelf indicates an expected call of RevokePersonalAccessTokenSelf.
func (mr *MockPersonalAccessTokensServiceMockRecorder) RevokePersonalAccessTokenSelf(options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokePersonalAccessTokenSelf", reflect.TypeOf((*MockPersonalAccessTokensService)(nil).RevokePersonalAccessTokenSelf), options...)
}

// RotatePersonalAccessToken mocks base method.
func (m *MockPersonalAccessTokensService) RotatePersonalAccessToken(token int, opt *gitlab.RotatePersonalAccessTokenOptions, options ...gitlab.RequestOptionFunc) (*gitlab.PersonalAccessToken, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{token, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RotatePersonalAccessToken", varargs...)
	ret0, _ := ret[0].(*gitlab.PersonalAccessToken)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RotatePersonalAccessToken indicates an expected call of RotatePersonalAccessToken.
func (mr *MockPersonalAccessTokensServiceMockRecorder) RotatePersonalAccessToken(token, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{token, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RotatePersonalAccessToken", reflect.TypeOf((*MockPersonalAccessTokensService)(nil).RotatePersonalAccessToken), varargs...)
}

// RotatePersonalAccessTokenByID mocks base method.
func (m *MockPersonalAccessTokensService) RotatePersonalAccessTokenByID(token int, opt *gitlab.RotatePersonalAccessTokenOptions, options ...gitlab.RequestOptionFunc) (*gitlab.PersonalAccessToken, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{token, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RotatePersonalAccessTokenByID", varargs...)
	ret0, _ := ret[0].(*gitlab.PersonalAccessToken)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RotatePersonalAccessTokenByID indicates an expected call of RotatePersonalAccessTokenByID.
func (mr *MockPersonalAccessTokensServiceMockRecorder) RotatePersonalAccessTokenByID(token, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{token, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RotatePersonalAccessTokenByID", reflect.TypeOf((*MockPersonalAccessTokensService)(nil).RotatePersonalAccessTokenByID), varargs...)
}

// RotatePersonalAccessTokenSelf mocks base method.
func (m *MockPersonalAccessTokensService) RotatePersonalAccessTokenSelf(opt *gitlab.RotatePersonalAccessTokenOptions, options ...gitlab.RequestOptionFunc) (*gitlab.PersonalAccessToken, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RotatePersonalAccessTokenSelf", varargs...)
	ret0, _ := ret[0].(*gitlab.PersonalAccessToken)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RotatePersonalAccessTokenSelf indicates an expected call of RotatePersonalAccessTokenSelf.
func (mr *MockPersonalAccessTokensServiceMockRecorder) RotatePersonalAccessTokenSelf(opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RotatePersonalAccessTokenSelf", reflect.TypeOf((*MockPersonalAccessTokensService)(nil).RotatePersonalAccessTokenSelf), varargs...)
}
