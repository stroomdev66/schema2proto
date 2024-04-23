// event_logging/v4/banner_complex_type.proto at 0:0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: event_logging/v4/banner_complex_type.proto

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

// Structure used to describe a banner presented to a user. Typically, it displays legal information relating to the system.
type BannerComplexType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This element can be used to supply any metadata relating to an object as long as it conforms to an allowed format/specification (defined outside this XML Schema). This can be used for adding metadata to the event after receipt.
	Meta []*AnyContentComplexType `protobuf:"bytes,1,rep,name=meta,proto3" json:"meta,omitempty"`
	// The type of the object in question and specific to the object type from the list above, e.g. a 'Resource' object may have a type such as 'image' or 'script'.
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	// An identifier for the object, e.g a document ID in a document management system. This ID is likely to be specific to the system that generated the event.
	Id string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// The name of the object, e.g. a filename.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Human readable description of what the object is.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// Any classification, protective marking or restrictions placed on the object, e.g. for commercially sensitive reports or user health records.
	Classification *ClassificationComplexType `protobuf:"bytes,6,opt,name=classification,proto3" json:"classification,omitempty"`
	// Any state information about the object, e.g. 'Archived'.
	State string `protobuf:"bytes,7,opt,name=state,proto3" json:"state,omitempty"`
	// Any groups associated with the object, e.g. group membership of a user account.
	Groups *GroupsComplexType `protobuf:"bytes,8,opt,name=groups,proto3" json:"groups,omitempty"`
	// The collection of permissions associated with the object, e.g. write access being granted to a list of named users.
	Permissions *BannerComplexType_PermissionsType `protobuf:"bytes,9,opt,name=permissions,proto3" json:"permissions,omitempty"`
	// Metadata tags that can be used for additional object tagging or categorisation. Object tagging allows for the labelling (or filtering) of objects using words that label, categorise or group similar items, using a taxonomy defined outside this schema. For example, an email could be tagged with tags like 'internal', 'spam', 'external', 'rich-content', etc.
	Tags *MetaDataTagsComplexType `protobuf:"bytes,10,opt,name=tags,proto3" json:"tags,omitempty"`
	// If the text of the banner is versioned then this records the version of the banner that is displayed to the user.
	Version string `protobuf:"bytes,11,opt,name=version,proto3" json:"version,omitempty"`
	// The message text displayed on the banner.
	Message string `protobuf:"bytes,12,opt,name=message,proto3" json:"message,omitempty"`
	// Any other event data that does not fit into a schema element but may be useful for the purpose of audit.
	Data []*DataComplexType `protobuf:"bytes,13,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *BannerComplexType) Reset() {
	*x = BannerComplexType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_logging_v4_banner_complex_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BannerComplexType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BannerComplexType) ProtoMessage() {}

func (x *BannerComplexType) ProtoReflect() protoreflect.Message {
	mi := &file_event_logging_v4_banner_complex_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BannerComplexType.ProtoReflect.Descriptor instead.
func (*BannerComplexType) Descriptor() ([]byte, []int) {
	return file_event_logging_v4_banner_complex_type_proto_rawDescGZIP(), []int{0}
}

func (x *BannerComplexType) GetMeta() []*AnyContentComplexType {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *BannerComplexType) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *BannerComplexType) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BannerComplexType) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BannerComplexType) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *BannerComplexType) GetClassification() *ClassificationComplexType {
	if x != nil {
		return x.Classification
	}
	return nil
}

func (x *BannerComplexType) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *BannerComplexType) GetGroups() *GroupsComplexType {
	if x != nil {
		return x.Groups
	}
	return nil
}

func (x *BannerComplexType) GetPermissions() *BannerComplexType_PermissionsType {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *BannerComplexType) GetTags() *MetaDataTagsComplexType {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *BannerComplexType) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *BannerComplexType) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *BannerComplexType) GetData() []*DataComplexType {
	if x != nil {
		return x.Data
	}
	return nil
}

// The collection of permissions associated with the object, e.g. write access being granted to a list of named users.
type BannerComplexType_PermissionsType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A permission rule associated with an object, e.g. read and write access being granted to a user.
	Permission []*BannerComplexType_PermissionsType_PermissionType `protobuf:"bytes,1,rep,name=permission,proto3" json:"permission,omitempty"`
}

func (x *BannerComplexType_PermissionsType) Reset() {
	*x = BannerComplexType_PermissionsType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_logging_v4_banner_complex_type_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BannerComplexType_PermissionsType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BannerComplexType_PermissionsType) ProtoMessage() {}

func (x *BannerComplexType_PermissionsType) ProtoReflect() protoreflect.Message {
	mi := &file_event_logging_v4_banner_complex_type_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BannerComplexType_PermissionsType.ProtoReflect.Descriptor instead.
func (*BannerComplexType_PermissionsType) Descriptor() ([]byte, []int) {
	return file_event_logging_v4_banner_complex_type_proto_rawDescGZIP(), []int{0, 0}
}

func (x *BannerComplexType_PermissionsType) GetPermission() []*BannerComplexType_PermissionsType_PermissionType {
	if x != nil {
		return x.Permission
	}
	return nil
}

// A permission rule associated with an object, e.g. read and write access being granted to a user.
type BannerComplexType_PermissionsType_PermissionType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A user that has been granted (or is prevented from having) some form of permission.
	User *UserComplexType `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// A named group of users that has been granted (or is prevented from having) some form of permission.
	Group *GroupComplexType `protobuf:"bytes,2,opt,name=group,proto3" json:"group,omitempty"`
	// The permission attributes that have been explicitly allowed.
	Allow []PermissionAttributeSimpleType `protobuf:"varint,11,rep,packed,name=allow,proto3,enum=event_logging.v4.PermissionAttributeSimpleType" json:"allow,omitempty"`
	// The permission attributes that have been explicitly denied.
	Deny []PermissionAttributeSimpleType `protobuf:"varint,12,rep,packed,name=deny,proto3,enum=event_logging.v4.PermissionAttributeSimpleType" json:"deny,omitempty"`
}

