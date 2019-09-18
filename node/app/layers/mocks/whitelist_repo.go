// Code generated by MockGen. DO NOT EDIT.
// Source: ./app/layers/repositories/whitelist/interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
)

// MockWhiteListRepo is a mock of Repo interface
type MockWhiteListRepo struct {
	ctrl     *gomock.Controller
	recorder *MockWhiteListRepoMockRecorder
}

// MockWhiteListRepoMockRecorder is the mock recorder for MockWhiteListRepo
type MockWhiteListRepoMockRecorder struct {
	mock *MockWhiteListRepo
}

// NewMockWhiteListRepo creates a new mock instance
func NewMockWhiteListRepo(ctrl *gomock.Controller) *MockWhiteListRepo {
	mock := &MockWhiteListRepo{ctrl: ctrl}
	mock.recorder = &MockWhiteListRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWhiteListRepo) EXPECT() *MockWhiteListRepoMockRecorder {
	return m.recorder
}