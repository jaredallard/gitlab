// Code generated by MockGen. DO NOT EDIT.
// Source: draftnotesservice_inf.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=draftnotesservice_inf.go -destination=mocks/draftnotesservice_inf.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gitlab "gitlab.com/gitlab-org/api/client-go"
	gomock "go.uber.org/mock/gomock"
)

// MockDraftNotesService is a mock of DraftNotesService interface.
type MockDraftNotesService struct {
	ctrl     *gomock.Controller
	recorder *MockDraftNotesServiceMockRecorder
	isgomock struct{}
}

// MockDraftNotesServiceMockRecorder is the mock recorder for MockDraftNotesService.
type MockDraftNotesServiceMockRecorder struct {
	mock *MockDraftNotesService
}

// NewMockDraftNotesService creates a new mock instance.
func NewMockDraftNotesService(ctrl *gomock.Controller) *MockDraftNotesService {
	mock := &MockDraftNotesService{ctrl: ctrl}
	mock.recorder = &MockDraftNotesServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDraftNotesService) EXPECT() *MockDraftNotesServiceMockRecorder {
	return m.recorder
}

// CreateDraftNote mocks base method.
func (m *MockDraftNotesService) CreateDraftNote(pid any, mergeRequest int, opt *gitlab.CreateDraftNoteOptions, options ...gitlab.RequestOptionFunc) (*gitlab.DraftNote, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, mergeRequest, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateDraftNote", varargs...)
	ret0, _ := ret[0].(*gitlab.DraftNote)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateDraftNote indicates an expected call of CreateDraftNote.
func (mr *MockDraftNotesServiceMockRecorder) CreateDraftNote(pid, mergeRequest, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, mergeRequest, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDraftNote", reflect.TypeOf((*MockDraftNotesService)(nil).CreateDraftNote), varargs...)
}

// DeleteDraftNote mocks base method.
func (m *MockDraftNotesService) DeleteDraftNote(pid any, mergeRequest, note int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, mergeRequest, note}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteDraftNote", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteDraftNote indicates an expected call of DeleteDraftNote.
func (mr *MockDraftNotesServiceMockRecorder) DeleteDraftNote(pid, mergeRequest, note any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, mergeRequest, note}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDraftNote", reflect.TypeOf((*MockDraftNotesService)(nil).DeleteDraftNote), varargs...)
}

// GetDraftNote mocks base method.
func (m *MockDraftNotesService) GetDraftNote(pid any, mergeRequest, note int, options ...gitlab.RequestOptionFunc) (*gitlab.DraftNote, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, mergeRequest, note}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDraftNote", varargs...)
	ret0, _ := ret[0].(*gitlab.DraftNote)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetDraftNote indicates an expected call of GetDraftNote.
func (mr *MockDraftNotesServiceMockRecorder) GetDraftNote(pid, mergeRequest, note any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, mergeRequest, note}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDraftNote", reflect.TypeOf((*MockDraftNotesService)(nil).GetDraftNote), varargs...)
}

// ListDraftNotes mocks base method.
func (m *MockDraftNotesService) ListDraftNotes(pid any, mergeRequest int, opt *gitlab.ListDraftNotesOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.DraftNote, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, mergeRequest, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListDraftNotes", varargs...)
	ret0, _ := ret[0].([]*gitlab.DraftNote)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListDraftNotes indicates an expected call of ListDraftNotes.
func (mr *MockDraftNotesServiceMockRecorder) ListDraftNotes(pid, mergeRequest, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, mergeRequest, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDraftNotes", reflect.TypeOf((*MockDraftNotesService)(nil).ListDraftNotes), varargs...)
}

// PublishAllDraftNotes mocks base method.
func (m *MockDraftNotesService) PublishAllDraftNotes(pid any, mergeRequest int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, mergeRequest}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PublishAllDraftNotes", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PublishAllDraftNotes indicates an expected call of PublishAllDraftNotes.
func (mr *MockDraftNotesServiceMockRecorder) PublishAllDraftNotes(pid, mergeRequest any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, mergeRequest}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishAllDraftNotes", reflect.TypeOf((*MockDraftNotesService)(nil).PublishAllDraftNotes), varargs...)
}

// PublishDraftNote mocks base method.
func (m *MockDraftNotesService) PublishDraftNote(pid any, mergeRequest, note int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, mergeRequest, note}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PublishDraftNote", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PublishDraftNote indicates an expected call of PublishDraftNote.
func (mr *MockDraftNotesServiceMockRecorder) PublishDraftNote(pid, mergeRequest, note any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, mergeRequest, note}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishDraftNote", reflect.TypeOf((*MockDraftNotesService)(nil).PublishDraftNote), varargs...)
}

// UpdateDraftNote mocks base method.
func (m *MockDraftNotesService) UpdateDraftNote(pid any, mergeRequest, note int, opt *gitlab.UpdateDraftNoteOptions, options ...gitlab.RequestOptionFunc) (*gitlab.DraftNote, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, mergeRequest, note, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateDraftNote", varargs...)
	ret0, _ := ret[0].(*gitlab.DraftNote)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UpdateDraftNote indicates an expected call of UpdateDraftNote.
func (mr *MockDraftNotesServiceMockRecorder) UpdateDraftNote(pid, mergeRequest, note, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, mergeRequest, note, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDraftNote", reflect.TypeOf((*MockDraftNotesService)(nil).UpdateDraftNote), varargs...)
}
