// Code generated by MockGen. DO NOT EDIT.
// Source: mergetrainsservice_inf.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=mergetrainsservice_inf.go -destination=mocks/mergetrainsservice_inf.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gitlab "gitlab.com/gitlab-org/api/client-go"
	gomock "go.uber.org/mock/gomock"
)

// MockMergeTrainsService is a mock of MergeTrainsService interface.
type MockMergeTrainsService struct {
	ctrl     *gomock.Controller
	recorder *MockMergeTrainsServiceMockRecorder
	isgomock struct{}
}

// MockMergeTrainsServiceMockRecorder is the mock recorder for MockMergeTrainsService.
type MockMergeTrainsServiceMockRecorder struct {
	mock *MockMergeTrainsService
}

// NewMockMergeTrainsService creates a new mock instance.
func NewMockMergeTrainsService(ctrl *gomock.Controller) *MockMergeTrainsService {
	mock := &MockMergeTrainsService{ctrl: ctrl}
	mock.recorder = &MockMergeTrainsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMergeTrainsService) EXPECT() *MockMergeTrainsServiceMockRecorder {
	return m.recorder
}

// AddMergeRequestToMergeTrain mocks base method.
func (m *MockMergeTrainsService) AddMergeRequestToMergeTrain(pid any, mergeRequest int, opts *gitlab.AddMergeRequestToMergeTrainOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.MergeTrain, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, mergeRequest, opts}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddMergeRequestToMergeTrain", varargs...)
	ret0, _ := ret[0].([]*gitlab.MergeTrain)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AddMergeRequestToMergeTrain indicates an expected call of AddMergeRequestToMergeTrain.
func (mr *MockMergeTrainsServiceMockRecorder) AddMergeRequestToMergeTrain(pid, mergeRequest, opts any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, mergeRequest, opts}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMergeRequestToMergeTrain", reflect.TypeOf((*MockMergeTrainsService)(nil).AddMergeRequestToMergeTrain), varargs...)
}

// GetMergeRequestOnAMergeTrain mocks base method.
func (m *MockMergeTrainsService) GetMergeRequestOnAMergeTrain(pid any, mergeRequest int, options ...gitlab.RequestOptionFunc) (*gitlab.MergeTrain, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, mergeRequest}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMergeRequestOnAMergeTrain", varargs...)
	ret0, _ := ret[0].(*gitlab.MergeTrain)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetMergeRequestOnAMergeTrain indicates an expected call of GetMergeRequestOnAMergeTrain.
func (mr *MockMergeTrainsServiceMockRecorder) GetMergeRequestOnAMergeTrain(pid, mergeRequest any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, mergeRequest}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMergeRequestOnAMergeTrain", reflect.TypeOf((*MockMergeTrainsService)(nil).GetMergeRequestOnAMergeTrain), varargs...)
}

// ListMergeRequestInMergeTrain mocks base method.
func (m *MockMergeTrainsService) ListMergeRequestInMergeTrain(pid any, targetBranch string, opts *gitlab.ListMergeTrainsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.MergeTrain, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, targetBranch, opts}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListMergeRequestInMergeTrain", varargs...)
	ret0, _ := ret[0].([]*gitlab.MergeTrain)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListMergeRequestInMergeTrain indicates an expected call of ListMergeRequestInMergeTrain.
func (mr *MockMergeTrainsServiceMockRecorder) ListMergeRequestInMergeTrain(pid, targetBranch, opts any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, targetBranch, opts}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMergeRequestInMergeTrain", reflect.TypeOf((*MockMergeTrainsService)(nil).ListMergeRequestInMergeTrain), varargs...)
}

// ListProjectMergeTrains mocks base method.
func (m *MockMergeTrainsService) ListProjectMergeTrains(pid any, opt *gitlab.ListMergeTrainsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.MergeTrain, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListProjectMergeTrains", varargs...)
	ret0, _ := ret[0].([]*gitlab.MergeTrain)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListProjectMergeTrains indicates an expected call of ListProjectMergeTrains.
func (mr *MockMergeTrainsServiceMockRecorder) ListProjectMergeTrains(pid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjectMergeTrains", reflect.TypeOf((*MockMergeTrainsService)(nil).ListProjectMergeTrains), varargs...)
}
