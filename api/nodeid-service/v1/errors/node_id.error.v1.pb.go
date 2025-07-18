// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v6.30.2
// source: api/nodeid-service/v1/errors/node_id.error.v1.proto

package errorv1

import (
	_ "github.com/go-kratos/kratos/v2/errors"
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

// ERROR 错误码用9位数字表示,其中前3位与服务序号保持一致,第4为使用0作为分隔占位,后5位用于表示具体错误码的枚举值
type ERROR int32

const (
	ERROR_UNKNOWN                       ERROR = 0         // 未知错误
	ERROR_S102_NO_AVAILABLE_ID          ERROR = 102000001 // 没有可用的节点ID
	ERROR_S102_RECORD_NOT_FOUNT         ERROR = 102000002 // 节点ID不存在
	ERROR_S102_RECORD_ALREADY_EXIST     ERROR = 102000003 // 节点ID已被使用
	ERROR_S102_HAS_BEEN_USED            ERROR = 102000004 // 节点ID已被使用
	ERROR_S102_NODE_ID_RENEWAL_FAILED   ERROR = 102000005 // 节点ID续订失败
	ERROR_S102_NODE_ID_INCORRECT        ERROR = 102000006 // 节点ID信息不正确
	ERROR_S102_ACCESS_TOKEN_INCORRECT   ERROR = 102000007 // 访问令牌不正确
	ERROR_S102_NODE_ID_STATUS_INCORRECT ERROR = 102000008 // 节点状态不正确
)

// Enum value maps for ERROR.
var (
	ERROR_name = map[int32]string{
		0:         "UNKNOWN",
		102000001: "S102_NO_AVAILABLE_ID",
		102000002: "S102_RECORD_NOT_FOUNT",
		102000003: "S102_RECORD_ALREADY_EXIST",
		102000004: "S102_HAS_BEEN_USED",
		102000005: "S102_NODE_ID_RENEWAL_FAILED",
		102000006: "S102_NODE_ID_INCORRECT",
		102000007: "S102_ACCESS_TOKEN_INCORRECT",
		102000008: "S102_NODE_ID_STATUS_INCORRECT",
	}
	ERROR_value = map[string]int32{
		"UNKNOWN":                       0,
		"S102_NO_AVAILABLE_ID":          102000001,
		"S102_RECORD_NOT_FOUNT":         102000002,
		"S102_RECORD_ALREADY_EXIST":     102000003,
		"S102_HAS_BEEN_USED":            102000004,
		"S102_NODE_ID_RENEWAL_FAILED":   102000005,
		"S102_NODE_ID_INCORRECT":        102000006,
		"S102_ACCESS_TOKEN_INCORRECT":   102000007,
		"S102_NODE_ID_STATUS_INCORRECT": 102000008,
	}
)

func (x ERROR) Enum() *ERROR {
	p := new(ERROR)
	*p = x
	return p
}

func (x ERROR) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ERROR) Descriptor() protoreflect.EnumDescriptor {
	return file_api_nodeid_service_v1_errors_node_id_error_v1_proto_enumTypes[0].Descriptor()
}

func (ERROR) Type() protoreflect.EnumType {
	return &file_api_nodeid_service_v1_errors_node_id_error_v1_proto_enumTypes[0]
}

func (x ERROR) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ERROR.Descriptor instead.
func (ERROR) EnumDescriptor() ([]byte, []int) {
	return file_api_nodeid_service_v1_errors_node_id_error_v1_proto_rawDescGZIP(), []int{0}
}

var File_api_nodeid_service_v1_errors_node_id_error_v1_proto protoreflect.FileDescriptor

