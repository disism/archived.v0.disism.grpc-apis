// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: proto/saved/v1alpha1/file.proto

package saved

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cid     string `protobuf:"bytes,1,opt,name=cid,proto3" json:"cid,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Size    uint64 `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	Caption string `protobuf:"bytes,4,opt,name=caption,proto3" json:"caption,omitempty"`
}

func (x *CreateFile) Reset() {
	*x = CreateFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_saved_v1alpha1_file_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFile) ProtoMessage() {}

func (x *CreateFile) ProtoReflect() protoreflect.Message {
	mi := &file_proto_saved_v1alpha1_file_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFile.ProtoReflect.Descriptor instead.
func (*CreateFile) Descriptor() ([]byte, []int) {
	return file_proto_saved_v1alpha1_file_proto_rawDescGZIP(), []int{0}
}

func (x *CreateFile) GetCid() string {
	if x != nil {
		return x.Cid
	}
	return ""
}

func (x *CreateFile) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateFile) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *CreateFile) GetCaption() string {
	if x != nil {
		return x.Caption
	}
	return ""
}

type CreateFilesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Files []*CreateFile `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *CreateFilesRequest) Reset() {
	*x = CreateFilesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_saved_v1alpha1_file_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFilesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFilesRequest) ProtoMessage() {}

func (x *CreateFilesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_saved_v1alpha1_file_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFilesRequest.ProtoReflect.Descriptor instead.
func (*CreateFilesRequest) Descriptor() ([]byte, []int) {
	return file_proto_saved_v1alpha1_file_proto_rawDescGZIP(), []int{1}
}

func (x *CreateFilesRequest) GetFiles() []*CreateFile {
	if x != nil {
		return x.Files
	}
	return nil
}

type CreateFilesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Files   []*FileData     `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
	Exists  []*FileWithDirs `protobuf:"bytes,2,rep,name=exists,proto3" json:"exists,omitempty"`
	Message string          `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreateFilesResponse) Reset() {
	*x = CreateFilesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_saved_v1alpha1_file_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFilesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFilesResponse) ProtoMessage() {}

func (x *CreateFilesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_saved_v1alpha1_file_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFilesResponse.ProtoReflect.Descriptor instead.
func (*CreateFilesResponse) Descriptor() ([]byte, []int) {
	return file_proto_saved_v1alpha1_file_proto_rawDescGZIP(), []int{2}
}

func (x *CreateFilesResponse) GetFiles() []*FileData {
	if x != nil {
		return x.Files
	}
	return nil
}

func (x *CreateFilesResponse) GetExists() []*FileWithDirs {
	if x != nil {
		return x.Exists
	}
	return nil
}

func (x *CreateFilesResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetFilesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PerPage   int32  `protobuf:"varint,1,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
	Page      int32  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Sort      string `protobuf:"bytes,3,opt,name=sort,proto3" json:"sort,omitempty"`
	Direction string `protobuf:"bytes,4,opt,name=direction,proto3" json:"direction,omitempty"`
}

func (x *GetFilesRequest) Reset() {
	*x = GetFilesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_saved_v1alpha1_file_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFilesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFilesRequest) ProtoMessage() {}

func (x *GetFilesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_saved_v1alpha1_file_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFilesRequest.ProtoReflect.Descriptor instead.
func (*GetFilesRequest) Descriptor() ([]byte, []int) {
	return file_proto_saved_v1alpha1_file_proto_rawDescGZIP(), []int{3}
}

func (x *GetFilesRequest) GetPerPage() int32 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

func (x *GetFilesRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetFilesRequest) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

func (x *GetFilesRequest) GetDirection() string {
	if x != nil {
		return x.Direction
	}
	return ""
}

type GetFilesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Files   []*FileData `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
	PerPage int32       `protobuf:"varint,2,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
	Page    int32       `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Message string      `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GetFilesResponse) Reset() {
	*x = GetFilesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_saved_v1alpha1_file_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFilesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFilesResponse) ProtoMessage() {}

func (x *GetFilesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_saved_v1alpha1_file_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFilesResponse.ProtoReflect.Descriptor instead.
func (*GetFilesResponse) Descriptor() ([]byte, []int) {
	return file_proto_saved_v1alpha1_file_proto_rawDescGZIP(), []int{4}
}

func (x *GetFilesResponse) GetFiles() []*FileData {
	if x != nil {
		return x.Files
	}
	return nil
}

func (x *GetFilesResponse) GetPerPage() int32 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

func (x *GetFilesResponse) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetFilesResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetFileRequest) Reset() {
	*x = GetFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_saved_v1alpha1_file_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFileRequest) ProtoMessage() {}

