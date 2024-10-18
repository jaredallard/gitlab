// Code generated by MockGen. DO NOT EDIT.
// Source: labelsservice_inf.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=labelsservice_inf.go -destination=mocks/labelsservice_inf.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gitlab "github.com/xanzy/go-gitlab"
	gomock "go.uber.org/mock/gomock"
)

// MockLabelsService is a mock of LabelsService interface.
type MockLabelsService struct {
	ctrl     *gomock.Controller
	recorder *MockLabelsServiceMockRecorder
	isgomock struct{}
}

// MockLabelsServiceMockRecorder is the mock recorder for MockLabelsService.
type MockLabelsServiceMockRecorder struct {
	mock *MockLabelsService
}

// NewMockLabelsService creates a new mock instance.
func NewMockLabelsService(ctrl *gomock.Controller) *MockLabelsService {
	mock := &MockLabelsService{ctrl: ctrl}
	mock.recorder = &MockLabelsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLabelsService) EXPECT() *MockLabelsServiceMockRecorder {
	return m.recorder
}

// CreateLabel mocks base method.
func (m *MockLabelsService) CreateLabel(pid any, opt *gitlab.CreateLabelOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Label, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateLabel", varargs...)
	ret0, _ := ret[0].(*gitlab.Label)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateLabel indicates an expected call of CreateLabel.
func (mr *MockLabelsServiceMockRecorder) CreateLabel(pid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLabel", reflect.TypeOf((*MockLabelsService)(nil).CreateLabel), varargs...)
}

// DeleteLabel mocks base method.
func (m *MockLabelsService) DeleteLabel(pid, lid any, opt *gitlab.DeleteLabelOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, lid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteLabel", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteLabel indicates an expected call of DeleteLabel.
func (mr *MockLabelsServiceMockRecorder) DeleteLabel(pid, lid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, lid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLabel", reflect.TypeOf((*MockLabelsService)(nil).DeleteLabel), varargs...)
}

// GetLabel mocks base method.
func (m *MockLabelsService) GetLabel(pid, labelID any, options ...gitlab.RequestOptionFunc) (*gitlab.Label, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, labelID}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetLabel", varargs...)
	ret0, _ := ret[0].(*gitlab.Label)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetLabel indicates an expected call of GetLabel.
func (mr *MockLabelsServiceMockRecorder) GetLabel(pid, labelID any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, labelID}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLabel", reflect.TypeOf((*MockLabelsService)(nil).GetLabel), varargs...)
}

// ListLabels mocks base method.
func (m *MockLabelsService) ListLabels(pid any, opt *gitlab.ListLabelsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.Label, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListLabels", varargs...)
	ret0, _ := ret[0].([]*gitlab.Label)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListLabels indicates an expected call of ListLabels.
func (mr *MockLabelsServiceMockRecorder) ListLabels(pid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLabels", reflect.TypeOf((*MockLabelsService)(nil).ListLabels), varargs...)
}

// PromoteLabel mocks base method.
func (m *MockLabelsService) PromoteLabel(pid, labelID any, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, labelID}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PromoteLabel", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PromoteLabel indicates an expected call of PromoteLabel.
func (mr *MockLabelsServiceMockRecorder) PromoteLabel(pid, labelID any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, labelID}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PromoteLabel", reflect.TypeOf((*MockLabelsService)(nil).PromoteLabel), varargs...)
}

// SubscribeToLabel mocks base method.
func (m *MockLabelsService) SubscribeToLabel(pid, labelID any, options ...gitlab.RequestOptionFunc) (*gitlab.Label, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, labelID}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SubscribeToLabel", varargs...)
	ret0, _ := ret[0].(*gitlab.Label)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SubscribeToLabel indicates an expected call of SubscribeToLabel.
func (mr *MockLabelsServiceMockRecorder) SubscribeToLabel(pid, labelID any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, labelID}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeToLabel", reflect.TypeOf((*MockLabelsService)(nil).SubscribeToLabel), varargs...)
}

// UnsubscribeFromLabel mocks base method.
func (m *MockLabelsService) UnsubscribeFromLabel(pid, labelID any, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, labelID}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UnsubscribeFromLabel", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnsubscribeFromLabel indicates an expected call of UnsubscribeFromLabel.
func (mr *MockLabelsServiceMockRecorder) UnsubscribeFromLabel(pid, labelID any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, labelID}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsubscribeFromLabel", reflect.TypeOf((*MockLabelsService)(nil).UnsubscribeFromLabel), varargs...)
}

// UpdateLabel mocks base method.
func (m *MockLabelsService) UpdateLabel(pid any, opt *gitlab.UpdateLabelOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Label, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateLabel", varargs...)
	ret0, _ := ret[0].(*gitlab.Label)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UpdateLabel indicates an expected call of UpdateLabel.
func (mr *MockLabelsServiceMockRecorder) UpdateLabel(pid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLabel", reflect.TypeOf((*MockLabelsService)(nil).UpdateLabel), varargs...)
}
