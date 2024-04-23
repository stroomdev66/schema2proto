// event_logging/v4/device_complex_type.proto at 0:0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: event_logging/v4/device_complex_type.proto

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

// Describes a device, e.g. a workstation, server or item of network infrastructure.
type DeviceComplexType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifier used to uniquely identify the device within the organisation's asset register/system. Also, this can be used to identify a device that does not have HostName/IPAddress/MACAddress.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// A descriptive name of the device, e.g. 'Sun Fire X4600', 'HP LaserJet 4+'.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// The security classification associated with this device.
	Classification *ClassificationComplexType `protobuf:"bytes,3,opt,name=classification,proto3" json:"classification,omitempty"`
	// The network host name of the device, e.g. someserver.somenet.org.uk. Ideally this field should always contain a fully qualified DNS name of the host.
	HostName string `protobuf:"bytes,4,opt,name=host_name,json=hostName,proto3" json:"host_name,omitempty"`
	// The network IP address of the device, e.g. 192.168.0.3. Ideally this should always be supplied.
	IpAddress string `protobuf:"bytes,5,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	// The Media Access Control (MAC) address of the device.
	MacAddress string `protobuf:"bytes,6,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	// The network port that is being used on the device, e.g. 443.
	Port int32 `protobuf:"varint,7,opt,name=port,proto3" json:"port,omitempty"`
	// Describes the geographic location of the device.
	Location *LocationComplexType `protobuf:"bytes,8,opt,name=location,proto3" json:"location,omitempty"`
	// Any other event data that does not fit into a schema element but may be useful for the purpose of audit.
	Data []*DataComplexType `protobuf:"bytes,9,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *DeviceComplexType) Reset() {
	*x = DeviceComplexType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_logging_v4_device_complex_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceComplexType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceComplexType) ProtoMessage() {}

func (x *DeviceComplexType) ProtoReflect() protoreflect.Message {
	mi := &file_event_logging_v4_device_complex_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceComplexType.ProtoReflect.Descriptor instead.
func (*DeviceComplexType) Descriptor() ([]byte, []int) {
	return file_event_logging_v4_device_complex_type_proto_rawDescGZIP(), []int{0}
}

func (x *DeviceComplexType) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeviceComplexType) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeviceComplexType) GetClassification() *ClassificationComplexType {
	if x != nil {
		return x.Classification
	}
	return nil
}

func (x *DeviceComplexType) GetHostName() string {
	if x != nil {
		return x.HostName
	}
	return ""
}

func (x *DeviceComplexType) GetIpAddress() string {
	if x != nil {
		return x.IpAddress
	}
	return ""
}

func (x *DeviceComplexType) GetMacAddress() string {
	if x != nil {
		return x.MacAddress
	}
	return ""
}

func (x *DeviceComplexType) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *DeviceComplexType) GetLocation() *LocationComplexType {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *DeviceComplexType) GetData() []*DataComplexType {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_event_logging_v4_device_complex_type_proto protoreflect.FileDescriptor

var file_event_logging_v4_device_complex_type_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f,
	0x76, 0x34, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x78, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x1a, 0x32,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x34,
	0x2f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x28, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e,
	0x67, 0x2f, 0x76, 0x34, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x78, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x34, 0x2f, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb5, 0x0d, 0x0a, 0x11, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x53, 0x0a, 0x0e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x2e, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x78, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0xe1, 0x09, 0x0a, 0x0a, 0x69, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0xc1, 0x09, 0xba, 0x48, 0xbd, 0x09, 0x72,
	0xba, 0x09, 0x32, 0xb7, 0x09, 0x28, 0x28, 0x32, 0x35, 0x5b, 0x30, 0x2d, 0x35, 0x5d, 0x7c, 0x32,
	0x5b, 0x30, 0x2d, 0x34, 0x5d, 0x5c, 0x64, 0x7c, 0x5b, 0x30, 0x2d, 0x31, 0x5d, 0x3f, 0x5c, 0x64,
	0x3f, 0x5c, 0x64, 0x29, 0x28, 0x5c, 0x2e, 0x28, 0x32, 0x35, 0x5b, 0x30, 0x2d, 0x35, 0x5d, 0x7c,
	0x32, 0x5b, 0x30, 0x2d, 0x34, 0x5d, 0x5c, 0x64, 0x7c, 0x5b, 0x30, 0x2d, 0x31, 0x5d, 0x3f, 0x5c,
	0x64, 0x3f, 0x5c, 0x64, 0x29, 0x29, 0x7b, 0x33, 0x7d, 0x29, 0x7c, 0x28, 0x3a, 0x3a, 0x29, 0x7c,
	0x28, 0x28, 0x5b, 0x30, 0x2d, 0x39, 0x61, 0x2d, 0x66, 0x5d, 0x7b, 0x31, 0x2c, 0x34, 0x7d, 0x3a,
	0x29, 0x28, 0x3a, 0x5b, 0x30, 0x2d, 0x39, 0x61, 0x2d, 0x66, 0x5d, 0x7b, 0x31, 0x2c, 0x34, 0x7d,
	0x29, 0x7b, 0x31, 0x2c, 0x36, 0x7d, 0x29, 0x7c, 0x28, 0x28, 0x5b, 0x30, 0x2d, 0x39, 0x61, 0x2d,
	0x66, 0x5d, 0x7b, 0x31, 0x2c, 0x34, 0x7d, 0x3a, 0x29, 0x7b, 0x31, 0x2c, 0x32, 0x7d, 0x28, 0x3a,
	0x5b, 0x30, 0x2d, 0x39, 0x61, 0x2d, 0x66, 0x5d, 0x7b, 0x31, 0x2c, 0x34, 0x7d, 0x29, 0x7b, 0x31,
	0x2c, 0x35, 0x7d, 0x29, 0x7c, 0x28, 0x28, 0x5b, 0x30, 0x2d, 0x39, 0x61, 0x2d, 0x66, 0x5d, 0x7b,
	0x31, 0x2c, 0x34, 0x7d, 0x3a, 0x29, 0x7b, 0x31, 0x2c, 0x33, 0x7d, 0x28, 0x3a, 0x5b, 0x30, 0x2d,
	0x39, 0x61, 0x2d, 0x66, 0x5d, 0x7b, 0x31, 0x2c, 0x34, 0x7d, 0x29, 0x7b, 0x31, 0x2c, 0x34, 0x7d,
	0x29, 0x7c, 0x28, 0x28, 0x5b, 0x30, 0x2d, 0x39, 0x61, 0x2d, 0x66, 0x5d, 0x7b, 0x31, 0x2c, 0x34,
	0x7d, 0x3a, 0x29, 0x7b, 0x31, 0x2c, 0x34, 0x7d, 0x28, 0x3a, 0x5b, 0x30, 0x2d, 0x39, 0x61, 0x2d,
	0x66, 0x5d, 0x7b, 0x31, 0x2c, 0x34, 0x7d, 0x29, 0x7b, 0x31, 0x2c, 0x33, 0x7d, 0x29, 0x7c, 0x28,
	0x28, 0x5b, 0x30, 0x2d, 0x39, 0x61, 0x2d, 0x66, 0x5d, 0x7b, 0x31, 0x2c, 0x34, 0x7d, 0x3a, 0x29,
	0x7b, 0x31, 0x2c, 0x35, 0x7d, 0x28, 0x3a, 0x5b, 0x30, 0x2d, 0x39, 0x61, 0x2d, 0x66, 0x5d, 0x7b,
	0x31, 0x2c, 0x34, 0x7d, 0x29, 0x7b, 0x31, 0x2c, 0x32, 0x7d, 0x29, 0x7c, 0x28, 0x28, 0x5b, 0x30,
	0x2d, 0x39, 0x61, 0x2d, 0x66, 0x5d, 0x7b, 0x31, 0x2c, 0x34, 0x7d, 0x3a, 0x29, 0x7b, 0x31, 0x2c,
	0x36, 0x7d, 0x28, 0x3a, 0x5b, 0x30, 0x2d, 0x39, 0x61, 0x2d, 0x66, 0x5d, 0x7b, 0x31, 0x2c, 0x34,
	0x7d, 0x29, 0x29, 0x7c, 0x28, 0x28, 0x5b, 0x30, 0x2d, 0x39, 0x61, 0x2d, 0x66, 0x5d, 0x7b, 0x31,
	0x2c, 0x34, 0x7d, 0x3a, 0x29, 0x7b, 0x31, 0x2c, 0x37, 0x7d, 0x3a, 0x29, 0x7c, 0x28, 0x3a, 0x28,
	0x3a, 0x5b, 0x30, 0x2d, 0x39, 0x61, 0x2d, 0x66, 0x5d, 0x7b, 0x31, 0x2c, 0x34, 0x7d, 0x29, 0x7b,
	0x31, 0x2c, 0x37, 0x7d, 0x29, 0x7c, 0x28, 0x28, 0x5b, 0x30, 0x2d, 0x39, 0x61, 0x2d, 0x66, 0x5d,
	0x7b, 0x31, 0x2c, 0x34, 0x7d, 0x3a, 0x29, 0x7b, 0x37, 0x7d, 0x5b, 0x30, 0x2d, 0x39, 0x61, 0x2d,
	0x66, 0x5d, 0x7b, 0x31, 0x2c, 0x34, 0x7d, 0x29, 0x7c, 0x28, 0x3a, 0x3a, 0x28, 0x32, 0x35, 0x5b,
	0x30, 0x2d, 0x35, 0x5d, 0x7c, 0x32, 0x5b, 0x30, 0x2d, 0x34, 0x5d, 0x5c, 0x64, 0x7c, 0x5b, 0x30,
	0x2d, 0x31, 0x5d, 0x3f, 0x5c, 0x64, 0x3f, 0x5c, 0x64, 0x29, 0x28, 0x5c, 0x2e, 0x28, 0x32, 0x35,
	0x5b, 0x30, 0x2d, 0x35, 0x5d, 0x7c, 0x32, 0x5b, 0x30, 0x2d, 0x34, 0x5d, 0x5c, 0x64, 0x7c, 0x5b,
	0x30, 0x2d, 0x31, 0x5d, 0x3f, 0x5c, 0x64, 0x3f, 0x5c, 0x64, 0x29, 0x29, 0x7b, 0x33, 0x7d, 0x29,
	0x7c, 0x28, 0x28, 0x5b, 0x30, 0x2d, 0x39, 0x61, 0x2d, 0x66, 0x5d, 0x7b, 0x31, 0x2c, 0x34, 0x7d,
	0x3a, 0x29, 0x28, 0x3a, 0x5b, 0x30, 0x2d, 0x39, 0x61, 0x2d, 0x66, 0x5d, 0x7b, 0x31, 0x2c, 0x34,
	0x7d, 0x29, 0x7b, 0x31, 0x2c, 0x34, 0x7d, 0x3a, 0x28, 0x32, 0x35, 0x5b, 0x30, 0x2d, 0x35, 0x5d,
	0x7c, 0x32, 0x5b, 0x30, 0x2d, 0x34, 0x5d, 0x5c, 0x64, 0x7c, 0x5b, 0x30, 0x2d, 0x31, 0x5d, 0x3f,
	0x5c, 0x64, 0x3f, 0x5c, 0x64, 0x29, 0x28, 0x5c, 0x2e, 0x28, 0x32, 0x35, 0x5b, 0x30, 0x2d, 0x35,
	0x5d, 0x7c, 0x32, 0x5b, 0x30, 0x2d, 0x34, 0x5d, 0x5c, 0x64, 0x7c, 0x5b, 0x30, 0x2d, 0x31, 0x5d,
	0x3f, 0x5c, 0x64, 0x3f, 0x5c, 0x64, 0x29, 0x29, 0x7b, 0x33, 0x7d, 0x29, 0x7c, 0x28, 0x28, 0x5b,
	0x30, 0x2d, 0x39, 0x61, 0x2d, 0x66, 0x5d, 0x7b, 0x31, 0x2c, 0x34, 0x7d, 0x3a, 0x29, 0x7b, 0x31,
	0x2c, 0x32, 0x7d, 0x28, 0x3a, 0x5b, 0x30, 0x2d, 0x39, 0x61, 0x2d, 0x66, 0x5d, 0x7b, 0x31, 0x2c,
	0x34, 0x7d, 0x29, 0x7b, 0x31, 0x2c, 0x33, 0x7d, 0x3a, 0x28, 0x32, 0x35, 0x5b, 0x30, 0x2d, 0x35,
	0x5d, 0x7c, 0x32, 0x5b, 0x30, 0x2d, 0x34, 0x5d, 0x5c, 0x64, 0x7c, 0x5b, 0x30, 0x2d, 0x31, 0x5d,
	0x3f, 0x5c, 0x64, 0x3f, 0x5c, 0x64, 0x29, 0x28, 0x5c, 0x2e, 0x28, 0x32, 0x35, 0x5b, 0x30, 0x2d,
	0x35, 0x5d, 0x7c, 0x32, 0x5b, 0x30, 0x2d, 0x34, 0x5d, 0x5c, 0x64, 0x7c, 0x5b, 0x30, 0x2d, 0x31,
	0x5d, 0x3f, 0x5c, 0x64, 0x3f, 0x5c, 0x64, 0x29, 0x29, 0x7b, 0x33, 0x7d, 0x29, 0x7c, 0x28, 0x28,
	0x5b, 0x30, 0x2d, 0x39, 0x61, 0x2d, 0x66, 0x5d, 0x7b, 0x31, 0x2c, 0x34, 0x7d, 0x3a, 0x29, 0x7b,
	0x31, 0x2c, 0x33, 0x7d, 0x28, 0x3a, 0x5b, 0x30, 0x2d, 0x39, 0x61, 0x2d, 0x66, 0x5d, 0x7b, 0x31,
	0x2c, 0x34, 0x7d, 0x29, 0x7b, 0x31, 0x2c, 0x32, 0x7d, 0x3a, 0x28, 0x32, 0x35, 0x5b, 0x30, 0x2d,
	0x35, 0x5d, 0x7c, 0x32, 0x5b, 0x30, 0x2d, 0x34, 0x5d, 0x5c, 0x64, 0x7c, 0x5b, 0x30, 0x2d, 0x31,
	0x5d, 0x3f, 0x5c, 0x64, 0x3f, 0x5c, 0x64, 0x29, 0x28, 0x5c, 0x2e, 0x28, 0x32, 0x35, 0x5b, 0x30,
	0x2d, 0x35, 0x5d, 0x7c, 0x32, 0x5b, 0x30, 0x2d, 0x34, 0x5d, 0x5c, 0x64, 0x7c, 0x5b, 0x30, 0x2d,
	0x31, 0x5d, 0x3f, 0x5c, 0x64, 0x3f, 0x5c, 0x64, 0x29, 0x29, 0x7b, 0x33, 0x7d, 0x29, 0x7c, 0x28,
	0x28, 0x5b, 0x30, 0x2d, 0x39, 0x61, 0x2d, 0x66, 0x5d, 0x7b, 0x31, 0x2c, 0x34, 0x7d, 0x3a, 0x29,
	0x7b, 0x31, 0x2c, 0x34, 0x7d, 0x28, 0x3a, 0x5b, 0x30, 0x2d, 0x39, 0x61, 0x2d, 0x66, 0x5d, 0x7b,
	0x31, 0x2c, 0x34, 0x7d, 0x29, 0x3a, 0x28, 0x32, 0x35, 0x5b, 0x30, 0x2d, 0x35, 0x5d, 0x7c, 0x32,
	0x5b, 0x30, 0x2d, 0x34, 0x5d, 0x5c, 0x64, 0x7c, 0x5b, 0x30, 0x2d, 0x31, 0x5d, 0x3f, 0x5c, 0x64,
	0x3f, 0x5c, 0x64, 0x29, 0x28, 0x5c, 0x2e, 0x28, 0x32, 0x35, 0x5b, 0x30, 0x2d, 0x35, 0x5d, 0x7c,
	0x32, 0x5b, 0x30, 0x2d, 0x34, 0x5d, 0x5c, 0x64, 0x7c, 0x5b, 0x30, 0x2d, 0x31, 0x5d, 0x3f, 0x5c,
	0x64, 0x3f, 0x5c, 0x64, 0x29, 0x29, 0x7b, 0x33, 0x7d, 0x29, 0x7c, 0x28, 0x28, 0x5b, 0x30, 0x2d,
	0x39, 0x61, 0x2d, 0x66, 0x5d, 0x7b, 0x31, 0x2c, 0x34, 0x7d, 0x3a, 0x29, 0x7b, 0x31, 0x2c, 0x35,
	0x7d, 0x3a, 0x28, 0x32, 0x35, 0x5b, 0x30, 0x2d, 0x35, 0x5d, 0x7c, 0x32, 0x5b, 0x30, 0x2d, 0x34,
	0x5d, 0x5c, 0x64, 0x7c, 0x5b, 0x30, 0x2d, 0x31, 0x5d, 0x3f, 0x5c, 0x64, 0x3f, 0x5c, 0x64, 0x29,
	0x28, 0x5c, 0x2e, 0x28, 0x32, 0x35, 0x5b, 0x30, 0x2d, 0x35, 0x5d, 0x7c, 0x32, 0x5b, 0x30, 0x2d,
	0x34, 0x5d, 0x5c, 0x64, 0x7c, 0x5b, 0x30, 0x2d, 0x31, 0x5d, 0x3f, 0x5c, 0x64, 0x3f, 0x5c, 0x64,
	0x29, 0x29, 0x7b, 0x33, 0x7d, 0x29, 0x7c, 0x28, 0x3a, 0x28, 0x3a, 0x5b, 0x30, 0x2d, 0x39, 0x61,
	0x2d, 0x66, 0x5d, 0x7b, 0x31, 0x2c, 0x34, 0x7d, 0x29, 0x7b, 0x31, 0x2c, 0x35, 0x7d, 0x3a, 0x28,
	0x32, 0x35, 0x5b, 0x30, 0x2d, 0x35, 0x5d, 0x7c, 0x32, 0x5b, 0x30, 0x2d, 0x34, 0x5d, 0x5c, 0x64,
	0x7c, 0x5b, 0x30, 0x2d, 0x31, 0x5d, 0x3f, 0x5c, 0x64, 0x3f, 0x5c, 0x64, 0x29, 0x28, 0x5c, 0x2e,
	0x28, 0x32, 0x35, 0x5b, 0x30, 0x2d, 0x35, 0x5d, 0x7c, 0x32, 0x5b, 0x30, 0x2d, 0x34, 0x5d, 0x5c,
	0x64, 0x7c, 0x5b, 0x30, 0x2d, 0x31, 0x5d, 0x3f, 0x5c, 0x64, 0x3f, 0x5c, 0x64, 0x29, 0x29, 0x7b,
	0x33, 0x7d, 0x29, 0x7c, 0x28, 0x28, 0x5b, 0x30, 0x2d, 0x39, 0x61, 0x2d, 0x66, 0x5d, 0x7b, 0x31,
	0x2c, 0x34, 0x7d, 0x3a, 0x29, 0x7b, 0x36, 0x7d, 0x28, 0x32, 0x35, 0x5b, 0x30, 0x2d, 0x35, 0x5d,
	0x7c, 0x32, 0x5b, 0x30, 0x2d, 0x34, 0x5d, 0x5c, 0x64, 0x7c, 0x5b, 0x30, 0x2d, 0x31, 0x5d, 0x3f,
	0x5c, 0x64, 0x3f, 0x5c, 0x64, 0x29, 0x28, 0x5c, 0x2e, 0x28, 0x32, 0x35, 0x5b, 0x30, 0x2d, 0x35,
	0x5d, 0x7c, 0x32, 0x5b, 0x30, 0x2d, 0x34, 0x5d, 0x5c, 0x64, 0x7c, 0x5b, 0x30, 0x2d, 0x31, 0x5d,
	0x3f, 0x5c, 0x64, 0x3f, 0x5c, 0x64, 0x29, 0x29, 0x7b, 0x33, 0x7d, 0x29, 0x52, 0x09, 0x69, 0x70,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x8d, 0x01, 0x0a, 0x0b, 0x6d, 0x61, 0x63, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x6c, 0xba,
	0x48, 0x69, 0x72, 0x67, 0x32, 0x65, 0x28, 0x5b, 0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37,
	0x38, 0x39, 0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x5d, 0x7b, 0x32, 0x7d, 0x28, 0x2d, 0x5b, 0x30,
	0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x5d,
	0x7b, 0x32, 0x7d, 0x29, 0x7b, 0x35, 0x7d, 0x29, 0x7c, 0x28, 0x5b, 0x30, 0x31, 0x32, 0x33, 0x34,
	0x35, 0x36, 0x37, 0x38, 0x39, 0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x5d, 0x7b, 0x32, 0x7d, 0x28,
	0x3a, 0x5b, 0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x41, 0x42, 0x43, 0x44,
	0x45, 0x46, 0x5d, 0x7b, 0x32, 0x7d, 0x29, 0x7b, 0x35, 0x7d, 0x29, 0x52, 0x0a, 0x6d, 0x61, 0x63,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x41, 0x0a, 0x08, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34,
	0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3f,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65, 0x42,
	0x08, 0xba, 0x48, 0x05, 0x92, 0x01, 0x02, 0x08, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42,
	0xd2, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f,
	0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x42, 0x16, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x63, 0x68, 0x71, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2d, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e,
	0x67, 0x2d, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c,
	0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x34, 0x3b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f,
	0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x76, 0x34, 0xa2, 0x02, 0x03, 0x45, 0x58, 0x58, 0xaa,
	0x02, 0x0f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x56,
	0x34, 0xca, 0x02, 0x0f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67,
	0x5c, 0x56, 0x34, 0xe2, 0x02, 0x1b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x67, 0x69,
	0x6e, 0x67, 0x5c, 0x56, 0x34, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x10, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67,
	0x3a, 0x3a, 0x56, 0x34, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_event_logging_v4_device_complex_type_proto_rawDescOnce sync.Once
	file_event_logging_v4_device_complex_type_proto_rawDescData = file_event_logging_v4_device_complex_type_proto_rawDesc
)

func file_event_logging_v4_device_complex_type_proto_rawDescGZIP() []byte {
	file_event_logging_v4_device_complex_type_proto_rawDescOnce.Do(func() {
		file_event_logging_v4_device_complex_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_event_logging_v4_device_complex_type_proto_rawDescData)
	})
	return file_event_logging_v4_device_complex_type_proto_rawDescData
}

var file_event_logging_v4_device_complex_type_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_event_logging_v4_device_complex_type_proto_goTypes = []interface{}{
	(*DeviceComplexType)(nil),         // 0: event_logging.v4.DeviceComplexType
	(*ClassificationComplexType)(nil), // 1: event_logging.v4.ClassificationComplexType
	(*LocationComplexType)(nil),       // 2: event_logging.v4.LocationComplexType
	(*DataComplexType)(nil),           // 3: event_logging.v4.DataComplexType
}
var file_event_logging_v4_device_complex_type_proto_depIdxs = []int32{
	1, // 0: event_logging.v4.DeviceComplexType.classification:type_name -> event_logging.v4.ClassificationComplexType
	2, // 1: event_logging.v4.DeviceComplexType.location:type_name -> event_logging.v4.LocationComplexType
	3, // 2: event_logging.v4.DeviceComplexType.data:type_name -> event_logging.v4.DataComplexType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_event_logging_v4_device_complex_type_proto_init() }
func file_event_logging_v4_device_complex_type_proto_init() {
	if File_event_logging_v4_device_complex_type_proto != nil {
		return
	}
	file_event_logging_v4_classification_complex_type_proto_init()
	file_event_logging_v4_data_complex_type_proto_init()
	file_event_logging_v4_location_complex_type_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_event_logging_v4_device_complex_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceComplexType); i {
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
			RawDescriptor: file_event_logging_v4_device_complex_type_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_event_logging_v4_device_complex_type_proto_goTypes,
		DependencyIndexes: file_event_logging_v4_device_complex_type_proto_depIdxs,
		MessageInfos:      file_event_logging_v4_device_complex_type_proto_msgTypes,
	}.Build()
	File_event_logging_v4_device_complex_type_proto = out.File
	file_event_logging_v4_device_complex_type_proto_rawDesc = nil
	file_event_logging_v4_device_complex_type_proto_goTypes = nil
	file_event_logging_v4_device_complex_type_proto_depIdxs = nil
}
