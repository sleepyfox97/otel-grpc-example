// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: src/api/uid/uid.proto

package api

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
	UIDService_Generate_FullMethodName = "/moa.api.uid.v1.UIDService/Generate"
)

// UIDServiceClient is the client API for UIDService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UIDServiceClient interface {
	Generate(ctx context.Context, in *GenerateRequest, opts ...grpc.CallOption) (*GenerateResponse, error)
}

type uIDServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUIDServiceClient(cc grpc.ClientConnInterface) UIDServiceClient {
	return &uIDServiceClient{cc}
}

func (c *uIDServiceClient) Generate(ctx context.Context, in *GenerateRequest, opts ...grpc.CallOption) (*GenerateResponse, error) {
	out := new(GenerateResponse)
	err := c.cc.Invoke(ctx, UIDService_Generate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UIDServiceServer is the server API for UIDService service.
// All implementations must embed UnimplementedUIDServiceServer
// for forward compatibility
type UIDServiceServer interface {
	Generate(context.Context, *GenerateRequest) (*GenerateResponse, error)
	mustEmbedUnimplementedUIDServiceServer()
}

// UnimplementedUIDServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUIDServiceServer struct {
}

func (UnimplementedUIDServiceServer) Generate(context.Context, *GenerateRequest) (*GenerateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Generate not implemented")
}
func (UnimplementedUIDServiceServer) mustEmbedUnimplementedUIDServiceServer() {}

// UnsafeUIDServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UIDServiceServer will
// result in compilation errors.
type UnsafeUIDServiceServer interface {
	mustEmbedUnimplementedUIDServiceServer()
}

func RegisterUIDServiceServer(s grpc.ServiceRegistrar, srv UIDServiceServer) {
	s.RegisterService(&UIDService_ServiceDesc, srv)
}

func _UIDService_Generate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UIDServiceServer).Generate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UIDService_Generate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UIDServiceServer).Generate(ctx, req.(*GenerateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UIDService_ServiceDesc is the grpc.ServiceDesc for UIDService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UIDService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "moa.api.uid.v1.UIDService",
	HandlerType: (*UIDServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Generate",
			Handler:    _UIDService_Generate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/api/uid/uid.proto",
}
