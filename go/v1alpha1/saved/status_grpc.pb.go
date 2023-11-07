// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: proto/saved/v1alpha1/status.proto

package saved

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

const (
	SavedService_GetActivateStatus_FullMethodName = "/proto.saved.v1alpha1.SavedService/GetActivateStatus"
	SavedService_Activate_FullMethodName          = "/proto.saved.v1alpha1.SavedService/Activate"
)

// SavedServiceClient is the client API for SavedService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SavedServiceClient interface {
	GetActivateStatus(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetActivateStatusResponse, error)
	Activate(ctx context.Context, in *ActivateRequest, opts ...grpc.CallOption) (*ActivateResponse, error)
}

type savedServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSavedServiceClient(cc grpc.ClientConnInterface) SavedServiceClient {
	return &savedServiceClient{cc}
}

func (c *savedServiceClient) GetActivateStatus(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetActivateStatusResponse, error) {
	out := new(GetActivateStatusResponse)
	err := c.cc.Invoke(ctx, SavedService_GetActivateStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *savedServiceClient) Activate(ctx context.Context, in *ActivateRequest, opts ...grpc.CallOption) (*ActivateResponse, error) {
	out := new(ActivateResponse)
	err := c.cc.Invoke(ctx, SavedService_Activate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SavedServiceServer is the server API for SavedService service.
// All implementations should embed UnimplementedSavedServiceServer
// for forward compatibility
type SavedServiceServer interface {
	GetActivateStatus(context.Context, *emptypb.Empty) (*GetActivateStatusResponse, error)
	Activate(context.Context, *ActivateRequest) (*ActivateResponse, error)
}

// UnimplementedSavedServiceServer should be embedded to have forward compatible implementations.
type UnimplementedSavedServiceServer struct {
}

func (UnimplementedSavedServiceServer) GetActivateStatus(context.Context, *emptypb.Empty) (*GetActivateStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActivateStatus not implemented")
}
func (UnimplementedSavedServiceServer) Activate(context.Context, *ActivateRequest) (*ActivateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Activate not implemented")
}

// UnsafeSavedServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SavedServiceServer will
// result in compilation errors.
type UnsafeSavedServiceServer interface {
	mustEmbedUnimplementedSavedServiceServer()
}

func RegisterSavedServiceServer(s grpc.ServiceRegistrar, srv SavedServiceServer) {
	s.RegisterService(&SavedService_ServiceDesc, srv)
}

func _SavedService_GetActivateStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SavedServiceServer).GetActivateStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SavedService_GetActivateStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SavedServiceServer).GetActivateStatus(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SavedService_Activate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SavedServiceServer).Activate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SavedService_Activate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SavedServiceServer).Activate(ctx, req.(*ActivateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SavedService_ServiceDesc is the grpc.ServiceDesc for SavedService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SavedService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.saved.v1alpha1.SavedService",
	HandlerType: (*SavedServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetActivateStatus",
			Handler:    _SavedService_GetActivateStatus_Handler,
		},
		{
			MethodName: "Activate",
			Handler:    _SavedService_Activate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/saved/v1alpha1/status.proto",
}