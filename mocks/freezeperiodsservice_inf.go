// Code generated by MockGen. DO NOT EDIT.
// Source: freezeperiodsservice_inf.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=freezeperiodsservice_inf.go -destination=mocks/freezeperiodsservice_inf.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gitlab "gitlab.com/gitlab-org/api/client-go"
	gomock "go.uber.org/mock/gomock"
)

// MockFreezePeriodsService is a mock of FreezePeriodsService interface.
type MockFreezePeriodsService struct {
	ctrl     *gomock.Controller
	recorder *MockFreezePeriodsServiceMockRecorder
	isgomock struct{}
}

// MockFreezePeriodsServiceMockRecorder is the mock recorder for MockFreezePeriodsService.
type MockFreezePeriodsServiceMockRecorder struct {
	mock *MockFreezePeriodsService
}

// NewMockFreezePeriodsService creates a new mock instance.
func NewMockFreezePeriodsService(ctrl *gomock.Controller) *MockFreezePeriodsService {
	mock := &MockFreezePeriodsService{ctrl: ctrl}
	mock.recorder = &MockFreezePeriodsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFreezePeriodsService) EXPECT() *MockFreezePeriodsServiceMockRecorder {
	return m.recorder
}

// CreateFreezePeriodOptions mocks base method.
func (m *MockFreezePeriodsService) CreateFreezePeriodOptions(pid any, opt *gitlab.CreateFreezePeriodOptions, options ...gitlab.RequestOptionFunc) (*gitlab.FreezePeriod, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateFreezePeriodOptions", varargs...)
	ret0, _ := ret[0].(*gitlab.FreezePeriod)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateFreezePeriodOptions indicates an expected call of CreateFreezePeriodOptions.
func (mr *MockFreezePeriodsServiceMockRecorder) CreateFreezePeriodOptions(pid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFreezePeriodOptions", reflect.TypeOf((*MockFreezePeriodsService)(nil).CreateFreezePeriodOptions), varargs...)
}

// DeleteFreezePeriod mocks base method.
func (m *MockFreezePeriodsService) DeleteFreezePeriod(pid any, freezePeriod int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, freezePeriod}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteFreezePeriod", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteFreezePeriod indicates an expected call of DeleteFreezePeriod.
func (mr *MockFreezePeriodsServiceMockRecorder) DeleteFreezePeriod(pid, freezePeriod any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, freezePeriod}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFreezePeriod", reflect.TypeOf((*MockFreezePeriodsService)(nil).DeleteFreezePeriod), varargs...)
}

// GetFreezePeriod mocks base method.
func (m *MockFreezePeriodsService) GetFreezePeriod(pid any, freezePeriod int, options ...gitlab.RequestOptionFunc) (*gitlab.FreezePeriod, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, freezePeriod}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetFreezePeriod", varargs...)
	ret0, _ := ret[0].(*gitlab.FreezePeriod)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetFreezePeriod indicates an expected call of GetFreezePeriod.
func (mr *MockFreezePeriodsServiceMockRecorder) GetFreezePeriod(pid, freezePeriod any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, freezePeriod}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFreezePeriod", reflect.TypeOf((*MockFreezePeriodsService)(nil).GetFreezePeriod), varargs...)
}

// ListFreezePeriods mocks base method.
func (m *MockFreezePeriodsService) ListFreezePeriods(pid any, opt *gitlab.ListFreezePeriodsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.FreezePeriod, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListFreezePeriods", varargs...)
	ret0, _ := ret[0].([]*gitlab.FreezePeriod)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListFreezePeriods indicates an expected call of ListFreezePeriods.
func (mr *MockFreezePeriodsServiceMockRecorder) ListFreezePeriods(pid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFreezePeriods", reflect.TypeOf((*MockFreezePeriodsService)(nil).ListFreezePeriods), varargs...)
}

// UpdateFreezePeriodOptions mocks base method.
func (m *MockFreezePeriodsService) UpdateFreezePeriodOptions(pid any, freezePeriod int, opt *gitlab.UpdateFreezePeriodOptions, options ...gitlab.RequestOptionFunc) (*gitlab.FreezePeriod, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, freezePeriod, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateFreezePeriodOptions", varargs...)
	ret0, _ := ret[0].(*gitlab.FreezePeriod)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UpdateFreezePeriodOptions indicates an expected call of UpdateFreezePeriodOptions.
func (mr *MockFreezePeriodsServiceMockRecorder) UpdateFreezePeriodOptions(pid, freezePeriod, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, freezePeriod, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFreezePeriodOptions", reflect.TypeOf((*MockFreezePeriodsService)(nil).UpdateFreezePeriodOptions), varargs...)
}
