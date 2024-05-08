// Code generated by MockGen. DO NOT EDIT.
// Source: ./gen/citilink/store/stock/v1/stock_api_grpc.pb.go

// Package mock_stockv1 is a generated GoMock package.
package mock_stockv1

import (
	context "context"
	reflect "reflect"

	stockv1 "go.citilink.cloud/grpc-skeleton/internal/specs/grpcclient/gen/citilink/store/stock/v1"
	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockStockAPIClient is a mock of StockAPIClient interface.
type MockStockAPIClient struct {
	ctrl     *gomock.Controller
	recorder *MockStockAPIClientMockRecorder
}

// MockStockAPIClientMockRecorder is the mock recorder for MockStockAPIClient.
type MockStockAPIClientMockRecorder struct {
	mock *MockStockAPIClient
}

// NewMockStockAPIClient creates a new mock instance.
func NewMockStockAPIClient(ctrl *gomock.Controller) *MockStockAPIClient {
	mock := &MockStockAPIClient{ctrl: ctrl}
	mock.recorder = &MockStockAPIClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStockAPIClient) EXPECT() *MockStockAPIClientMockRecorder {
	return m.recorder
}

// FindAllBySpaceId mocks base method.
func (m *MockStockAPIClient) FindAllBySpaceId(ctx context.Context, in *stockv1.FindAllBySpaceIdRequest, opts ...grpc.CallOption) (*stockv1.FindAllBySpaceIdResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindAllBySpaceId", varargs...)
	ret0, _ := ret[0].(*stockv1.FindAllBySpaceIdResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllBySpaceId indicates an expected call of FindAllBySpaceId.
func (mr *MockStockAPIClientMockRecorder) FindAllBySpaceId(ctx, in interface{}, opts ...interface{}) *StockAPIClientFindAllBySpaceIdCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllBySpaceId", reflect.TypeOf((*MockStockAPIClient)(nil).FindAllBySpaceId), varargs...)
	return &StockAPIClientFindAllBySpaceIdCall{Call: call}
}

// StockAPIClientFindAllBySpaceIdCall wrap *gomock.Call
type StockAPIClientFindAllBySpaceIdCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StockAPIClientFindAllBySpaceIdCall) Return(arg0 *stockv1.FindAllBySpaceIdResponse, arg1 error) *StockAPIClientFindAllBySpaceIdCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StockAPIClientFindAllBySpaceIdCall) Do(f func(context.Context, *stockv1.FindAllBySpaceIdRequest, ...grpc.CallOption) (*stockv1.FindAllBySpaceIdResponse, error)) *StockAPIClientFindAllBySpaceIdCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StockAPIClientFindAllBySpaceIdCall) DoAndReturn(f func(context.Context, *stockv1.FindAllBySpaceIdRequest, ...grpc.CallOption) (*stockv1.FindAllBySpaceIdResponse, error)) *StockAPIClientFindAllBySpaceIdCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockStockAPIServer is a mock of StockAPIServer interface.
type MockStockAPIServer struct {
	ctrl     *gomock.Controller
	recorder *MockStockAPIServerMockRecorder
}

// MockStockAPIServerMockRecorder is the mock recorder for MockStockAPIServer.
type MockStockAPIServerMockRecorder struct {
	mock *MockStockAPIServer
}

// NewMockStockAPIServer creates a new mock instance.
func NewMockStockAPIServer(ctrl *gomock.Controller) *MockStockAPIServer {
	mock := &MockStockAPIServer{ctrl: ctrl}
	mock.recorder = &MockStockAPIServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStockAPIServer) EXPECT() *MockStockAPIServerMockRecorder {
	return m.recorder
}

