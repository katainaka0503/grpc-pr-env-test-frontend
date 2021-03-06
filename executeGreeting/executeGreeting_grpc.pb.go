// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: executeGreeting/executeGreeting.proto

package executeGreeting

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

// ExecuteGreetingClient is the client API for ExecuteGreeting service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExecuteGreetingClient interface {
	// Sends a greeting
	ExecuteGreeting(ctx context.Context, in *ExecuteGreetingRequest, opts ...grpc.CallOption) (*ExecuteGreetingReply, error)
}

type executeGreetingClient struct {
	cc grpc.ClientConnInterface
}

func NewExecuteGreetingClient(cc grpc.ClientConnInterface) ExecuteGreetingClient {
	return &executeGreetingClient{cc}
}

func (c *executeGreetingClient) ExecuteGreeting(ctx context.Context, in *ExecuteGreetingRequest, opts ...grpc.CallOption) (*ExecuteGreetingReply, error) {
	out := new(ExecuteGreetingReply)
	err := c.cc.Invoke(ctx, "/helloworld.ExecuteGreeting/ExecuteGreeting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExecuteGreetingServer is the server API for ExecuteGreeting service.
// All implementations must embed UnimplementedExecuteGreetingServer
// for forward compatibility
type ExecuteGreetingServer interface {
	// Sends a greeting
	ExecuteGreeting(context.Context, *ExecuteGreetingRequest) (*ExecuteGreetingReply, error)
	mustEmbedUnimplementedExecuteGreetingServer()
}

// UnimplementedExecuteGreetingServer must be embedded to have forward compatible implementations.
type UnimplementedExecuteGreetingServer struct {
}

func (UnimplementedExecuteGreetingServer) ExecuteGreeting(context.Context, *ExecuteGreetingRequest) (*ExecuteGreetingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteGreeting not implemented")
}
func (UnimplementedExecuteGreetingServer) mustEmbedUnimplementedExecuteGreetingServer() {}

// UnsafeExecuteGreetingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExecuteGreetingServer will
// result in compilation errors.
type UnsafeExecuteGreetingServer interface {
	mustEmbedUnimplementedExecuteGreetingServer()
}

func RegisterExecuteGreetingServer(s grpc.ServiceRegistrar, srv ExecuteGreetingServer) {
	s.RegisterService(&ExecuteGreeting_ServiceDesc, srv)
}

func _ExecuteGreeting_ExecuteGreeting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteGreetingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecuteGreetingServer).ExecuteGreeting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.ExecuteGreeting/ExecuteGreeting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecuteGreetingServer).ExecuteGreeting(ctx, req.(*ExecuteGreetingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExecuteGreeting_ServiceDesc is the grpc.ServiceDesc for ExecuteGreeting service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExecuteGreeting_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.ExecuteGreeting",
	HandlerType: (*ExecuteGreetingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExecuteGreeting",
			Handler:    _ExecuteGreeting_ExecuteGreeting_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "executeGreeting/executeGreeting.proto",
}
