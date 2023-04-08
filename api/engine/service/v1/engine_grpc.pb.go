// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: engine.proto

package v1

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

const (
	Engine_CreateOrder_FullMethodName = "/api.engine.service.v1.Engine/CreateOrder"
	Engine_CancelOrder_FullMethodName = "/api.engine.service.v1.Engine/CancelOrder"
	Engine_AddSymbol_FullMethodName   = "/api.engine.service.v1.Engine/AddSymbol"
	Engine_GetEngine_FullMethodName   = "/api.engine.service.v1.Engine/GetEngine"
)

// EngineClient is the client API for Engine service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EngineClient interface {
	CreateOrder(ctx context.Context, in *AddOrderReq, opts ...grpc.CallOption) (*AddOrderReply, error)
	CancelOrder(ctx context.Context, in *CancelOrderReq, opts ...grpc.CallOption) (*CancelOrderReply, error)
	AddSymbol(ctx context.Context, in *AddSymbolReq, opts ...grpc.CallOption) (*AddSymbolReply, error)
	GetEngine(ctx context.Context, in *CloseSymbolReq, opts ...grpc.CallOption) (*CloseSymbolReply, error)
}

type engineClient struct {
	cc grpc.ClientConnInterface
}

func NewEngineClient(cc grpc.ClientConnInterface) EngineClient {
	return &engineClient{cc}
}

func (c *engineClient) CreateOrder(ctx context.Context, in *AddOrderReq, opts ...grpc.CallOption) (*AddOrderReply, error) {
	out := new(AddOrderReply)
	err := c.cc.Invoke(ctx, Engine_CreateOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *engineClient) CancelOrder(ctx context.Context, in *CancelOrderReq, opts ...grpc.CallOption) (*CancelOrderReply, error) {
	out := new(CancelOrderReply)
	err := c.cc.Invoke(ctx, Engine_CancelOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *engineClient) AddSymbol(ctx context.Context, in *AddSymbolReq, opts ...grpc.CallOption) (*AddSymbolReply, error) {
	out := new(AddSymbolReply)
	err := c.cc.Invoke(ctx, Engine_AddSymbol_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *engineClient) GetEngine(ctx context.Context, in *CloseSymbolReq, opts ...grpc.CallOption) (*CloseSymbolReply, error) {
	out := new(CloseSymbolReply)
	err := c.cc.Invoke(ctx, Engine_GetEngine_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EngineServer is the server API for Engine service.
// All implementations must embed UnimplementedEngineServer
// for forward compatibility
type EngineServer interface {
	CreateOrder(context.Context, *AddOrderReq) (*AddOrderReply, error)
	CancelOrder(context.Context, *CancelOrderReq) (*CancelOrderReply, error)
	AddSymbol(context.Context, *AddSymbolReq) (*AddSymbolReply, error)
	GetEngine(context.Context, *CloseSymbolReq) (*CloseSymbolReply, error)
	mustEmbedUnimplementedEngineServer()
}

// UnimplementedEngineServer must be embedded to have forward compatible implementations.
type UnimplementedEngineServer struct {
}

func (UnimplementedEngineServer) CreateOrder(context.Context, *AddOrderReq) (*AddOrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedEngineServer) CancelOrder(context.Context, *CancelOrderReq) (*CancelOrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelOrder not implemented")
}
func (UnimplementedEngineServer) AddSymbol(context.Context, *AddSymbolReq) (*AddSymbolReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSymbol not implemented")
}
func (UnimplementedEngineServer) GetEngine(context.Context, *CloseSymbolReq) (*CloseSymbolReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEngine not implemented")
}
func (UnimplementedEngineServer) mustEmbedUnimplementedEngineServer() {}

// UnsafeEngineServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EngineServer will
// result in compilation errors.
type UnsafeEngineServer interface {
	mustEmbedUnimplementedEngineServer()
}

func RegisterEngineServer(s grpc.ServiceRegistrar, srv EngineServer) {
	s.RegisterService(&Engine_ServiceDesc, srv)
}

func _Engine_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Engine_CreateOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).CreateOrder(ctx, req.(*AddOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Engine_CancelOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).CancelOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Engine_CancelOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).CancelOrder(ctx, req.(*CancelOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Engine_AddSymbol_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSymbolReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).AddSymbol(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Engine_AddSymbol_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).AddSymbol(ctx, req.(*AddSymbolReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Engine_GetEngine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseSymbolReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).GetEngine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Engine_GetEngine_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).GetEngine(ctx, req.(*CloseSymbolReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Engine_ServiceDesc is the grpc.ServiceDesc for Engine service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Engine_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.engine.service.v1.Engine",
	HandlerType: (*EngineServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrder",
			Handler:    _Engine_CreateOrder_Handler,
		},
		{
			MethodName: "CancelOrder",
			Handler:    _Engine_CancelOrder_Handler,
		},
		{
			MethodName: "AddSymbol",
			Handler:    _Engine_AddSymbol_Handler,
		},
		{
			MethodName: "GetEngine",
			Handler:    _Engine_GetEngine_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "engine.proto",
}