func (x *GetFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_saved_v1alpha1_file_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFileRequest.ProtoReflect.Descriptor instead.
func (*GetFileRequest) Descriptor() ([]byte, []int) {
	return file_proto_saved_v1alpha1_file_proto_rawDescGZIP(), []int{5}
}

func (x *GetFileRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File    *FileData `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
	Message string    `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GetFileResponse) Reset() {
	*x = GetFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_saved_v1alpha1_file_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFileResponse) ProtoMessage() {}

func (x *GetFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_saved_v1alpha1_file_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFileResponse.ProtoReflect.Descriptor instead.
func (*GetFileResponse) Descriptor() ([]byte, []int) {
	return file_proto_saved_v1alpha1_file_proto_rawDescGZIP(), []int{6}
}

func (x *GetFileResponse) GetFile() *FileData {
	if x != nil {
		return x.File
	}
	return nil
}

func (x *GetFileResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DeleteFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteFileRequest) Reset() {
	*x = DeleteFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_saved_v1alpha1_file_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFileRequest) ProtoMessage() {}

func (x *DeleteFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_saved_v1alpha1_file_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFileRequest.ProtoReflect.Descriptor instead.
func (*DeleteFileRequest) Descriptor() ([]byte, []int) {
	return file_proto_saved_v1alpha1_file_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteFileRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteFileResponse) Reset() {
	*x = DeleteFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_saved_v1alpha1_file_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFileResponse) ProtoMessage() {}

func (x *DeleteFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_saved_v1alpha1_file_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFileResponse.ProtoReflect.Descriptor instead.
func (*DeleteFileResponse) Descriptor() ([]byte, []int) {
	return file_proto_saved_v1alpha1_file_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteFileResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_saved_v1alpha1_file_proto protoreflect.FileDescriptor

var file_proto_saved_v1alpha1_file_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x61, 0x76, 0x65, 0x64, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x61, 0x76, 0x65, 0x64, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x61, 0x76, 0x65, 0x64,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x60, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69,
	0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x63, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x61, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x61, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x54, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x05,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x61, 0x76, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x42, 0x06, 0xba,
	0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x22, 0xa1, 0x01, 0x0a,
	0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x61, 0x76, 0x65,
	0x64, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x3a, 0x0a, 0x06, 0x65, 0x78,
	0x69, 0x73, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x73, 0x61, 0x76, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x57, 0x69, 0x74, 0x68, 0x44, 0x69, 0x72, 0x73, 0x52, 0x06,
	0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x72, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x91, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x73, 0x61, 0x76, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x46, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x12,
	0x19, 0x0a, 0x08, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x70, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x25, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x5f, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x32, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x61, 0x76, 0x65, 0x64, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x28, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2e, 0x0a, 0x12, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xfa, 0x03, 0x0a, 0x0b, 0x46,
	0x69, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7f, 0x0a, 0x0b, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x28, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x73, 0x61, 0x76, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x61, 0x76, 0x65,
	0x64, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a, 0x22, 0x10, 0x2f, 0x5f, 0x73, 0x61, 0x76,
	0x65, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x73, 0x0a, 0x08, 0x47,
	0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x73, 0x61, 0x76, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x61, 0x76, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10,
	0x2f, 0x5f, 0x73, 0x61, 0x76, 0x65, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x12, 0x75, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x24, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x61, 0x76, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x61, 0x76, 0x65, 0x64, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17,
	0x12, 0x15, 0x2f, 0x5f, 0x73, 0x61, 0x76, 0x65, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x7e, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x61,
	0x76, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x61, 0x76, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17,
	0x2a, 0x15, 0x2f, 0x5f, 0x73, 0x61, 0x76, 0x65, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x10, 0x5a, 0x0e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x73, 0x61, 0x76, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_saved_v1alpha1_file_proto_rawDescOnce sync.Once
	file_proto_saved_v1alpha1_file_proto_rawDescData = file_proto_saved_v1alpha1_file_proto_rawDesc
)

func file_proto_saved_v1alpha1_file_proto_rawDescGZIP() []byte {
	file_proto_saved_v1alpha1_file_proto_rawDescOnce.Do(func() {
		file_proto_saved_v1alpha1_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_saved_v1alpha1_file_proto_rawDescData)
	})
	return file_proto_saved_v1alpha1_file_proto_rawDescData
}

var file_proto_saved_v1alpha1_file_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_saved_v1alpha1_file_proto_goTypes = []interface{}{
	(*CreateFile)(nil),          // 0: proto.saved.v1alpha1.CreateFile
	(*CreateFilesRequest)(nil),  // 1: proto.saved.v1alpha1.CreateFilesRequest
	(*CreateFilesResponse)(nil), // 2: proto.saved.v1alpha1.CreateFilesResponse
	(*GetFilesRequest)(nil),     // 3: proto.saved.v1alpha1.GetFilesRequest
	(*GetFilesResponse)(nil),    // 4: proto.saved.v1alpha1.GetFilesResponse
	(*GetFileRequest)(nil),      // 5: proto.saved.v1alpha1.GetFileRequest
	(*GetFileResponse)(nil),     // 6: proto.saved.v1alpha1.GetFileResponse
	(*DeleteFileRequest)(nil),   // 7: proto.saved.v1alpha1.DeleteFileRequest
	(*DeleteFileResponse)(nil),  // 8: proto.saved.v1alpha1.DeleteFileResponse
	(*FileData)(nil),            // 9: proto.saved.v1alpha1.FileData
	(*FileWithDirs)(nil),        // 10: proto.saved.v1alpha1.FileWithDirs
}
var file_proto_saved_v1alpha1_file_proto_depIdxs = []int32{
	0,  // 0: proto.saved.v1alpha1.CreateFilesRequest.files:type_name -> proto.saved.v1alpha1.CreateFile
	9,  // 1: proto.saved.v1alpha1.CreateFilesResponse.files:type_name -> proto.saved.v1alpha1.FileData
	10, // 2: proto.saved.v1alpha1.CreateFilesResponse.exists:type_name -> proto.saved.v1alpha1.FileWithDirs
	9,  // 3: proto.saved.v1alpha1.GetFilesResponse.files:type_name -> proto.saved.v1alpha1.FileData
	9,  // 4: proto.saved.v1alpha1.GetFileResponse.file:type_name -> proto.saved.v1alpha1.FileData
	1,  // 5: proto.saved.v1alpha1.FileService.CreateFiles:input_type -> proto.saved.v1alpha1.CreateFilesRequest
	3,  // 6: proto.saved.v1alpha1.FileService.GetFiles:input_type -> proto.saved.v1alpha1.GetFilesRequest
	5,  // 7: proto.saved.v1alpha1.FileService.GetFile:input_type -> proto.saved.v1alpha1.GetFileRequest
	7,  // 8: proto.saved.v1alpha1.FileService.DeleteFile:input_type -> proto.saved.v1alpha1.DeleteFileRequest
	2,  // 9: proto.saved.v1alpha1.FileService.CreateFiles:output_type -> proto.saved.v1alpha1.CreateFilesResponse
	4,  // 10: proto.saved.v1alpha1.FileService.GetFiles:output_type -> proto.saved.v1alpha1.GetFilesResponse
	6,  // 11: proto.saved.v1alpha1.FileService.GetFile:output_type -> proto.saved.v1alpha1.GetFileResponse
	8,  // 12: proto.saved.v1alpha1.FileService.DeleteFile:output_type -> proto.saved.v1alpha1.DeleteFileResponse
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_proto_saved_v1alpha1_file_proto_init() }
func file_proto_saved_v1alpha1_file_proto_init() {
	if File_proto_saved_v1alpha1_file_proto != nil {
		return
	}
	file_proto_saved_v1alpha1_data_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_saved_v1alpha1_file_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFile); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_saved_v1alpha1_file_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFilesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_saved_v1alpha1_file_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFilesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_saved_v1alpha1_file_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFilesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_saved_v1alpha1_file_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFilesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_saved_v1alpha1_file_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFileRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_saved_v1alpha1_file_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFileResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_saved_v1alpha1_file_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFileRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_saved_v1alpha1_file_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFileResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_saved_v1alpha1_file_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_saved_v1alpha1_file_proto_goTypes,
		DependencyIndexes: file_proto_saved_v1alpha1_file_proto_depIdxs,
		MessageInfos:      file_proto_saved_v1alpha1_file_proto_msgTypes,
	}.Build()
	File_proto_saved_v1alpha1_file_proto = out.File
	file_proto_saved_v1alpha1_file_proto_rawDesc = nil
	file_proto_saved_v1alpha1_file_proto_goTypes = nil
	file_proto_saved_v1alpha1_file_proto_depIdxs = nil
}
