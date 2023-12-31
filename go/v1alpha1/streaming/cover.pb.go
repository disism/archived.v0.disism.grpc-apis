// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: proto/streaming/v1alpha1/cover.proto

package streaming

import (
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

type CreateCoverRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileId uint64 `protobuf:"varint,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	Width  int32  `protobuf:"varint,2,opt,name=width,proto3" json:"width,omitempty"`
	Height int32  `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *CreateCoverRequest) Reset() {
	*x = CreateCoverRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_streaming_v1alpha1_cover_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCoverRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCoverRequest) ProtoMessage() {}

func (x *CreateCoverRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_streaming_v1alpha1_cover_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCoverRequest.ProtoReflect.Descriptor instead.
func (*CreateCoverRequest) Descriptor() ([]byte, []int) {
	return file_proto_streaming_v1alpha1_cover_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCoverRequest) GetFileId() uint64 {
	if x != nil {
		return x.FileId
	}
	return 0
}

func (x *CreateCoverRequest) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *CreateCoverRequest) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

type CreateCoverResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string     `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Cover   *CoverData `protobuf:"bytes,2,opt,name=cover,proto3" json:"cover,omitempty"`
}

func (x *CreateCoverResponse) Reset() {
	*x = CreateCoverResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_streaming_v1alpha1_cover_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCoverResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCoverResponse) ProtoMessage() {}

func (x *CreateCoverResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_streaming_v1alpha1_cover_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCoverResponse.ProtoReflect.Descriptor instead.
func (*CreateCoverResponse) Descriptor() ([]byte, []int) {
	return file_proto_streaming_v1alpha1_cover_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCoverResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CreateCoverResponse) GetCover() *CoverData {
	if x != nil {
		return x.Cover
	}
	return nil
}

var File_proto_streaming_v1alpha1_cover_proto protoreflect.FileDescriptor

var file_proto_streaming_v1alpha1_cover_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e,
	0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x5b, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x76,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x22, 0x6a, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x39, 0x0a, 0x05, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x76, 0x65,
	0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x05, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x32, 0x9c, 0x01, 0x0a,
	0x0c, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x8b, 0x01,
	0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x2c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x76,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x19, 0x3a, 0x01, 0x2a, 0x22, 0x14, 0x2f, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69,
	0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x42, 0x14, 0x5a, 0x12, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e,
	0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_streaming_v1alpha1_cover_proto_rawDescOnce sync.Once
	file_proto_streaming_v1alpha1_cover_proto_rawDescData = file_proto_streaming_v1alpha1_cover_proto_rawDesc
)

func file_proto_streaming_v1alpha1_cover_proto_rawDescGZIP() []byte {
	file_proto_streaming_v1alpha1_cover_proto_rawDescOnce.Do(func() {
		file_proto_streaming_v1alpha1_cover_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_streaming_v1alpha1_cover_proto_rawDescData)
	})
	return file_proto_streaming_v1alpha1_cover_proto_rawDescData
}

var file_proto_streaming_v1alpha1_cover_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_streaming_v1alpha1_cover_proto_goTypes = []interface{}{
	(*CreateCoverRequest)(nil),  // 0: proto.streaming.v1alpha1.CreateCoverRequest
	(*CreateCoverResponse)(nil), // 1: proto.streaming.v1alpha1.CreateCoverResponse
	(*CoverData)(nil),           // 2: proto.streaming.v1alpha1.CoverData
}
var file_proto_streaming_v1alpha1_cover_proto_depIdxs = []int32{
	2, // 0: proto.streaming.v1alpha1.CreateCoverResponse.cover:type_name -> proto.streaming.v1alpha1.CoverData
	0, // 1: proto.streaming.v1alpha1.CoverService.CreateCover:input_type -> proto.streaming.v1alpha1.CreateCoverRequest
	1, // 2: proto.streaming.v1alpha1.CoverService.CreateCover:output_type -> proto.streaming.v1alpha1.CreateCoverResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_streaming_v1alpha1_cover_proto_init() }
func file_proto_streaming_v1alpha1_cover_proto_init() {
	if File_proto_streaming_v1alpha1_cover_proto != nil {
		return
	}
	file_proto_streaming_v1alpha1_data_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_streaming_v1alpha1_cover_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCoverRequest); i {
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
		file_proto_streaming_v1alpha1_cover_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCoverResponse); i {
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
			RawDescriptor: file_proto_streaming_v1alpha1_cover_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_streaming_v1alpha1_cover_proto_goTypes,
		DependencyIndexes: file_proto_streaming_v1alpha1_cover_proto_depIdxs,
		MessageInfos:      file_proto_streaming_v1alpha1_cover_proto_msgTypes,
	}.Build()
	File_proto_streaming_v1alpha1_cover_proto = out.File
	file_proto_streaming_v1alpha1_cover_proto_rawDesc = nil
	file_proto_streaming_v1alpha1_cover_proto_goTypes = nil
	file_proto_streaming_v1alpha1_cover_proto_depIdxs = nil
}
