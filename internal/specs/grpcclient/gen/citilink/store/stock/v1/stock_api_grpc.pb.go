// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: citilink/store/stock/v1/stock_api.proto

package stockv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// StockAPIClient is the client API for StockAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StockAPIClient interface {
	// Находит все склады в пространстве.
	FindAllBySpaceId(ctx context.Context, in *FindAllBySpaceIdRequest, opts ...grpc.CallOption) (*FindAllBySpaceIdResponse, error)
}

type stockAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewStockAPIClient(cc grpc.ClientConnInterface) StockAPIClient {
	return &stockAPIClient{cc}
}

func (c *stockAPIClient) FindAllBySpaceId(ctx context.Context, in *FindAllBySpaceIdRequest, opts ...grpc.CallOption) (*FindAllBySpaceIdResponse, error) {
	out := new(FindAllBySpaceIdResponse)
	err := c.cc.Invoke(ctx, "/citilink.store.stock.v1.StockAPI/FindAllBySpaceId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StockAPIServer is the server API for StockAPI service.
// All implementations must embed UnimplementedStockAPIServer
// for forward compatibility
type StockAPIServer interface {
	// Находит все склады в пространстве.
	FindAllBySpaceId(context.Context, *FindAllBySpaceIdRequest) (*FindAllBySpaceIdResponse, error)
	mustEmbedUnimplementedStockAPIServer()
}

// UnimplementedStockAPIServer must be embedded to have forward compatible implementations.
type UnimplementedStockAPIServer struct {
}

func (UnimplementedStockAPIServer) FindAllBySpaceId(context.Context, *FindAllBySpaceIdRequest) (*FindAllBySpaceIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllBySpaceId not implemented")
}
func (UnimplementedStockAPIServer) mustEmbedUnimplementedStockAPIServer() {}

// UnsafeStockAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StockAPIServer will
// result in compilation errors.
type UnsafeStockAPIServer interface {
	mustEmbedUnimplementedStockAPIServer()
}

func RegisterStockAPIServer(s grpc.ServiceRegistrar, srv StockAPIServer) {
	s.RegisterService(&StockAPI_ServiceDesc, srv)
}

func _StockAPI_FindAllBySpaceId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllBySpaceIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockAPIServer).FindAllBySpaceId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/citilink.store.stock.v1.StockAPI/FindAllBySpaceId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockAPIServer).FindAllBySpaceId(ctx, req.(*FindAllBySpaceIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StockAPI_ServiceDesc is the grpc.ServiceDesc for StockAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StockAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "citilink.store.stock.v1.StockAPI",
	HandlerType: (*StockAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindAllBySpaceId",
			Handler:    _StockAPI_FindAllBySpaceId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "citilink/store/stock/v1/stock_api.proto",
}
