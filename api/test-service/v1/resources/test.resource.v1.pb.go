// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.6
// source: api/test-service/v1/resources/test.resource.v1.proto

package resourcev1

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

// PingReq request
type PingReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// message
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PingReq) Reset() {
	*x = PingReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_test_service_v1_resources_test_resource_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingReq) ProtoMessage() {}

func (x *PingReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_test_service_v1_resources_test_resource_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingReq.ProtoReflect.Descriptor instead.
func (*PingReq) Descriptor() ([]byte, []int) {
	return file_api_test_service_v1_resources_test_resource_v1_proto_rawDescGZIP(), []int{0}
}

func (x *PingReq) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// PingResp reply
type PingResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// message
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PingResp) Reset() {
	*x = PingResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_test_service_v1_resources_test_resource_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResp) ProtoMessage() {}

func (x *PingResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_test_service_v1_resources_test_resource_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResp.ProtoReflect.Descriptor instead.
func (*PingResp) Descriptor() ([]byte, []int) {
	return file_api_test_service_v1_resources_test_resource_v1_proto_rawDescGZIP(), []int{1}
}

func (x *PingResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_api_test_service_v1_resources_test_resource_v1_proto protoreflect.FileDescriptor

var file_api_test_service_v1_resources_test_resource_v1_proto_rawDesc = []byte{
	0x0a, 0x34, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x73, 0x61, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x69, 0x64, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x29, 0x0a, 0x07, 0x50, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x12, 0x1e, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x24, 0x0a, 0x08, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x92, 0x01, 0x0a, 0x1f, 0x73, 0x61,
	0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x69, 0x64, 0x5f, 0x74, 0x65,
	0x73, 0x74, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x76, 0x31, 0x42, 0x1b, 0x53,
	0x61, 0x61, 0x73, 0x41, 0x70, 0x69, 0x4e, 0x6f, 0x64, 0x65, 0x69, 0x64, 0x54, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x56, 0x31, 0x50, 0x01, 0x5a, 0x50, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x2d, 0x73, 0x61, 0x61, 0x73, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x69, 0x64, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_test_service_v1_resources_test_resource_v1_proto_rawDescOnce sync.Once
	file_api_test_service_v1_resources_test_resource_v1_proto_rawDescData = file_api_test_service_v1_resources_test_resource_v1_proto_rawDesc
)

func file_api_test_service_v1_resources_test_resource_v1_proto_rawDescGZIP() []byte {
	file_api_test_service_v1_resources_test_resource_v1_proto_rawDescOnce.Do(func() {
		file_api_test_service_v1_resources_test_resource_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_test_service_v1_resources_test_resource_v1_proto_rawDescData)
	})
	return file_api_test_service_v1_resources_test_resource_v1_proto_rawDescData
}

var file_api_test_service_v1_resources_test_resource_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_test_service_v1_resources_test_resource_v1_proto_goTypes = []any{
	(*PingReq)(nil),  // 0: saas.api.nodeid_test.resourcev1.PingReq
	(*PingResp)(nil), // 1: saas.api.nodeid_test.resourcev1.PingResp
}
var file_api_test_service_v1_resources_test_resource_v1_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_test_service_v1_resources_test_resource_v1_proto_init() }
func file_api_test_service_v1_resources_test_resource_v1_proto_init() {
	if File_api_test_service_v1_resources_test_resource_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_test_service_v1_resources_test_resource_v1_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*PingReq); i {
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
		file_api_test_service_v1_resources_test_resource_v1_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*PingResp); i {
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
			RawDescriptor: file_api_test_service_v1_resources_test_resource_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_test_service_v1_resources_test_resource_v1_proto_goTypes,
		DependencyIndexes: file_api_test_service_v1_resources_test_resource_v1_proto_depIdxs,
		MessageInfos:      file_api_test_service_v1_resources_test_resource_v1_proto_msgTypes,
	}.Build()
	File_api_test_service_v1_resources_test_resource_v1_proto = out.File
	file_api_test_service_v1_resources_test_resource_v1_proto_rawDesc = nil
	file_api_test_service_v1_resources_test_resource_v1_proto_goTypes = nil
	file_api_test_service_v1_resources_test_resource_v1_proto_depIdxs = nil
}
