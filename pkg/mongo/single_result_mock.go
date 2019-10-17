// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nomkhonwaan/myblog/pkg/mongo (interfaces: SingleResult)

// Package mongo is a generated GoMock package.
package mongo

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockSingleResult is a mock of SingleResult interface
type MockSingleResult struct {
	ctrl     *gomock.Controller
	recorder *MockSingleResultMockRecorder
}

// MockSingleResultMockRecorder is the mock recorder for MockSingleResult
type MockSingleResultMockRecorder struct {
	mock *MockSingleResult
}

// NewMockSingleResult creates a new mock instance
func NewMockSingleResult(ctrl *gomock.Controller) *MockSingleResult {
	mock := &MockSingleResult{ctrl: ctrl}
	mock.recorder = &MockSingleResultMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSingleResult) EXPECT() *MockSingleResultMockRecorder {
	return m.recorder
}

// Decode mocks base method
func (m *MockSingleResult) Decode(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decode", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Decode indicates an expected call of Decode
func (mr *MockSingleResultMockRecorder) Decode(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decode", reflect.TypeOf((*MockSingleResult)(nil).Decode), arg0)
}