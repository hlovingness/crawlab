// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.20.1
// source: services/dependency_service_v2.proto

package grpc

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

var File_services_dependency_service_v2_proto protoreflect.FileDescriptor

var file_services_dependency_service_v2_proto_rawDesc = []byte{
	0x0a, 0x24, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x64, 0x65, 0x70, 0x65, 0x6e,
	0x64, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x76, 0x32,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x2c, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2f, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65,
	0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x76, 0x32, 0x5f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x32, 0xc9, 0x02, 0x0a, 0x13, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x56, 0x32, 0x12, 0x48, 0x0a, 0x07, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x12, 0x29, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x65, 0x70, 0x65,
	0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x56,
	0x32, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x28, 0x01, 0x12, 0x40, 0x0a, 0x04, 0x53, 0x79, 0x6e, 0x63, 0x12, 0x26, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x56, 0x32, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x07, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c,
	0x12, 0x29, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e,
	0x63, 0x69, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x56, 0x32, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30,
	0x01, 0x12, 0x5a, 0x0a, 0x15, 0x55, 0x6e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x44, 0x65,
	0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x12, 0x2b, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x56, 0x32, 0x55, 0x6e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x08, 0x5a,
	0x06, 0x2e, 0x3b, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_services_dependency_service_v2_proto_goTypes = []interface{}{
	(*DependenciesServiceV2ConnectRequest)(nil),   // 0: grpc.DependenciesServiceV2ConnectRequest
	(*DependenciesServiceV2SyncRequest)(nil),      // 1: grpc.DependenciesServiceV2SyncRequest
	(*DependenciesServiceV2InstallRequest)(nil),   // 2: grpc.DependenciesServiceV2InstallRequest
	(*DependenciesServiceV2UninstallRequest)(nil), // 3: grpc.DependenciesServiceV2UninstallRequest
	(*Response)(nil), // 4: grpc.Response
}
var file_services_dependency_service_v2_proto_depIdxs = []int32{
	0, // 0: grpc.DependencyServiceV2.Connect:input_type -> grpc.DependenciesServiceV2ConnectRequest
	1, // 1: grpc.DependencyServiceV2.Sync:input_type -> grpc.DependenciesServiceV2SyncRequest
	2, // 2: grpc.DependencyServiceV2.Install:input_type -> grpc.DependenciesServiceV2InstallRequest
	3, // 3: grpc.DependencyServiceV2.UninstallDependencies:input_type -> grpc.DependenciesServiceV2UninstallRequest
	4, // 4: grpc.DependencyServiceV2.Connect:output_type -> grpc.Response
	4, // 5: grpc.DependencyServiceV2.Sync:output_type -> grpc.Response
	4, // 6: grpc.DependencyServiceV2.Install:output_type -> grpc.Response
	4, // 7: grpc.DependencyServiceV2.UninstallDependencies:output_type -> grpc.Response
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_services_dependency_service_v2_proto_init() }
func file_services_dependency_service_v2_proto_init() {
	if File_services_dependency_service_v2_proto != nil {
		return
	}
	file_entity_dependencies_service_v2_request_proto_init()
	file_entity_response_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_services_dependency_service_v2_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_dependency_service_v2_proto_goTypes,
		DependencyIndexes: file_services_dependency_service_v2_proto_depIdxs,
	}.Build()
	File_services_dependency_service_v2_proto = out.File
	file_services_dependency_service_v2_proto_rawDesc = nil
	file_services_dependency_service_v2_proto_goTypes = nil
	file_services_dependency_service_v2_proto_depIdxs = nil
}
