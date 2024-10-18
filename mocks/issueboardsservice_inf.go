// Code generated by MockGen. DO NOT EDIT.
// Source: issueboardsservice_inf.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=issueboardsservice_inf.go -destination=mocks/issueboardsservice_inf.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gitlab "github.com/xanzy/go-gitlab"
	gomock "go.uber.org/mock/gomock"
)

// MockIssueBoardsService is a mock of IssueBoardsService interface.
type MockIssueBoardsService struct {
	ctrl     *gomock.Controller
	recorder *MockIssueBoardsServiceMockRecorder
	isgomock struct{}
}

// MockIssueBoardsServiceMockRecorder is the mock recorder for MockIssueBoardsService.
type MockIssueBoardsServiceMockRecorder struct {
	mock *MockIssueBoardsService
}

// NewMockIssueBoardsService creates a new mock instance.
func NewMockIssueBoardsService(ctrl *gomock.Controller) *MockIssueBoardsService {
	mock := &MockIssueBoardsService{ctrl: ctrl}
	mock.recorder = &MockIssueBoardsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIssueBoardsService) EXPECT() *MockIssueBoardsServiceMockRecorder {
	return m.recorder
}

// CreateIssueBoard mocks base method.
func (m *MockIssueBoardsService) CreateIssueBoard(pid any, opt *gitlab.CreateIssueBoardOptions, options ...gitlab.RequestOptionFunc) (*gitlab.IssueBoard, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateIssueBoard", varargs...)
	ret0, _ := ret[0].(*gitlab.IssueBoard)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateIssueBoard indicates an expected call of CreateIssueBoard.
func (mr *MockIssueBoardsServiceMockRecorder) CreateIssueBoard(pid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateIssueBoard", reflect.TypeOf((*MockIssueBoardsService)(nil).CreateIssueBoard), varargs...)
}

// CreateIssueBoardList mocks base method.
func (m *MockIssueBoardsService) CreateIssueBoardList(pid any, board int, opt *gitlab.CreateIssueBoardListOptions, options ...gitlab.RequestOptionFunc) (*gitlab.BoardList, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, board, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateIssueBoardList", varargs...)
	ret0, _ := ret[0].(*gitlab.BoardList)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateIssueBoardList indicates an expected call of CreateIssueBoardList.
func (mr *MockIssueBoardsServiceMockRecorder) CreateIssueBoardList(pid, board, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, board, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateIssueBoardList", reflect.TypeOf((*MockIssueBoardsService)(nil).CreateIssueBoardList), varargs...)
}

// DeleteIssueBoard mocks base method.
func (m *MockIssueBoardsService) DeleteIssueBoard(pid any, board int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, board}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteIssueBoard", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteIssueBoard indicates an expected call of DeleteIssueBoard.
func (mr *MockIssueBoardsServiceMockRecorder) DeleteIssueBoard(pid, board any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, board}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteIssueBoard", reflect.TypeOf((*MockIssueBoardsService)(nil).DeleteIssueBoard), varargs...)
}

// DeleteIssueBoardList mocks base method.
func (m *MockIssueBoardsService) DeleteIssueBoardList(pid any, board, list int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, board, list}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteIssueBoardList", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteIssueBoardList indicates an expected call of DeleteIssueBoardList.
func (mr *MockIssueBoardsServiceMockRecorder) DeleteIssueBoardList(pid, board, list any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, board, list}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteIssueBoardList", reflect.TypeOf((*MockIssueBoardsService)(nil).DeleteIssueBoardList), varargs...)
}

// GetIssueBoard mocks base method.
func (m *MockIssueBoardsService) GetIssueBoard(pid any, board int, options ...gitlab.RequestOptionFunc) (*gitlab.IssueBoard, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, board}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetIssueBoard", varargs...)
	ret0, _ := ret[0].(*gitlab.IssueBoard)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetIssueBoard indicates an expected call of GetIssueBoard.
func (mr *MockIssueBoardsServiceMockRecorder) GetIssueBoard(pid, board any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, board}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIssueBoard", reflect.TypeOf((*MockIssueBoardsService)(nil).GetIssueBoard), varargs...)
}

// GetIssueBoardList mocks base method.
func (m *MockIssueBoardsService) GetIssueBoardList(pid any, board, list int, options ...gitlab.RequestOptionFunc) (*gitlab.BoardList, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, board, list}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetIssueBoardList", varargs...)
	ret0, _ := ret[0].(*gitlab.BoardList)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetIssueBoardList indicates an expected call of GetIssueBoardList.
func (mr *MockIssueBoardsServiceMockRecorder) GetIssueBoardList(pid, board, list any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, board, list}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIssueBoardList", reflect.TypeOf((*MockIssueBoardsService)(nil).GetIssueBoardList), varargs...)
}

// GetIssueBoardLists mocks base method.
func (m *MockIssueBoardsService) GetIssueBoardLists(pid any, board int, opt *gitlab.GetIssueBoardListsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.BoardList, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, board, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetIssueBoardLists", varargs...)
	ret0, _ := ret[0].([]*gitlab.BoardList)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetIssueBoardLists indicates an expected call of GetIssueBoardLists.
func (mr *MockIssueBoardsServiceMockRecorder) GetIssueBoardLists(pid, board, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, board, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIssueBoardLists", reflect.TypeOf((*MockIssueBoardsService)(nil).GetIssueBoardLists), varargs...)
}

// ListIssueBoards mocks base method.
func (m *MockIssueBoardsService) ListIssueBoards(pid any, opt *gitlab.ListIssueBoardsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.IssueBoard, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListIssueBoards", varargs...)
	ret0, _ := ret[0].([]*gitlab.IssueBoard)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListIssueBoards indicates an expected call of ListIssueBoards.
func (mr *MockIssueBoardsServiceMockRecorder) ListIssueBoards(pid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListIssueBoards", reflect.TypeOf((*MockIssueBoardsService)(nil).ListIssueBoards), varargs...)
}

// UpdateIssueBoard mocks base method.
func (m *MockIssueBoardsService) UpdateIssueBoard(pid any, board int, opt *gitlab.UpdateIssueBoardOptions, options ...gitlab.RequestOptionFunc) (*gitlab.IssueBoard, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, board, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateIssueBoard", varargs...)
	ret0, _ := ret[0].(*gitlab.IssueBoard)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UpdateIssueBoard indicates an expected call of UpdateIssueBoard.
func (mr *MockIssueBoardsServiceMockRecorder) UpdateIssueBoard(pid, board, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, board, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateIssueBoard", reflect.TypeOf((*MockIssueBoardsService)(nil).UpdateIssueBoard), varargs...)
}

// UpdateIssueBoardList mocks base method.
func (m *MockIssueBoardsService) UpdateIssueBoardList(pid any, board, list int, opt *gitlab.UpdateIssueBoardListOptions, options ...gitlab.RequestOptionFunc) (*gitlab.BoardList, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, board, list, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateIssueBoardList", varargs...)
	ret0, _ := ret[0].(*gitlab.BoardList)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UpdateIssueBoardList indicates an expected call of UpdateIssueBoardList.
func (mr *MockIssueBoardsServiceMockRecorder) UpdateIssueBoardList(pid, board, list, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, board, list, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateIssueBoardList", reflect.TypeOf((*MockIssueBoardsService)(nil).UpdateIssueBoardList), varargs...)
}
