// Code generated by MockGen. DO NOT EDIT.
// Source: notificationsettingsservice_inf.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=notificationsettingsservice_inf.go -destination=mocks/notificationsettingsservice_inf.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gitlab "github.com/xanzy/go-gitlab"
	gomock "go.uber.org/mock/gomock"
)

// MockNotificationSettingsService is a mock of NotificationSettingsService interface.
type MockNotificationSettingsService struct {
	ctrl     *gomock.Controller
	recorder *MockNotificationSettingsServiceMockRecorder
	isgomock struct{}
}

// MockNotificationSettingsServiceMockRecorder is the mock recorder for MockNotificationSettingsService.
type MockNotificationSettingsServiceMockRecorder struct {
	mock *MockNotificationSettingsService
}

// NewMockNotificationSettingsService creates a new mock instance.
func NewMockNotificationSettingsService(ctrl *gomock.Controller) *MockNotificationSettingsService {
	mock := &MockNotificationSettingsService{ctrl: ctrl}
	mock.recorder = &MockNotificationSettingsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNotificationSettingsService) EXPECT() *MockNotificationSettingsServiceMockRecorder {
	return m.recorder
}

// GetGlobalSettings mocks base method.
func (m *MockNotificationSettingsService) GetGlobalSettings(options ...gitlab.RequestOptionFunc) (*gitlab.NotificationSettings, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetGlobalSettings", varargs...)
	ret0, _ := ret[0].(*gitlab.NotificationSettings)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetGlobalSettings indicates an expected call of GetGlobalSettings.
func (mr *MockNotificationSettingsServiceMockRecorder) GetGlobalSettings(options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGlobalSettings", reflect.TypeOf((*MockNotificationSettingsService)(nil).GetGlobalSettings), options...)
}

// GetSettingsForGroup mocks base method.
func (m *MockNotificationSettingsService) GetSettingsForGroup(gid any, options ...gitlab.RequestOptionFunc) (*gitlab.NotificationSettings, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{gid}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSettingsForGroup", varargs...)
	ret0, _ := ret[0].(*gitlab.NotificationSettings)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetSettingsForGroup indicates an expected call of GetSettingsForGroup.
func (mr *MockNotificationSettingsServiceMockRecorder) GetSettingsForGroup(gid any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{gid}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSettingsForGroup", reflect.TypeOf((*MockNotificationSettingsService)(nil).GetSettingsForGroup), varargs...)
}

// GetSettingsForProject mocks base method.
func (m *MockNotificationSettingsService) GetSettingsForProject(pid any, options ...gitlab.RequestOptionFunc) (*gitlab.NotificationSettings, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSettingsForProject", varargs...)
	ret0, _ := ret[0].(*gitlab.NotificationSettings)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetSettingsForProject indicates an expected call of GetSettingsForProject.
func (mr *MockNotificationSettingsServiceMockRecorder) GetSettingsForProject(pid any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSettingsForProject", reflect.TypeOf((*MockNotificationSettingsService)(nil).GetSettingsForProject), varargs...)
}

// UpdateGlobalSettings mocks base method.
func (m *MockNotificationSettingsService) UpdateGlobalSettings(opt *gitlab.NotificationSettingsOptions, options ...gitlab.RequestOptionFunc) (*gitlab.NotificationSettings, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateGlobalSettings", varargs...)
	ret0, _ := ret[0].(*gitlab.NotificationSettings)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UpdateGlobalSettings indicates an expected call of UpdateGlobalSettings.
func (mr *MockNotificationSettingsServiceMockRecorder) UpdateGlobalSettings(opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGlobalSettings", reflect.TypeOf((*MockNotificationSettingsService)(nil).UpdateGlobalSettings), varargs...)
}

// UpdateSettingsForGroup mocks base method.
func (m *MockNotificationSettingsService) UpdateSettingsForGroup(gid any, opt *gitlab.NotificationSettingsOptions, options ...gitlab.RequestOptionFunc) (*gitlab.NotificationSettings, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{gid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateSettingsForGroup", varargs...)
	ret0, _ := ret[0].(*gitlab.NotificationSettings)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UpdateSettingsForGroup indicates an expected call of UpdateSettingsForGroup.
func (mr *MockNotificationSettingsServiceMockRecorder) UpdateSettingsForGroup(gid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{gid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSettingsForGroup", reflect.TypeOf((*MockNotificationSettingsService)(nil).UpdateSettingsForGroup), varargs...)
}

// UpdateSettingsForProject mocks base method.
func (m *MockNotificationSettingsService) UpdateSettingsForProject(pid any, opt *gitlab.NotificationSettingsOptions, options ...gitlab.RequestOptionFunc) (*gitlab.NotificationSettings, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateSettingsForProject", varargs...)
	ret0, _ := ret[0].(*gitlab.NotificationSettings)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UpdateSettingsForProject indicates an expected call of UpdateSettingsForProject.
func (mr *MockNotificationSettingsServiceMockRecorder) UpdateSettingsForProject(pid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSettingsForProject", reflect.TypeOf((*MockNotificationSettingsService)(nil).UpdateSettingsForProject), varargs...)
}
