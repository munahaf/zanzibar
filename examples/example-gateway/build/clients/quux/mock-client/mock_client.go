// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/uber/zanzibar/examples/example-gateway/clients/quux (interfaces: IClient)

// Package clientmock is a generated GoMock package.
package clientmock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	base "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients-idl/clients/foo/base/base"
)

// MockIClient is a mock of IClient interface.
type MockIClient struct {
	ctrl     *gomock.Controller
	recorder *MockIClientMockRecorder
}

// MockIClientMockRecorder is the mock recorder for MockIClient.
type MockIClientMockRecorder struct {
	mock *MockIClient
}

// NewMockIClient creates a new mock instance.
func NewMockIClient(ctrl *gomock.Controller) *MockIClient {
	mock := &MockIClient{ctrl: ctrl}
	mock.recorder = &MockIClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIClient) EXPECT() *MockIClientMockRecorder {
	return m.recorder
}

// DropMessages mocks base method.
func (m *MockIClient) DropMessages(arg0, arg1 *base.Message) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DropMessages", arg0, arg1)
}

// DropMessages indicates an expected call of DropMessages.
func (mr *MockIClientMockRecorder) DropMessages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DropMessages", reflect.TypeOf((*MockIClient)(nil).DropMessages), arg0, arg1)
}

// EchoMessage mocks base method.
func (m *MockIClient) EchoMessage(arg0 *base.Message) *base.Message {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EchoMessage", arg0)
	ret0, _ := ret[0].(*base.Message)
	return ret0
}

// EchoMessage indicates an expected call of EchoMessage.
func (mr *MockIClientMockRecorder) EchoMessage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EchoMessage", reflect.TypeOf((*MockIClient)(nil).EchoMessage), arg0)
}

// EchoString mocks base method.
func (m *MockIClient) EchoString(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EchoString", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// EchoString indicates an expected call of EchoString.
func (mr *MockIClientMockRecorder) EchoString(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EchoString", reflect.TypeOf((*MockIClient)(nil).EchoString), arg0)
}
