// Code generated by MockGen. DO NOT EDIT.
// Source: projectimportexportservice_inf.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=projectimportexportservice_inf.go -destination=mocks/projectimportexportservice_inf.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	io "io"
	reflect "reflect"

	gitlab "github.com/xanzy/go-gitlab"
	gomock "go.uber.org/mock/gomock"
)

// MockProjectImportExportService is a mock of ProjectImportExportService interface.
type MockProjectImportExportService struct {
	ctrl     *gomock.Controller
	recorder *MockProjectImportExportServiceMockRecorder
}

// MockProjectImportExportServiceMockRecorder is the mock recorder for MockProjectImportExportService.
type MockProjectImportExportServiceMockRecorder struct {
	mock *MockProjectImportExportService
}

// NewMockProjectImportExportService creates a new mock instance.
func NewMockProjectImportExportService(ctrl *gomock.Controller) *MockProjectImportExportService {
	mock := &MockProjectImportExportService{ctrl: ctrl}
	mock.recorder = &MockProjectImportExportServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectImportExportService) EXPECT() *MockProjectImportExportServiceMockRecorder {
	return m.recorder
}

// ExportDownload mocks base method.
func (m *MockProjectImportExportService) ExportDownload(pid any, options ...gitlab.RequestOptionFunc) ([]byte, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ExportDownload", varargs...)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ExportDownload indicates an expected call of ExportDownload.
func (mr *MockProjectImportExportServiceMockRecorder) ExportDownload(pid any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportDownload", reflect.TypeOf((*MockProjectImportExportService)(nil).ExportDownload), varargs...)
}

// ExportStatus mocks base method.
func (m *MockProjectImportExportService) ExportStatus(pid any, options ...gitlab.RequestOptionFunc) (*gitlab.ExportStatus, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ExportStatus", varargs...)
	ret0, _ := ret[0].(*gitlab.ExportStatus)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ExportStatus indicates an expected call of ExportStatus.
func (mr *MockProjectImportExportServiceMockRecorder) ExportStatus(pid any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportStatus", reflect.TypeOf((*MockProjectImportExportService)(nil).ExportStatus), varargs...)
}

// ImportFromFile mocks base method.
func (m *MockProjectImportExportService) ImportFromFile(archive io.Reader, opt *gitlab.ImportFileOptions, options ...gitlab.RequestOptionFunc) (*gitlab.ImportStatus, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{archive, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ImportFromFile", varargs...)
	ret0, _ := ret[0].(*gitlab.ImportStatus)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ImportFromFile indicates an expected call of ImportFromFile.
func (mr *MockProjectImportExportServiceMockRecorder) ImportFromFile(archive, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{archive, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportFromFile", reflect.TypeOf((*MockProjectImportExportService)(nil).ImportFromFile), varargs...)
}

// ImportStatus mocks base method.
func (m *MockProjectImportExportService) ImportStatus(pid any, options ...gitlab.RequestOptionFunc) (*gitlab.ImportStatus, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ImportStatus", varargs...)
	ret0, _ := ret[0].(*gitlab.ImportStatus)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ImportStatus indicates an expected call of ImportStatus.
func (mr *MockProjectImportExportServiceMockRecorder) ImportStatus(pid any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportStatus", reflect.TypeOf((*MockProjectImportExportService)(nil).ImportStatus), varargs...)
}

// ScheduleExport mocks base method.
func (m *MockProjectImportExportService) ScheduleExport(pid any, opt *gitlab.ScheduleExportOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ScheduleExport", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ScheduleExport indicates an expected call of ScheduleExport.
func (mr *MockProjectImportExportServiceMockRecorder) ScheduleExport(pid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScheduleExport", reflect.TypeOf((*MockProjectImportExportService)(nil).ScheduleExport), varargs...)
}
