// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.6
// source: api/uuid-service/v1/enums/uuid.enum.v1.proto

package enumv1

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

type UUIDInitEnum_UuidInit int32

const (
	UUIDInitEnum_UNSPECIFIED UUIDInitEnum_UuidInit = 0 // 未指定
)

// Enum value maps for UUIDInitEnum_UuidInit.
var (
	UUIDInitEnum_UuidInit_name = map[int32]string{
		0: "UNSPECIFIED",
	}
	UUIDInitEnum_UuidInit_value = map[string]int32{
		"UNSPECIFIED": 0,
	}
)

func (x UUIDInitEnum_UuidInit) Enum() *UUIDInitEnum_UuidInit {
	p := new(UUIDInitEnum_UuidInit)
	*p = x
	return p
}

func (x UUIDInitEnum_UuidInit) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UUIDInitEnum_UuidInit) Descriptor() protoreflect.EnumDescriptor {
	return file_api_uuid_service_v1_enums_uuid_enum_v1_proto_enumTypes[0].Descriptor()
}

func (UUIDInitEnum_UuidInit) Type() protoreflect.EnumType {
	return &file_api_uuid_service_v1_enums_uuid_enum_v1_proto_enumTypes[0]
}

func (x UUIDInitEnum_UuidInit) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UUIDInitEnum_UuidInit.Descriptor instead.
func (UUIDInitEnum_UuidInit) EnumDescriptor() ([]byte, []int) {
	return file_api_uuid_service_v1_enums_uuid_enum_v1_proto_rawDescGZIP(), []int{0, 0}
}

type UUIDInitEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UUIDInitEnum) Reset() {
	*x = UUIDInitEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_uuid_service_v1_enums_uuid_enum_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UUIDInitEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UUIDInitEnum) ProtoMessage() {}

func (x *UUIDInitEnum) ProtoReflect() protoreflect.Message {
	mi := &file_api_uuid_service_v1_enums_uuid_enum_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UUIDInitEnum.ProtoReflect.Descriptor instead.
func (*UUIDInitEnum) Descriptor() ([]byte, []int) {
	return file_api_uuid_service_v1_enums_uuid_enum_v1_proto_rawDescGZIP(), []int{0}
}

var File_api_uuid_service_v1_enums_uuid_enum_v1_proto protoreflect.FileDescriptor

var file_api_uuid_service_v1_enums_uuid_enum_v1_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x75, 0x69, 0x64, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x75, 0x75, 0x69, 0x64,
	0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14,
	0x73, 0x61, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x75, 0x69, 0x64, 0x2e, 0x65, 0x6e,
	0x75, 0x6d, 0x76, 0x31, 0x22, 0x2b, 0x0a, 0x0c, 0x55, 0x55, 0x49, 0x44, 0x49, 0x6e, 0x69, 0x74,
	0x45, 0x6e, 0x75, 0x6d, 0x22, 0x1b, 0x0a, 0x08, 0x55, 0x75, 0x69, 0x64, 0x49, 0x6e, 0x69, 0x74,
	0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x42, 0x73, 0x0a, 0x14, 0x73, 0x61, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x75,
	0x69, 0x64, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x76, 0x31, 0x42, 0x11, 0x53, 0x61, 0x61, 0x73, 0x41,
	0x70, 0x69, 0x55, 0x75, 0x69, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x31, 0x50, 0x01, 0x5a, 0x46,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x2d, 0x73, 0x61, 0x61, 0x73, 0x2f, 0x75, 0x75, 0x69, 0x64, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x75, 0x69, 0x64, 0x2d, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b,
	0x65, 0x6e, 0x75, 0x6d, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_uuid_service_v1_enums_uuid_enum_v1_proto_rawDescOnce sync.Once
	file_api_uuid_service_v1_enums_uuid_enum_v1_proto_rawDescData = file_api_uuid_service_v1_enums_uuid_enum_v1_proto_rawDesc
)

func file_api_uuid_service_v1_enums_uuid_enum_v1_proto_rawDescGZIP() []byte {
	file_api_uuid_service_v1_enums_uuid_enum_v1_proto_rawDescOnce.Do(func() {
		file_api_uuid_service_v1_enums_uuid_enum_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_uuid_service_v1_enums_uuid_enum_v1_proto_rawDescData)
	})
	return file_api_uuid_service_v1_enums_uuid_enum_v1_proto_rawDescData
}

var file_api_uuid_service_v1_enums_uuid_enum_v1_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_uuid_service_v1_enums_uuid_enum_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_uuid_service_v1_enums_uuid_enum_v1_proto_goTypes = []any{
	(UUIDInitEnum_UuidInit)(0), // 0: saas.api.uuid.enumv1.UUIDInitEnum.UuidInit
	(*UUIDInitEnum)(nil),       // 1: saas.api.uuid.enumv1.UUIDInitEnum
}
var file_api_uuid_service_v1_enums_uuid_enum_v1_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_uuid_service_v1_enums_uuid_enum_v1_proto_init() }
func file_api_uuid_service_v1_enums_uuid_enum_v1_proto_init() {
	if File_api_uuid_service_v1_enums_uuid_enum_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_uuid_service_v1_enums_uuid_enum_v1_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*UUIDInitEnum); i {
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
			RawDescriptor: file_api_uuid_service_v1_enums_uuid_enum_v1_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_uuid_service_v1_enums_uuid_enum_v1_proto_goTypes,
		DependencyIndexes: file_api_uuid_service_v1_enums_uuid_enum_v1_proto_depIdxs,
		EnumInfos:         file_api_uuid_service_v1_enums_uuid_enum_v1_proto_enumTypes,
		MessageInfos:      file_api_uuid_service_v1_enums_uuid_enum_v1_proto_msgTypes,
	}.Build()
	File_api_uuid_service_v1_enums_uuid_enum_v1_proto = out.File
	file_api_uuid_service_v1_enums_uuid_enum_v1_proto_rawDesc = nil
	file_api_uuid_service_v1_enums_uuid_enum_v1_proto_goTypes = nil
	file_api_uuid_service_v1_enums_uuid_enum_v1_proto_depIdxs = nil
}
