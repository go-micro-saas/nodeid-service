// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.6
// source: api/nodeid-service/v1/services/node_id.service.v1.proto

package servicev1

import (
	resources "github.com/go-micro-saas/nodeid-service/api/nodeid-service/v1/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_api_nodeid_service_v1_services_node_id_service_v1_proto protoreflect.FileDescriptor

var file_api_nodeid_service_v1_services_node_id_service_v1_proto_rawDesc = []byte{
	0x0a, 0x37, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x69, 0x64, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x73, 0x61, 0x61, 0x73, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x69, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x39, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x69, 0x64, 0x2d, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xdd, 0x05,
	0x0a, 0x0b, 0x53, 0x72, 0x76, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x56, 0x31, 0x12, 0x78, 0x0a,
	0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x23, 0x2e, 0x73, 0x61, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x69, 0x64, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x76, 0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x24, 0x2e, 0x73, 0x61, 0x61,
	0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x69, 0x64, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x12, 0x1d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x69, 0x64, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x7b, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x7d, 0x12, 0x98, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2d, 0x2e, 0x73, 0x61, 0x61,
	0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x69, 0x64, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x2e, 0x2e, 0x73, 0x61, 0x61, 0x73,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x69, 0x64, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x21, 0x12, 0x1f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x69,
	0x64, 0x2f, 0x67, 0x65, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x69, 0x6e,
	0x66, 0x6f, 0x12, 0x84, 0x01, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64,
	0x12, 0x28, 0x2e, 0x73, 0x61, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x64, 0x65,
	0x69, 0x64, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x29, 0x2e, 0x73, 0x61, 0x61,
	0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x69, 0x64, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x49,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x12, 0x1a, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x69, 0x64, 0x2f, 0x67, 0x65,
	0x74, 0x2d, 0x6e, 0x6f, 0x64, 0x65, 0x2d, 0x69, 0x64, 0x12, 0x97, 0x01, 0x0a, 0x0d, 0x52, 0x65,
	0x6e, 0x65, 0x77, 0x61, 0x6c, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x2c, 0x2e, 0x73, 0x61,
	0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x69, 0x64, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x61, 0x6c,
	0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x2d, 0x2e, 0x73, 0x61, 0x61, 0x73,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x69, 0x64, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x61, 0x6c, 0x4e, 0x6f,
	0x64, 0x65, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23,
	0x3a, 0x01, 0x2a, 0x22, 0x1e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x64,
	0x65, 0x69, 0x64, 0x2f, 0x72, 0x65, 0x6e, 0x65, 0x77, 0x61, 0x6c, 0x2d, 0x6e, 0x6f, 0x64, 0x65,
	0x2d, 0x69, 0x64, 0x12, 0x97, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x4e,
	0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x2c, 0x2e, 0x73, 0x61, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x69, 0x64, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x1a, 0x2d, 0x2e, 0x73, 0x61, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e,
	0x6f, 0x64, 0x65, 0x69, 0x64, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x3a, 0x01, 0x2a, 0x1a, 0x1e, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x69, 0x64, 0x2f, 0x72, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x2d, 0x6e, 0x6f, 0x64, 0x65, 0x2d, 0x69, 0x64, 0x42, 0x87, 0x01,
	0x0a, 0x19, 0x73, 0x61, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x69,
	0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x76, 0x31, 0x42, 0x16, 0x53, 0x61, 0x61,
	0x73, 0x41, 0x70, 0x69, 0x4e, 0x6f, 0x64, 0x65, 0x69, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x56, 0x31, 0x50, 0x01, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2d, 0x73, 0x61, 0x61, 0x73, 0x2f,
	0x6e, 0x6f, 0x64, 0x65, 0x69, 0x64, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x69, 0x64, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3b, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_api_nodeid_service_v1_services_node_id_service_v1_proto_goTypes = []any{
	(*resources.PingReq)(nil),            // 0: saas.api.nodeid.resourcev1.PingReq
	(*resources.GetServiceInfoReq)(nil),  // 1: saas.api.nodeid.resourcev1.GetServiceInfoReq
	(*resources.GetNodeIdReq)(nil),       // 2: saas.api.nodeid.resourcev1.GetNodeIdReq
	(*resources.RenewalNodeIdReq)(nil),   // 3: saas.api.nodeid.resourcev1.RenewalNodeIdReq
	(*resources.ReleaseNodeIdReq)(nil),   // 4: saas.api.nodeid.resourcev1.ReleaseNodeIdReq
	(*resources.PingResp)(nil),           // 5: saas.api.nodeid.resourcev1.PingResp
	(*resources.GetServiceInfoResp)(nil), // 6: saas.api.nodeid.resourcev1.GetServiceInfoResp
	(*resources.GetNodeIdResp)(nil),      // 7: saas.api.nodeid.resourcev1.GetNodeIdResp
	(*resources.RenewalNodeIdResp)(nil),  // 8: saas.api.nodeid.resourcev1.RenewalNodeIdResp
	(*resources.ReleaseNodeIdResp)(nil),  // 9: saas.api.nodeid.resourcev1.ReleaseNodeIdResp
}
var file_api_nodeid_service_v1_services_node_id_service_v1_proto_depIdxs = []int32{
	0, // 0: saas.api.nodeid.servicev1.SrvNodeIDV1.Ping:input_type -> saas.api.nodeid.resourcev1.PingReq
	1, // 1: saas.api.nodeid.servicev1.SrvNodeIDV1.GetServiceInfo:input_type -> saas.api.nodeid.resourcev1.GetServiceInfoReq
	2, // 2: saas.api.nodeid.servicev1.SrvNodeIDV1.GetNodeId:input_type -> saas.api.nodeid.resourcev1.GetNodeIdReq
	3, // 3: saas.api.nodeid.servicev1.SrvNodeIDV1.RenewalNodeId:input_type -> saas.api.nodeid.resourcev1.RenewalNodeIdReq
	4, // 4: saas.api.nodeid.servicev1.SrvNodeIDV1.ReleaseNodeId:input_type -> saas.api.nodeid.resourcev1.ReleaseNodeIdReq
	5, // 5: saas.api.nodeid.servicev1.SrvNodeIDV1.Ping:output_type -> saas.api.nodeid.resourcev1.PingResp
	6, // 6: saas.api.nodeid.servicev1.SrvNodeIDV1.GetServiceInfo:output_type -> saas.api.nodeid.resourcev1.GetServiceInfoResp
	7, // 7: saas.api.nodeid.servicev1.SrvNodeIDV1.GetNodeId:output_type -> saas.api.nodeid.resourcev1.GetNodeIdResp
	8, // 8: saas.api.nodeid.servicev1.SrvNodeIDV1.RenewalNodeId:output_type -> saas.api.nodeid.resourcev1.RenewalNodeIdResp
	9, // 9: saas.api.nodeid.servicev1.SrvNodeIDV1.ReleaseNodeId:output_type -> saas.api.nodeid.resourcev1.ReleaseNodeIdResp
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_nodeid_service_v1_services_node_id_service_v1_proto_init() }
func file_api_nodeid_service_v1_services_node_id_service_v1_proto_init() {
	if File_api_nodeid_service_v1_services_node_id_service_v1_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_nodeid_service_v1_services_node_id_service_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_nodeid_service_v1_services_node_id_service_v1_proto_goTypes,
		DependencyIndexes: file_api_nodeid_service_v1_services_node_id_service_v1_proto_depIdxs,
	}.Build()
	File_api_nodeid_service_v1_services_node_id_service_v1_proto = out.File
	file_api_nodeid_service_v1_services_node_id_service_v1_proto_rawDesc = nil
	file_api_nodeid_service_v1_services_node_id_service_v1_proto_goTypes = nil
	file_api_nodeid_service_v1_services_node_id_service_v1_proto_depIdxs = nil
}
