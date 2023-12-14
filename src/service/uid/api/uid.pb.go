// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: src/api/uid/uid.proto

package api

import (
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

type GenerateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *GenerateResponse) Reset() {
	*x = GenerateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_api_uid_uid_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateResponse) ProtoMessage() {}

func (x *GenerateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_api_uid_uid_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateResponse.ProtoReflect.Descriptor instead.
func (*GenerateResponse) Descriptor() ([]byte, []int) {
	return file_src_api_uid_uid_proto_rawDescGZIP(), []int{0}
}

func (x *GenerateResponse) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type GenerateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GenerateRequest) Reset() {
	*x = GenerateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_api_uid_uid_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateRequest) ProtoMessage() {}

func (x *GenerateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_api_uid_uid_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateRequest.ProtoReflect.Descriptor instead.
func (*GenerateRequest) Descriptor() ([]byte, []int) {
	return file_src_api_uid_uid_proto_rawDescGZIP(), []int{1}
}

var File_src_api_uid_uid_proto protoreflect.FileDescriptor

var file_src_api_uid_uid_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x72, 0x63, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x69, 0x64, 0x2f, 0x75, 0x69,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6d, 0x6f, 0x61, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x75, 0x69, 0x64, 0x2e, 0x76, 0x31, 0x22, 0x22, 0x0a, 0x10, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x11, 0x0a, 0x0f, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0x5d,
	0x0a, 0x0a, 0x55, 0x49, 0x44, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x08,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x2e, 0x6d, 0x6f, 0x61, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x75, 0x69, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6d, 0x6f, 0x61, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x75, 0x69, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x3f, 0x5a,
	0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6c, 0x65, 0x65,
	0x70, 0x79, 0x66, 0x6f, 0x78, 0x39, 0x37, 0x2f, 0x6f, 0x74, 0x65, 0x6c, 0x2d, 0x67, 0x72, 0x70,
	0x63, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x75, 0x69, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_api_uid_uid_proto_rawDescOnce sync.Once
	file_src_api_uid_uid_proto_rawDescData = file_src_api_uid_uid_proto_rawDesc
)

func file_src_api_uid_uid_proto_rawDescGZIP() []byte {
	file_src_api_uid_uid_proto_rawDescOnce.Do(func() {
		file_src_api_uid_uid_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_api_uid_uid_proto_rawDescData)
	})
	return file_src_api_uid_uid_proto_rawDescData
}

var file_src_api_uid_uid_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_src_api_uid_uid_proto_goTypes = []interface{}{
	(*GenerateResponse)(nil), // 0: moa.api.uid.v1.GenerateResponse
	(*GenerateRequest)(nil),  // 1: moa.api.uid.v1.GenerateRequest
}
var file_src_api_uid_uid_proto_depIdxs = []int32{
	1, // 0: moa.api.uid.v1.UIDService.Generate:input_type -> moa.api.uid.v1.GenerateRequest
	0, // 1: moa.api.uid.v1.UIDService.Generate:output_type -> moa.api.uid.v1.GenerateResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_src_api_uid_uid_proto_init() }
func file_src_api_uid_uid_proto_init() {
	if File_src_api_uid_uid_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_api_uid_uid_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateResponse); i {
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
		file_src_api_uid_uid_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateRequest); i {
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
			RawDescriptor: file_src_api_uid_uid_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_src_api_uid_uid_proto_goTypes,
		DependencyIndexes: file_src_api_uid_uid_proto_depIdxs,
		MessageInfos:      file_src_api_uid_uid_proto_msgTypes,
	}.Build()
	File_src_api_uid_uid_proto = out.File
	file_src_api_uid_uid_proto_rawDesc = nil
	file_src_api_uid_uid_proto_goTypes = nil
	file_src_api_uid_uid_proto_depIdxs = nil
}
