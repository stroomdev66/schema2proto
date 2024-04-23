// event_logging/v4/network_event_action_complex_type.proto at 0:0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: event_logging/v4/network_event_action_complex_type.proto

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

// Describes an event involving some form of activity on a computer network.
type NetworkEventActionComplexType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The action of a server binding a network socket to a port and IP address.
	Bind *NetworkOutcomeComplexType `protobuf:"bytes,1,opt,name=bind,proto3" json:"bind,omitempty"`
	// The action of a client system establishing a connection with a server.
	Connect *NetworkOutcomeComplexType `protobuf:"bytes,2,opt,name=connect,proto3" json:"connect,omitempty"`
	// The action of opening an unnamed socket that is bound to an address.
	Open *NetworkOutcomeComplexType `protobuf:"bytes,3,opt,name=open,proto3" json:"open,omitempty"`
	// The action of closing an open socket or connection.
	Close *NetworkOutcomeComplexType `protobuf:"bytes,4,opt,name=close,proto3" json:"close,omitempty"`
	// The action of sending data on a socket.
	Send *NetworkOutcomeComplexType `protobuf:"bytes,5,opt,name=send,proto3" json:"send,omitempty"`
	// The action of receiving data on a socket.
	Receive *NetworkOutcomeComplexType `protobuf:"bytes,6,opt,name=receive,proto3" json:"receive,omitempty"`
	// The action of making a socket listen for connections.
	Listen *NetworkOutcomeComplexType `protobuf:"bytes,7,opt,name=listen,proto3" json:"listen,omitempty"`
	// The action of network traffic being permitted by an Access Control List (ACL).
	Permit *NetworkOutcomeComplexType `protobuf:"bytes,8,opt,name=permit,proto3" json:"permit,omitempty"`
	// The action of network traffic being denied by an Access Control List (ACL).
	Deny *NetworkOutcomeComplexType `protobuf:"bytes,9,opt,name=deny,proto3" json:"deny,omitempty"`
}

func (x *NetworkEventActionComplexType) Reset() {
	*x = NetworkEventActionComplexType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_logging_v4_network_event_action_complex_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkEventActionComplexType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkEventActionComplexType) ProtoMessage() {}

func (x *NetworkEventActionComplexType) ProtoReflect() protoreflect.Message {
	mi := &file_event_logging_v4_network_event_action_complex_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkEventActionComplexType.ProtoReflect.Descriptor instead.
func (*NetworkEventActionComplexType) Descriptor() ([]byte, []int) {
	return file_event_logging_v4_network_event_action_complex_type_proto_rawDescGZIP(), []int{0}
}

func (x *NetworkEventActionComplexType) GetBind() *NetworkOutcomeComplexType {
	if x != nil {
		return x.Bind
	}
	return nil
}

func (x *NetworkEventActionComplexType) GetConnect() *NetworkOutcomeComplexType {
	if x != nil {
		return x.Connect
	}
	return nil
}

func (x *NetworkEventActionComplexType) GetOpen() *NetworkOutcomeComplexType {
	if x != nil {
		return x.Open
	}
	return nil
}

func (x *NetworkEventActionComplexType) GetClose() *NetworkOutcomeComplexType {
	if x != nil {
		return x.Close
	}
	return nil
}

func (x *NetworkEventActionComplexType) GetSend() *NetworkOutcomeComplexType {
	if x != nil {
		return x.Send
	}
	return nil
}

func (x *NetworkEventActionComplexType) GetReceive() *NetworkOutcomeComplexType {
	if x != nil {
		return x.Receive
	}
	return nil
}

func (x *NetworkEventActionComplexType) GetListen() *NetworkOutcomeComplexType {
	if x != nil {
		return x.Listen
	}
	return nil
}

func (x *NetworkEventActionComplexType) GetPermit() *NetworkOutcomeComplexType {
	if x != nil {
		return x.Permit
	}
	return nil
}

func (x *NetworkEventActionComplexType) GetDeny() *NetworkOutcomeComplexType {
	if x != nil {
		return x.Deny
	}
	return nil
}

var File_event_logging_v4_network_event_action_complex_type_proto protoreflect.FileDescriptor

var file_event_logging_v4_network_event_action_complex_type_proto_rawDesc = []byte{
	0x0a, 0x38, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f,
	0x76, 0x34, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x1a, 0x33, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x34, 0x2f, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x6f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x5f, 0x63,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc6,
	0x05, 0x0a, 0x1d, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x47, 0x0a, 0x04, 0x62, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x34, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65, 0x42, 0x06, 0xba, 0x48, 0x03,
	0xc8, 0x01, 0x01, 0x52, 0x04, 0x62, 0x69, 0x6e, 0x64, 0x12, 0x4d, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x2e, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x47, 0x0a, 0x04, 0x6f, 0x70, 0x65, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c,
	0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54,
	0x79, 0x70, 0x65, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x04, 0x6f, 0x70, 0x65,
	0x6e, 0x12, 0x49, 0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2b, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x34, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4f, 0x75, 0x74, 0x63, 0x6f,
	0x6d, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65, 0x42, 0x06, 0xba,
	0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x04,
	0x73, 0x65, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x2e, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52,
	0x04, 0x73, 0x65, 0x6e, 0x64, 0x12, 0x4d, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c,
	0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54,
	0x79, 0x70, 0x65, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x07, 0x72, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x65, 0x12, 0x4b, 0x0a, 0x06, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67,
	0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4f,
	0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x79, 0x70,
	0x65, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x06, 0x6c, 0x69, 0x73, 0x74, 0x65,
	0x6e, 0x12, 0x4b, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2b, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x34, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4f, 0x75, 0x74, 0x63,
	0x6f, 0x6d, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65, 0x42, 0x06,
	0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x06, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x74, 0x12, 0x47,
	0x0a, 0x04, 0x64, 0x65, 0x6e, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x2e,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01,
	0x01, 0x52, 0x04, 0x64, 0x65, 0x6e, 0x79, 0x42, 0xde, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34,
	0x42, 0x22, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x63, 0x68, 0x71, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2d, 0x6c, 0x6f,
	0x67, 0x67, 0x69, 0x6e, 0x67, 0x2d, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x34, 0x3b, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x76, 0x34, 0xa2, 0x02, 0x03,
	0x45, 0x58, 0x58, 0xaa, 0x02, 0x0f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x67, 0x69,
	0x6e, 0x67, 0x2e, 0x56, 0x34, 0xca, 0x02, 0x0f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67,
	0x67, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x34, 0xe2, 0x02, 0x1b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c,
	0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x34, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x10, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67,
	0x67, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x56, 0x34, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_event_logging_v4_network_event_action_complex_type_proto_rawDescOnce sync.Once
	file_event_logging_v4_network_event_action_complex_type_proto_rawDescData = file_event_logging_v4_network_event_action_complex_type_proto_rawDesc
)

func file_event_logging_v4_network_event_action_complex_type_proto_rawDescGZIP() []byte {
	file_event_logging_v4_network_event_action_complex_type_proto_rawDescOnce.Do(func() {
		file_event_logging_v4_network_event_action_complex_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_event_logging_v4_network_event_action_complex_type_proto_rawDescData)
	})
	return file_event_logging_v4_network_event_action_complex_type_proto_rawDescData
}

