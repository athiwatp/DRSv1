// Code generated by MockGen. DO NOT EDIT.
// Source: ./app/layers/repositories/stellar/interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	horizon "github.com/stellar/go/protocols/horizon"
	reflect "reflect"
)

// MockStellarRepo is a mock of Repo interface
type MockStellarRepo struct {
	ctrl     *gomock.Controller
	recorder *MockStellarRepoMockRecorder
}

// MockStellarRepoMockRecorder is the mock recorder for MockStellarRepo
type MockStellarRepoMockRecorder struct {
	mock *MockStellarRepo
}

// NewMockStellarRepo creates a new mock instance
func NewMockStellarRepo(ctrl *gomock.Controller) *MockStellarRepo {
	mock := &MockStellarRepo{ctrl: ctrl}
	mock.recorder = &MockStellarRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStellarRepo) EXPECT() *MockStellarRepoMockRecorder {
	return m.recorder
}

// BuildSetupTx mocks base method
func (m *MockStellarRepo) BuildSetupTx(drsAccount *horizon.Account, peggedValue, peggedCurrency, assetName, creditOwnerAddress string) (string, string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildSetupTx", drsAccount, peggedValue, peggedCurrency, assetName, creditOwnerAddress)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// BuildSetupTx indicates an expected call of BuildSetupTx
func (mr *MockStellarRepoMockRecorder) BuildSetupTx(drsAccount, peggedValue, peggedCurrency, assetName, creditOwnerAddress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildSetupTx", reflect.TypeOf((*MockStellarRepo)(nil).BuildSetupTx), drsAccount, peggedValue, peggedCurrency, assetName, creditOwnerAddress)
}

// BuildMintTx mocks base method
func (m *MockStellarRepo) BuildMintTx(drsAccount *horizon.Account, amount, assetName, issuerAddress, distributorAddress string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildMintTx", drsAccount, amount, assetName, issuerAddress, distributorAddress)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildMintTx indicates an expected call of BuildMintTx
func (mr *MockStellarRepoMockRecorder) BuildMintTx(drsAccount, amount, assetName, issuerAddress, distributorAddress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildMintTx", reflect.TypeOf((*MockStellarRepo)(nil).BuildMintTx), drsAccount, amount, assetName, issuerAddress, distributorAddress)
}

// LoadAccount mocks base method
func (m *MockStellarRepo) LoadAccount(stellarAddress string) (*horizon.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadAccount", stellarAddress)
	ret0, _ := ret[0].(*horizon.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadAccount indicates an expected call of LoadAccount
func (mr *MockStellarRepoMockRecorder) LoadAccount(stellarAddress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadAccount", reflect.TypeOf((*MockStellarRepo)(nil).LoadAccount), stellarAddress)
}

// SubmitTransaction mocks base method
func (m *MockStellarRepo) SubmitTransaction(txB64 string) (*horizon.TransactionSuccess, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitTransaction", txB64)
	ret0, _ := ret[0].(*horizon.TransactionSuccess)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitTransaction indicates an expected call of SubmitTransaction
func (mr *MockStellarRepoMockRecorder) SubmitTransaction(txB64 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitTransaction", reflect.TypeOf((*MockStellarRepo)(nil).SubmitTransaction), txB64)
}
