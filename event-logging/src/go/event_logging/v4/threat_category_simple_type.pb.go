// event_logging/v4/threat_category_simple_type.proto at 0:0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: event_logging/v4/threat_category_simple_type.proto

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

// The types of threat assigned to detected malware.
type ThreatCategorySimpleType int32

const (
	// Default
	ThreatCategorySimpleType_THREAT_CATEGORY_SIMPLE_TYPE_UNSPECIFIED  ThreatCategorySimpleType = 0
	ThreatCategorySimpleType_THREAT_CATEGORY_SIMPLE_TYPE_VIRUS        ThreatCategorySimpleType = 1
	ThreatCategorySimpleType_THREAT_CATEGORY_SIMPLE_TYPE_WORM         ThreatCategorySimpleType = 2
	ThreatCategorySimpleType_THREAT_CATEGORY_SIMPLE_TYPE_HACKING_TOOL ThreatCategorySimpleType = 3
	ThreatCategorySimpleType_THREAT_CATEGORY_SIMPLE_TYPE_SPYWARE      ThreatCategorySimpleType = 4
	ThreatCategorySimpleType_THREAT_CATEGORY_SIMPLE_TYPE_TROJAN_HORSE ThreatCategorySimpleType = 5
	ThreatCategorySimpleType_THREAT_CATEGORY_SIMPLE_TYPE_ADWARE       ThreatCategorySimpleType = 6
	ThreatCategorySimpleType_THREAT_CATEGORY_SIMPLE_TYPE_OTHER        ThreatCategorySimpleType = 7
)

// Enum value maps for ThreatCategorySimpleType.
var (
	ThreatCategorySimpleType_name = map[int32]string{
		0: "THREAT_CATEGORY_SIMPLE_TYPE_UNSPECIFIED",
		1: "THREAT_CATEGORY_SIMPLE_TYPE_VIRUS",
		2: "THREAT_CATEGORY_SIMPLE_TYPE_WORM",
		3: "THREAT_CATEGORY_SIMPLE_TYPE_HACKING_TOOL",
		4: "THREAT_CATEGORY_SIMPLE_TYPE_SPYWARE",
		5: "THREAT_CATEGORY_SIMPLE_TYPE_TROJAN_HORSE",
		6: "THREAT_CATEGORY_SIMPLE_TYPE_ADWARE",
		7: "THREAT_CATEGORY_SIMPLE_TYPE_OTHER",
	}
	ThreatCategorySimpleType_value = map[string]int32{
		"THREAT_CATEGORY_SIMPLE_TYPE_UNSPECIFIED":  0,
		"THREAT_CATEGORY_SIMPLE_TYPE_VIRUS":        1,
		"THREAT_CATEGORY_SIMPLE_TYPE_WORM":         2,
		"THREAT_CATEGORY_SIMPLE_TYPE_HACKING_TOOL": 3,
		"THREAT_CATEGORY_SIMPLE_TYPE_SPYWARE":      4,
		"THREAT_CATEGORY_SIMPLE_TYPE_TROJAN_HORSE": 5,
		"THREAT_CATEGORY_SIMPLE_TYPE_ADWARE":       6,
		"THREAT_CATEGORY_SIMPLE_TYPE_OTHER":        7,
	}
)

func (x ThreatCategorySimpleType) Enum() *ThreatCategorySimpleType {
	p := new(ThreatCategorySimpleType)
	*p = x
	return p
}

func (x ThreatCategorySimpleType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ThreatCategorySimpleType) Descriptor() protoreflect.EnumDescriptor {
	return file_event_logging_v4_threat_category_simple_type_proto_enumTypes[0].Descriptor()
}

func (ThreatCategorySimpleType) Type() protoreflect.EnumType {
	return &file_event_logging_v4_threat_category_simple_type_proto_enumTypes[0]
}