var file_api_nodeid_service_v1_errors_node_id_error_v1_proto_rawDesc = []byte{
	0x0a, 0x33, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x69, 0x64, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x6e,
	0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x73, 0x61, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x6e, 0x6f, 0x64, 0x65, 0x69, 0x64, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x76, 0x31, 0x1a, 0x13,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2a, 0xd5, 0x02, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x12, 0x11, 0x0a,
	0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x1a, 0x04, 0xa8, 0x45, 0xf4, 0x03,
	0x12, 0x21, 0x0a, 0x14, 0x53, 0x31, 0x30, 0x32, 0x5f, 0x4e, 0x4f, 0x5f, 0x41, 0x56, 0x41, 0x49,
	0x4c, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x49, 0x44, 0x10, 0x81, 0xcb, 0xd1, 0x30, 0x1a, 0x04, 0xa8,
	0x45, 0x90, 0x03, 0x12, 0x22, 0x0a, 0x15, 0x53, 0x31, 0x30, 0x32, 0x5f, 0x52, 0x45, 0x43, 0x4f,
	0x52, 0x44, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x82, 0xcb, 0xd1,
	0x30, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x26, 0x0a, 0x19, 0x53, 0x31, 0x30, 0x32, 0x5f,
	0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x45,
	0x58, 0x49, 0x53, 0x54, 0x10, 0x83, 0xcb, 0xd1, 0x30, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12,
	0x1f, 0x0a, 0x12, 0x53, 0x31, 0x30, 0x32, 0x5f, 0x48, 0x41, 0x53, 0x5f, 0x42, 0x45, 0x45, 0x4e,
	0x5f, 0x55, 0x53, 0x45, 0x44, 0x10, 0x84, 0xcb, 0xd1, 0x30, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03,
	0x12, 0x28, 0x0a, 0x1b, 0x53, 0x31, 0x30, 0x32, 0x5f, 0x4e, 0x4f, 0x44, 0x45, 0x5f, 0x49, 0x44,
	0x5f, 0x52, 0x45, 0x4e, 0x45, 0x57, 0x41, 0x4c, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10,
	0x85, 0xcb, 0xd1, 0x30, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x23, 0x0a, 0x16, 0x53, 0x31,
	0x30, 0x32, 0x5f, 0x4e, 0x4f, 0x44, 0x45, 0x5f, 0x49, 0x44, 0x5f, 0x49, 0x4e, 0x43, 0x4f, 0x52,
	0x52, 0x45, 0x43, 0x54, 0x10, 0x86, 0xcb, 0xd1, 0x30, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12,
	0x28, 0x0a, 0x1b, 0x53, 0x31, 0x30, 0x32, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x54,
	0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x49, 0x4e, 0x43, 0x4f, 0x52, 0x52, 0x45, 0x43, 0x54, 0x10, 0x87,
	0xcb, 0xd1, 0x30, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x2a, 0x0a, 0x1d, 0x53, 0x31, 0x30,
	0x32, 0x5f, 0x4e, 0x4f, 0x44, 0x45, 0x5f, 0x49, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x49, 0x4e, 0x43, 0x4f, 0x52, 0x52, 0x45, 0x43, 0x54, 0x10, 0x88, 0xcb, 0xd1, 0x30, 0x1a,
	0x04, 0xa8, 0x45, 0x90, 0x03, 0x1a, 0x04, 0xa0, 0x45, 0xf4, 0x03, 0x42, 0x7f, 0x0a, 0x17, 0x73,
	0x61, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x69, 0x64, 0x2e, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x76, 0x31, 0x42, 0x14, 0x53, 0x61, 0x61, 0x73, 0x41, 0x70, 0x69, 0x4e,
	0x6f, 0x64, 0x65, 0x69, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x56, 0x31, 0x50, 0x01, 0x5a, 0x4c,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x2d, 0x73, 0x61, 0x61, 0x73, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x69, 0x64, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x6f, 0x64, 0x65,
	0x69, 0x64, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_nodeid_service_v1_errors_node_id_error_v1_proto_rawDescOnce sync.Once
	file_api_nodeid_service_v1_errors_node_id_error_v1_proto_rawDescData = file_api_nodeid_service_v1_errors_node_id_error_v1_proto_rawDesc
)

func file_api_nodeid_service_v1_errors_node_id_error_v1_proto_rawDescGZIP() []byte {
	file_api_nodeid_service_v1_errors_node_id_error_v1_proto_rawDescOnce.Do(func() {
		file_api_nodeid_service_v1_errors_node_id_error_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_nodeid_service_v1_errors_node_id_error_v1_proto_rawDescData)
	})
	return file_api_nodeid_service_v1_errors_node_id_error_v1_proto_rawDescData
}

var file_api_nodeid_service_v1_errors_node_id_error_v1_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_nodeid_service_v1_errors_node_id_error_v1_proto_goTypes = []any{
	(ERROR)(0), // 0: saas.api.nodeid.errorv1.ERROR
}
var file_api_nodeid_service_v1_errors_node_id_error_v1_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_nodeid_service_v1_errors_node_id_error_v1_proto_init() }
func file_api_nodeid_service_v1_errors_node_id_error_v1_proto_init() {
	if File_api_nodeid_service_v1_errors_node_id_error_v1_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_nodeid_service_v1_errors_node_id_error_v1_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_nodeid_service_v1_errors_node_id_error_v1_proto_goTypes,
		DependencyIndexes: file_api_nodeid_service_v1_errors_node_id_error_v1_proto_depIdxs,
		EnumInfos:         file_api_nodeid_service_v1_errors_node_id_error_v1_proto_enumTypes,
	}.Build()
	File_api_nodeid_service_v1_errors_node_id_error_v1_proto = out.File
	file_api_nodeid_service_v1_errors_node_id_error_v1_proto_rawDesc = nil
	file_api_nodeid_service_v1_errors_node_id_error_v1_proto_goTypes = nil
	file_api_nodeid_service_v1_errors_node_id_error_v1_proto_depIdxs = nil
}
