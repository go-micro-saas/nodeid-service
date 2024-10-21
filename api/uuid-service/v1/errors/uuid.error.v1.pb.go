// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.6
// source: api/uuid-service/v1/errors/uuid.error.v1.proto

package errorv1

import (
	_ "github.com/go-kratos/kratos/v2/errors"
	errors "github.com/go-micro-saas/nodeid-service/api/nodeid-service/v1/errors"
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

// ERROR 复用node_id的错误码，详情请查看node_id的错误码；统一在node_id定义
type ERROR int32

const (
	ERROR_UNKNOWN ERROR = 0 // 未知错误
)

// Enum value maps for ERROR.
var (
	ERROR_name = map[int32]string{
		0: "UNKNOWN",
	}
	ERROR_value = map[string]int32{
		"UNKNOWN": 0,
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
	return file_api_uuid_service_v1_errors_uuid_error_v1_proto_enumTypes[0].Descriptor()
}

func (ERROR) Type() protoreflect.EnumType {
	return &file_api_uuid_service_v1_errors_uuid_error_v1_proto_enumTypes[0]
}

func (x ERROR) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ERROR.Descriptor instead.
func (ERROR) EnumDescriptor() ([]byte, []int) {
	return file_api_uuid_service_v1_errors_uuid_error_v1_proto_rawDescGZIP(), []int{0}
}

type README struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// error 复用node_id的错误码，详情请查看node_id的错误码；统一在node_id定义
	Error errors.ERROR `protobuf:"varint,1,opt,name=error,proto3,enum=saas.api.nodeid.errorv1.ERROR" json:"error,omitempty"`
}

func (x *README) Reset() {
	*x = README{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_uuid_service_v1_errors_uuid_error_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *README) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*README) ProtoMessage() {}

func (x *README) ProtoReflect() protoreflect.Message {
	mi := &file_api_uuid_service_v1_errors_uuid_error_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use README.ProtoReflect.Descriptor instead.
func (*README) Descriptor() ([]byte, []int) {
	return file_api_uuid_service_v1_errors_uuid_error_v1_proto_rawDescGZIP(), []int{0}
}

func (x *README) GetError() errors.ERROR {
	if x != nil {
		return x.Error
	}
	return errors.ERROR(0)
}

var File_api_uuid_service_v1_errors_uuid_error_v1_proto protoreflect.FileDescriptor

var file_api_uuid_service_v1_errors_uuid_error_v1_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x75, 0x69, 0x64, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x75, 0x75, 0x69,
	0x64, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x15, 0x73, 0x61, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x75, 0x69, 0x64, 0x2e,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x76, 0x31, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x61, 0x70,
	0x69, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x69, 0x64, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x5f,
	0x69, 0x64, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x3e, 0x0a, 0x06, 0x52, 0x45, 0x41, 0x44, 0x4d, 0x45, 0x12, 0x34, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x73, 0x61, 0x61,
	0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x69, 0x64, 0x2e, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x76, 0x31, 0x2e, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x2a, 0x20, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x12, 0x11, 0x0a, 0x07, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x1a, 0x04, 0xa8, 0x45, 0xf4, 0x03, 0x1a, 0x04, 0xa0,
	0x45, 0xf4, 0x03, 0x42, 0x79, 0x0a, 0x15, 0x73, 0x61, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x75, 0x75, 0x69, 0x64, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x76, 0x31, 0x42, 0x12, 0x53, 0x61,
	0x61, 0x73, 0x41, 0x70, 0x69, 0x55, 0x75, 0x69, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x56, 0x31,
	0x50, 0x01, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x6f, 0x2d, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2d, 0x73, 0x61, 0x61, 0x73, 0x2f, 0x6e, 0x6f, 0x64,
	0x65, 0x69, 0x64, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x75, 0x75, 0x69, 0x64, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_uuid_service_v1_errors_uuid_error_v1_proto_rawDescOnce sync.Once
	file_api_uuid_service_v1_errors_uuid_error_v1_proto_rawDescData = file_api_uuid_service_v1_errors_uuid_error_v1_proto_rawDesc
)

func file_api_uuid_service_v1_errors_uuid_error_v1_proto_rawDescGZIP() []byte {
	file_api_uuid_service_v1_errors_uuid_error_v1_proto_rawDescOnce.Do(func() {
		file_api_uuid_service_v1_errors_uuid_error_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_uuid_service_v1_errors_uuid_error_v1_proto_rawDescData)
	})
	return file_api_uuid_service_v1_errors_uuid_error_v1_proto_rawDescData
}

var file_api_uuid_service_v1_errors_uuid_error_v1_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_uuid_service_v1_errors_uuid_error_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_uuid_service_v1_errors_uuid_error_v1_proto_goTypes = []any{
	(ERROR)(0),        // 0: saas.api.uuid.errorv1.ERROR
	(*README)(nil),    // 1: saas.api.uuid.errorv1.README
	(errors.ERROR)(0), // 2: saas.api.nodeid.errorv1.ERROR
}
var file_api_uuid_service_v1_errors_uuid_error_v1_proto_depIdxs = []int32{
	2, // 0: saas.api.uuid.errorv1.README.error:type_name -> saas.api.nodeid.errorv1.ERROR
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_uuid_service_v1_errors_uuid_error_v1_proto_init() }
func file_api_uuid_service_v1_errors_uuid_error_v1_proto_init() {
	if File_api_uuid_service_v1_errors_uuid_error_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_uuid_service_v1_errors_uuid_error_v1_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*README); i {
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
			RawDescriptor: file_api_uuid_service_v1_errors_uuid_error_v1_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_uuid_service_v1_errors_uuid_error_v1_proto_goTypes,
		DependencyIndexes: file_api_uuid_service_v1_errors_uuid_error_v1_proto_depIdxs,
		EnumInfos:         file_api_uuid_service_v1_errors_uuid_error_v1_proto_enumTypes,
		MessageInfos:      file_api_uuid_service_v1_errors_uuid_error_v1_proto_msgTypes,
	}.Build()
	File_api_uuid_service_v1_errors_uuid_error_v1_proto = out.File
	file_api_uuid_service_v1_errors_uuid_error_v1_proto_rawDesc = nil
	file_api_uuid_service_v1_errors_uuid_error_v1_proto_goTypes = nil
	file_api_uuid_service_v1_errors_uuid_error_v1_proto_depIdxs = nil
}
