// event_logging/v4/process_type_simple_type.proto at 0:0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: event_logging/v4/process_type_simple_type.proto

package event_loggingv4

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
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

// The types of a process event, e.g. a background process running within an application.
type ProcessTypeSimpleType int32

const (
	// Default
	ProcessTypeSimpleType_PROCESS_TYPE_SIMPLE_TYPE_UNSPECIFIED ProcessTypeSimpleType = 0
	ProcessTypeSimpleType_PROCESS_TYPE_SIMPLE_TYPE_OS          ProcessTypeSimpleType = 1
	ProcessTypeSimpleType_PROCESS_TYPE_SIMPLE_TYPE_SERVICE     ProcessTypeSimpleType = 2
	ProcessTypeSimpleType_PROCESS_TYPE_SIMPLE_TYPE_APPLICATION ProcessTypeSimpleType = 3
)

// Enum value maps for ProcessTypeSimpleType.
var (
	ProcessTypeSimpleType_name = map[int32]string{
		0: "PROCESS_TYPE_SIMPLE_TYPE_UNSPECIFIED",
		1: "PROCESS_TYPE_SIMPLE_TYPE_OS",
		2: "PROCESS_TYPE_SIMPLE_TYPE_SERVICE",
		3: "PROCESS_TYPE_SIMPLE_TYPE_APPLICATION",
	}
	ProcessTypeSimpleType_value = map[string]int32{
		"PROCESS_TYPE_SIMPLE_TYPE_UNSPECIFIED": 0,
		"PROCESS_TYPE_SIMPLE_TYPE_OS":          1,
		"PROCESS_TYPE_SIMPLE_TYPE_SERVICE":     2,
		"PROCESS_TYPE_SIMPLE_TYPE_APPLICATION": 3,
	}
)

func (x ProcessTypeSimpleType) Enum() *ProcessTypeSimpleType {
	p := new(ProcessTypeSimpleType)
	*p = x
	return p
}

func (x ProcessTypeSimpleType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProcessTypeSimpleType) Descriptor() protoreflect.EnumDescriptor {
	return file_event_logging_v4_process_type_simple_type_proto_enumTypes[0].Descriptor()
}

func (ProcessTypeSimpleType) Type() protoreflect.EnumType {
	return &file_event_logging_v4_process_type_simple_type_proto_enumTypes[0]
}

func (x ProcessTypeSimpleType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProcessTypeSimpleType.Descriptor instead.
func (ProcessTypeSimpleType) EnumDescriptor() ([]byte, []int) {
	return file_event_logging_v4_process_type_simple_type_proto_rawDescGZIP(), []int{0}
}

var File_event_logging_v4_process_type_simple_type_proto protoreflect.FileDescriptor

var file_event_logging_v4_process_type_simple_type_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f,
	0x76, 0x34, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f,
	0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x10, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x34, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2a, 0xb2, 0x01, 0x0a, 0x15, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x54, 0x79, 0x70, 0x65,
	0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x24, 0x50, 0x52,
	0x4f, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x49, 0x4d, 0x50, 0x4c,
	0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x1f, 0x0a, 0x1b, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x49, 0x4d, 0x50, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x4f, 0x53, 0x10, 0x01, 0x12, 0x24, 0x0a, 0x20, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x49, 0x4d, 0x50, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x10, 0x02, 0x12, 0x28, 0x0a, 0x24, 0x50,
	0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x49, 0x4d, 0x50,
	0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x50, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x10, 0x03, 0x42, 0xd6, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x42, 0x1a,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x54, 0x79, 0x70, 0x65, 0x53, 0x69, 0x6d, 0x70, 0x6c,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x63, 0x68, 0x71, 0x2f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x2d, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2d, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67,
	0x2f, 0x76, 0x34, 0x3b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e,
	0x67, 0x76, 0x34, 0xa2, 0x02, 0x03, 0x45, 0x58, 0x58, 0xaa, 0x02, 0x0f, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x34, 0xca, 0x02, 0x0f, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x34, 0xe2, 0x02, 0x1b,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x34, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x10, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x56, 0x34, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_event_logging_v4_process_type_simple_type_proto_rawDescOnce sync.Once
	file_event_logging_v4_process_type_simple_type_proto_rawDescData = file_event_logging_v4_process_type_simple_type_proto_rawDesc
)

func file_event_logging_v4_process_type_simple_type_proto_rawDescGZIP() []byte {
	file_event_logging_v4_process_type_simple_type_proto_rawDescOnce.Do(func() {
		file_event_logging_v4_process_type_simple_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_event_logging_v4_process_type_simple_type_proto_rawDescData)
	})
	return file_event_logging_v4_process_type_simple_type_proto_rawDescData
}

var file_event_logging_v4_process_type_simple_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_event_logging_v4_process_type_simple_type_proto_goTypes = []interface{}{
	(ProcessTypeSimpleType)(0), // 0: event_logging.v4.ProcessTypeSimpleType
}
var file_event_logging_v4_process_type_simple_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_event_logging_v4_process_type_simple_type_proto_init() }
func file_event_logging_v4_process_type_simple_type_proto_init() {
	if File_event_logging_v4_process_type_simple_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_event_logging_v4_process_type_simple_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_event_logging_v4_process_type_simple_type_proto_goTypes,
		DependencyIndexes: file_event_logging_v4_process_type_simple_type_proto_depIdxs,
		EnumInfos:         file_event_logging_v4_process_type_simple_type_proto_enumTypes,
	}.Build()
	File_event_logging_v4_process_type_simple_type_proto = out.File
	file_event_logging_v4_process_type_simple_type_proto_rawDesc = nil
	file_event_logging_v4_process_type_simple_type_proto_goTypes = nil
	file_event_logging_v4_process_type_simple_type_proto_depIdxs = nil
}
