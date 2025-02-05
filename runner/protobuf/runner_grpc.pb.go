// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: runner.proto

package protobuf

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	RunnerService_RequestStream_FullMethodName = "/protobuf.RunnerService/RequestStream"
)

// RunnerServiceClient is the client API for RunnerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RunnerServiceClient interface {
	// Requests a stream from a lecture hall
	RequestStream(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (*StreamResponse, error)
}

type runnerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRunnerServiceClient(cc grpc.ClientConnInterface) RunnerServiceClient {
	return &runnerServiceClient{cc}
}

func (c *runnerServiceClient) RequestStream(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (*StreamResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StreamResponse)
	err := c.cc.Invoke(ctx, RunnerService_RequestStream_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RunnerServiceServer is the server API for RunnerService service.
// All implementations must embed UnimplementedRunnerServiceServer
// for forward compatibility.
type RunnerServiceServer interface {
	// Requests a stream from a lecture hall
	RequestStream(context.Context, *StreamRequest) (*StreamResponse, error)
	mustEmbedUnimplementedRunnerServiceServer()
}

// UnimplementedRunnerServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRunnerServiceServer struct{}

func (UnimplementedRunnerServiceServer) RequestStream(context.Context, *StreamRequest) (*StreamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestStream not implemented")
}
func (UnimplementedRunnerServiceServer) mustEmbedUnimplementedRunnerServiceServer() {}
func (UnimplementedRunnerServiceServer) testEmbeddedByValue()                       {}

// UnsafeRunnerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RunnerServiceServer will
// result in compilation errors.
type UnsafeRunnerServiceServer interface {
	mustEmbedUnimplementedRunnerServiceServer()
}

func RegisterRunnerServiceServer(s grpc.ServiceRegistrar, srv RunnerServiceServer) {
	// If the following call pancis, it indicates UnimplementedRunnerServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RunnerService_ServiceDesc, srv)
}

func _RunnerService_RequestStream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunnerServiceServer).RequestStream(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RunnerService_RequestStream_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunnerServiceServer).RequestStream(ctx, req.(*StreamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RunnerService_ServiceDesc is the grpc.ServiceDesc for RunnerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RunnerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.RunnerService",
	HandlerType: (*RunnerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestStream",
			Handler:    _RunnerService_RequestStream_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "runner.proto",
}

const (
	RunnerManagerService_Register_FullMethodName = "/protobuf.RunnerManagerService/Register"
)

// RunnerManagerServiceClient is the client API for RunnerManagerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// RunnerManagerService service defines communication from runners to gocast
type RunnerManagerServiceClient interface {
	// Register is a request to the server to join the runners pool.
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
}

type runnerManagerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRunnerManagerServiceClient(cc grpc.ClientConnInterface) RunnerManagerServiceClient {
	return &runnerManagerServiceClient{cc}
}

func (c *runnerManagerServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, RunnerManagerService_Register_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RunnerManagerServiceServer is the server API for RunnerManagerService service.
// All implementations must embed UnimplementedRunnerManagerServiceServer
// for forward compatibility.
//
// RunnerManagerService service defines communication from runners to gocast
type RunnerManagerServiceServer interface {
	// Register is a request to the server to join the runners pool.
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	mustEmbedUnimplementedRunnerManagerServiceServer()
}

// UnimplementedRunnerManagerServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRunnerManagerServiceServer struct{}

func (UnimplementedRunnerManagerServiceServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedRunnerManagerServiceServer) mustEmbedUnimplementedRunnerManagerServiceServer() {}
func (UnimplementedRunnerManagerServiceServer) testEmbeddedByValue()                              {}

// UnsafeRunnerManagerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RunnerManagerServiceServer will
// result in compilation errors.
type UnsafeRunnerManagerServiceServer interface {
	mustEmbedUnimplementedRunnerManagerServiceServer()
}

func RegisterRunnerManagerServiceServer(s grpc.ServiceRegistrar, srv RunnerManagerServiceServer) {
	// If the following call pancis, it indicates UnimplementedRunnerManagerServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RunnerManagerService_ServiceDesc, srv)
}

func _RunnerManagerService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunnerManagerServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RunnerManagerService_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunnerManagerServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RunnerManagerService_ServiceDesc is the grpc.ServiceDesc for RunnerManagerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RunnerManagerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.RunnerManagerService",
	HandlerType: (*RunnerManagerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _RunnerManagerService_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "runner.proto",
}
