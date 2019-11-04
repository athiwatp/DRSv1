// Code generated by MockGen. DO NOT EDIT.
// Source: ./layers/logic/interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	entity "gitlab.com/velo-labs/cen/cmd/gvel/entity"
	reflect "reflect"
)

// MockLogic is a mock of Logic interface
type MockLogic struct {
	ctrl     *gomock.Controller
	recorder *MockLogicMockRecorder
}

// MockLogicMockRecorder is the mock recorder for MockLogic
type MockLogicMockRecorder struct {
	mock *MockLogic
}

// NewMockLogic creates a new mock instance
func NewMockLogic(ctrl *gomock.Controller) *MockLogic {
	mock := &MockLogic{ctrl: ctrl}
	mock.recorder = &MockLogicMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLogic) EXPECT() *MockLogicMockRecorder {
	return m.recorder
}

// Init mocks base method
func (m *MockLogic) Init(configFilePath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", configFilePath)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init
func (mr *MockLogicMockRecorder) Init(configFilePath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockLogic)(nil).Init), configFilePath)
}

// CreateAccount mocks base method
func (m *MockLogic) CreateAccount(input *entity.CreateAccountInput) (*entity.CreateAccountOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", input)
	ret0, _ := ret[0].(*entity.CreateAccountOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount
func (mr *MockLogicMockRecorder) CreateAccount(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockLogic)(nil).CreateAccount), input)
}

// ListAccount mocks base method
func (m *MockLogic) ListAccount() (*[]entity.StellarAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAccount")
	ret0, _ := ret[0].(*[]entity.StellarAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccount indicates an expected call of ListAccount
func (mr *MockLogicMockRecorder) ListAccount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccount", reflect.TypeOf((*MockLogic)(nil).ListAccount))
}

// SetDefaultAccount mocks base method
func (m *MockLogic) SetDefaultAccount(input *entity.SetDefaultAccountInput) (*entity.SetDefaultAccountOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDefaultAccount", input)
	ret0, _ := ret[0].(*entity.SetDefaultAccountOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetDefaultAccount indicates an expected call of SetDefaultAccount
func (mr *MockLogicMockRecorder) SetDefaultAccount(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDefaultAccount", reflect.TypeOf((*MockLogic)(nil).SetDefaultAccount), input)
}

// SetupCredit mocks base method
func (m *MockLogic) SetupCredit(input *entity.SetupCreditInput) (*entity.SetupCreditOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetupCredit", input)
	ret0, _ := ret[0].(*entity.SetupCreditOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetupCredit indicates an expected call of SetupCredit
func (mr *MockLogicMockRecorder) SetupCredit(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetupCredit", reflect.TypeOf((*MockLogic)(nil).SetupCredit), input)
}

// ImportAccount mocks base method
func (m *MockLogic) ImportAccount(input *entity.ImportAccountInput) (*entity.ImportAccountOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImportAccount", input)
	ret0, _ := ret[0].(*entity.ImportAccountOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImportAccount indicates an expected call of ImportAccount
func (mr *MockLogicMockRecorder) ImportAccount(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportAccount", reflect.TypeOf((*MockLogic)(nil).ImportAccount), input)
}

// MintCredit mocks base method
func (m *MockLogic) MintCredit(input *entity.MintCreditInput) (*entity.MintCreditOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MintCredit", input)
	ret0, _ := ret[0].(*entity.MintCreditOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MintCredit indicates an expected call of MintCredit
func (mr *MockLogicMockRecorder) MintCredit(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MintCredit", reflect.TypeOf((*MockLogic)(nil).MintCredit), input)
}

// RedeemCredit mocks base method
func (m *MockLogic) RedeemCredit(input *entity.RedeemCreditInput) (*entity.RedeemCreditOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RedeemCredit", input)
	ret0, _ := ret[0].(*entity.RedeemCreditOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RedeemCredit indicates an expected call of RedeemCredit
func (mr *MockLogicMockRecorder) RedeemCredit(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RedeemCredit", reflect.TypeOf((*MockLogic)(nil).RedeemCredit), input)
}

// GetExchangeRate mocks base method
func (m *MockLogic) GetExchangeRate(input *entity.GetExchangeRateInput) (*entity.GetExchangeRateOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExchangeRate", input)
	ret0, _ := ret[0].(*entity.GetExchangeRateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExchangeRate indicates an expected call of GetExchangeRate
func (mr *MockLogicMockRecorder) GetExchangeRate(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExchangeRate", reflect.TypeOf((*MockLogic)(nil).GetExchangeRate), input)
}
