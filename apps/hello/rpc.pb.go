// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: apps/hello/pb/rpc.proto

package hello

import (
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

var File_apps_hello_pb_rpc_proto protoreflect.FileDescriptor

var file_apps_hello_pb_rpc_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2f, 0x70, 0x62, 0x2f,
	0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x1a, 0x19, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2f, 0x70, 0x62, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x3c, 0x0a, 0x03, 0x52,
	0x50, 0x43, 0x12, 0x35, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x13,
	0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x75, 0x79, 0x61, 0x73, 0x72, 0x2f, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x2d, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_apps_hello_pb_rpc_proto_goTypes = []interface{}{
	(*HelloRequest)(nil),  // 0: hello.HelloRequest
	(*HelloResponse)(nil), // 1: hello.HelloResponse
}
var file_apps_hello_pb_rpc_proto_depIdxs = []int32{
	0, // 0: hello.RPC.SayHello:input_type -> hello.HelloRequest
	1, // 1: hello.RPC.SayHello:output_type -> hello.HelloResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_apps_hello_pb_rpc_proto_init() }
func file_apps_hello_pb_rpc_proto_init() {
	if File_apps_hello_pb_rpc_proto != nil {
		return
	}
	file_apps_hello_pb_model_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_apps_hello_pb_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_hello_pb_rpc_proto_goTypes,
		DependencyIndexes: file_apps_hello_pb_rpc_proto_depIdxs,
	}.Build()
	File_apps_hello_pb_rpc_proto = out.File
	file_apps_hello_pb_rpc_proto_rawDesc = nil
	file_apps_hello_pb_rpc_proto_goTypes = nil
	file_apps_hello_pb_rpc_proto_depIdxs = nil
}