func (x ThreatCategorySimpleType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ThreatCategorySimpleType.Descriptor instead.
func (ThreatCategorySimpleType) EnumDescriptor() ([]byte, []int) {
	return file_event_logging_v4_threat_category_simple_type_proto_rawDescGZIP(), []int{0}
}

var File_event_logging_v4_threat_category_simple_type_proto protoreflect.FileDescriptor

var file_event_logging_v4_threat_category_simple_type_proto_rawDesc = []byte{
	0x0a, 0x32, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f,
	0x76, 0x34, 0x2f, 0x74, 0x68, 0x72, 0x65, 0x61, 0x74, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x5f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2a, 0xe8, 0x02, 0x0a, 0x18, 0x54, 0x68, 0x72, 0x65, 0x61, 0x74, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x2b, 0x0a, 0x27, 0x54, 0x48, 0x52, 0x45, 0x41, 0x54, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47,
	0x4f, 0x52, 0x59, 0x5f, 0x53, 0x49, 0x4d, 0x50, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x25, 0x0a,
	0x21, 0x54, 0x48, 0x52, 0x45, 0x41, 0x54, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59,
	0x5f, 0x53, 0x49, 0x4d, 0x50, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x56, 0x49, 0x52,
	0x55, 0x53, 0x10, 0x01, 0x12, 0x24, 0x0a, 0x20, 0x54, 0x48, 0x52, 0x45, 0x41, 0x54, 0x5f, 0x43,
	0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x53, 0x49, 0x4d, 0x50, 0x4c, 0x45, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x57, 0x4f, 0x52, 0x4d, 0x10, 0x02, 0x12, 0x2c, 0x0a, 0x28, 0x54, 0x48,
	0x52, 0x45, 0x41, 0x54, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x53, 0x49,
	0x4d, 0x50, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x48, 0x41, 0x43, 0x4b, 0x49, 0x4e,
	0x47, 0x5f, 0x54, 0x4f, 0x4f, 0x4c, 0x10, 0x03, 0x12, 0x27, 0x0a, 0x23, 0x54, 0x48, 0x52, 0x45,
	0x41, 0x54, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x53, 0x49, 0x4d, 0x50,
	0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x50, 0x59, 0x57, 0x41, 0x52, 0x45, 0x10,
	0x04, 0x12, 0x2c, 0x0a, 0x28, 0x54, 0x48, 0x52, 0x45, 0x41, 0x54, 0x5f, 0x43, 0x41, 0x54, 0x45,
	0x47, 0x4f, 0x52, 0x59, 0x5f, 0x53, 0x49, 0x4d, 0x50, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x54, 0x52, 0x4f, 0x4a, 0x41, 0x4e, 0x5f, 0x48, 0x4f, 0x52, 0x53, 0x45, 0x10, 0x05, 0x12,
	0x26, 0x0a, 0x22, 0x54, 0x48, 0x52, 0x45, 0x41, 0x54, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f,
	0x52, 0x59, 0x5f, 0x53, 0x49, 0x4d, 0x50, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41,
	0x44, 0x57, 0x41, 0x52, 0x45, 0x10, 0x06, 0x12, 0x25, 0x0a, 0x21, 0x54, 0x48, 0x52, 0x45, 0x41,
	0x54, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x53, 0x49, 0x4d, 0x50, 0x4c,
	0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x10, 0x07, 0x42, 0xe1,
	0x01, 0x0a, 0x1c, 0x75, 0x6b, 0x2e, 0x67, 0x6f, 0x76, 0x2e, 0x67, 0x63, 0x68, 0x71, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x42,
	0x1d, 0x54, 0x68, 0x72, 0x65, 0x61, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x53,
	0x69, 0x6d, 0x70, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x63, 0x68,
	0x71, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2d, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2d,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67,
	0x67, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x34, 0x3b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f,
	0x67, 0x67, 0x69, 0x6e, 0x67, 0x76, 0x34, 0xa2, 0x02, 0x03, 0x45, 0x58, 0x58, 0xaa, 0x02, 0x0f,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x34, 0xca,
	0x02, 0x0f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x5c, 0x56,
	0x34, 0xe2, 0x02, 0x1b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67,
	0x5c, 0x56, 0x34, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x10, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x3a, 0x3a,
	0x56, 0x34, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_event_logging_v4_threat_category_simple_type_proto_rawDescOnce sync.Once
	file_event_logging_v4_threat_category_simple_type_proto_rawDescData = file_event_logging_v4_threat_category_simple_type_proto_rawDesc
)

func file_event_logging_v4_threat_category_simple_type_proto_rawDescGZIP() []byte {
	file_event_logging_v4_threat_category_simple_type_proto_rawDescOnce.Do(func() {
		file_event_logging_v4_threat_category_simple_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_event_logging_v4_threat_category_simple_type_proto_rawDescData)
	})
	return file_event_logging_v4_threat_category_simple_type_proto_rawDescData
}

var file_event_logging_v4_threat_category_simple_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_event_logging_v4_threat_category_simple_type_proto_goTypes = []interface{}{
	(ThreatCategorySimpleType)(0), // 0: event_logging.v4.ThreatCategorySimpleType
}
var file_event_logging_v4_threat_category_simple_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_event_logging_v4_threat_category_simple_type_proto_init() }
func file_event_logging_v4_threat_category_simple_type_proto_init() {
	if File_event_logging_v4_threat_category_simple_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_event_logging_v4_threat_category_simple_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_event_logging_v4_threat_category_simple_type_proto_goTypes,
		DependencyIndexes: file_event_logging_v4_threat_category_simple_type_proto_depIdxs,
		EnumInfos:         file_event_logging_v4_threat_category_simple_type_proto_enumTypes,
	}.Build()
	File_event_logging_v4_threat_category_simple_type_proto = out.File
	file_event_logging_v4_threat_category_simple_type_proto_rawDesc = nil
	file_event_logging_v4_threat_category_simple_type_proto_goTypes = nil
	file_event_logging_v4_threat_category_simple_type_proto_depIdxs = nil
}
