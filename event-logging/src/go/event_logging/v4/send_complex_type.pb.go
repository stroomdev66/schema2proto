// event_logging/v4/send_complex_type.proto at 0:0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: event_logging/v4/send_complex_type.proto

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

// Describes the action of sending something (e.g. a file, data, object, etc.) from a source location/application/system/user to a destination.
type SendComplexType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The initiator(s) of the object or resource sent or received. An initiator can be a user and/or device.
	Source *SendComplexType_SourceType `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	// The destination(s) of the object or resource sent or received. A destination can be a user and/or device.
	Destination *SendComplexType_DestinationType `protobuf:"bytes,2,opt,name=destination,proto3" json:"destination,omitempty"`
	// If the network action is related to message transfer from one place to another then this element describes the message.
	Payload *MultiObjectComplexType `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	// Used to determine if the action was successful. If omitted it is assumed that the event was successful and was permitted.
	Outcome *OutcomeComplexType `protobuf:"bytes,4,opt,name=outcome,proto3" json:"outcome,omitempty"`
	// Any other event data that does not fit into a schema element but may be useful for the purpose of audit.
	Data []*DataComplexType `protobuf:"bytes,5,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *SendComplexType) Reset() {
	*x = SendComplexType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_logging_v4_send_complex_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendComplexType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendComplexType) ProtoMessage() {}

func (x *SendComplexType) ProtoReflect() protoreflect.Message {
	mi := &file_event_logging_v4_send_complex_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendComplexType.ProtoReflect.Descriptor instead.
func (*SendComplexType) Descriptor() ([]byte, []int) {
	return file_event_logging_v4_send_complex_type_proto_rawDescGZIP(), []int{0}
}

func (x *SendComplexType) GetSource() *SendComplexType_SourceType {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *SendComplexType) GetDestination() *SendComplexType_DestinationType {
	if x != nil {
		return x.Destination
	}
	return nil
}

func (x *SendComplexType) GetPayload() *MultiObjectComplexType {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *SendComplexType) GetOutcome() *OutcomeComplexType {
	if x != nil {
		return x.Outcome
	}
	return nil
}

func (x *SendComplexType) GetData() []*DataComplexType {
	if x != nil {
		return x.Data
	}
	return nil
}

// The initiator(s) of the object or resource sent or received. An initiator can be a user and/or device.
type SendComplexType_SourceType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChoiceWrapper []*SendComplexType_SourceType_ChoiceWrapperSourceType `protobuf:"bytes,1,rep,name=choice_wrapper,json=choiceWrapper,proto3" json:"choice_wrapper,omitempty"`
}

func (x *SendComplexType_SourceType) Reset() {
	*x = SendComplexType_SourceType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_logging_v4_send_complex_type_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendComplexType_SourceType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendComplexType_SourceType) ProtoMessage() {}

func (x *SendComplexType_SourceType) ProtoReflect() protoreflect.Message {
	mi := &file_event_logging_v4_send_complex_type_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendComplexType_SourceType.ProtoReflect.Descriptor instead.
func (*SendComplexType_SourceType) Descriptor() ([]byte, []int) {
	return file_event_logging_v4_send_complex_type_proto_rawDescGZIP(), []int{0, 0}
}

func (x *SendComplexType_SourceType) GetChoiceWrapper() []*SendComplexType_SourceType_ChoiceWrapperSourceType {
	if x != nil {
		return x.ChoiceWrapper
	}
	return nil
}

// The destination(s) of the object or resource sent or received. A destination can be a user and/or device.
type SendComplexType_DestinationType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChoiceWrapper []*SendComplexType_DestinationType_ChoiceWrapperDestinationType `protobuf:"bytes,1,rep,name=choice_wrapper,json=choiceWrapper,proto3" json:"choice_wrapper,omitempty"`
}

func (x *SendComplexType_DestinationType) Reset() {
	*x = SendComplexType_DestinationType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_logging_v4_send_complex_type_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendComplexType_DestinationType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendComplexType_DestinationType) ProtoMessage() {}

func (x *SendComplexType_DestinationType) ProtoReflect() protoreflect.Message {
	mi := &file_event_logging_v4_send_complex_type_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendComplexType_DestinationType.ProtoReflect.Descriptor instead.
func (*SendComplexType_DestinationType) Descriptor() ([]byte, []int) {
	return file_event_logging_v4_send_complex_type_proto_rawDescGZIP(), []int{0, 1}
}

func (x *SendComplexType_DestinationType) GetChoiceWrapper() []*SendComplexType_DestinationType_ChoiceWrapperDestinationType {
	if x != nil {
		return x.ChoiceWrapper
	}
	return nil
}

type SendComplexType_SourceType_ChoiceWrapperSourceType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The user that sent the payload.
	User *UserComplexType `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// The device that sent the payload.
	Device *DeviceComplexType `protobuf:"bytes,2,opt,name=device,proto3" json:"device,omitempty"`
}

