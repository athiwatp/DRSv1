// Code generated by MockGen. DO NOT EDIT.
// Source: ../../libs/client/interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	keypair "github.com/stellar/go/keypair"
	grpc "github.com/velo-protocol/DRSv1/grpc"
	client "github.com/velo-protocol/DRSv1/libs/client"
	txnbuild "github.com/velo-protocol/DRSv1/libs/txnbuild"
	reflect "reflect"
)

// MockVeloClient is a mock of ClientInterface interface
type MockVeloClient struct {
	ctrl     *gomock.Controller
	recorder *MockVeloClientMockRecorder
}

// MockVeloClientMockRecorder is the mock recorder for MockVeloClient
type MockVeloClientMockRecorder struct {
	mock *MockVeloClient
}

// NewMockVeloClient creates a new mock instance
func NewMockVeloClient(ctrl *gomock.Controller) *MockVeloClient {
	mock := &MockVeloClient{ctrl: ctrl}
	mock.recorder = &MockVeloClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockVeloClient) EXPECT() *MockVeloClientMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockVeloClient) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockVeloClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockVeloClient)(nil).Close))
}

// SetKeyPair mocks base method
func (m *MockVeloClient) SetKeyPair(keyPair *keypair.Full) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetKeyPair", keyPair)
}

// SetKeyPair indicates an expected call of SetKeyPair
func (mr *MockVeloClientMockRecorder) SetKeyPair(keyPair interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetKeyPair", reflect.TypeOf((*MockVeloClient)(nil).SetKeyPair), keyPair)
}

// Whitelist mocks base method
func (m *MockVeloClient) Whitelist(ctx context.Context, veloOp txnbuild.Whitelist) (client.WhitelistResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Whitelist", ctx, veloOp)
	ret0, _ := ret[0].(client.WhitelistResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Whitelist indicates an expected call of Whitelist
func (mr *MockVeloClientMockRecorder) Whitelist(ctx, veloOp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Whitelist", reflect.TypeOf((*MockVeloClient)(nil).Whitelist), ctx, veloOp)
}

// SetupCredit mocks base method
func (m *MockVeloClient) SetupCredit(ctx context.Context, veloOp txnbuild.SetupCredit) (client.SetupCreditResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetupCredit", ctx, veloOp)
	ret0, _ := ret[0].(client.SetupCreditResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetupCredit indicates an expected call of SetupCredit
func (mr *MockVeloClientMockRecorder) SetupCredit(ctx, veloOp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetupCredit", reflect.TypeOf((*MockVeloClient)(nil).SetupCredit), ctx, veloOp)
}

// PriceUpdate mocks base method
func (m *MockVeloClient) PriceUpdate(ctx context.Context, veloOp txnbuild.PriceUpdate) (client.PriceUpdateResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PriceUpdate", ctx, veloOp)
	ret0, _ := ret[0].(client.PriceUpdateResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PriceUpdate indicates an expected call of PriceUpdate
func (mr *MockVeloClientMockRecorder) PriceUpdate(ctx, veloOp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PriceUpdate", reflect.TypeOf((*MockVeloClient)(nil).PriceUpdate), ctx, veloOp)
}

// MintCredit mocks base method
func (m *MockVeloClient) MintCredit(ctx context.Context, veloOp txnbuild.MintCredit) (client.MintCreditResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MintCredit", ctx, veloOp)
	ret0, _ := ret[0].(client.MintCreditResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MintCredit indicates an expected call of MintCredit
func (mr *MockVeloClientMockRecorder) MintCredit(ctx, veloOp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MintCredit", reflect.TypeOf((*MockVeloClient)(nil).MintCredit), ctx, veloOp)
}

// RedeemCredit mocks base method
func (m *MockVeloClient) RedeemCredit(ctx context.Context, veloOp txnbuild.RedeemCredit) (client.RedeemCreditResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RedeemCredit", ctx, veloOp)
	ret0, _ := ret[0].(client.RedeemCreditResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RedeemCredit indicates an expected call of RedeemCredit
func (mr *MockVeloClientMockRecorder) RedeemCredit(ctx, veloOp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RedeemCredit", reflect.TypeOf((*MockVeloClient)(nil).RedeemCredit), ctx, veloOp)
}

// RebalanceReserve mocks base method
func (m *MockVeloClient) RebalanceReserve(ctx context.Context, veloOp txnbuild.RebalanceReserve) (client.RebalanceReserveResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RebalanceReserve", ctx, veloOp)
	ret0, _ := ret[0].(client.RebalanceReserveResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RebalanceReserve indicates an expected call of RebalanceReserve
func (mr *MockVeloClientMockRecorder) RebalanceReserve(ctx, veloOp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RebalanceReserve", reflect.TypeOf((*MockVeloClient)(nil).RebalanceReserve), ctx, veloOp)
}

// GetExchangeRate mocks base method
func (m *MockVeloClient) GetExchangeRate(ctx context.Context, request *grpc.GetExchangeRateRequest) (*grpc.GetExchangeRateReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExchangeRate", ctx, request)
	ret0, _ := ret[0].(*grpc.GetExchangeRateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExchangeRate indicates an expected call of GetExchangeRate
func (mr *MockVeloClientMockRecorder) GetExchangeRate(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExchangeRate", reflect.TypeOf((*MockVeloClient)(nil).GetExchangeRate), ctx, request)
}

// GetCollateralHealthCheck mocks base method
func (m *MockVeloClient) GetCollateralHealthCheck(ctx context.Context, request *grpc.GetCollateralHealthCheckRequest) (*grpc.GetCollateralHealthCheckReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCollateralHealthCheck", ctx, request)
	ret0, _ := ret[0].(*grpc.GetCollateralHealthCheckReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCollateralHealthCheck indicates an expected call of GetCollateralHealthCheck
func (mr *MockVeloClientMockRecorder) GetCollateralHealthCheck(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCollateralHealthCheck", reflect.TypeOf((*MockVeloClient)(nil).GetCollateralHealthCheck), ctx, request)
}
