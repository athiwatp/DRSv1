// Code generated by MockGen. DO NOT EDIT.
// Source: ../grpc/velo_node.pb.go

// Package mock_grpc is a generated GoMock package.
package mock_grpc

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	grpc "github.com/velo-protocol/DRSv1/grpc"
	grpc0 "google.golang.org/grpc"
	reflect "reflect"
)

// MockVeloNodeClient is a mock of VeloNodeClient interface
type MockVeloNodeClient struct {
	ctrl     *gomock.Controller
	recorder *MockVeloNodeClientMockRecorder
}

// MockVeloNodeClientMockRecorder is the mock recorder for MockVeloNodeClient
type MockVeloNodeClientMockRecorder struct {
	mock *MockVeloNodeClient
}

// NewMockVeloNodeClient creates a new mock instance
func NewMockVeloNodeClient(ctrl *gomock.Controller) *MockVeloNodeClient {
	mock := &MockVeloNodeClient{ctrl: ctrl}
	mock.recorder = &MockVeloNodeClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockVeloNodeClient) EXPECT() *MockVeloNodeClientMockRecorder {
	return m.recorder
}

// SubmitVeloTx mocks base method
func (m *MockVeloNodeClient) SubmitVeloTx(ctx context.Context, in *grpc.VeloTxRequest, opts ...grpc0.CallOption) (*grpc.VeloTxReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SubmitVeloTx", varargs...)
	ret0, _ := ret[0].(*grpc.VeloTxReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitVeloTx indicates an expected call of SubmitVeloTx
func (mr *MockVeloNodeClientMockRecorder) SubmitVeloTx(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitVeloTx", reflect.TypeOf((*MockVeloNodeClient)(nil).SubmitVeloTx), varargs...)
}

// GetExchangeRate mocks base method
func (m *MockVeloNodeClient) GetExchangeRate(ctx context.Context, in *grpc.GetExchangeRateRequest, opts ...grpc0.CallOption) (*grpc.GetExchangeRateReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetExchangeRate", varargs...)
	ret0, _ := ret[0].(*grpc.GetExchangeRateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExchangeRate indicates an expected call of GetExchangeRate
func (mr *MockVeloNodeClientMockRecorder) GetExchangeRate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExchangeRate", reflect.TypeOf((*MockVeloNodeClient)(nil).GetExchangeRate), varargs...)
}

// GetCollateralHealthCheck mocks base method
func (m *MockVeloNodeClient) GetCollateralHealthCheck(ctx context.Context, in *grpc.GetCollateralHealthCheckRequest, opts ...grpc0.CallOption) (*grpc.GetCollateralHealthCheckReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCollateralHealthCheck", varargs...)
	ret0, _ := ret[0].(*grpc.GetCollateralHealthCheckReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCollateralHealthCheck indicates an expected call of GetCollateralHealthCheck
func (mr *MockVeloNodeClientMockRecorder) GetCollateralHealthCheck(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCollateralHealthCheck", reflect.TypeOf((*MockVeloNodeClient)(nil).GetCollateralHealthCheck), varargs...)
}

// MockVeloNodeServer is a mock of VeloNodeServer interface
type MockVeloNodeServer struct {
	ctrl     *gomock.Controller
	recorder *MockVeloNodeServerMockRecorder
}

// MockVeloNodeServerMockRecorder is the mock recorder for MockVeloNodeServer
type MockVeloNodeServerMockRecorder struct {
	mock *MockVeloNodeServer
}

// NewMockVeloNodeServer creates a new mock instance
func NewMockVeloNodeServer(ctrl *gomock.Controller) *MockVeloNodeServer {
	mock := &MockVeloNodeServer{ctrl: ctrl}
	mock.recorder = &MockVeloNodeServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockVeloNodeServer) EXPECT() *MockVeloNodeServerMockRecorder {
	return m.recorder
}

// SubmitVeloTx mocks base method
func (m *MockVeloNodeServer) SubmitVeloTx(arg0 context.Context, arg1 *grpc.VeloTxRequest) (*grpc.VeloTxReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitVeloTx", arg0, arg1)
	ret0, _ := ret[0].(*grpc.VeloTxReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitVeloTx indicates an expected call of SubmitVeloTx
func (mr *MockVeloNodeServerMockRecorder) SubmitVeloTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitVeloTx", reflect.TypeOf((*MockVeloNodeServer)(nil).SubmitVeloTx), arg0, arg1)
}

// GetExchangeRate mocks base method
func (m *MockVeloNodeServer) GetExchangeRate(arg0 context.Context, arg1 *grpc.GetExchangeRateRequest) (*grpc.GetExchangeRateReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExchangeRate", arg0, arg1)
	ret0, _ := ret[0].(*grpc.GetExchangeRateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExchangeRate indicates an expected call of GetExchangeRate
func (mr *MockVeloNodeServerMockRecorder) GetExchangeRate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExchangeRate", reflect.TypeOf((*MockVeloNodeServer)(nil).GetExchangeRate), arg0, arg1)
}

// GetCollateralHealthCheck mocks base method
func (m *MockVeloNodeServer) GetCollateralHealthCheck(arg0 context.Context, arg1 *grpc.GetCollateralHealthCheckRequest) (*grpc.GetCollateralHealthCheckReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCollateralHealthCheck", arg0, arg1)
	ret0, _ := ret[0].(*grpc.GetCollateralHealthCheckReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCollateralHealthCheck indicates an expected call of GetCollateralHealthCheck
func (mr *MockVeloNodeServerMockRecorder) GetCollateralHealthCheck(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCollateralHealthCheck", reflect.TypeOf((*MockVeloNodeServer)(nil).GetCollateralHealthCheck), arg0, arg1)
}
