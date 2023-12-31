// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: proto/account/v1alpha1/public.proto

package account

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetVersionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GetVersionResponse) Reset() {
	*x = GetVersionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_v1alpha1_public_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVersionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVersionResponse) ProtoMessage() {}

func (x *GetVersionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_v1alpha1_public_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVersionResponse.ProtoReflect.Descriptor instead.
func (*GetVersionResponse) Descriptor() ([]byte, []int) {
	return file_proto_account_v1alpha1_public_proto_rawDescGZIP(), []int{0}
}

func (x *GetVersionResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetWebfingerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resource string `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (x *GetWebfingerRequest) Reset() {
	*x = GetWebfingerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_v1alpha1_public_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWebfingerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWebfingerRequest) ProtoMessage() {}

func (x *GetWebfingerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_v1alpha1_public_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWebfingerRequest.ProtoReflect.Descriptor instead.
func (*GetWebfingerRequest) Descriptor() ([]byte, []int) {
	return file_proto_account_v1alpha1_public_proto_rawDescGZIP(), []int{1}
}

func (x *GetWebfingerRequest) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

type GetWebfingerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subject string                       `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	Aliases []string                     `protobuf:"bytes,2,rep,name=aliases,proto3" json:"aliases,omitempty"`
	Links   []*GetWebfingerResponse_Link `protobuf:"bytes,3,rep,name=links,proto3" json:"links,omitempty"`
}

func (x *GetWebfingerResponse) Reset() {
	*x = GetWebfingerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_v1alpha1_public_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWebfingerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWebfingerResponse) ProtoMessage() {}

func (x *GetWebfingerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_v1alpha1_public_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWebfingerResponse.ProtoReflect.Descriptor instead.
func (*GetWebfingerResponse) Descriptor() ([]byte, []int) {
	return file_proto_account_v1alpha1_public_proto_rawDescGZIP(), []int{2}
}

func (x *GetWebfingerResponse) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *GetWebfingerResponse) GetAliases() []string {
	if x != nil {
		return x.Aliases
	}
	return nil
}

func (x *GetWebfingerResponse) GetLinks() []*GetWebfingerResponse_Link {
	if x != nil {
		return x.Links
	}
	return nil
}

type GetWebfingerResponse_Link struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rel  string `protobuf:"bytes,1,opt,name=rel,proto3" json:"rel,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Href string `protobuf:"bytes,3,opt,name=href,proto3" json:"href,omitempty"`
}

func (x *GetWebfingerResponse_Link) Reset() {
	*x = GetWebfingerResponse_Link{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_v1alpha1_public_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWebfingerResponse_Link) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWebfingerResponse_Link) ProtoMessage() {}

func (x *GetWebfingerResponse_Link) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_v1alpha1_public_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWebfingerResponse_Link.ProtoReflect.Descriptor instead.
func (*GetWebfingerResponse_Link) Descriptor() ([]byte, []int) {
	return file_proto_account_v1alpha1_public_proto_rawDescGZIP(), []int{2, 0}
}

func (x *GetWebfingerResponse_Link) GetRel() string {
	if x != nil {
		return x.Rel
	}
	return ""
}

func (x *GetWebfingerResponse_Link) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *GetWebfingerResponse_Link) GetHref() string {
	if x != nil {
		return x.Href
	}
	return ""
}

var File_proto_account_v1alpha1_public_proto protoreflect.FileDescriptor

var file_proto_account_v1alpha1_public_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x31, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x57,
	0x65, 0x62, 0x66, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0xd5, 0x01, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x57, 0x65, 0x62, 0x66, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x65, 0x73, 0x12, 0x47, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x6b,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x57, 0x65, 0x62, 0x66, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x6b,
	0x73, 0x1a, 0x40, 0x0a, 0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x68, 0x72, 0x65, 0x66, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68,
	0x72, 0x65, 0x66, 0x32, 0x8c, 0x02, 0x0a, 0x0d, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6f, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12,
	0x15, 0x2f, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x89, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x57, 0x65,
	0x62, 0x66, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x12, 0x2b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x57, 0x65, 0x62, 0x66, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x57, 0x65, 0x62, 0x66, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x2e, 0x77, 0x65,
	0x6c, 0x6c, 0x2d, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2f, 0x77, 0x65, 0x62, 0x66, 0x69, 0x6e, 0x67,
	0x65, 0x72, 0x42, 0x12, 0x5a, 0x10, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_account_v1alpha1_public_proto_rawDescOnce sync.Once
	file_proto_account_v1alpha1_public_proto_rawDescData = file_proto_account_v1alpha1_public_proto_rawDesc
)

func file_proto_account_v1alpha1_public_proto_rawDescGZIP() []byte {
	file_proto_account_v1alpha1_public_proto_rawDescOnce.Do(func() {
		file_proto_account_v1alpha1_public_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_account_v1alpha1_public_proto_rawDescData)
	})
	return file_proto_account_v1alpha1_public_proto_rawDescData
}

var file_proto_account_v1alpha1_public_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_account_v1alpha1_public_proto_goTypes = []interface{}{
	(*GetVersionResponse)(nil),        // 0: proto.account.v1alpha1.GetVersionResponse
	(*GetWebfingerRequest)(nil),       // 1: proto.account.v1alpha1.GetWebfingerRequest
	(*GetWebfingerResponse)(nil),      // 2: proto.account.v1alpha1.GetWebfingerResponse
	(*GetWebfingerResponse_Link)(nil), // 3: proto.account.v1alpha1.GetWebfingerResponse.Link
	(*emptypb.Empty)(nil),             // 4: google.protobuf.Empty
}
var file_proto_account_v1alpha1_public_proto_depIdxs = []int32{
	3, // 0: proto.account.v1alpha1.GetWebfingerResponse.links:type_name -> proto.account.v1alpha1.GetWebfingerResponse.Link
	4, // 1: proto.account.v1alpha1.PublicService.GetVersion:input_type -> google.protobuf.Empty
	1, // 2: proto.account.v1alpha1.PublicService.GetWebfinger:input_type -> proto.account.v1alpha1.GetWebfingerRequest
	0, // 3: proto.account.v1alpha1.PublicService.GetVersion:output_type -> proto.account.v1alpha1.GetVersionResponse
	2, // 4: proto.account.v1alpha1.PublicService.GetWebfinger:output_type -> proto.account.v1alpha1.GetWebfingerResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_account_v1alpha1_public_proto_init() }
func file_proto_account_v1alpha1_public_proto_init() {
	if File_proto_account_v1alpha1_public_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_account_v1alpha1_public_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVersionResponse); i {
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
		file_proto_account_v1alpha1_public_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWebfingerRequest); i {
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
		file_proto_account_v1alpha1_public_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWebfingerResponse); i {
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
		file_proto_account_v1alpha1_public_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWebfingerResponse_Link); i {
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
			RawDescriptor: file_proto_account_v1alpha1_public_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_account_v1alpha1_public_proto_goTypes,
		DependencyIndexes: file_proto_account_v1alpha1_public_proto_depIdxs,
		MessageInfos:      file_proto_account_v1alpha1_public_proto_msgTypes,
	}.Build()
	File_proto_account_v1alpha1_public_proto = out.File
	file_proto_account_v1alpha1_public_proto_rawDesc = nil
	file_proto_account_v1alpha1_public_proto_goTypes = nil
	file_proto_account_v1alpha1_public_proto_depIdxs = nil
}
