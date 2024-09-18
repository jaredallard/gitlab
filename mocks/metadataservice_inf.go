// Code generated by MockGen. DO NOT EDIT.
// Source: metadataservice_inf.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=metadataservice_inf.go -destination=mocks/metadataservice_inf.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gitlab "github.com/xanzy/go-gitlab"
	gomock "go.uber.org/mock/gomock"
)

// MockMetadataService is a mock of MetadataService interface.
type MockMetadataService struct {
	ctrl     *gomock.Controller
	recorder *MockMetadataServiceMockRecorder
}

// MockMetadataServiceMockRecorder is the mock recorder for MockMetadataService.
type MockMetadataServiceMockRecorder struct {
	mock *MockMetadataService
}

// NewMockMetadataService creates a new mock instance.
func NewMockMetadataService(ctrl *gomock.Controller) *MockMetadataService {
	mock := &MockMetadataService{ctrl: ctrl}
	mock.recorder = &MockMetadataServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMetadataService) EXPECT() *MockMetadataServiceMockRecorder {
	return m.recorder
}

// GetMetadata mocks base method.
func (m *MockMetadataService) GetMetadata(options ...gitlab.RequestOptionFunc) (*gitlab.Metadata, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMetadata", varargs...)
	ret0, _ := ret[0].(*gitlab.Metadata)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetMetadata indicates an expected call of GetMetadata.
func (mr *MockMetadataServiceMockRecorder) GetMetadata(options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetadata", reflect.TypeOf((*MockMetadataService)(nil).GetMetadata), options...)
}