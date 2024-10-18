// Code generated by MockGen. DO NOT EDIT.
// Source: projectclustersservice_inf.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=projectclustersservice_inf.go -destination=mocks/projectclustersservice_inf.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gitlab "github.com/xanzy/go-gitlab"
	gomock "go.uber.org/mock/gomock"
)

// MockProjectClustersService is a mock of ProjectClustersService interface.
type MockProjectClustersService struct {
	ctrl     *gomock.Controller
	recorder *MockProjectClustersServiceMockRecorder
	isgomock struct{}
}

// MockProjectClustersServiceMockRecorder is the mock recorder for MockProjectClustersService.
type MockProjectClustersServiceMockRecorder struct {
	mock *MockProjectClustersService
}

// NewMockProjectClustersService creates a new mock instance.
func NewMockProjectClustersService(ctrl *gomock.Controller) *MockProjectClustersService {
	mock := &MockProjectClustersService{ctrl: ctrl}
	mock.recorder = &MockProjectClustersServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectClustersService) EXPECT() *MockProjectClustersServiceMockRecorder {
	return m.recorder
}

// AddCluster mocks base method.
func (m *MockProjectClustersService) AddCluster(pid any, opt *gitlab.AddClusterOptions, options ...gitlab.RequestOptionFunc) (*gitlab.ProjectCluster, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddCluster", varargs...)
	ret0, _ := ret[0].(*gitlab.ProjectCluster)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AddCluster indicates an expected call of AddCluster.
func (mr *MockProjectClustersServiceMockRecorder) AddCluster(pid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCluster", reflect.TypeOf((*MockProjectClustersService)(nil).AddCluster), varargs...)
}

// DeleteCluster mocks base method.
func (m *MockProjectClustersService) DeleteCluster(pid any, cluster int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, cluster}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteCluster", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteCluster indicates an expected call of DeleteCluster.
func (mr *MockProjectClustersServiceMockRecorder) DeleteCluster(pid, cluster any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, cluster}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCluster", reflect.TypeOf((*MockProjectClustersService)(nil).DeleteCluster), varargs...)
}

// EditCluster mocks base method.
func (m *MockProjectClustersService) EditCluster(pid any, cluster int, opt *gitlab.EditClusterOptions, options ...gitlab.RequestOptionFunc) (*gitlab.ProjectCluster, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, cluster, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EditCluster", varargs...)
	ret0, _ := ret[0].(*gitlab.ProjectCluster)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// EditCluster indicates an expected call of EditCluster.
func (mr *MockProjectClustersServiceMockRecorder) EditCluster(pid, cluster, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, cluster, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditCluster", reflect.TypeOf((*MockProjectClustersService)(nil).EditCluster), varargs...)
}

// GetCluster mocks base method.
func (m *MockProjectClustersService) GetCluster(pid any, cluster int, options ...gitlab.RequestOptionFunc) (*gitlab.ProjectCluster, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, cluster}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCluster", varargs...)
	ret0, _ := ret[0].(*gitlab.ProjectCluster)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetCluster indicates an expected call of GetCluster.
func (mr *MockProjectClustersServiceMockRecorder) GetCluster(pid, cluster any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, cluster}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCluster", reflect.TypeOf((*MockProjectClustersService)(nil).GetCluster), varargs...)
}

// ListClusters mocks base method.
func (m *MockProjectClustersService) ListClusters(pid any, options ...gitlab.RequestOptionFunc) ([]*gitlab.ProjectCluster, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListClusters", varargs...)
	ret0, _ := ret[0].([]*gitlab.ProjectCluster)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListClusters indicates an expected call of ListClusters.
func (mr *MockProjectClustersServiceMockRecorder) ListClusters(pid any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListClusters", reflect.TypeOf((*MockProjectClustersService)(nil).ListClusters), varargs...)
}