func (x *BannerComplexType_PermissionsType_PermissionType) Reset() {
	*x = BannerComplexType_PermissionsType_PermissionType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_logging_v4_banner_complex_type_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BannerComplexType_PermissionsType_PermissionType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BannerComplexType_PermissionsType_PermissionType) ProtoMessage() {}

func (x *BannerComplexType_PermissionsType_PermissionType) ProtoReflect() protoreflect.Message {
	mi := &file_event_logging_v4_banner_complex_type_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BannerComplexType_PermissionsType_PermissionType.ProtoReflect.Descriptor instead.
func (*BannerComplexType_PermissionsType_PermissionType) Descriptor() ([]byte, []int) {
	return file_event_logging_v4_banner_complex_type_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *BannerComplexType_PermissionsType_PermissionType) GetUser() *UserComplexType {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *BannerComplexType_PermissionsType_PermissionType) GetGroup() *GroupComplexType {
	if x != nil {
		return x.Group
	}
	return nil
}

func (x *BannerComplexType_PermissionsType_PermissionType) GetAllow() []PermissionAttributeSimpleType {
	if x != nil {
		return x.Allow
	}
	return nil
}

func (x *BannerComplexType_PermissionsType_PermissionType) GetDeny() []PermissionAttributeSimpleType {
	if x != nil {
		return x.Deny
	}
	return nil
}

var File_event_logging_v4_banner_complex_type_proto protoreflect.FileDescriptor

var file_event_logging_v4_banner_complex_type_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f,
	0x76, 0x34, 0x2f, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x78, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x1a, 0x2f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x34,
	0x2f, 0x61, 0x6e, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x78, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x32, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x76,
	0x34, 0x2f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69,
	0x6e, 0x67, 0x2f, 0x76, 0x34, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x78, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x34, 0x2f,
	0x6d, 0x65, 0x74, 0x61, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x74, 0x61, 0x67, 0x73, 0x5f, 0x63,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x37, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67,
	0x2f, 0x76, 0x34, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x61,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x34, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xab, 0x08, 0x0a, 0x11, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65, 0x12, 0x45, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f,
	0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x2e, 0x41, 0x6e, 0x79, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65, 0x42, 0x08,
	0xba, 0x48, 0x05, 0x92, 0x01, 0x02, 0x08, 0x00, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x53, 0x0a, 0x0e, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2b, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x34, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0e, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x12, 0x55, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f,
	0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x2e, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x43,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3d, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f,
	0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74,
	0x61, 0x54, 0x61, 0x67, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3f, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65, 0x42, 0x08, 0xba, 0x48, 0x05,
	0x92, 0x01, 0x02, 0x08, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0xc1, 0x03, 0x0a, 0x0f,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x6c, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x2e, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x42, 0x08, 0xba, 0x48, 0x05, 0x92, 0x01, 0x02, 0x08,
	0x01, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0xbf, 0x02,
	0x0a, 0x0e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x3d, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x34, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x79, 0x70,
	0x65, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12,
	0x40, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x34, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x79,
	0x70, 0x65, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x56, 0x0a, 0x05, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0e,
	0x32, 0x2f, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x34, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x42, 0x0f, 0xba, 0x48, 0x0c, 0x92, 0x01, 0x09, 0x08, 0x00, 0x22, 0x05, 0x82, 0x01, 0x02,
	0x10, 0x01, 0x52, 0x05, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x54, 0x0a, 0x04, 0x64, 0x65, 0x6e,
	0x79, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f,
	0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x53, 0x69,
	0x6d, 0x70, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x0f, 0xba, 0x48, 0x0c, 0x92, 0x01, 0x09,
	0x08, 0x00, 0x22, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x04, 0x64, 0x65, 0x6e, 0x79, 0x42,
	0xda, 0x01, 0x0a, 0x1c, 0x75, 0x6b, 0x2e, 0x67, 0x6f, 0x76, 0x2e, 0x67, 0x63, 0x68, 0x71, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34,
	0x42, 0x16, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54,
	0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x63, 0x68, 0x71, 0x2f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x2d, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2d, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x76,
	0x34, 0x3b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x76,
	0x34, 0xa2, 0x02, 0x03, 0x45, 0x58, 0x58, 0xaa, 0x02, 0x0f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c,
	0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x34, 0xca, 0x02, 0x0f, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x34, 0xe2, 0x02, 0x1b, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x34, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x10, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x56, 0x34, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_event_logging_v4_banner_complex_type_proto_rawDescOnce sync.Once
	file_event_logging_v4_banner_complex_type_proto_rawDescData = file_event_logging_v4_banner_complex_type_proto_rawDesc
)

