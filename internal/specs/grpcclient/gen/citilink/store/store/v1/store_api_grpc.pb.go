// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: citilink/store/store/v1/store_api.proto

package storev1

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

// StoreAPIClient is the client API for StoreAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StoreAPIClient interface {
	// Возвращает все магазины. Метод не учитывает заполненность точки (need_to_fill)
	FindAll(ctx context.Context, in *FindAllRequest, opts ...grpc.CallOption) (*FindAllResponse, error)
	// Возвращает отфильтрованные магазины. Если фильтров нет - возвращает все магазины
	// по умолчанию возвращает все не скрытые (isHidden = false) и заполненные (need_to_fill == false) точки
	Filter(ctx context.Context, in *FilterRequest, opts ...grpc.CallOption) (*FilterResponse, error)
	// Возвращает список идентификаторов пространства.
	GetSpaceIds(ctx context.Context, in *GetSpaceIdsRequest, opts ...grpc.CallOption) (*GetSpaceIdsResponse, error)
	// Возвращает список терминальных идентификаторов пространства.
	GetTerminalSpaceIds(ctx context.Context, in *GetTerminalSpaceIdsRequest, opts ...grpc.CallOption) (*GetTerminalSpaceIdsResponse, error)
}

type storeAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewStoreAPIClient(cc grpc.ClientConnInterface) StoreAPIClient {
	return &storeAPIClient{cc}
}

func (c *storeAPIClient) FindAll(ctx context.Context, in *FindAllRequest, opts ...grpc.CallOption) (*FindAllResponse, error) {
	out := new(FindAllResponse)
	err := c.cc.Invoke(ctx, "/citilink.store.store.v1.StoreAPI/FindAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeAPIClient) Filter(ctx context.Context, in *FilterRequest, opts ...grpc.CallOption) (*FilterResponse, error) {
	out := new(FilterResponse)
	err := c.cc.Invoke(ctx, "/citilink.store.store.v1.StoreAPI/Filter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeAPIClient) GetSpaceIds(ctx context.Context, in *GetSpaceIdsRequest, opts ...grpc.CallOption) (*GetSpaceIdsResponse, error) {
	out := new(GetSpaceIdsResponse)
	err := c.cc.Invoke(ctx, "/citilink.store.store.v1.StoreAPI/GetSpaceIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeAPIClient) GetTerminalSpaceIds(ctx context.Context, in *GetTerminalSpaceIdsRequest, opts ...grpc.CallOption) (*GetTerminalSpaceIdsResponse, error) {
	out := new(GetTerminalSpaceIdsResponse)
	err := c.cc.Invoke(ctx, "/citilink.store.store.v1.StoreAPI/GetTerminalSpaceIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoreAPIServer is the server API for StoreAPI service.
// All implementations must embed UnimplementedStoreAPIServer
// for forward compatibility
type StoreAPIServer interface {
	// Возвращает все магазины. Метод не учитывает заполненность точки (need_to_fill)
	FindAll(context.Context, *FindAllRequest) (*FindAllResponse, error)
	// Возвращает отфильтрованные магазины. Если фильтров нет - возвращает все магазины
	// по умолчанию возвращает все не скрытые (isHidden = false) и заполненные (need_to_fill == false) точки
	Filter(context.Context, *FilterRequest) (*FilterResponse, error)
	// Возвращает список идентификаторов пространства.
	GetSpaceIds(context.Context, *GetSpaceIdsRequest) (*GetSpaceIdsResponse, error)
	// Возвращает список терминальных идентификаторов пространства.
	GetTerminalSpaceIds(context.Context, *GetTerminalSpaceIdsRequest) (*GetTerminalSpaceIdsResponse, error)
	mustEmbedUnimplementedStoreAPIServer()
}

// UnimplementedStoreAPIServer must be embedded to have forward compatible implementations.
type UnimplementedStoreAPIServer struct {
}

func (UnimplementedStoreAPIServer) FindAll(context.Context, *FindAllRequest) (*FindAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAll not implemented")
}
func (UnimplementedStoreAPIServer) Filter(context.Context, *FilterRequest) (*FilterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Filter not implemented")
}
func (UnimplementedStoreAPIServer) GetSpaceIds(context.Context, *GetSpaceIdsRequest) (*GetSpaceIdsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSpaceIds not implemented")
}
func (UnimplementedStoreAPIServer) GetTerminalSpaceIds(context.Context, *GetTerminalSpaceIdsRequest) (*GetTerminalSpaceIdsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTerminalSpaceIds not implemented")
}
func (UnimplementedStoreAPIServer) mustEmbedUnimplementedStoreAPIServer() {}

// UnsafeStoreAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StoreAPIServer will
// result in compilation errors.
type UnsafeStoreAPIServer interface {
	mustEmbedUnimplementedStoreAPIServer()
}

func RegisterStoreAPIServer(s grpc.ServiceRegistrar, srv StoreAPIServer) {
	s.RegisterService(&StoreAPI_ServiceDesc, srv)
}

func _StoreAPI_FindAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreAPIServer).FindAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/citilink.store.store.v1.StoreAPI/FindAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreAPIServer).FindAll(ctx, req.(*FindAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreAPI_Filter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreAPIServer).Filter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/citilink.store.store.v1.StoreAPI/Filter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreAPIServer).Filter(ctx, req.(*FilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreAPI_GetSpaceIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSpaceIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreAPIServer).GetSpaceIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/citilink.store.store.v1.StoreAPI/GetSpaceIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreAPIServer).GetSpaceIds(ctx, req.(*GetSpaceIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreAPI_GetTerminalSpaceIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTerminalSpaceIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreAPIServer).GetTerminalSpaceIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/citilink.store.store.v1.StoreAPI/GetTerminalSpaceIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreAPIServer).GetTerminalSpaceIds(ctx, req.(*GetTerminalSpaceIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StoreAPI_ServiceDesc is the grpc.ServiceDesc for StoreAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StoreAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "citilink.store.store.v1.StoreAPI",
	HandlerType: (*StoreAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindAll",
			Handler:    _StoreAPI_FindAll_Handler,
		},
		{
			MethodName: "Filter",
			Handler:    _StoreAPI_Filter_Handler,
		},
		{
			MethodName: "GetSpaceIds",
			Handler:    _StoreAPI_GetSpaceIds_Handler,
		},
		{
			MethodName: "GetTerminalSpaceIds",
			Handler:    _StoreAPI_GetTerminalSpaceIds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "citilink/store/store/v1/store_api.proto",
}