// Code generated by MockGen. DO NOT EDIT.
// Source: clusteragentsservice_inf.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=clusteragentsservice_inf.go -destination=mocks/clusteragentsservice_inf.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gitlab "gitlab.com/gitlab-org/api/client-go"
	gomock "go.uber.org/mock/gomock"
)

// MockClusterAgentsService is a mock of ClusterAgentsService interface.
type MockClusterAgentsService struct {
	ctrl     *gomock.Controller
	recorder *MockClusterAgentsServiceMockRecorder
	isgomock struct{}
}

// MockClusterAgentsServiceMockRecorder is the mock recorder for MockClusterAgentsService.
type MockClusterAgentsServiceMockRecorder struct {
	mock *MockClusterAgentsService
}

// NewMockClusterAgentsService creates a new mock instance.
func NewMockClusterAgentsService(ctrl *gomock.Controller) *MockClusterAgentsService {
	mock := &MockClusterAgentsService{ctrl: ctrl}
	mock.recorder = &MockClusterAgentsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterAgentsService) EXPECT() *MockClusterAgentsServiceMockRecorder {
	return m.recorder
}

// CreateAgentToken mocks base method.
func (m *MockClusterAgentsService) CreateAgentToken(pid any, aid int, opt *gitlab.CreateAgentTokenOptions, options ...gitlab.RequestOptionFunc) (*gitlab.AgentToken, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, aid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateAgentToken", varargs...)
	ret0, _ := ret[0].(*gitlab.AgentToken)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateAgentToken indicates an expected call of CreateAgentToken.
func (mr *MockClusterAgentsServiceMockRecorder) CreateAgentToken(pid, aid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, aid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAgentToken", reflect.TypeOf((*MockClusterAgentsService)(nil).CreateAgentToken), varargs...)
}

// DeleteAgent mocks base method.
func (m *MockClusterAgentsService) DeleteAgent(pid any, id int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, id}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAgent", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAgent indicates an expected call of DeleteAgent.
func (mr *MockClusterAgentsServiceMockRecorder) DeleteAgent(pid, id any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, id}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAgent", reflect.TypeOf((*MockClusterAgentsService)(nil).DeleteAgent), varargs...)
}

// GetAgent mocks base method.
func (m *MockClusterAgentsService) GetAgent(pid any, id int, options ...gitlab.RequestOptionFunc) (*gitlab.Agent, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, id}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAgent", varargs...)
	ret0, _ := ret[0].(*gitlab.Agent)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAgent indicates an expected call of GetAgent.
func (mr *MockClusterAgentsServiceMockRecorder) GetAgent(pid, id any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, id}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAgent", reflect.TypeOf((*MockClusterAgentsService)(nil).GetAgent), varargs...)
}

// GetAgentToken mocks base method.
func (m *MockClusterAgentsService) GetAgentToken(pid any, aid, id int, options ...gitlab.RequestOptionFunc) (*gitlab.AgentToken, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, aid, id}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAgentToken", varargs...)
	ret0, _ := ret[0].(*gitlab.AgentToken)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAgentToken indicates an expected call of GetAgentToken.
func (mr *MockClusterAgentsServiceMockRecorder) GetAgentToken(pid, aid, id any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, aid, id}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAgentToken", reflect.TypeOf((*MockClusterAgentsService)(nil).GetAgentToken), varargs...)
}

// ListAgentTokens mocks base method.
func (m *MockClusterAgentsService) ListAgentTokens(pid any, aid int, opt *gitlab.ListAgentTokensOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.AgentToken, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, aid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAgentTokens", varargs...)
	ret0, _ := ret[0].([]*gitlab.AgentToken)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListAgentTokens indicates an expected call of ListAgentTokens.
func (mr *MockClusterAgentsServiceMockRecorder) ListAgentTokens(pid, aid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, aid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAgentTokens", reflect.TypeOf((*MockClusterAgentsService)(nil).ListAgentTokens), varargs...)
}

// ListAgents mocks base method.
func (m *MockClusterAgentsService) ListAgents(pid any, opt *gitlab.ListAgentsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.Agent, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAgents", varargs...)
	ret0, _ := ret[0].([]*gitlab.Agent)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListAgents indicates an expected call of ListAgents.
func (mr *MockClusterAgentsServiceMockRecorder) ListAgents(pid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAgents", reflect.TypeOf((*MockClusterAgentsService)(nil).ListAgents), varargs...)
}

// RegisterAgent mocks base method.
func (m *MockClusterAgentsService) RegisterAgent(pid any, opt *gitlab.RegisterAgentOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Agent, *gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, opt}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RegisterAgent", varargs...)
	ret0, _ := ret[0].(*gitlab.Agent)
	ret1, _ := ret[1].(*gitlab.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RegisterAgent indicates an expected call of RegisterAgent.
func (mr *MockClusterAgentsServiceMockRecorder) RegisterAgent(pid, opt any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, opt}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterAgent", reflect.TypeOf((*MockClusterAgentsService)(nil).RegisterAgent), varargs...)
}

// RevokeAgentToken mocks base method.
func (m *MockClusterAgentsService) RevokeAgentToken(pid any, aid, id int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{pid, aid, id}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RevokeAgentToken", varargs...)
	ret0, _ := ret[0].(*gitlab.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RevokeAgentToken indicates an expected call of RevokeAgentToken.
func (mr *MockClusterAgentsServiceMockRecorder) RevokeAgentToken(pid, aid, id any, options ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{pid, aid, id}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeAgentToken", reflect.TypeOf((*MockClusterAgentsService)(nil).RevokeAgentToken), varargs...)
}
