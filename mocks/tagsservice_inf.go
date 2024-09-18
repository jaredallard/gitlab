// Code generated by MockGen. DO NOT EDIT.
// Source: tagsservice_inf.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=tagsservice_inf.go -destination=mocks/tagsservice_inf.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gitlab "github.com/xanzy/go-gitlab"
	gomock "go.uber.org/mock/gomock"
)

// MockTagsService is a mock of TagsService interface.
type MockTagsService struct {
	ctrl     *gomock.Controller
	recorder *MockTagsServiceMockRecorder
}

// MockTagsServiceMockRecorder is the mock recorder for MockTagsService.
type MockTagsServiceMockRecorder struct {
	mock *MockTagsService
}

// NewMockTagsService creates a new mock instance.
func NewMockTagsService(ctrl *gomock.Controller) *MockTagsService {
	mock := &MockTagsService{ctrl: ctrl}
	mock.recorder = &MockTagsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTagsService) EXPECT() *MockTagsServiceMockRecorder {
	return m.recorder
}

// CreateReleaseNote mocks base method.
func (m *MockTagsService) CreateReleaseNote(pid any, tag string, opt *gitlab.CreateReleaseNoteOptions, options ...gitlab.RequestOptionFunc) (*gitlab.ReleaseNote, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, tag, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateReleaseNote", varargs...)
	ret0, _ := ret[0].(*gitlab.ReleaseNote)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateReleaseNote indicates an expected call of CreateReleaseNote.
func (mr *MockTagsServiceMockRecorder) CreateReleaseNote(pid, tag, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, tag, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateReleaseNote", reflect.TypeOf((*MockTagsService)(nil).CreateReleaseNote), varargs...)
}

// CreateTag mocks base method.
func (m *MockTagsService) CreateTag(pid any, opt *gitlab.CreateTagOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Tag, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateTag", varargs...)
	ret0, _ := ret[0].(*gitlab.Tag)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateTag indicates an expected call of CreateTag.
func (mr *MockTagsServiceMockRecorder) CreateTag(pid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTag", reflect.TypeOf((*MockTagsService)(nil).CreateTag), varargs...)
}

// DeleteTag mocks base method.
func (m *MockTagsService) DeleteTag(pid any, tag string, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, tag}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteTag", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteTag indicates an expected call of DeleteTag.
func (mr *MockTagsServiceMockRecorder) DeleteTag(pid, tag any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, tag}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTag", reflect.TypeOf((*MockTagsService)(nil).DeleteTag), varargs...)
}

// GetTag mocks base method.
func (m *MockTagsService) GetTag(pid any, tag string, options ...gitlab.RequestOptionFunc) (*gitlab.Tag, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, tag}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTag", varargs...)
	ret0, _ := ret[0].(*gitlab.Tag)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTag indicates an expected call of GetTag.
func (mr *MockTagsServiceMockRecorder) GetTag(pid, tag any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, tag}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTag", reflect.TypeOf((*MockTagsService)(nil).GetTag), varargs...)
}

// ListTags mocks base method.
func (m *MockTagsService) ListTags(pid any, opt *gitlab.ListTagsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.Tag, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTags", varargs...)
	ret0, _ := ret[0].([]*gitlab.Tag)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListTags indicates an expected call of ListTags.
func (mr *MockTagsServiceMockRecorder) ListTags(pid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTags", reflect.TypeOf((*MockTagsService)(nil).ListTags), varargs...)
}

// UpdateReleaseNote mocks base method.
func (m *MockTagsService) UpdateReleaseNote(pid any, tag string, opt *gitlab.UpdateReleaseNoteOptions, options ...gitlab.RequestOptionFunc) (*gitlab.ReleaseNote, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, tag, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateReleaseNote", varargs...)
	ret0, _ := ret[0].(*gitlab.ReleaseNote)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UpdateReleaseNote indicates an expected call of UpdateReleaseNote.
func (mr *MockTagsServiceMockRecorder) UpdateReleaseNote(pid, tag, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, tag, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateReleaseNote", reflect.TypeOf((*MockTagsService)(nil).UpdateReleaseNote), varargs...)
}