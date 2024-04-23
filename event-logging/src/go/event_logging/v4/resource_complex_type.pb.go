// event_logging/v4/resource_complex_type.proto at 0:0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: event_logging/v4/resource_complex_type.proto

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

// Describes a resource within a website or web application, such as an HTML file, image file or script, along with the details of that resource such as size or response codes. It can represent both successful and failed access to the resource object.
type ResourceComplexType struct {
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
	Permissions *ResourceComplexType_PermissionsType `protobuf:"bytes,9,opt,name=permissions,proto3" json:"permissions,omitempty"`
	// Metadata tags that can be used for additional object tagging or categorisation. Object tagging allows for the labelling (or filtering) of objects using words that label, categorise or group similar items, using a taxonomy defined outside this schema. For example, an email could be tagged with tags like 'internal', 'spam', 'external', 'rich-content', etc.
	Tags *MetaDataTagsComplexType `protobuf:"bytes,10,opt,name=tags,proto3" json:"tags,omitempty"`
	// The title of the resource or of the object the resource presents.
	Title string `protobuf:"bytes,11,opt,name=title,proto3" json:"title,omitempty"`
	// The URL of the resource the event relates to
	Url string `protobuf:"bytes,12,opt,name=url,proto3" json:"url,omitempty"`
	// The URL of the resource that referred to the URL of this event
	Referrer string `protobuf:"bytes,13,opt,name=referrer,proto3" json:"referrer,omitempty"`
	// The session identifier or token used to identify a session or series of related message exchanges.
	SessionId string `protobuf:"bytes,14,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	// The HTTP method, e.g. GET, POST, DELETE, PUT etc
	HttpMethod string `protobuf:"bytes,15,opt,name=http_method,json=httpMethod,proto3" json:"http_method,omitempty"`
	// The HTTP version, e.g. 1.1
	HttpVersion string `protobuf:"bytes,16,opt,name=http_version,json=httpVersion,proto3" json:"http_version,omitempty"`
	// This is a string provided by the initiating software agent used to identify itself, its application type, operating system, software vendor or software version. This string typically appears as a field in a request message with a field header name of 'User-Agent'.
	UserAgent string `protobuf:"bytes,17,opt,name=user_agent,json=userAgent,proto3" json:"user_agent,omitempty"`
	// The size in bytes received, including the request and HTTP headers.
	InboundSize uint32 `protobuf:"varint,18,opt,name=inbound_size,json=inboundSize,proto3" json:"inbound_size,omitempty"`
	// The size in bytes of the incoming data, EXCLUDING HTTP headers.
	InboundContentSize uint32 `protobuf:"varint,19,opt,name=inbound_content_size,json=inboundContentSize,proto3" json:"inbound_content_size,omitempty"`
	// The HTTP request header.
	InboundHeader string `protobuf:"bytes,20,opt,name=inbound_header,json=inboundHeader,proto3" json:"inbound_header,omitempty"`
	// The size in bytes of the outgoing data, including HTTP headers.
	OutboundSize uint32 `protobuf:"varint,21,opt,name=outbound_size,json=outboundSize,proto3" json:"outbound_size,omitempty"`
	// The size in bytes of the outgoing data, EXCLUDING HTTP headers.
	OutboundContentSize uint32 `protobuf:"varint,22,opt,name=outbound_content_size,json=outboundContentSize,proto3" json:"outbound_content_size,omitempty"`
	// The HTTP response header.
	OutboundHeader string `protobuf:"bytes,23,opt,name=outbound_header,json=outboundHeader,proto3" json:"outbound_header,omitempty"`
	// The number of microseconds the server took to handle the request.
	RequestTime uint32 `protobuf:"varint,24,opt,name=request_time,json=requestTime,proto3" json:"request_time,omitempty"`
	// The connection status of the client connection.
	ConnectionStatus string `protobuf:"bytes,25,opt,name=connection_status,json=connectionStatus,proto3" json:"connection_status,omitempty"`
	// The status code of the original request.
	InitialResponseCode string `protobuf:"bytes,26,opt,name=initial_response_code,json=initialResponseCode,proto3" json:"initial_response_code,omitempty"`
	// The final status code of the request, after any internal redirections may have taken place.
	ResponseCode string `protobuf:"bytes,27,opt,name=response_code,json=responseCode,proto3" json:"response_code,omitempty"`
	// The Internet Media Type identifying the file format of the resource provided (format of request or response body). This string typically appears in the 'Content-Type' field of a Request or Response Header.
	MimeType string `protobuf:"bytes,28,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	// The category of a web page or resource where a categorisation engine is used, e.g. News, Search Engine, Social Media, etc.
	Category string `protobuf:"bytes,29,opt,name=category,proto3" json:"category,omitempty"`
	// Any other event data that does not fit into a schema element but may be useful for the purpose of audit.
	Data []*DataComplexType `protobuf:"bytes,30,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ResourceComplexType) Reset() {
	*x = ResourceComplexType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_logging_v4_resource_complex_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceComplexType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceComplexType) ProtoMessage() {}

func (x *ResourceComplexType) ProtoReflect() protoreflect.Message {
	mi := &file_event_logging_v4_resource_complex_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceComplexType.ProtoReflect.Descriptor instead.
func (*ResourceComplexType) Descriptor() ([]byte, []int) {
	return file_event_logging_v4_resource_complex_type_proto_rawDescGZIP(), []int{0}
}

func (x *ResourceComplexType) GetMeta() []*AnyContentComplexType {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *ResourceComplexType) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ResourceComplexType) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ResourceComplexType) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResourceComplexType) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ResourceComplexType) GetClassification() *ClassificationComplexType {
	if x != nil {
		return x.Classification
	}
	return nil
}

func (x *ResourceComplexType) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *ResourceComplexType) GetGroups() *GroupsComplexType {
	if x != nil {
		return x.Groups
	}
	return nil
}

func (x *ResourceComplexType) GetPermissions() *ResourceComplexType_PermissionsType {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *ResourceComplexType) GetTags() *MetaDataTagsComplexType {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *ResourceComplexType) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ResourceComplexType) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ResourceComplexType) GetReferrer() string {
	if x != nil {
		return x.Referrer
	}
	return ""
}

func (x *ResourceComplexType) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *ResourceComplexType) GetHttpMethod() string {
	if x != nil {
		return x.HttpMethod
	}
	return ""
}

func (x *ResourceComplexType) GetHttpVersion() string {
	if x != nil {
		return x.HttpVersion
	}
	return ""
}

func (x *ResourceComplexType) GetUserAgent() string {
	if x != nil {
		return x.UserAgent
	}
	return ""
}

func (x *ResourceComplexType) GetInboundSize() uint32 {
	if x != nil {
		return x.InboundSize
	}
	return 0
}

func (x *ResourceComplexType) GetInboundContentSize() uint32 {
	if x != nil {
		return x.InboundContentSize
	}
	return 0
}

func (x *ResourceComplexType) GetInboundHeader() string {
	if x != nil {
		return x.InboundHeader
	}
	return ""
}

func (x *ResourceComplexType) GetOutboundSize() uint32 {
	if x != nil {
		return x.OutboundSize
	}
	return 0
}

func (x *ResourceComplexType) GetOutboundContentSize() uint32 {
	if x != nil {
		return x.OutboundContentSize
	}
	return 0
}

func (x *ResourceComplexType) GetOutboundHeader() string {
	if x != nil {
		return x.OutboundHeader
	}
	return ""
}

func (x *ResourceComplexType) GetRequestTime() uint32 {
	if x != nil {
		return x.RequestTime
	}
	return 0
}

func (x *ResourceComplexType) GetConnectionStatus() string {
	if x != nil {
		return x.ConnectionStatus
	}
	return ""
}

func (x *ResourceComplexType) GetInitialResponseCode() string {
	if x != nil {
		return x.InitialResponseCode
	}
	return ""
}

func (x *ResourceComplexType) GetResponseCode() string {
	if x != nil {
		return x.ResponseCode
	}
	return ""
}

func (x *ResourceComplexType) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

func (x *ResourceComplexType) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *ResourceComplexType) GetData() []*DataComplexType {
	if x != nil {
		return x.Data
	}
	return nil
}

// The collection of permissions associated with the object, e.g. write access being granted to a list of named users.
type ResourceComplexType_PermissionsType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A permission rule associated with an object, e.g. read and write access being granted to a user.
	Permission []*ResourceComplexType_PermissionsType_PermissionType `protobuf:"bytes,1,rep,name=permission,proto3" json:"permission,omitempty"`
}

func (x *ResourceComplexType_PermissionsType) Reset() {
	*x = ResourceComplexType_PermissionsType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_logging_v4_resource_complex_type_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceComplexType_PermissionsType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceComplexType_PermissionsType) ProtoMessage() {}

func (x *ResourceComplexType_PermissionsType) ProtoReflect() protoreflect.Message {
	mi := &file_event_logging_v4_resource_complex_type_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceComplexType_PermissionsType.ProtoReflect.Descriptor instead.
func (*ResourceComplexType_PermissionsType) Descriptor() ([]byte, []int) {
	return file_event_logging_v4_resource_complex_type_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ResourceComplexType_PermissionsType) GetPermission() []*ResourceComplexType_PermissionsType_PermissionType {
	if x != nil {
		return x.Permission
	}
	return nil
}

// A permission rule associated with an object, e.g. read and write access being granted to a user.
type ResourceComplexType_PermissionsType_PermissionType struct {
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

func (x *ResourceComplexType_PermissionsType_PermissionType) Reset() {
	*x = ResourceComplexType_PermissionsType_PermissionType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_logging_v4_resource_complex_type_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceComplexType_PermissionsType_PermissionType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceComplexType_PermissionsType_PermissionType) ProtoMessage() {}

func (x *ResourceComplexType_PermissionsType_PermissionType) ProtoReflect() protoreflect.Message {
	mi := &file_event_logging_v4_resource_complex_type_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceComplexType_PermissionsType_PermissionType.ProtoReflect.Descriptor instead.
func (*ResourceComplexType_PermissionsType_PermissionType) Descriptor() ([]byte, []int) {
	return file_event_logging_v4_resource_complex_type_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *ResourceComplexType_PermissionsType_PermissionType) GetUser() *UserComplexType {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *ResourceComplexType_PermissionsType_PermissionType) GetGroup() *GroupComplexType {
	if x != nil {
		return x.Group
	}
	return nil
}

func (x *ResourceComplexType_PermissionsType_PermissionType) GetAllow() []PermissionAttributeSimpleType {
	if x != nil {
		return x.Allow
	}
	return nil
}

func (x *ResourceComplexType_PermissionsType_PermissionType) GetDeny() []PermissionAttributeSimpleType {
	if x != nil {
		return x.Deny
	}
	return nil
}

var File_event_logging_v4_resource_complex_type_proto protoreflect.FileDescriptor

var file_event_logging_v4_resource_complex_type_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f,
	0x76, 0x34, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x78, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34,
	0x1a, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f,
	0x76, 0x34, 0x2f, 0x61, 0x6e, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x63,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x32, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67,
	0x2f, 0x76, 0x34, 0x2f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67,
	0x67, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x34, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x63, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x78, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x32, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x76,
	0x34, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x74, 0x61, 0x67, 0x73,
	0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x37, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69,
	0x6e, 0x67, 0x2f, 0x76, 0x34, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x73, 0x69, 0x6d, 0x70, 0x6c,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x34, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xd0, 0x0d, 0x0a, 0x13, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65, 0x12, 0x45, 0x0a, 0x04, 0x6d,
	0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x2e, 0x41, 0x6e, 0x79,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x79,
	0x70, 0x65, 0x42, 0x08, 0xba, 0x48, 0x05, 0x92, 0x01, 0x02, 0x08, 0x00, 0x52, 0x04, 0x6d, 0x65,
	0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x53, 0x0a, 0x0e,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67,
	0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x0e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f,
	0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x73, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x12, 0x57, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x2e, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65,
	0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3d, 0x0a,
	0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x2e, 0x4d,
	0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x54, 0x61, 0x67, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x78, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x72,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x72,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x68, 0x74, 0x74, 0x70, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x12, 0x21, 0x0a, 0x0c, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x68, 0x74, 0x74, 0x70, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x0c, 0x69, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xba, 0x48, 0x04, 0x2a, 0x02, 0x28,
	0x00, 0x52, 0x0b, 0x69, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x39,
	0x0a, 0x14, 0x69, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xba, 0x48,
	0x04, 0x2a, 0x02, 0x28, 0x00, 0x52, 0x12, 0x69, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x62,
	0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x69, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x12, 0x2c, 0x0a, 0x0d, 0x6f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xba, 0x48, 0x04, 0x2a, 0x02, 0x28, 0x00,
	0x52, 0x0c, 0x6f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x3b,
	0x0a, 0x15, 0x6f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xba,
	0x48, 0x04, 0x2a, 0x02, 0x28, 0x00, 0x52, 0x13, 0x6f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x6f,
	0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x17,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x18, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xba, 0x48, 0x04, 0x2a,
	0x02, 0x28, 0x00, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x2b, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x32, 0x0a,
	0x15, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x69, 0x6d, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18,
	0x1d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12,
	0x3f, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x1e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65,
	0x42, 0x08, 0xba, 0x48, 0x05, 0x92, 0x01, 0x02, 0x08, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x1a, 0xc3, 0x03, 0x0a, 0x0f, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x6e, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x44, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65, 0x2e,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x54, 0x79, 0x70, 0x65, 0x2e,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x42, 0x08,
	0xba, 0x48, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x1a, 0xbf, 0x02, 0x0a, 0x0e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3d, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f,
	0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x40, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f,
	0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01,
	0x01, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x56, 0x0a, 0x05, 0x61, 0x6c, 0x6c, 0x6f,
	0x77, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f,
	0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x53, 0x69,
	0x6d, 0x70, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x0f, 0xba, 0x48, 0x0c, 0x92, 0x01, 0x09,
	0x08, 0x00, 0x22, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x05, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x12, 0x54, 0x0a, 0x04, 0x64, 0x65, 0x6e, 0x79, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x2f,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x34, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42,
	0x0f, 0xba, 0x48, 0x0c, 0x92, 0x01, 0x09, 0x08, 0x00, 0x22, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01,
	0x52, 0x04, 0x64, 0x65, 0x6e, 0x79, 0x42, 0xe1, 0x01, 0x0a, 0x21, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x2d, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x42, 0x18, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x79, 0x70,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x63, 0x68, 0x71, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2d,
	0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2d, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x34, 0x3b,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x76, 0x34, 0xa2,
	0x02, 0x03, 0x45, 0x58, 0x58, 0xaa, 0x02, 0x0f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67,
	0x67, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x34, 0xca, 0x02, 0x0f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c,
	0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x34, 0xe2, 0x02, 0x1b, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x34, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x10, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c,
	0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x56, 0x34, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_event_logging_v4_resource_complex_type_proto_rawDescOnce sync.Once
	file_event_logging_v4_resource_complex_type_proto_rawDescData = file_event_logging_v4_resource_complex_type_proto_rawDesc
)

func file_event_logging_v4_resource_complex_type_proto_rawDescGZIP() []byte {
	file_event_logging_v4_resource_complex_type_proto_rawDescOnce.Do(func() {
		file_event_logging_v4_resource_complex_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_event_logging_v4_resource_complex_type_proto_rawDescData)
	})
	return file_event_logging_v4_resource_complex_type_proto_rawDescData
}

var file_event_logging_v4_resource_complex_type_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_event_logging_v4_resource_complex_type_proto_goTypes = []interface{}{
	(*ResourceComplexType)(nil),                                // 0: event_logging.v4.ResourceComplexType
	(*ResourceComplexType_PermissionsType)(nil),                // 1: event_logging.v4.ResourceComplexType.PermissionsType
	(*ResourceComplexType_PermissionsType_PermissionType)(nil), // 2: event_logging.v4.ResourceComplexType.PermissionsType.PermissionType
	(*AnyContentComplexType)(nil),                              // 3: event_logging.v4.AnyContentComplexType
	(*ClassificationComplexType)(nil),                          // 4: event_logging.v4.ClassificationComplexType
	(*GroupsComplexType)(nil),                                  // 5: event_logging.v4.GroupsComplexType
	(*MetaDataTagsComplexType)(nil),                            // 6: event_logging.v4.MetaDataTagsComplexType
	(*DataComplexType)(nil),                                    // 7: event_logging.v4.DataComplexType
	(*UserComplexType)(nil),                                    // 8: event_logging.v4.UserComplexType
	(*GroupComplexType)(nil),                                   // 9: event_logging.v4.GroupComplexType
	(PermissionAttributeSimpleType)(0),                         // 10: event_logging.v4.PermissionAttributeSimpleType
}
var file_event_logging_v4_resource_complex_type_proto_depIdxs = []int32{
	3,  // 0: event_logging.v4.ResourceComplexType.meta:type_name -> event_logging.v4.AnyContentComplexType
	4,  // 1: event_logging.v4.ResourceComplexType.classification:type_name -> event_logging.v4.ClassificationComplexType
	5,  // 2: event_logging.v4.ResourceComplexType.groups:type_name -> event_logging.v4.GroupsComplexType
	1,  // 3: event_logging.v4.ResourceComplexType.permissions:type_name -> event_logging.v4.ResourceComplexType.PermissionsType
	6,  // 4: event_logging.v4.ResourceComplexType.tags:type_name -> event_logging.v4.MetaDataTagsComplexType
	7,  // 5: event_logging.v4.ResourceComplexType.data:type_name -> event_logging.v4.DataComplexType
	2,  // 6: event_logging.v4.ResourceComplexType.PermissionsType.permission:type_name -> event_logging.v4.ResourceComplexType.PermissionsType.PermissionType
	8,  // 7: event_logging.v4.ResourceComplexType.PermissionsType.PermissionType.user:type_name -> event_logging.v4.UserComplexType
	9,  // 8: event_logging.v4.ResourceComplexType.PermissionsType.PermissionType.group:type_name -> event_logging.v4.GroupComplexType
	10, // 9: event_logging.v4.ResourceComplexType.PermissionsType.PermissionType.allow:type_name -> event_logging.v4.PermissionAttributeSimpleType
	10, // 10: event_logging.v4.ResourceComplexType.PermissionsType.PermissionType.deny:type_name -> event_logging.v4.PermissionAttributeSimpleType
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_event_logging_v4_resource_complex_type_proto_init() }
func file_event_logging_v4_resource_complex_type_proto_init() {
	if File_event_logging_v4_resource_complex_type_proto != nil {
		return
	}
	file_event_logging_v4_any_content_complex_type_proto_init()
	file_event_logging_v4_classification_complex_type_proto_init()
	file_event_logging_v4_data_complex_type_proto_init()
	file_event_logging_v4_meta_data_tags_complex_type_proto_init()
	file_event_logging_v4_permission_attribute_simple_type_proto_init()
	file_event_logging_v4_user_complex_type_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_event_logging_v4_resource_complex_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceComplexType); i {
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
		file_event_logging_v4_resource_complex_type_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceComplexType_PermissionsType); i {
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
		file_event_logging_v4_resource_complex_type_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceComplexType_PermissionsType_PermissionType); i {
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
			RawDescriptor: file_event_logging_v4_resource_complex_type_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_event_logging_v4_resource_complex_type_proto_goTypes,
		DependencyIndexes: file_event_logging_v4_resource_complex_type_proto_depIdxs,
		MessageInfos:      file_event_logging_v4_resource_complex_type_proto_msgTypes,
	}.Build()
	File_event_logging_v4_resource_complex_type_proto = out.File
	file_event_logging_v4_resource_complex_type_proto_rawDesc = nil
	file_event_logging_v4_resource_complex_type_proto_goTypes = nil
	file_event_logging_v4_resource_complex_type_proto_depIdxs = nil
}
