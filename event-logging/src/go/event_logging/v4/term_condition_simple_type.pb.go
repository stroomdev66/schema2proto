// event_logging/v4/term_condition_simple_type.proto at 0:0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: event_logging/v4/term_condition_simple_type.proto

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

// The types of term used in query predicates.
type TermConditionSimpleType int32

const (
	// Default
	TermConditionSimpleType_TERM_CONDITION_SIMPLE_TYPE_UNSPECIFIED           TermConditionSimpleType = 0
	TermConditionSimpleType_TERM_CONDITION_SIMPLE_TYPE_EXISTS                TermConditionSimpleType = 1
	TermConditionSimpleType_TERM_CONDITION_SIMPLE_TYPE_NOT_EXISTS            TermConditionSimpleType = 2
	TermConditionSimpleType_TERM_CONDITION_SIMPLE_TYPE_CONTAINS              TermConditionSimpleType = 3
	TermConditionSimpleType_TERM_CONDITION_SIMPLE_TYPE_NOT_CONTAINS          TermConditionSimpleType = 4
	TermConditionSimpleType_TERM_CONDITION_SIMPLE_TYPE_EMPTY                 TermConditionSimpleType = 5
	TermConditionSimpleType_TERM_CONDITION_SIMPLE_TYPE_NOT_EMPTY             TermConditionSimpleType = 6
	TermConditionSimpleType_TERM_CONDITION_SIMPLE_TYPE_EQUALS                TermConditionSimpleType = 7
	TermConditionSimpleType_TERM_CONDITION_SIMPLE_TYPE_NOT_EQUALS            TermConditionSimpleType = 8
	TermConditionSimpleType_TERM_CONDITION_SIMPLE_TYPE_GREATER_THAN          TermConditionSimpleType = 9
	TermConditionSimpleType_TERM_CONDITION_SIMPLE_TYPE_GREATER_THAN_EQUAL_TO TermConditionSimpleType = 10
	TermConditionSimpleType_TERM_CONDITION_SIMPLE_TYPE_LESS_THAN             TermConditionSimpleType = 11
	TermConditionSimpleType_TERM_CONDITION_SIMPLE_TYPE_LESS_THAN_EQUAL_TO    TermConditionSimpleType = 12
	TermConditionSimpleType_TERM_CONDITION_SIMPLE_TYPE_STARTS_WITH           TermConditionSimpleType = 13
	TermConditionSimpleType_TERM_CONDITION_SIMPLE_TYPE_NOT_STARTS_WITH       TermConditionSimpleType = 14
	TermConditionSimpleType_TERM_CONDITION_SIMPLE_TYPE_REGEX                 TermConditionSimpleType = 15
	TermConditionSimpleType_TERM_CONDITION_SIMPLE_TYPE_NOT_REGEX             TermConditionSimpleType = 16
)

// Enum value maps for TermConditionSimpleType.
var (
	TermConditionSimpleType_name = map[int32]string{
		0:  "TERM_CONDITION_SIMPLE_TYPE_UNSPECIFIED",
		1:  "TERM_CONDITION_SIMPLE_TYPE_EXISTS",
		2:  "TERM_CONDITION_SIMPLE_TYPE_NOT_EXISTS",
		3:  "TERM_CONDITION_SIMPLE_TYPE_CONTAINS",
		4:  "TERM_CONDITION_SIMPLE_TYPE_NOT_CONTAINS",
		5:  "TERM_CONDITION_SIMPLE_TYPE_EMPTY",
		6:  "TERM_CONDITION_SIMPLE_TYPE_NOT_EMPTY",
		7:  "TERM_CONDITION_SIMPLE_TYPE_EQUALS",
		8:  "TERM_CONDITION_SIMPLE_TYPE_NOT_EQUALS",
		9:  "TERM_CONDITION_SIMPLE_TYPE_GREATER_THAN",
		10: "TERM_CONDITION_SIMPLE_TYPE_GREATER_THAN_EQUAL_TO",
		11: "TERM_CONDITION_SIMPLE_TYPE_LESS_THAN",
		12: "TERM_CONDITION_SIMPLE_TYPE_LESS_THAN_EQUAL_TO",
		13: "TERM_CONDITION_SIMPLE_TYPE_STARTS_WITH",
		14: "TERM_CONDITION_SIMPLE_TYPE_NOT_STARTS_WITH",
		15: "TERM_CONDITION_SIMPLE_TYPE_REGEX",
		16: "TERM_CONDITION_SIMPLE_TYPE_NOT_REGEX",
	}
	TermConditionSimpleType_value = map[string]int32{
		"TERM_CONDITION_SIMPLE_TYPE_UNSPECIFIED":           0,
		"TERM_CONDITION_SIMPLE_TYPE_EXISTS":                1,
		"TERM_CONDITION_SIMPLE_TYPE_NOT_EXISTS":            2,
		"TERM_CONDITION_SIMPLE_TYPE_CONTAINS":              3,
		"TERM_CONDITION_SIMPLE_TYPE_NOT_CONTAINS":          4,
		"TERM_CONDITION_SIMPLE_TYPE_EMPTY":                 5,
		"TERM_CONDITION_SIMPLE_TYPE_NOT_EMPTY":             6,
		"TERM_CONDITION_SIMPLE_TYPE_EQUALS":                7,
		"TERM_CONDITION_SIMPLE_TYPE_NOT_EQUALS":            8,
		"TERM_CONDITION_SIMPLE_TYPE_GREATER_THAN":          9,
		"TERM_CONDITION_SIMPLE_TYPE_GREATER_THAN_EQUAL_TO": 10,
		"TERM_CONDITION_SIMPLE_TYPE_LESS_THAN":             11,
		"TERM_CONDITION_SIMPLE_TYPE_LESS_THAN_EQUAL_TO":    12,
		"TERM_CONDITION_SIMPLE_TYPE_STARTS_WITH":           13,
		"TERM_CONDITION_SIMPLE_TYPE_NOT_STARTS_WITH":       14,
		"TERM_CONDITION_SIMPLE_TYPE_REGEX":                 15,
		"TERM_CONDITION_SIMPLE_TYPE_NOT_REGEX":             16,
	}
)