func (x *SendComplexType_SourceType_ChoiceWrapperSourceType) Reset() {
	*x = SendComplexType_SourceType_ChoiceWrapperSourceType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_logging_v4_send_complex_type_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendComplexType_SourceType_ChoiceWrapperSourceType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendComplexType_SourceType_ChoiceWrapperSourceType) ProtoMessage() {}

func (x *SendComplexType_SourceType_ChoiceWrapperSourceType) ProtoReflect() protoreflect.Message {
	mi := &file_event_logging_v4_send_complex_type_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendComplexType_SourceType_ChoiceWrapperSourceType.ProtoReflect.Descriptor instead.
func (*SendComplexType_SourceType_ChoiceWrapperSourceType) Descriptor() ([]byte, []int) {
	return file_event_logging_v4_send_complex_type_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *SendComplexType_SourceType_ChoiceWrapperSourceType) GetUser() *UserComplexType {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *SendComplexType_SourceType_ChoiceWrapperSourceType) GetDevice() *DeviceComplexType {
	if x != nil {
		return x.Device
	}
	return nil
}

type SendComplexType_DestinationType_ChoiceWrapperDestinationType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The user that the payload is being sent to
	User *UserComplexType `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// The device that the payload is being sent to
	Device *DeviceComplexType `protobuf:"bytes,2,opt,name=device,proto3" json:"device,omitempty"`
}

func (x *SendComplexType_DestinationType_ChoiceWrapperDestinationType) Reset() {
	*x = SendComplexType_DestinationType_ChoiceWrapperDestinationType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_logging_v4_send_complex_type_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendComplexType_DestinationType_ChoiceWrapperDestinationType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendComplexType_DestinationType_ChoiceWrapperDestinationType) ProtoMessage() {}

func (x *SendComplexType_DestinationType_ChoiceWrapperDestinationType) ProtoReflect() protoreflect.Message {
	mi := &file_event_logging_v4_send_complex_type_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendComplexType_DestinationType_ChoiceWrapperDestinationType.ProtoReflect.Descriptor instead.
func (*SendComplexType_DestinationType_ChoiceWrapperDestinationType) Descriptor() ([]byte, []int) {
	return file_event_logging_v4_send_complex_type_proto_rawDescGZIP(), []int{0, 1, 0}
}

func (x *SendComplexType_DestinationType_ChoiceWrapperDestinationType) GetUser() *UserComplexType {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *SendComplexType_DestinationType_ChoiceWrapperDestinationType) GetDevice() *DeviceComplexType {
	if x != nil {
		return x.Device
	}
	return nil
}

var File_event_logging_v4_send_complex_type_proto protoreflect.FileDescriptor

var file_event_logging_v4_send_complex_type_proto_rawDesc = []byte{
	0x0a, 0x28, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f,
	0x76, 0x34, 0x2f, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x1a, 0x28, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x34, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f,
	0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x34, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x30, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e,
	0x67, 0x2f, 0x76, 0x34, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67,
	0x69, 0x6e, 0x67, 0x2f, 0x76, 0x34, 0x2f, 0x6f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x5f, 0x63,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x28, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67,
	0x2f, 0x76, 0x34, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe1, 0x07, 0x0a, 0x0f, 0x53, 0x65, 0x6e,
	0x64, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65, 0x12, 0x4c, 0x0a, 0x06,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x2e,
	0x53, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65, 0x2e,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8,
	0x01, 0x01, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x5b, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x31, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x34, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x79,
	0x70, 0x65, 0x2e, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x42, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x2e, 0x4d, 0x75, 0x6c, 0x74,
	0x69, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x3e, 0x0a, 0x07, 0x6f,
	0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x2e,
	0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x12, 0x3f, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65, 0x42, 0x08, 0xba, 0x48,
	0x05, 0x92, 0x01, 0x02, 0x08, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0xa3, 0x02, 0x0a,
	0x0a, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x75, 0x0a, 0x0e, 0x63,
	0x68, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x44, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x78, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x2e, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x08, 0xba, 0x48, 0x05, 0x92, 0x01,
	0x02, 0x08, 0x01, 0x52, 0x0d, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x57, 0x72, 0x61, 0x70, 0x70,
	0x65, 0x72, 0x1a, 0x9d, 0x01, 0x0a, 0x17, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x57, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3d,
	0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65, 0x42,
	0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x43, 0x0a,
	0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34,
	0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x79,
	0x70, 0x65, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x1a, 0xb7, 0x02, 0x0a, 0x0f, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x7f, 0x0a, 0x0e, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65,
	0x5f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x4e,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x34, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x79, 0x70,
	0x65, 0x2e, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x2e, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x44,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x42, 0x08,
	0xba, 0x48, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x0d, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65,
	0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x1a, 0xa2, 0x01, 0x0a, 0x1c, 0x43, 0x68, 0x6f, 0x69,
	0x63, 0x65, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3d, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c,
	0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01,
	0x01, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x43, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f,
	0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65, 0x42, 0x06, 0xba, 0x48,
	0x03, 0xc8, 0x01, 0x01, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x42, 0xd0, 0x01, 0x0a,
	0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x34, 0x42, 0x14, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x78, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x63, 0x68, 0x71, 0x2f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x2d, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2d, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e,
	0x67, 0x2f, 0x76, 0x34, 0x3b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69,
	0x6e, 0x67, 0x76, 0x34, 0xa2, 0x02, 0x03, 0x45, 0x58, 0x58, 0xaa, 0x02, 0x0f, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x34, 0xca, 0x02, 0x0f, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x34, 0xe2, 0x02,
	0x1b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x34,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x10, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x56, 0x34, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_event_logging_v4_send_complex_type_proto_rawDescOnce sync.Once
	file_event_logging_v4_send_complex_type_proto_rawDescData = file_event_logging_v4_send_complex_type_proto_rawDesc
)

func file_event_logging_v4_send_complex_type_proto_rawDescGZIP() []byte {
	file_event_logging_v4_send_complex_type_proto_rawDescOnce.Do(func() {
		file_event_logging_v4_send_complex_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_event_logging_v4_send_complex_type_proto_rawDescData)
	})
	return file_event_logging_v4_send_complex_type_proto_rawDescData
}

var file_event_logging_v4_send_complex_type_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_event_logging_v4_send_complex_type_proto_goTypes = []interface{}{
	(*SendComplexType)(nil),                                              // 0: event_logging.v4.SendComplexType
	(*SendComplexType_SourceType)(nil),                                   // 1: event_logging.v4.SendComplexType.SourceType
	(*SendComplexType_DestinationType)(nil),                              // 2: event_logging.v4.SendComplexType.DestinationType
	(*SendComplexType_SourceType_ChoiceWrapperSourceType)(nil),           // 3: event_logging.v4.SendComplexType.SourceType.ChoiceWrapperSourceType
	(*SendComplexType_DestinationType_ChoiceWrapperDestinationType)(nil), // 4: event_logging.v4.SendComplexType.DestinationType.ChoiceWrapperDestinationType
	(*MultiObjectComplexType)(nil),                                       // 5: event_logging.v4.MultiObjectComplexType
	(*OutcomeComplexType)(nil),                                           // 6: event_logging.v4.OutcomeComplexType
	(*DataComplexType)(nil),                                              // 7: event_logging.v4.DataComplexType
	(*UserComplexType)(nil),                                              // 8: event_logging.v4.UserComplexType
	(*DeviceComplexType)(nil),                                            // 9: event_logging.v4.DeviceComplexType
}
var file_event_logging_v4_send_complex_type_proto_depIdxs = []int32{
	1,  // 0: event_logging.v4.SendComplexType.source:type_name -> event_logging.v4.SendComplexType.SourceType
	2,  // 1: event_logging.v4.SendComplexType.destination:type_name -> event_logging.v4.SendComplexType.DestinationType
	5,  // 2: event_logging.v4.SendComplexType.payload:type_name -> event_logging.v4.MultiObjectComplexType
	6,  // 3: event_logging.v4.SendComplexType.outcome:type_name -> event_logging.v4.OutcomeComplexType
	7,  // 4: event_logging.v4.SendComplexType.data:type_name -> event_logging.v4.DataComplexType
	3,  // 5: event_logging.v4.SendComplexType.SourceType.choice_wrapper:type_name -> event_logging.v4.SendComplexType.SourceType.ChoiceWrapperSourceType
	4,  // 6: event_logging.v4.SendComplexType.DestinationType.choice_wrapper:type_name -> event_logging.v4.SendComplexType.DestinationType.ChoiceWrapperDestinationType
	8,  // 7: event_logging.v4.SendComplexType.SourceType.ChoiceWrapperSourceType.user:type_name -> event_logging.v4.UserComplexType
	9,  // 8: event_logging.v4.SendComplexType.SourceType.ChoiceWrapperSourceType.device:type_name -> event_logging.v4.DeviceComplexType
	8,  // 9: event_logging.v4.SendComplexType.DestinationType.ChoiceWrapperDestinationType.user:type_name -> event_logging.v4.UserComplexType
	9,  // 10: event_logging.v4.SendComplexType.DestinationType.ChoiceWrapperDestinationType.device:type_name -> event_logging.v4.DeviceComplexType
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_event_logging_v4_send_complex_type_proto_init() }
func file_event_logging_v4_send_complex_type_proto_init() {
	if File_event_logging_v4_send_complex_type_proto != nil {
		return
	}
	file_event_logging_v4_data_complex_type_proto_init()
	file_event_logging_v4_device_complex_type_proto_init()
	file_event_logging_v4_multi_object_complex_type_proto_init()
	file_event_logging_v4_outcome_complex_type_proto_init()
	file_event_logging_v4_user_complex_type_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_event_logging_v4_send_complex_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendComplexType); i {
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
		file_event_logging_v4_send_complex_type_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendComplexType_SourceType); i {
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
		file_event_logging_v4_send_complex_type_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendComplexType_DestinationType); i {
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
		file_event_logging_v4_send_complex_type_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendComplexType_SourceType_ChoiceWrapperSourceType); i {
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
		file_event_logging_v4_send_complex_type_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendComplexType_DestinationType_ChoiceWrapperDestinationType); i {
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
			RawDescriptor: file_event_logging_v4_send_complex_type_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_event_logging_v4_send_complex_type_proto_goTypes,
		DependencyIndexes: file_event_logging_v4_send_complex_type_proto_depIdxs,
		MessageInfos:      file_event_logging_v4_send_complex_type_proto_msgTypes,
	}.Build()
	File_event_logging_v4_send_complex_type_proto = out.File
	file_event_logging_v4_send_complex_type_proto_rawDesc = nil
	file_event_logging_v4_send_complex_type_proto_goTypes = nil
	file_event_logging_v4_send_complex_type_proto_depIdxs = nil
}
