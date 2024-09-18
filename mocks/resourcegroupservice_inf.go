// Code generated by MockGen. DO NOT EDIT.
// Source: resourcegroupservice_inf.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=resourcegroupservice_inf.go -destination=mocks/resourcegroupservice_inf.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gitlab "github.com/xanzy/go-gitlab"
	gomock "go.uber.org/mock/gomock"
)

// MockResourceGroupService is a mock of ResourceGroupService interface.
type MockResourceGroupService struct {
	ctrl     *gomock.Controller
	recorder *MockResourceGroupServiceMockRecorder
}

// MockResourceGroupServiceMockRecorder is the mock recorder for MockResourceGroupService.
type MockResourceGroupServiceMockRecorder struct {
	mock *MockResourceGroupService
}

// NewMockResourceGroupService creates a new mock instance.
func NewMockResourceGroupService(ctrl *gomock.Controller) *MockResourceGroupService {
	mock := &MockResourceGroupService{ctrl: ctrl}
	mock.recorder = &MockResourceGroupServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResourceGroupService) EXPECT() *MockResourceGroupServiceMockRecorder {
	return m.recorder
}

// EditAnExistingResourceGroup mocks base method.
func (m *MockResourceGroupService) EditAnExistingResourceGroup(pid any, key string, opts *gitlab.EditAnExistingResourceGroupOptions, options ...gitlab.RequestOptionFunc) (*gitlab.ResourceGroup, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, key, opts}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EditAnExistingResourceGroup", varargs...)
	ret0, _ := ret[0].(*gitlab.ResourceGroup)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// EditAnExistingResourceGroup indicates an expected call of EditAnExistingResourceGroup.
func (mr *MockResourceGroupServiceMockRecorder) EditAnExistingResourceGroup(pid, key, opts any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, key, opts}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditAnExistingResourceGroup", reflect.TypeOf((*MockResourceGroupService)(nil).EditAnExistingResourceGroup), varargs...)
}

// GetASpecificResourceGroup mocks base method.
func (m *MockResourceGroupService) GetASpecificResourceGroup(pid any, key string, options ...gitlab.RequestOptionFunc) (*gitlab.ResourceGroup, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, key}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetASpecificResourceGroup", varargs...)
	ret0, _ := ret[0].(*gitlab.ResourceGroup)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetASpecificResourceGroup indicates an expected call of GetASpecificResourceGroup.
func (mr *MockResourceGroupServiceMockRecorder) GetASpecificResourceGroup(pid, key any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, key}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetASpecificResourceGroup", reflect.TypeOf((*MockResourceGroupService)(nil).GetASpecificResourceGroup), varargs...)
}

// GetAllResourceGroupsForAProject mocks base method.
func (m *MockResourceGroupService) GetAllResourceGroupsForAProject(pid any, options ...gitlab.RequestOptionFunc) ([]*gitlab.ResourceGroup, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAllResourceGroupsForAProject", varargs...)
	ret0, _ := ret[0].([]*gitlab.ResourceGroup)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAllResourceGroupsForAProject indicates an expected call of GetAllResourceGroupsForAProject.
func (mr *MockResourceGroupServiceMockRecorder) GetAllResourceGroupsForAProject(pid any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllResourceGroupsForAProject", reflect.TypeOf((*MockResourceGroupService)(nil).GetAllResourceGroupsForAProject), varargs...)
}

// ListUpcomingJobsForASpecificResourceGroup mocks base method.
func (m *MockResourceGroupService) ListUpcomingJobsForASpecificResourceGroup(pid any, key string, options ...gitlab.RequestOptionFunc) ([]*gitlab.Job, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, key}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListUpcomingJobsForASpecificResourceGroup", varargs...)
	ret0, _ := ret[0].([]*gitlab.Job)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListUpcomingJobsForASpecificResourceGroup indicates an expected call of ListUpcomingJobsForASpecificResourceGroup.
func (mr *MockResourceGroupServiceMockRecorder) ListUpcomingJobsForASpecificResourceGroup(pid, key any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, key}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUpcomingJobsForASpecificResourceGroup", reflect.TypeOf((*MockResourceGroupService)(nil).ListUpcomingJobsForASpecificResourceGroup), varargs...)
}
