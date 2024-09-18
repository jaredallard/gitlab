// Code generated by MockGen. DO NOT EDIT.
// Source: releasesservice_inf.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=releasesservice_inf.go -destination=mocks/releasesservice_inf.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gitlab "github.com/xanzy/go-gitlab"
	gomock "go.uber.org/mock/gomock"
)

// MockReleasesService is a mock of ReleasesService interface.
type MockReleasesService struct {
	ctrl     *gomock.Controller
	recorder *MockReleasesServiceMockRecorder
}

// MockReleasesServiceMockRecorder is the mock recorder for MockReleasesService.
type MockReleasesServiceMockRecorder struct {
	mock *MockReleasesService
}

// NewMockReleasesService creates a new mock instance.
func NewMockReleasesService(ctrl *gomock.Controller) *MockReleasesService {
	mock := &MockReleasesService{ctrl: ctrl}
	mock.recorder = &MockReleasesServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReleasesService) EXPECT() *MockReleasesServiceMockRecorder {
	return m.recorder
}

// CreateRelease mocks base method.
func (m *MockReleasesService) CreateRelease(pid any, opts *gitlab.CreateReleaseOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Release, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opts}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateRelease", varargs...)
	ret0, _ := ret[0].(*gitlab.Release)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateRelease indicates an expected call of CreateRelease.
func (mr *MockReleasesServiceMockRecorder) CreateRelease(pid, opts any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opts}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRelease", reflect.TypeOf((*MockReleasesService)(nil).CreateRelease), varargs...)
}

// DeleteRelease mocks base method.
func (m *MockReleasesService) DeleteRelease(pid any, tagName string, options ...gitlab.RequestOptionFunc) (*gitlab.Release, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, tagName}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteRelease", varargs...)
	ret0, _ := ret[0].(*gitlab.Release)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// DeleteRelease indicates an expected call of DeleteRelease.
func (mr *MockReleasesServiceMockRecorder) DeleteRelease(pid, tagName any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, tagName}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRelease", reflect.TypeOf((*MockReleasesService)(nil).DeleteRelease), varargs...)
}

// GetLatestRelease mocks base method.
func (m *MockReleasesService) GetLatestRelease(pid any, options ...gitlab.RequestOptionFunc) (*gitlab.Release, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetLatestRelease", varargs...)
	ret0, _ := ret[0].(*gitlab.Release)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetLatestRelease indicates an expected call of GetLatestRelease.
func (mr *MockReleasesServiceMockRecorder) GetLatestRelease(pid any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestRelease", reflect.TypeOf((*MockReleasesService)(nil).GetLatestRelease), varargs...)
}

// GetRelease mocks base method.
func (m *MockReleasesService) GetRelease(pid any, tagName string, options ...gitlab.RequestOptionFunc) (*gitlab.Release, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, tagName}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRelease", varargs...)
	ret0, _ := ret[0].(*gitlab.Release)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetRelease indicates an expected call of GetRelease.
func (mr *MockReleasesServiceMockRecorder) GetRelease(pid, tagName any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, tagName}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRelease", reflect.TypeOf((*MockReleasesService)(nil).GetRelease), varargs...)
}

// ListReleases mocks base method.
func (m *MockReleasesService) ListReleases(pid any, opt *gitlab.ListReleasesOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.Release, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListReleases", varargs...)
	ret0, _ := ret[0].([]*gitlab.Release)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListReleases indicates an expected call of ListReleases.
func (mr *MockReleasesServiceMockRecorder) ListReleases(pid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListReleases", reflect.TypeOf((*MockReleasesService)(nil).ListReleases), varargs...)
}

// UpdateRelease mocks base method.
func (m *MockReleasesService) UpdateRelease(pid any, tagName string, opts *gitlab.UpdateReleaseOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Release, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, tagName, opts}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateRelease", varargs...)
	ret0, _ := ret[0].(*gitlab.Release)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UpdateRelease indicates an expected call of UpdateRelease.
func (mr *MockReleasesServiceMockRecorder) UpdateRelease(pid, tagName, opts any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, tagName, opts}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRelease", reflect.TypeOf((*MockReleasesService)(nil).UpdateRelease), varargs...)
}