var file_event_logging_v4_network_event_action_complex_type_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_event_logging_v4_network_event_action_complex_type_proto_goTypes = []interface{}{
	(*NetworkEventActionComplexType)(nil), // 0: event_logging.v4.NetworkEventActionComplexType
	(*NetworkOutcomeComplexType)(nil),     // 1: event_logging.v4.NetworkOutcomeComplexType
}
var file_event_logging_v4_network_event_action_complex_type_proto_depIdxs = []int32{
	1, // 0: event_logging.v4.NetworkEventActionComplexType.bind:type_name -> event_logging.v4.NetworkOutcomeComplexType
	1, // 1: event_logging.v4.NetworkEventActionComplexType.connect:type_name -> event_logging.v4.NetworkOutcomeComplexType
	1, // 2: event_logging.v4.NetworkEventActionComplexType.open:type_name -> event_logging.v4.NetworkOutcomeComplexType
	1, // 3: event_logging.v4.NetworkEventActionComplexType.close:type_name -> event_logging.v4.NetworkOutcomeComplexType
	1, // 4: event_logging.v4.NetworkEventActionComplexType.send:type_name -> event_logging.v4.NetworkOutcomeComplexType
	1, // 5: event_logging.v4.NetworkEventActionComplexType.receive:type_name -> event_logging.v4.NetworkOutcomeComplexType
	1, // 6: event_logging.v4.NetworkEventActionComplexType.listen:type_name -> event_logging.v4.NetworkOutcomeComplexType
	1, // 7: event_logging.v4.NetworkEventActionComplexType.permit:type_name -> event_logging.v4.NetworkOutcomeComplexType
	1, // 8: event_logging.v4.NetworkEventActionComplexType.deny:type_name -> event_logging.v4.NetworkOutcomeComplexType
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_event_logging_v4_network_event_action_complex_type_proto_init() }
func file_event_logging_v4_network_event_action_complex_type_proto_init() {
	if File_event_logging_v4_network_event_action_complex_type_proto != nil {
		return
	}
	file_event_logging_v4_network_outcome_complex_type_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_event_logging_v4_network_event_action_complex_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkEventActionComplexType); i {
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
			RawDescriptor: file_event_logging_v4_network_event_action_complex_type_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_event_logging_v4_network_event_action_complex_type_proto_goTypes,
		DependencyIndexes: file_event_logging_v4_network_event_action_complex_type_proto_depIdxs,
		MessageInfos:      file_event_logging_v4_network_event_action_complex_type_proto_msgTypes,
	}.Build()
	File_event_logging_v4_network_event_action_complex_type_proto = out.File
	file_event_logging_v4_network_event_action_complex_type_proto_rawDesc = nil
	file_event_logging_v4_network_event_action_complex_type_proto_goTypes = nil
	file_event_logging_v4_network_event_action_complex_type_proto_depIdxs = nil
}