// FindAllBySpaceId mocks base method.
func (m *MockStockAPIServer) FindAllBySpaceId(arg0 context.Context, arg1 *stockv1.FindAllBySpaceIdRequest) (*stockv1.FindAllBySpaceIdResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllBySpaceId", arg0, arg1)
	ret0, _ := ret[0].(*stockv1.FindAllBySpaceIdResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllBySpaceId indicates an expected call of FindAllBySpaceId.
func (mr *MockStockAPIServerMockRecorder) FindAllBySpaceId(arg0, arg1 interface{}) *StockAPIServerFindAllBySpaceIdCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllBySpaceId", reflect.TypeOf((*MockStockAPIServer)(nil).FindAllBySpaceId), arg0, arg1)
	return &StockAPIServerFindAllBySpaceIdCall{Call: call}
}

// StockAPIServerFindAllBySpaceIdCall wrap *gomock.Call
type StockAPIServerFindAllBySpaceIdCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StockAPIServerFindAllBySpaceIdCall) Return(arg0 *stockv1.FindAllBySpaceIdResponse, arg1 error) *StockAPIServerFindAllBySpaceIdCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StockAPIServerFindAllBySpaceIdCall) Do(f func(context.Context, *stockv1.FindAllBySpaceIdRequest) (*stockv1.FindAllBySpaceIdResponse, error)) *StockAPIServerFindAllBySpaceIdCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StockAPIServerFindAllBySpaceIdCall) DoAndReturn(f func(context.Context, *stockv1.FindAllBySpaceIdRequest) (*stockv1.FindAllBySpaceIdResponse, error)) *StockAPIServerFindAllBySpaceIdCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// mustEmbedUnimplementedStockAPIServer mocks base method.
func (m *MockStockAPIServer) mustEmbedUnimplementedStockAPIServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedStockAPIServer")
}

// mustEmbedUnimplementedStockAPIServer indicates an expected call of mustEmbedUnimplementedStockAPIServer.
func (mr *MockStockAPIServerMockRecorder) mustEmbedUnimplementedStockAPIServer() *StockAPIServermustEmbedUnimplementedStockAPIServerCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedStockAPIServer", reflect.TypeOf((*MockStockAPIServer)(nil).mustEmbedUnimplementedStockAPIServer))
	return &StockAPIServermustEmbedUnimplementedStockAPIServerCall{Call: call}
}

// StockAPIServermustEmbedUnimplementedStockAPIServerCall wrap *gomock.Call
type StockAPIServermustEmbedUnimplementedStockAPIServerCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *StockAPIServermustEmbedUnimplementedStockAPIServerCall) Return() *StockAPIServermustEmbedUnimplementedStockAPIServerCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *StockAPIServermustEmbedUnimplementedStockAPIServerCall) Do(f func()) *StockAPIServermustEmbedUnimplementedStockAPIServerCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *StockAPIServermustEmbedUnimplementedStockAPIServerCall) DoAndReturn(f func()) *StockAPIServermustEmbedUnimplementedStockAPIServerCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockUnsafeStockAPIServer is a mock of UnsafeStockAPIServer interface.
type MockUnsafeStockAPIServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeStockAPIServerMockRecorder
}

// MockUnsafeStockAPIServerMockRecorder is the mock recorder for MockUnsafeStockAPIServer.
type MockUnsafeStockAPIServerMockRecorder struct {
	mock *MockUnsafeStockAPIServer
}

// NewMockUnsafeStockAPIServer creates a new mock instance.
func NewMockUnsafeStockAPIServer(ctrl *gomock.Controller) *MockUnsafeStockAPIServer {
	mock := &MockUnsafeStockAPIServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeStockAPIServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeStockAPIServer) EXPECT() *MockUnsafeStockAPIServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedStockAPIServer mocks base method.
func (m *MockUnsafeStockAPIServer) mustEmbedUnimplementedStockAPIServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedStockAPIServer")
}

// mustEmbedUnimplementedStockAPIServer indicates an expected call of mustEmbedUnimplementedStockAPIServer.
func (mr *MockUnsafeStockAPIServerMockRecorder) mustEmbedUnimplementedStockAPIServer() *UnsafeStockAPIServermustEmbedUnimplementedStockAPIServerCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedStockAPIServer", reflect.TypeOf((*MockUnsafeStockAPIServer)(nil).mustEmbedUnimplementedStockAPIServer))
	return &UnsafeStockAPIServermustEmbedUnimplementedStockAPIServerCall{Call: call}
}

// UnsafeStockAPIServermustEmbedUnimplementedStockAPIServerCall wrap *gomock.Call
type UnsafeStockAPIServermustEmbedUnimplementedStockAPIServerCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *UnsafeStockAPIServermustEmbedUnimplementedStockAPIServerCall) Return() *UnsafeStockAPIServermustEmbedUnimplementedStockAPIServerCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *UnsafeStockAPIServermustEmbedUnimplementedStockAPIServerCall) Do(f func()) *UnsafeStockAPIServermustEmbedUnimplementedStockAPIServerCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *UnsafeStockAPIServermustEmbedUnimplementedStockAPIServerCall) DoAndReturn(f func()) *UnsafeStockAPIServermustEmbedUnimplementedStockAPIServerCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}