// event_logging/v4/copy_move_outcome_complex_type.proto at 0:0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: event_logging/v4/copy_move_outcome_complex_type.proto

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

// Used to describe the outcome of a copy or move event, including the reason for any failure.
type CopyMoveOutcomeComplexType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If the outcome of an event was successful then 'true', 'false' otherwise. Can be omitted if true as this is the default. The main exception to this default would be if there were many varied Descriptions for the success criteria and such Descriptions could provide additional context to the event.
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	// If an action was permitted then 'true', 'false' otherwise. Can be omitted if true as this is the default. The main exception to this default would be if there were many varied Descriptions for the success criteria and such Descriptions could provide additional context to the event.
	Permitted bool `protobuf:"varint,2,opt,name=permitted,proto3" json:"permitted,omitempty"`
	// A description of the authorisation service that was used to decide if the action was permitted.
	AuthService *CopyMoveOutcomeComplexType_AuthServiceType `protobuf:"bytes,3,opt,name=auth_service,json=authService,proto3" json:"auth_service,omitempty"`
	// Human readable text that describes the outcome.
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// Used to explain the reason for failure, e.g. Device full.
	Reason CopyMoveOutcomeReasonSimpleType `protobuf:"varint,5,opt,name=reason,proto3,enum=event_logging.v4.CopyMoveOutcomeReasonSimpleType" json:"reason,omitempty"`
	// Any other event data that does not fit into a schema element but may be useful for the purpose of audit.
	Data []*DataComplexType `protobuf:"bytes,6,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *CopyMoveOutcomeComplexType) Reset() {
	*x = CopyMoveOutcomeComplexType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_logging_v4_copy_move_outcome_complex_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CopyMoveOutcomeComplexType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CopyMoveOutcomeComplexType) ProtoMessage() {}

func (x *CopyMoveOutcomeComplexType) ProtoReflect() protoreflect.Message {
	mi := &file_event_logging_v4_copy_move_outcome_complex_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CopyMoveOutcomeComplexType.ProtoReflect.Descriptor instead.
func (*CopyMoveOutcomeComplexType) Descriptor() ([]byte, []int) {
	return file_event_logging_v4_copy_move_outcome_complex_type_proto_rawDescGZIP(), []int{0}
}

func (x *CopyMoveOutcomeComplexType) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CopyMoveOutcomeComplexType) GetPermitted() bool {
	if x != nil {
		return x.Permitted
	}
	return false
}

func (x *CopyMoveOutcomeComplexType) GetAuthService() *CopyMoveOutcomeComplexType_AuthServiceType {
	if x != nil {
		return x.AuthService
	}
	return nil
}

func (x *CopyMoveOutcomeComplexType) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CopyMoveOutcomeComplexType) GetReason() CopyMoveOutcomeReasonSimpleType {
	if x != nil {
		return x.Reason
	}
	return CopyMoveOutcomeReasonSimpleType_COPY_MOVE_OUTCOME_REASON_SIMPLE_TYPE_UNSPECIFIED
}

func (x *CopyMoveOutcomeComplexType) GetData() []*DataComplexType {
	if x != nil {
		return x.Data
	}
	return nil
}

// A description of the authorisation service that was used to decide if the action was permitted.
type CopyMoveOutcomeComplexType_AuthServiceType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An identifier for the authorisation service, usually a URI string.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The number of seconds a system is allowed to cache this authorisation before it needs to be checked again.
	CacheTimeout uint32 `protobuf:"varint,2,opt,name=cache_timeout,json=cacheTimeout,proto3" json:"cache_timeout,omitempty"`
}

func (x *CopyMoveOutcomeComplexType_AuthServiceType) Reset() {
	*x = CopyMoveOutcomeComplexType_AuthServiceType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_logging_v4_copy_move_outcome_complex_type_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CopyMoveOutcomeComplexType_AuthServiceType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CopyMoveOutcomeComplexType_AuthServiceType) ProtoMessage() {}

func (x *CopyMoveOutcomeComplexType_AuthServiceType) ProtoReflect() protoreflect.Message {
	mi := &file_event_logging_v4_copy_move_outcome_complex_type_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CopyMoveOutcomeComplexType_AuthServiceType.ProtoReflect.Descriptor instead.
func (*CopyMoveOutcomeComplexType_AuthServiceType) Descriptor() ([]byte, []int) {
	return file_event_logging_v4_copy_move_outcome_complex_type_proto_rawDescGZIP(), []int{0, 0}
}

func (x *CopyMoveOutcomeComplexType_AuthServiceType) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CopyMoveOutcomeComplexType_AuthServiceType) GetCacheTimeout() uint32 {
	if x != nil {
		return x.CacheTimeout
	}
	return 0
}

var File_event_logging_v4_copy_move_outcome_complex_type_proto protoreflect.FileDescriptor

var file_event_logging_v4_copy_move_outcome_complex_type_proto_rawDesc = []byte{
	0x0a, 0x35, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f,
	0x76, 0x34, 0x2f, 0x63, 0x6f, 0x70, 0x79, 0x5f, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x6f, 0x75, 0x74,
	0x63, 0x6f, 0x6d, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c,
	0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x1a, 0x3b, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x34, 0x2f, 0x63, 0x6f, 0x70, 0x79,
	0x5f, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x6f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x5f, 0x72, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x5f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f,
	0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x34, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x63, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x78, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbe, 0x03,
	0x0a, 0x1a, 0x43, 0x6f, 0x70, 0x79, 0x4d, 0x6f, 0x76, 0x65, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d,
	0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x74,
	0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x74, 0x74, 0x65, 0x64, 0x12, 0x5f, 0x0a, 0x0c, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x2e, 0x43, 0x6f,
	0x70, 0x79, 0x4d, 0x6f, 0x76, 0x65, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x53, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x31, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f,
	0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x2e, 0x43, 0x6f, 0x70, 0x79, 0x4d,
	0x6f, 0x76, 0x65, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x08, 0xba, 0x48, 0x05, 0x82,
	0x01, 0x02, 0x10, 0x01, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x3f, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65, 0x42, 0x08, 0xba,
	0x48, 0x05, 0x92, 0x01, 0x02, 0x08, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x4f, 0x0a,
	0x0f, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x2c, 0x0a, 0x0d, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xba, 0x48, 0x04, 0x2a, 0x02, 0x20, 0x00,
	0x52, 0x0c, 0x63, 0x61, 0x63, 0x68, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x42, 0xe8,
	0x01, 0x0a, 0x21, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2d, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x34, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x34, 0x42, 0x1f, 0x43, 0x6f, 0x70, 0x79, 0x4d, 0x6f, 0x76, 0x65, 0x4f, 0x75,
	0x74, 0x63, 0x6f, 0x6d, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x63, 0x68, 0x71, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2d, 0x6c,
	0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2d, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x34, 0x3b, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x76, 0x34, 0xa2, 0x02,
	0x03, 0x45, 0x58, 0x58, 0xaa, 0x02, 0x0f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x67,
	0x69, 0x6e, 0x67, 0x2e, 0x56, 0x34, 0xca, 0x02, 0x0f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x6f,
	0x67, 0x67, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x34, 0xe2, 0x02, 0x1b, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x34, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x10, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x6f,
	0x67, 0x67, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x56, 0x34, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_event_logging_v4_copy_move_outcome_complex_type_proto_rawDescOnce sync.Once
	file_event_logging_v4_copy_move_outcome_complex_type_proto_rawDescData = file_event_logging_v4_copy_move_outcome_complex_type_proto_rawDesc
)

func file_event_logging_v4_copy_move_outcome_complex_type_proto_rawDescGZIP() []byte {
	file_event_logging_v4_copy_move_outcome_complex_type_proto_rawDescOnce.Do(func() {
		file_event_logging_v4_copy_move_outcome_complex_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_event_logging_v4_copy_move_outcome_complex_type_proto_rawDescData)
	})
	return file_event_logging_v4_copy_move_outcome_complex_type_proto_rawDescData
}

var file_event_logging_v4_copy_move_outcome_complex_type_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_event_logging_v4_copy_move_outcome_complex_type_proto_goTypes = []interface{}{
	(*CopyMoveOutcomeComplexType)(nil),                 // 0: event_logging.v4.CopyMoveOutcomeComplexType
	(*CopyMoveOutcomeComplexType_AuthServiceType)(nil), // 1: event_logging.v4.CopyMoveOutcomeComplexType.AuthServiceType
	(CopyMoveOutcomeReasonSimpleType)(0),               // 2: event_logging.v4.CopyMoveOutcomeReasonSimpleType
	(*DataComplexType)(nil),                            // 3: event_logging.v4.DataComplexType
}
var file_event_logging_v4_copy_move_outcome_complex_type_proto_depIdxs = []int32{
	1, // 0: event_logging.v4.CopyMoveOutcomeComplexType.auth_service:type_name -> event_logging.v4.CopyMoveOutcomeComplexType.AuthServiceType
	2, // 1: event_logging.v4.CopyMoveOutcomeComplexType.reason:type_name -> event_logging.v4.CopyMoveOutcomeReasonSimpleType
	3, // 2: event_logging.v4.CopyMoveOutcomeComplexType.data:type_name -> event_logging.v4.DataComplexType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_event_logging_v4_copy_move_outcome_complex_type_proto_init() }
func file_event_logging_v4_copy_move_outcome_complex_type_proto_init() {
	if File_event_logging_v4_copy_move_outcome_complex_type_proto != nil {
		return
	}
	file_event_logging_v4_copy_move_outcome_reason_simple_type_proto_init()
	file_event_logging_v4_data_complex_type_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_event_logging_v4_copy_move_outcome_complex_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CopyMoveOutcomeComplexType); i {
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
		file_event_logging_v4_copy_move_outcome_complex_type_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CopyMoveOutcomeComplexType_AuthServiceType); i {
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
			RawDescriptor: file_event_logging_v4_copy_move_outcome_complex_type_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_event_logging_v4_copy_move_outcome_complex_type_proto_goTypes,
		DependencyIndexes: file_event_logging_v4_copy_move_outcome_complex_type_proto_depIdxs,
		MessageInfos:      file_event_logging_v4_copy_move_outcome_complex_type_proto_msgTypes,
	}.Build()
	File_event_logging_v4_copy_move_outcome_complex_type_proto = out.File
	file_event_logging_v4_copy_move_outcome_complex_type_proto_rawDesc = nil
	file_event_logging_v4_copy_move_outcome_complex_type_proto_goTypes = nil
	file_event_logging_v4_copy_move_outcome_complex_type_proto_depIdxs = nil
}
