// Code generated by MockGen. DO NOT EDIT.
// Source: featuresservice_inf.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=featuresservice_inf.go -destination=mocks/featuresservice_inf.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gitlab "gitlab.com/gitlab-org/api/client-go"
	gomock "go.uber.org/mock/gomock"
)

// MockFeaturesService is a mock of FeaturesService interface.
type MockFeaturesService struct {
	ctrl     *gomock.Controller
	recorder *MockFeaturesServiceMockRecorder
	isgomock struct{}
}

// MockFeaturesServiceMockRecorder is the mock recorder for MockFeaturesService.
type MockFeaturesServiceMockRecorder struct {
	mock *MockFeaturesService
}

// NewMockFeaturesService creates a new mock instance.
func NewMockFeaturesService(ctrl *gomock.Controller) *MockFeaturesService {
	mock := &MockFeaturesService{ctrl: ctrl}
	mock.recorder = &MockFeaturesServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFeaturesService) EXPECT() *MockFeaturesServiceMockRecorder {
	return m.recorder
}

// ListFeatures mocks base method.
func (m *MockFeaturesService) ListFeatures(options ...gitlab.RequestOptionFunc) ([]*gitlab.Feature, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListFeatures", varargs...)
	ret0, _ := ret[0].([]*gitlab.Feature)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListFeatures indicates an expected call of ListFeatures.
func (mr *MockFeaturesServiceMockRecorder) ListFeatures(options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFeatures", reflect.TypeOf((*MockFeaturesService)(nil).ListFeatures), options...)
}

// SetFeatureFlag mocks base method.
func (m *MockFeaturesService) SetFeatureFlag(name string, value any, options ...gitlab.RequestOptionFunc) (*gitlab.Feature, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{name, value}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetFeatureFlag", varargs...)
	ret0, _ := ret[0].(*gitlab.Feature)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SetFeatureFlag indicates an expected call of SetFeatureFlag.
func (mr *MockFeaturesServiceMockRecorder) SetFeatureFlag(name, value any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{name, value}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFeatureFlag", reflect.TypeOf((*MockFeaturesService)(nil).SetFeatureFlag), varargs...)
}
