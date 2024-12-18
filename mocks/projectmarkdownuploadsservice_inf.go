// Code generated by MockGen. DO NOT EDIT.
// Source: projectmarkdownuploadsservice_inf.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=projectmarkdownuploadsservice_inf.go -destination=mocks/projectmarkdownuploadsservice_inf.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	io "io"
	reflect "reflect"

	gitlab "gitlab.com/gitlab-org/api/client-go"
	gomock "go.uber.org/mock/gomock"
)

// MockProjectMarkdownUploadsService is a mock of ProjectMarkdownUploadsService interface.
type MockProjectMarkdownUploadsService struct {
	ctrl     *gomock.Controller
	recorder *MockProjectMarkdownUploadsServiceMockRecorder
	isgomock struct{}
}

// MockProjectMarkdownUploadsServiceMockRecorder is the mock recorder for MockProjectMarkdownUploadsService.
type MockProjectMarkdownUploadsServiceMockRecorder struct {
	mock *MockProjectMarkdownUploadsService
}

// NewMockProjectMarkdownUploadsService creates a new mock instance.
func NewMockProjectMarkdownUploadsService(ctrl *gomock.Controller) *MockProjectMarkdownUploadsService {
	mock := &MockProjectMarkdownUploadsService{ctrl: ctrl}
	mock.recorder = &MockProjectMarkdownUploadsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectMarkdownUploadsService) EXPECT() *MockProjectMarkdownUploadsServiceMockRecorder {
	return m.recorder
}

// DeleteProjectMarkdownUploadByID mocks base method.
func (m *MockProjectMarkdownUploadsService) DeleteProjectMarkdownUploadByID(pid any, uploadID int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, uploadID}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteProjectMarkdownUploadByID", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteProjectMarkdownUploadByID indicates an expected call of DeleteProjectMarkdownUploadByID.
func (mr *MockProjectMarkdownUploadsServiceMockRecorder) DeleteProjectMarkdownUploadByID(pid, uploadID any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, uploadID}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProjectMarkdownUploadByID", reflect.TypeOf((*MockProjectMarkdownUploadsService)(nil).DeleteProjectMarkdownUploadByID), varargs...)
}

// DeleteProjectMarkdownUploadBySecretAndFilename mocks base method.
func (m *MockProjectMarkdownUploadsService) DeleteProjectMarkdownUploadBySecretAndFilename(pid any, secret, filename string, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, secret, filename}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteProjectMarkdownUploadBySecretAndFilename", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteProjectMarkdownUploadBySecretAndFilename indicates an expected call of DeleteProjectMarkdownUploadBySecretAndFilename.
func (mr *MockProjectMarkdownUploadsServiceMockRecorder) DeleteProjectMarkdownUploadBySecretAndFilename(pid, secret, filename any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, secret, filename}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProjectMarkdownUploadBySecretAndFilename", reflect.TypeOf((*MockProjectMarkdownUploadsService)(nil).DeleteProjectMarkdownUploadBySecretAndFilename), varargs...)
}

// DownloadProjectMarkdownUploadByID mocks base method.
func (m *MockProjectMarkdownUploadsService) DownloadProjectMarkdownUploadByID(pid any, uploadID int, options ...gitlab.RequestOptionFunc) ([]byte, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, uploadID}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DownloadProjectMarkdownUploadByID", varargs...)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// DownloadProjectMarkdownUploadByID indicates an expected call of DownloadProjectMarkdownUploadByID.
func (mr *MockProjectMarkdownUploadsServiceMockRecorder) DownloadProjectMarkdownUploadByID(pid, uploadID any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, uploadID}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadProjectMarkdownUploadByID", reflect.TypeOf((*MockProjectMarkdownUploadsService)(nil).DownloadProjectMarkdownUploadByID), varargs...)
}

// DownloadProjectMarkdownUploadBySecretAndFilename mocks base method.
func (m *MockProjectMarkdownUploadsService) DownloadProjectMarkdownUploadBySecretAndFilename(pid any, secret, filename string, options ...gitlab.RequestOptionFunc) ([]byte, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, secret, filename}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DownloadProjectMarkdownUploadBySecretAndFilename", varargs...)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// DownloadProjectMarkdownUploadBySecretAndFilename indicates an expected call of DownloadProjectMarkdownUploadBySecretAndFilename.
func (mr *MockProjectMarkdownUploadsServiceMockRecorder) DownloadProjectMarkdownUploadBySecretAndFilename(pid, secret, filename any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, secret, filename}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadProjectMarkdownUploadBySecretAndFilename", reflect.TypeOf((*MockProjectMarkdownUploadsService)(nil).DownloadProjectMarkdownUploadBySecretAndFilename), varargs...)
}

// ListProjectMarkdownUploads mocks base method.
func (m *MockProjectMarkdownUploadsService) ListProjectMarkdownUploads(pid any, options ...gitlab.RequestOptionFunc) ([]*gitlab.ProjectMarkdownUpload, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListProjectMarkdownUploads", varargs...)
	ret0, _ := ret[0].([]*gitlab.ProjectMarkdownUpload)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListProjectMarkdownUploads indicates an expected call of ListProjectMarkdownUploads.
func (mr *MockProjectMarkdownUploadsServiceMockRecorder) ListProjectMarkdownUploads(pid any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjectMarkdownUploads", reflect.TypeOf((*MockProjectMarkdownUploadsService)(nil).ListProjectMarkdownUploads), varargs...)
}

// UploadProjectMarkdown mocks base method.
func (m *MockProjectMarkdownUploadsService) UploadProjectMarkdown(pid any, content io.Reader, options ...gitlab.RequestOptionFunc) (*gitlab.ProjectMarkdownUploadedFile, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, content}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UploadProjectMarkdown", varargs...)
	ret0, _ := ret[0].(*gitlab.ProjectMarkdownUploadedFile)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UploadProjectMarkdown indicates an expected call of UploadProjectMarkdown.
func (mr *MockProjectMarkdownUploadsServiceMockRecorder) UploadProjectMarkdown(pid, content any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, content}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadProjectMarkdown", reflect.TypeOf((*MockProjectMarkdownUploadsService)(nil).UploadProjectMarkdown), varargs...)
}
