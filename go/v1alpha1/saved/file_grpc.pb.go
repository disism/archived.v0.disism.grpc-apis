// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: proto/saved/v1alpha1/file.proto

package saved

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
	FileService_CreateFiles_FullMethodName = "/proto.saved.v1alpha1.FileService/CreateFiles"
	FileService_GetFiles_FullMethodName    = "/proto.saved.v1alpha1.FileService/GetFiles"
	FileService_GetFile_FullMethodName     = "/proto.saved.v1alpha1.FileService/GetFile"
	FileService_DeleteFile_FullMethodName  = "/proto.saved.v1alpha1.FileService/DeleteFile"
)

// FileServiceClient is the client API for FileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileServiceClient interface {
	CreateFiles(ctx context.Context, in *CreateFilesRequest, opts ...grpc.CallOption) (*CreateFilesResponse, error)
	GetFiles(ctx context.Context, in *GetFilesRequest, opts ...grpc.CallOption) (*GetFilesResponse, error)
	GetFile(ctx context.Context, in *GetFileRequest, opts ...grpc.CallOption) (*GetFileResponse, error)
	DeleteFile(ctx context.Context, in *DeleteFileRequest, opts ...grpc.CallOption) (*DeleteFileResponse, error)
}

type fileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileServiceClient(cc grpc.ClientConnInterface) FileServiceClient {
	return &fileServiceClient{cc}
}

func (c *fileServiceClient) CreateFiles(ctx context.Context, in *CreateFilesRequest, opts ...grpc.CallOption) (*CreateFilesResponse, error) {
	out := new(CreateFilesResponse)
	err := c.cc.Invoke(ctx, FileService_CreateFiles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) GetFiles(ctx context.Context, in *GetFilesRequest, opts ...grpc.CallOption) (*GetFilesResponse, error) {
	out := new(GetFilesResponse)
	err := c.cc.Invoke(ctx, FileService_GetFiles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) GetFile(ctx context.Context, in *GetFileRequest, opts ...grpc.CallOption) (*GetFileResponse, error) {
	out := new(GetFileResponse)
	err := c.cc.Invoke(ctx, FileService_GetFile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) DeleteFile(ctx context.Context, in *DeleteFileRequest, opts ...grpc.CallOption) (*DeleteFileResponse, error) {
	out := new(DeleteFileResponse)
	err := c.cc.Invoke(ctx, FileService_DeleteFile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileServiceServer is the server API for FileService service.
// All implementations should embed UnimplementedFileServiceServer
// for forward compatibility
type FileServiceServer interface {
	CreateFiles(context.Context, *CreateFilesRequest) (*CreateFilesResponse, error)
	GetFiles(context.Context, *GetFilesRequest) (*GetFilesResponse, error)
	GetFile(context.Context, *GetFileRequest) (*GetFileResponse, error)
	DeleteFile(context.Context, *DeleteFileRequest) (*DeleteFileResponse, error)
}

// UnimplementedFileServiceServer should be embedded to have forward compatible implementations.
type UnimplementedFileServiceServer struct {
}

func (UnimplementedFileServiceServer) CreateFiles(context.Context, *CreateFilesRequest) (*CreateFilesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFiles not implemented")
}
func (UnimplementedFileServiceServer) GetFiles(context.Context, *GetFilesRequest) (*GetFilesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFiles not implemented")
}
func (UnimplementedFileServiceServer) GetFile(context.Context, *GetFileRequest) (*GetFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFile not implemented")
}
func (UnimplementedFileServiceServer) DeleteFile(context.Context, *DeleteFileRequest) (*DeleteFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFile not implemented")
}

// UnsafeFileServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileServiceServer will
// result in compilation errors.
type UnsafeFileServiceServer interface {
	mustEmbedUnimplementedFileServiceServer()
}

func RegisterFileServiceServer(s grpc.ServiceRegistrar, srv FileServiceServer) {
	s.RegisterService(&FileService_ServiceDesc, srv)
}

func _FileService_CreateFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFilesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).CreateFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileService_CreateFiles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).CreateFiles(ctx, req.(*CreateFilesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_GetFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFilesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).GetFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileService_GetFiles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).GetFiles(ctx, req.(*GetFilesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_GetFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).GetFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileService_GetFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).GetFile(ctx, req.(*GetFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_DeleteFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).DeleteFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileService_DeleteFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).DeleteFile(ctx, req.(*DeleteFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FileService_ServiceDesc is the grpc.ServiceDesc for FileService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.saved.v1alpha1.FileService",
	HandlerType: (*FileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFiles",
			Handler:    _FileService_CreateFiles_Handler,
		},
		{
			MethodName: "GetFiles",
			Handler:    _FileService_GetFiles_Handler,
		},
		{
			MethodName: "GetFile",
			Handler:    _FileService_GetFile_Handler,
		},
		{
			MethodName: "DeleteFile",
			Handler:    _FileService_DeleteFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/saved/v1alpha1/file.proto",
}
