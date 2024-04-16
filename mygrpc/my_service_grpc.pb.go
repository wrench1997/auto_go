// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.0--rc3
// source: mygrpc/my_service.proto

package mygrpc

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

// MyServiceClient is the client API for MyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MyServiceClient interface {
	Authenticate(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	PressKey(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Click(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ReleaseKey(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type myServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMyServiceClient(cc grpc.ClientConnInterface) MyServiceClient {
	return &myServiceClient{cc}
}

func (c *myServiceClient) Authenticate(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/myservice.MyService/Authenticate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myServiceClient) PressKey(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/myservice.MyService/PressKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myServiceClient) Click(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/myservice.MyService/Click", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myServiceClient) ReleaseKey(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/myservice.MyService/ReleaseKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MyServiceServer is the server API for MyService service.
// All implementations must embed UnimplementedMyServiceServer
// for forward compatibility
type MyServiceServer interface {
	Authenticate(context.Context, *AuthRequest) (*emptypb.Empty, error)
	PressKey(context.Context, *KeyRequest) (*emptypb.Empty, error)
	Click(context.Context, *KeyRequest) (*emptypb.Empty, error)
	ReleaseKey(context.Context, *KeyRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedMyServiceServer()
}

// UnimplementedMyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMyServiceServer struct {
}

func (UnimplementedMyServiceServer) Authenticate(context.Context, *AuthRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authenticate not implemented")
}
func (UnimplementedMyServiceServer) PressKey(context.Context, *KeyRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PressKey not implemented")
}
func (UnimplementedMyServiceServer) Click(context.Context, *KeyRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Click not implemented")
}
func (UnimplementedMyServiceServer) ReleaseKey(context.Context, *KeyRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReleaseKey not implemented")
}
func (UnimplementedMyServiceServer) mustEmbedUnimplementedMyServiceServer() {}

// UnsafeMyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MyServiceServer will
// result in compilation errors.
type UnsafeMyServiceServer interface {
	mustEmbedUnimplementedMyServiceServer()
}

func RegisterMyServiceServer(s grpc.ServiceRegistrar, srv MyServiceServer) {
	s.RegisterService(&MyService_ServiceDesc, srv)
}

func _MyService_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyServiceServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/myservice.MyService/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyServiceServer).Authenticate(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MyService_PressKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyServiceServer).PressKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/myservice.MyService/PressKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyServiceServer).PressKey(ctx, req.(*KeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MyService_Click_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyServiceServer).Click(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/myservice.MyService/Click",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyServiceServer).Click(ctx, req.(*KeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MyService_ReleaseKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyServiceServer).ReleaseKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/myservice.MyService/ReleaseKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyServiceServer).ReleaseKey(ctx, req.(*KeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MyService_ServiceDesc is the grpc.ServiceDesc for MyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "myservice.MyService",
	HandlerType: (*MyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authenticate",
			Handler:    _MyService_Authenticate_Handler,
		},
		{
			MethodName: "PressKey",
			Handler:    _MyService_PressKey_Handler,
		},
		{
			MethodName: "Click",
			Handler:    _MyService_Click_Handler,
		},
		{
			MethodName: "ReleaseKey",
			Handler:    _MyService_ReleaseKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mygrpc/my_service.proto",
}
