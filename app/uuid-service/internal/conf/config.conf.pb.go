// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.6
// source: app/uuid-service/internal/conf/config.conf.proto

package conf

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type ServiceConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UuidService *ServiceConfig_UuidService `protobuf:"bytes,2,opt,name=uuid_service,json=uuidService,proto3" json:"uuid_service,omitempty"`
}

func (x *ServiceConfig) Reset() {
	*x = ServiceConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_uuid_service_internal_conf_config_conf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceConfig) ProtoMessage() {}

func (x *ServiceConfig) ProtoReflect() protoreflect.Message {
	mi := &file_app_uuid_service_internal_conf_config_conf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceConfig.ProtoReflect.Descriptor instead.
func (*ServiceConfig) Descriptor() ([]byte, []int) {
	return file_app_uuid_service_internal_conf_config_conf_proto_rawDescGZIP(), []int{0}
}

func (x *ServiceConfig) GetUuidService() *ServiceConfig_UuidService {
	if x != nil {
		return x.UuidService
	}
	return nil
}

type ServiceConfig_UuidService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstanceId   string            `protobuf:"bytes,1,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	InstanceName string            `protobuf:"bytes,2,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	Metadata     map[string]string `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ServiceConfig_UuidService) Reset() {
	*x = ServiceConfig_UuidService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_uuid_service_internal_conf_config_conf_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceConfig_UuidService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceConfig_UuidService) ProtoMessage() {}

func (x *ServiceConfig_UuidService) ProtoReflect() protoreflect.Message {
	mi := &file_app_uuid_service_internal_conf_config_conf_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceConfig_UuidService.ProtoReflect.Descriptor instead.
func (*ServiceConfig_UuidService) Descriptor() ([]byte, []int) {
	return file_app_uuid_service_internal_conf_config_conf_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ServiceConfig_UuidService) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *ServiceConfig_UuidService) GetInstanceName() string {
	if x != nil {
		return x.InstanceName
	}
	return ""
}

func (x *ServiceConfig_UuidService) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_app_uuid_service_internal_conf_config_conf_proto protoreflect.FileDescriptor

var file_app_uuid_service_internal_conf_config_conf_proto_rawDesc = []byte{
	0x0a, 0x30, 0x61, 0x70, 0x70, 0x2f, 0x75, 0x75, 0x69, 0x64, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x12, 0x73, 0x61, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x75, 0x69,
	0x64, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xd6, 0x02, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x50, 0x0a, 0x0c, 0x75, 0x75, 0x69, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x73, 0x61, 0x61, 0x73, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x75, 0x75, 0x69, 0x64, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x55, 0x75, 0x69, 0x64, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x0b, 0x75, 0x75, 0x69, 0x64, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x1a, 0xf2, 0x01, 0x0a, 0x0b, 0x55, 0x75, 0x69, 0x64, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10,
	0x01, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x23, 0x0a,
	0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x57, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x73, 0x61, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x75, 0x75, 0x69, 0x64, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x55, 0x75, 0x69, 0x64, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x72, 0x0a, 0x12, 0x73, 0x61, 0x61, 0x73,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x75, 0x69, 0x64, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x42, 0x0f,
	0x53, 0x61, 0x61, 0x73, 0x41, 0x70, 0x69, 0x55, 0x75, 0x69, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x50,
	0x01, 0x5a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f,
	0x2d, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2d, 0x73, 0x61, 0x61, 0x73, 0x2f, 0x75, 0x75, 0x69, 0x64,
	0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x75, 0x75, 0x69,
	0x64, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_uuid_service_internal_conf_config_conf_proto_rawDescOnce sync.Once
	file_app_uuid_service_internal_conf_config_conf_proto_rawDescData = file_app_uuid_service_internal_conf_config_conf_proto_rawDesc
)

func file_app_uuid_service_internal_conf_config_conf_proto_rawDescGZIP() []byte {
	file_app_uuid_service_internal_conf_config_conf_proto_rawDescOnce.Do(func() {
		file_app_uuid_service_internal_conf_config_conf_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_uuid_service_internal_conf_config_conf_proto_rawDescData)
	})
	return file_app_uuid_service_internal_conf_config_conf_proto_rawDescData
}

var file_app_uuid_service_internal_conf_config_conf_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_app_uuid_service_internal_conf_config_conf_proto_goTypes = []any{
	(*ServiceConfig)(nil),             // 0: saas.api.uuid.conf.ServiceConfig
	(*ServiceConfig_UuidService)(nil), // 1: saas.api.uuid.conf.ServiceConfig.UuidService
	nil,                               // 2: saas.api.uuid.conf.ServiceConfig.UuidService.MetadataEntry
}
var file_app_uuid_service_internal_conf_config_conf_proto_depIdxs = []int32{
	1, // 0: saas.api.uuid.conf.ServiceConfig.uuid_service:type_name -> saas.api.uuid.conf.ServiceConfig.UuidService
	2, // 1: saas.api.uuid.conf.ServiceConfig.UuidService.metadata:type_name -> saas.api.uuid.conf.ServiceConfig.UuidService.MetadataEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_app_uuid_service_internal_conf_config_conf_proto_init() }
func file_app_uuid_service_internal_conf_config_conf_proto_init() {
	if File_app_uuid_service_internal_conf_config_conf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_uuid_service_internal_conf_config_conf_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ServiceConfig); i {
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
		file_app_uuid_service_internal_conf_config_conf_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ServiceConfig_UuidService); i {
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
			RawDescriptor: file_app_uuid_service_internal_conf_config_conf_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_app_uuid_service_internal_conf_config_conf_proto_goTypes,
		DependencyIndexes: file_app_uuid_service_internal_conf_config_conf_proto_depIdxs,
		MessageInfos:      file_app_uuid_service_internal_conf_config_conf_proto_msgTypes,
	}.Build()
	File_app_uuid_service_internal_conf_config_conf_proto = out.File
	file_app_uuid_service_internal_conf_config_conf_proto_rawDesc = nil
	file_app_uuid_service_internal_conf_config_conf_proto_goTypes = nil
	file_app_uuid_service_internal_conf_config_conf_proto_depIdxs = nil
}
