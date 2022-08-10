// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: service_1.proto

package service_1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// Service_1Client is the client API for Service_1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Service_1Client interface {
	GetAll(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Status, error)
}

type service_1Client struct {
	cc grpc.ClientConnInterface
}

func NewService_1Client(cc grpc.ClientConnInterface) Service_1Client {
	return &service_1Client{cc}
}

func (c *service_1Client) GetAll(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/service_1.Service_1/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Service_1Server is the server API for Service_1 service.
// All implementations must embed UnimplementedService_1Server
// for forward compatibility
type Service_1Server interface {
	GetAll(context.Context, *emptypb.Empty) (*Status, error)
	mustEmbedUnimplementedService_1Server()
}

// UnimplementedService_1Server must be embedded to have forward compatible implementations.
type UnimplementedService_1Server struct {
}

func (UnimplementedService_1Server) GetAll(context.Context, *emptypb.Empty) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedService_1Server) mustEmbedUnimplementedService_1Server() {}

// UnsafeService_1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Service_1Server will
// result in compilation errors.
type UnsafeService_1Server interface {
	mustEmbedUnimplementedService_1Server()
}

func RegisterService_1Server(s grpc.ServiceRegistrar, srv Service_1Server) {
	s.RegisterService(&Service_1_ServiceDesc, srv)
}

func _Service_1_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Service_1Server).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service_1.Service_1/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Service_1Server).GetAll(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_1_ServiceDesc is the grpc.ServiceDesc for Service_1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service_1.Service_1",
	HandlerType: (*Service_1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAll",
			Handler:    _Service_1_GetAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_1.proto",
}