func file_event_logging_v4_banner_complex_type_proto_rawDescGZIP() []byte {
	file_event_logging_v4_banner_complex_type_proto_rawDescOnce.Do(func() {
		file_event_logging_v4_banner_complex_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_event_logging_v4_banner_complex_type_proto_rawDescData)
	})
	return file_event_logging_v4_banner_complex_type_proto_rawDescData
}

var file_event_logging_v4_banner_complex_type_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_event_logging_v4_banner_complex_type_proto_goTypes = []interface{}{
	(*BannerComplexType)(nil),                                // 0: event_logging.v4.BannerComplexType
	(*BannerComplexType_PermissionsType)(nil),                // 1: event_logging.v4.BannerComplexType.PermissionsType
	(*BannerComplexType_PermissionsType_PermissionType)(nil), // 2: event_logging.v4.BannerComplexType.PermissionsType.PermissionType
	(*AnyContentComplexType)(nil),                            // 3: event_logging.v4.AnyContentComplexType
	(*ClassificationComplexType)(nil),                        // 4: event_logging.v4.ClassificationComplexType
	(*GroupsComplexType)(nil),                                // 5: event_logging.v4.GroupsComplexType
	(*MetaDataTagsComplexType)(nil),                          // 6: event_logging.v4.MetaDataTagsComplexType
	(*DataComplexType)(nil),                                  // 7: event_logging.v4.DataComplexType
	(*UserComplexType)(nil),                                  // 8: event_logging.v4.UserComplexType
	(*GroupComplexType)(nil),                                 // 9: event_logging.v4.GroupComplexType
	(PermissionAttributeSimpleType)(0),                       // 10: event_logging.v4.PermissionAttributeSimpleType
}
var file_event_logging_v4_banner_complex_type_proto_depIdxs = []int32{
	3,  // 0: event_logging.v4.BannerComplexType.meta:type_name -> event_logging.v4.AnyContentComplexType
	4,  // 1: event_logging.v4.BannerComplexType.classification:type_name -> event_logging.v4.ClassificationComplexType
	5,  // 2: event_logging.v4.BannerComplexType.groups:type_name -> event_logging.v4.GroupsComplexType
	1,  // 3: event_logging.v4.BannerComplexType.permissions:type_name -> event_logging.v4.BannerComplexType.PermissionsType
	6,  // 4: event_logging.v4.BannerComplexType.tags:type_name -> event_logging.v4.MetaDataTagsComplexType
	7,  // 5: event_logging.v4.BannerComplexType.data:type_name -> event_logging.v4.DataComplexType
	2,  // 6: event_logging.v4.BannerComplexType.PermissionsType.permission:type_name -> event_logging.v4.BannerComplexType.PermissionsType.PermissionType
	8,  // 7: event_logging.v4.BannerComplexType.PermissionsType.PermissionType.user:type_name -> event_logging.v4.UserComplexType
	9,  // 8: event_logging.v4.BannerComplexType.PermissionsType.PermissionType.group:type_name -> event_logging.v4.GroupComplexType
	10, // 9: event_logging.v4.BannerComplexType.PermissionsType.PermissionType.allow:type_name -> event_logging.v4.PermissionAttributeSimpleType
	10, // 10: event_logging.v4.BannerComplexType.PermissionsType.PermissionType.deny:type_name -> event_logging.v4.PermissionAttributeSimpleType
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_event_logging_v4_banner_complex_type_proto_init() }
func file_event_logging_v4_banner_complex_type_proto_init() {
	if File_event_logging_v4_banner_complex_type_proto != nil {
		return
	}
	file_event_logging_v4_any_content_complex_type_proto_init()
	file_event_logging_v4_classification_complex_type_proto_init()
	file_event_logging_v4_data_complex_type_proto_init()
	file_event_logging_v4_meta_data_tags_complex_type_proto_init()
	file_event_logging_v4_permission_attribute_simple_type_proto_init()
	file_event_logging_v4_user_complex_type_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_event_logging_v4_banner_complex_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BannerComplexType); i {
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
		file_event_logging_v4_banner_complex_type_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BannerComplexType_PermissionsType); i {
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
		file_event_logging_v4_banner_complex_type_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BannerComplexType_PermissionsType_PermissionType); i {
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
			RawDescriptor: file_event_logging_v4_banner_complex_type_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_event_logging_v4_banner_complex_type_proto_goTypes,
		DependencyIndexes: file_event_logging_v4_banner_complex_type_proto_depIdxs,
		MessageInfos:      file_event_logging_v4_banner_complex_type_proto_msgTypes,
	}.Build()
	File_event_logging_v4_banner_complex_type_proto = out.File
	file_event_logging_v4_banner_complex_type_proto_rawDesc = nil
	file_event_logging_v4_banner_complex_type_proto_goTypes = nil
	file_event_logging_v4_banner_complex_type_proto_depIdxs = nil
}
