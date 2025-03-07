// Code generated by MockGen. DO NOT EDIT.
// Source: geonodesservice_inf.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=geonodesservice_inf.go -destination=mocks/geonodesservice_inf.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gitlab "gitlab.com/gitlab-org/api/client-go"
	gomock "go.uber.org/mock/gomock"
)

// MockGeoNodesService is a mock of GeoNodesService interface.
type MockGeoNodesService struct {
	ctrl     *gomock.Controller
	recorder *MockGeoNodesServiceMockRecorder
	isgomock struct{}
}

// MockGeoNodesServiceMockRecorder is the mock recorder for MockGeoNodesService.
type MockGeoNodesServiceMockRecorder struct {
	mock *MockGeoNodesService
}

// NewMockGeoNodesService creates a new mock instance.
func NewMockGeoNodesService(ctrl *gomock.Controller) *MockGeoNodesService {
	mock := &MockGeoNodesService{ctrl: ctrl}
	mock.recorder = &MockGeoNodesServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGeoNodesService) EXPECT() *MockGeoNodesServiceMockRecorder {
	return m.recorder
}

// CreateGeoNode mocks base method.
func (m *MockGeoNodesService) CreateGeoNode(opt *gitlab.CreateGeoNodesOptions, options ...gitlab.RequestOptionFunc) (*gitlab.GeoNode, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateGeoNode", varargs...)
	ret0, _ := ret[0].(*gitlab.GeoNode)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateGeoNode indicates an expected call of CreateGeoNode.
func (mr *MockGeoNodesServiceMockRecorder) CreateGeoNode(opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGeoNode", reflect.TypeOf((*MockGeoNodesService)(nil).CreateGeoNode), varargs...)
}

// DeleteGeoNode mocks base method.
func (m *MockGeoNodesService) DeleteGeoNode(id int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{id}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteGeoNode", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteGeoNode indicates an expected call of DeleteGeoNode.
func (mr *MockGeoNodesServiceMockRecorder) DeleteGeoNode(id any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{id}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGeoNode", reflect.TypeOf((*MockGeoNodesService)(nil).DeleteGeoNode), varargs...)
}

// EditGeoNode mocks base method.
func (m *MockGeoNodesService) EditGeoNode(id int, opt *gitlab.UpdateGeoNodesOptions, options ...gitlab.RequestOptionFunc) (*gitlab.GeoNode, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{id, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EditGeoNode", varargs...)
	ret0, _ := ret[0].(*gitlab.GeoNode)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// EditGeoNode indicates an expected call of EditGeoNode.
func (mr *MockGeoNodesServiceMockRecorder) EditGeoNode(id, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{id, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditGeoNode", reflect.TypeOf((*MockGeoNodesService)(nil).EditGeoNode), varargs...)
}

// GetGeoNode mocks base method.
func (m *MockGeoNodesService) GetGeoNode(id int, options ...gitlab.RequestOptionFunc) (*gitlab.GeoNode, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{id}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetGeoNode", varargs...)
	ret0, _ := ret[0].(*gitlab.GeoNode)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetGeoNode indicates an expected call of GetGeoNode.
func (mr *MockGeoNodesServiceMockRecorder) GetGeoNode(id any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{id}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGeoNode", reflect.TypeOf((*MockGeoNodesService)(nil).GetGeoNode), varargs...)
}

// ListGeoNodes mocks base method.
func (m *MockGeoNodesService) ListGeoNodes(opt *gitlab.ListGeoNodesOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.GeoNode, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListGeoNodes", varargs...)
	ret0, _ := ret[0].([]*gitlab.GeoNode)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListGeoNodes indicates an expected call of ListGeoNodes.
func (mr *MockGeoNodesServiceMockRecorder) ListGeoNodes(opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGeoNodes", reflect.TypeOf((*MockGeoNodesService)(nil).ListGeoNodes), varargs...)
}

// RepairGeoNode mocks base method.
func (m *MockGeoNodesService) RepairGeoNode(id int, options ...gitlab.RequestOptionFunc) (*gitlab.GeoNode, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{id}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RepairGeoNode", varargs...)
	ret0, _ := ret[0].(*gitlab.GeoNode)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RepairGeoNode indicates an expected call of RepairGeoNode.
func (mr *MockGeoNodesServiceMockRecorder) RepairGeoNode(id any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{id}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RepairGeoNode", reflect.TypeOf((*MockGeoNodesService)(nil).RepairGeoNode), varargs...)
}

// RetrieveStatusOfAllGeoNodes mocks base method.
func (m *MockGeoNodesService) RetrieveStatusOfAllGeoNodes(options ...gitlab.RequestOptionFunc) ([]*gitlab.GeoNodeStatus, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RetrieveStatusOfAllGeoNodes", varargs...)
	ret0, _ := ret[0].([]*gitlab.GeoNodeStatus)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RetrieveStatusOfAllGeoNodes indicates an expected call of RetrieveStatusOfAllGeoNodes.
func (mr *MockGeoNodesServiceMockRecorder) RetrieveStatusOfAllGeoNodes(options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveStatusOfAllGeoNodes", reflect.TypeOf((*MockGeoNodesService)(nil).RetrieveStatusOfAllGeoNodes), options...)
}

// RetrieveStatusOfGeoNode mocks base method.
func (m *MockGeoNodesService) RetrieveStatusOfGeoNode(id int, options ...gitlab.RequestOptionFunc) (*gitlab.GeoNodeStatus, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{id}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RetrieveStatusOfGeoNode", varargs...)
	ret0, _ := ret[0].(*gitlab.GeoNodeStatus)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RetrieveStatusOfGeoNode indicates an expected call of RetrieveStatusOfGeoNode.
func (mr *MockGeoNodesServiceMockRecorder) RetrieveStatusOfGeoNode(id any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{id}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveStatusOfGeoNode", reflect.TypeOf((*MockGeoNodesService)(nil).RetrieveStatusOfGeoNode), varargs...)
}