func (x TermConditionSimpleType) Enum() *TermConditionSimpleType {
	p := new(TermConditionSimpleType)
	*p = x
	return p
}

func (x TermConditionSimpleType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TermConditionSimpleType) Descriptor() protoreflect.EnumDescriptor {
	return file_event_logging_v4_term_condition_simple_type_proto_enumTypes[0].Descriptor()
}

func (TermConditionSimpleType) Type() protoreflect.EnumType {
	return &file_event_logging_v4_term_condition_simple_type_proto_enumTypes[0]
}

func (x TermConditionSimpleType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TermConditionSimpleType.Descriptor instead.
func (TermConditionSimpleType) EnumDescriptor() ([]byte, []int) {
	return file_event_logging_v4_term_condition_simple_type_proto_rawDescGZIP(), []int{0}
}

var File_event_logging_v4_term_condition_simple_type_proto protoreflect.FileDescriptor

var file_event_logging_v4_term_condition_simple_type_proto_rawDesc = []byte{
	0x0a, 0x31, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f,
	0x76, 0x34, 0x2f, 0x74, 0x65, 0x72, 0x6d, 0x5f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x10, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x34, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2a, 0xfb, 0x05, 0x0a, 0x17, 0x54, 0x65, 0x72, 0x6d, 0x43, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2a,
	0x0a, 0x26, 0x54, 0x45, 0x52, 0x4d, 0x5f, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x53, 0x49, 0x4d, 0x50, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x25, 0x0a, 0x21, 0x54, 0x45,
	0x52, 0x4d, 0x5f, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x49, 0x4d,
	0x50, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x10,
	0x01, 0x12, 0x29, 0x0a, 0x25, 0x54, 0x45, 0x52, 0x4d, 0x5f, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x49, 0x4d, 0x50, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x10, 0x02, 0x12, 0x27, 0x0a, 0x23,
	0x54, 0x45, 0x52, 0x4d, 0x5f, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53,
	0x49, 0x4d, 0x50, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x41,
	0x49, 0x4e, 0x53, 0x10, 0x03, 0x12, 0x2b, 0x0a, 0x27, 0x54, 0x45, 0x52, 0x4d, 0x5f, 0x43, 0x4f,
	0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x49, 0x4d, 0x50, 0x4c, 0x45, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x41, 0x49, 0x4e, 0x53,
	0x10, 0x04, 0x12, 0x24, 0x0a, 0x20, 0x54, 0x45, 0x52, 0x4d, 0x5f, 0x43, 0x4f, 0x4e, 0x44, 0x49,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x49, 0x4d, 0x50, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x45, 0x4d, 0x50, 0x54, 0x59, 0x10, 0x05, 0x12, 0x28, 0x0a, 0x24, 0x54, 0x45, 0x52, 0x4d,
	0x5f, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x49, 0x4d, 0x50, 0x4c,
	0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x4d, 0x50, 0x54, 0x59,
	0x10, 0x06, 0x12, 0x25, 0x0a, 0x21, 0x54, 0x45, 0x52, 0x4d, 0x5f, 0x43, 0x4f, 0x4e, 0x44, 0x49,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x49, 0x4d, 0x50, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x53, 0x10, 0x07, 0x12, 0x29, 0x0a, 0x25, 0x54, 0x45, 0x52,
	0x4d, 0x5f, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x49, 0x4d, 0x50,
	0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x51, 0x55, 0x41,
	0x4c, 0x53, 0x10, 0x08, 0x12, 0x2b, 0x0a, 0x27, 0x54, 0x45, 0x52, 0x4d, 0x5f, 0x43, 0x4f, 0x4e,
	0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x49, 0x4d, 0x50, 0x4c, 0x45, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x47, 0x52, 0x45, 0x41, 0x54, 0x45, 0x52, 0x5f, 0x54, 0x48, 0x41, 0x4e, 0x10,
	0x09, 0x12, 0x34, 0x0a, 0x30, 0x54, 0x45, 0x52, 0x4d, 0x5f, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x49, 0x4d, 0x50, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x47, 0x52, 0x45, 0x41, 0x54, 0x45, 0x52, 0x5f, 0x54, 0x48, 0x41, 0x4e, 0x5f, 0x45, 0x51, 0x55,
	0x41, 0x4c, 0x5f, 0x54, 0x4f, 0x10, 0x0a, 0x12, 0x28, 0x0a, 0x24, 0x54, 0x45, 0x52, 0x4d, 0x5f,
	0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x49, 0x4d, 0x50, 0x4c, 0x45,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4c, 0x45, 0x53, 0x53, 0x5f, 0x54, 0x48, 0x41, 0x4e, 0x10,
	0x0b, 0x12, 0x31, 0x0a, 0x2d, 0x54, 0x45, 0x52, 0x4d, 0x5f, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x49, 0x4d, 0x50, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x4c, 0x45, 0x53, 0x53, 0x5f, 0x54, 0x48, 0x41, 0x4e, 0x5f, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x5f,
	0x54, 0x4f, 0x10, 0x0c, 0x12, 0x2a, 0x0a, 0x26, 0x54, 0x45, 0x52, 0x4d, 0x5f, 0x43, 0x4f, 0x4e,
	0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x49, 0x4d, 0x50, 0x4c, 0x45, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x53, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x10, 0x0d,
	0x12, 0x2e, 0x0a, 0x2a, 0x54, 0x45, 0x52, 0x4d, 0x5f, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x53, 0x49, 0x4d, 0x50, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e,
	0x4f, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x53, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x10, 0x0e,
	0x12, 0x24, 0x0a, 0x20, 0x54, 0x45, 0x52, 0x4d, 0x5f, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x53, 0x49, 0x4d, 0x50, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52,
	0x45, 0x47, 0x45, 0x58, 0x10, 0x0f, 0x12, 0x28, 0x0a, 0x24, 0x54, 0x45, 0x52, 0x4d, 0x5f, 0x43,
	0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x49, 0x4d, 0x50, 0x4c, 0x45, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x52, 0x45, 0x47, 0x45, 0x58, 0x10, 0x10,
	0x42, 0xd8, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c,
	0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x42, 0x1c, 0x54, 0x65, 0x72, 0x6d, 0x43,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x63, 0x68, 0x71, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x2d, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2d, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x34,
	0x3b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x76, 0x34,
	0xa2, 0x02, 0x03, 0x45, 0x58, 0x58, 0xaa, 0x02, 0x0f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x6f,
	0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x34, 0xca, 0x02, 0x0f, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x34, 0xe2, 0x02, 0x1b, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x34, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x10, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x56, 0x34, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_event_logging_v4_term_condition_simple_type_proto_rawDescOnce sync.Once
	file_event_logging_v4_term_condition_simple_type_proto_rawDescData = file_event_logging_v4_term_condition_simple_type_proto_rawDesc
)

func file_event_logging_v4_term_condition_simple_type_proto_rawDescGZIP() []byte {
	file_event_logging_v4_term_condition_simple_type_proto_rawDescOnce.Do(func() {
		file_event_logging_v4_term_condition_simple_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_event_logging_v4_term_condition_simple_type_proto_rawDescData)
	})
	return file_event_logging_v4_term_condition_simple_type_proto_rawDescData
}

var file_event_logging_v4_term_condition_simple_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_event_logging_v4_term_condition_simple_type_proto_goTypes = []interface{}{
	(TermConditionSimpleType)(0), // 0: event_logging.v4.TermConditionSimpleType
}
var file_event_logging_v4_term_condition_simple_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_event_logging_v4_term_condition_simple_type_proto_init() }
func file_event_logging_v4_term_condition_simple_type_proto_init() {
	if File_event_logging_v4_term_condition_simple_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_event_logging_v4_term_condition_simple_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_event_logging_v4_term_condition_simple_type_proto_goTypes,
		DependencyIndexes: file_event_logging_v4_term_condition_simple_type_proto_depIdxs,
		EnumInfos:         file_event_logging_v4_term_condition_simple_type_proto_enumTypes,
	}.Build()
	File_event_logging_v4_term_condition_simple_type_proto = out.File
	file_event_logging_v4_term_condition_simple_type_proto_rawDesc = nil
	file_event_logging_v4_term_condition_simple_type_proto_goTypes = nil
	file_event_logging_v4_term_condition_simple_type_proto_depIdxs = nil
}
