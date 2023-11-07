// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: proto/streaming/v1alpha1/cover.proto

package streaming

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
	CoverService_CreateCover_FullMethodName = "/proto.streaming.v1alpha1.CoverService/CreateCover"
)

// CoverServiceClient is the client API for CoverService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CoverServiceClient interface {
	CreateCover(ctx context.Context, in *CreateCoverRequest, opts ...grpc.CallOption) (*CreateCoverResponse, error)
}

type coverServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCoverServiceClient(cc grpc.ClientConnInterface) CoverServiceClient {
	return &coverServiceClient{cc}
}

func (c *coverServiceClient) CreateCover(ctx context.Context, in *CreateCoverRequest, opts ...grpc.CallOption) (*CreateCoverResponse, error) {
	out := new(CreateCoverResponse)
	err := c.cc.Invoke(ctx, CoverService_CreateCover_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoverServiceServer is the server API for CoverService service.
// All implementations should embed UnimplementedCoverServiceServer
// for forward compatibility
type CoverServiceServer interface {
	CreateCover(context.Context, *CreateCoverRequest) (*CreateCoverResponse, error)
}

// UnimplementedCoverServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCoverServiceServer struct {
}

func (UnimplementedCoverServiceServer) CreateCover(context.Context, *CreateCoverRequest) (*CreateCoverResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCover not implemented")
}

// UnsafeCoverServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CoverServiceServer will
// result in compilation errors.
type UnsafeCoverServiceServer interface {
	mustEmbedUnimplementedCoverServiceServer()
}

func RegisterCoverServiceServer(s grpc.ServiceRegistrar, srv CoverServiceServer) {
	s.RegisterService(&CoverService_ServiceDesc, srv)
}

func _CoverService_CreateCover_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCoverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoverServiceServer).CreateCover(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CoverService_CreateCover_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoverServiceServer).CreateCover(ctx, req.(*CreateCoverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CoverService_ServiceDesc is the grpc.ServiceDesc for CoverService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CoverService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.streaming.v1alpha1.CoverService",
	HandlerType: (*CoverServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCover",
			Handler:    _CoverService_CreateCover_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/streaming/v1alpha1/cover.proto",
}