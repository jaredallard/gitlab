// Code generated by MockGen. DO NOT EDIT.
// Source: namespacesservice_inf.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=namespacesservice_inf.go -destination=mocks/namespacesservice_inf.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gitlab "github.com/xanzy/go-gitlab"
	gomock "go.uber.org/mock/gomock"
)

// MockNamespacesService is a mock of NamespacesService interface.
type MockNamespacesService struct {
	ctrl     *gomock.Controller
	recorder *MockNamespacesServiceMockRecorder
}

// MockNamespacesServiceMockRecorder is the mock recorder for MockNamespacesService.
type MockNamespacesServiceMockRecorder struct {
	mock *MockNamespacesService
}

// NewMockNamespacesService creates a new mock instance.
func NewMockNamespacesService(ctrl *gomock.Controller) *MockNamespacesService {
	mock := &MockNamespacesService{ctrl: ctrl}
	mock.recorder = &MockNamespacesServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNamespacesService) EXPECT() *MockNamespacesServiceMockRecorder {
	return m.recorder
}

// GetNamespace mocks base method.
func (m *MockNamespacesService) GetNamespace(id any, options ...gitlab.RequestOptionFunc) (*gitlab.Namespace, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{id}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetNamespace", varargs...)
	ret0, _ := ret[0].(*gitlab.Namespace)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetNamespace indicates an expected call of GetNamespace.
func (mr *MockNamespacesServiceMockRecorder) GetNamespace(id any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{id}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNamespace", reflect.TypeOf((*MockNamespacesService)(nil).GetNamespace), varargs...)
}

// ListNamespaces mocks base method.
func (m *MockNamespacesService) ListNamespaces(opt *gitlab.ListNamespacesOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.Namespace, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListNamespaces", varargs...)
	ret0, _ := ret[0].([]*gitlab.Namespace)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListNamespaces indicates an expected call of ListNamespaces.
func (mr *MockNamespacesServiceMockRecorder) ListNamespaces(opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNamespaces", reflect.TypeOf((*MockNamespacesService)(nil).ListNamespaces), varargs...)
}

// NamespaceExists mocks base method.
func (m *MockNamespacesService) NamespaceExists(id any, opt *gitlab.NamespaceExistsOptions, options ...gitlab.RequestOptionFunc) (*gitlab.NamespaceExistance, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{id, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "NamespaceExists", varargs...)
	ret0, _ := ret[0].(*gitlab.NamespaceExistance)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// NamespaceExists indicates an expected call of NamespaceExists.
func (mr *MockNamespacesServiceMockRecorder) NamespaceExists(id, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{id, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NamespaceExists", reflect.TypeOf((*MockNamespacesService)(nil).NamespaceExists), varargs...)
}

// SearchNamespace mocks base method.
func (m *MockNamespacesService) SearchNamespace(query string, options ...gitlab.RequestOptionFunc) ([]*gitlab.Namespace, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{query}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SearchNamespace", varargs...)
	ret0, _ := ret[0].([]*gitlab.Namespace)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SearchNamespace indicates an expected call of SearchNamespace.
func (mr *MockNamespacesServiceMockRecorder) SearchNamespace(query any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{query}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchNamespace", reflect.TypeOf((*MockNamespacesService)(nil).SearchNamespace), varargs...)
}