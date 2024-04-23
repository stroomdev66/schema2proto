// event_logging/v4/query_complex_type.proto at 0:0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: event_logging/v4/query_complex_type.proto

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

type QueryComplexType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An identifier to uniquely identify the query that was executed. This may be used to link the execution of a query (i.e. Search/Query/Id) with the results that are persisted and viewed at another time (i.e. View/SearchResults/Query/Id).
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The name of the query that was executed. This may be used to link the execution of a query with the results that are persisted and viewed at another time.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// A human-readable description of what the query is searching for.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// A complex boolean tree or operators and terms that describes the query.
	Advanced *QueryComplexType_AdvancedType `protobuf:"bytes,4,opt,name=advanced,proto3" json:"advanced,omitempty"`
	// A simple representation of a query using includes and excludes terms. This is suitable for simple filtered lists, e.g. for a list of names excluding "John,Bob".
	Simple *QueryComplexType_SimpleType `protobuf:"bytes,5,opt,name=simple,proto3" json:"simple,omitempty"`
	// The raw query in the query language used by the application executing the query, e.g. SQL, xpath, etc.
	Raw string `protobuf:"bytes,6,opt,name=raw,proto3" json:"raw,omitempty"`
}

func (x *QueryComplexType) Reset() {
	*x = QueryComplexType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_logging_v4_query_complex_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryComplexType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryComplexType) ProtoMessage() {}

func (x *QueryComplexType) ProtoReflect() protoreflect.Message {
	mi := &file_event_logging_v4_query_complex_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryComplexType.ProtoReflect.Descriptor instead.
func (*QueryComplexType) Descriptor() ([]byte, []int) {
	return file_event_logging_v4_query_complex_type_proto_rawDescGZIP(), []int{0}
}

func (x *QueryComplexType) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *QueryComplexType) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *QueryComplexType) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *QueryComplexType) GetAdvanced() *QueryComplexType_AdvancedType {
	if x != nil {
		return x.Advanced
	}
	return nil
}

func (x *QueryComplexType) GetSimple() *QueryComplexType_SimpleType {
	if x != nil {
		return x.Simple
	}
	return nil
}

func (x *QueryComplexType) GetRaw() string {
	if x != nil {
		return x.Raw
	}
	return ""
}

// A complex boolean tree or operators and terms that describes the query.
type QueryComplexType_AdvancedType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChoiceWrapper []*QueryComplexType_AdvancedType_ChoiceWrapperAdvancedType `protobuf:"bytes,1,rep,name=choice_wrapper,json=choiceWrapper,proto3" json:"choice_wrapper,omitempty"`
}

func (x *QueryComplexType_AdvancedType) Reset() {
	*x = QueryComplexType_AdvancedType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_logging_v4_query_complex_type_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryComplexType_AdvancedType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryComplexType_AdvancedType) ProtoMessage() {}

func (x *QueryComplexType_AdvancedType) ProtoReflect() protoreflect.Message {
	mi := &file_event_logging_v4_query_complex_type_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryComplexType_AdvancedType.ProtoReflect.Descriptor instead.
func (*QueryComplexType_AdvancedType) Descriptor() ([]byte, []int) {
	return file_event_logging_v4_query_complex_type_proto_rawDescGZIP(), []int{0, 0}
}

func (x *QueryComplexType_AdvancedType) GetChoiceWrapper() []*QueryComplexType_AdvancedType_ChoiceWrapperAdvancedType {
	if x != nil {
		return x.ChoiceWrapper
	}
	return nil
}

// A simple representation of a query using includes and excludes terms. This is suitable for simple filtered lists, e.g. for a list of names excluding "John,Bob".
type QueryComplexType_SimpleType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Values to include in the query results, typically delimited by a comma.
	Include string `protobuf:"bytes,1,opt,name=include,proto3" json:"include,omitempty"`
	// Values to exclude in the query results, typically delimited by a comma.
	Exclude string `protobuf:"bytes,2,opt,name=exclude,proto3" json:"exclude,omitempty"`
}

func (x *QueryComplexType_SimpleType) Reset() {
	*x = QueryComplexType_SimpleType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_logging_v4_query_complex_type_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryComplexType_SimpleType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryComplexType_SimpleType) ProtoMessage() {}

func (x *QueryComplexType_SimpleType) ProtoReflect() protoreflect.Message {
	mi := &file_event_logging_v4_query_complex_type_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryComplexType_SimpleType.ProtoReflect.Descriptor instead.
func (*QueryComplexType_SimpleType) Descriptor() ([]byte, []int) {
	return file_event_logging_v4_query_complex_type_proto_rawDescGZIP(), []int{0, 1}
}

func (x *QueryComplexType_SimpleType) GetInclude() string {
	if x != nil {
		return x.Include
	}
	return ""
}

func (x *QueryComplexType_SimpleType) GetExclude() string {
	if x != nil {
		return x.Exclude
	}
	return ""
}

type QueryComplexType_AdvancedType_ChoiceWrapperAdvancedType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Describes a field to search and what to search for using name, condition and value, e.g. Title Contains 'Fox' or Title Equals 'The Quick Brown Fox'.
	Term *TermComplexType `protobuf:"bytes,1,opt,name=term,proto3" json:"term,omitempty"`
	// All of the enclosed items are to be treated as being AND'd together.
	And *AndComplexType `protobuf:"bytes,2,opt,name=and,proto3" json:"and,omitempty"`
	// All of the enclosed items are to be treated as being OR'd together.
	Or *OrComplexType `protobuf:"bytes,3,opt,name=or,proto3" json:"or,omitempty"`
	// The enclosed structure operators and terms are negated.
	Not *NotComplexType `protobuf:"bytes,4,opt,name=not,proto3" json:"not,omitempty"`
}

func (x *QueryComplexType_AdvancedType_ChoiceWrapperAdvancedType) Reset() {
	*x = QueryComplexType_AdvancedType_ChoiceWrapperAdvancedType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_logging_v4_query_complex_type_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryComplexType_AdvancedType_ChoiceWrapperAdvancedType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryComplexType_AdvancedType_ChoiceWrapperAdvancedType) ProtoMessage() {}

func (x *QueryComplexType_AdvancedType_ChoiceWrapperAdvancedType) ProtoReflect() protoreflect.Message {
	mi := &file_event_logging_v4_query_complex_type_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryComplexType_AdvancedType_ChoiceWrapperAdvancedType.ProtoReflect.Descriptor instead.
func (*QueryComplexType_AdvancedType_ChoiceWrapperAdvancedType) Descriptor() ([]byte, []int) {
	return file_event_logging_v4_query_complex_type_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *QueryComplexType_AdvancedType_ChoiceWrapperAdvancedType) GetTerm() *TermComplexType {
	if x != nil {
		return x.Term
	}
	return nil
}

func (x *QueryComplexType_AdvancedType_ChoiceWrapperAdvancedType) GetAnd() *AndComplexType {
	if x != nil {
		return x.And
	}
	return nil
}

func (x *QueryComplexType_AdvancedType_ChoiceWrapperAdvancedType) GetOr() *OrComplexType {
	if x != nil {
		return x.Or
	}
	return nil
}

func (x *QueryComplexType_AdvancedType_ChoiceWrapperAdvancedType) GetNot() *NotComplexType {
	if x != nil {
		return x.Not
	}
	return nil
}

var File_event_logging_v4_query_complex_type_proto protoreflect.FileDescriptor

var file_event_logging_v4_query_complex_type_proto_rawDesc = []byte{
	0x0a, 0x29, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f,
	0x76, 0x34, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x1a, 0x29, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x34, 0x2f,
	0x6c, 0x6f, 0x67, 0x69, 0x63, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f,
	0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x34, 0x2f, 0x74, 0x65, 0x72, 0x6d, 0x5f,
	0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xbb, 0x05, 0x0a, 0x10, 0x51, 0x75, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4b, 0x0a, 0x08, 0x61, 0x64,
	0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x2e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65,
	0x2e, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x61,
	0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x12, 0x45, 0x0a, 0x06, 0x73, 0x69, 0x6d, 0x70, 0x6c,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f,
	0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x53, 0x69, 0x6d, 0x70,
	0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x72, 0x61, 0x77, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x61, 0x77,
	0x1a, 0xf8, 0x02, 0x0a, 0x0c, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x7a, 0x0a, 0x0e, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x77, 0x72, 0x61, 0x70,
	0x70, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x49, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x2e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x41, 0x64,
	0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x43, 0x68, 0x6f, 0x69, 0x63,
	0x65, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64,
	0x54, 0x79, 0x70, 0x65, 0x42, 0x08, 0xba, 0x48, 0x05, 0x92, 0x01, 0x02, 0x08, 0x00, 0x52, 0x0d,
	0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x1a, 0xeb, 0x01,
	0x0a, 0x19, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x41,
	0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x35, 0x0a, 0x04, 0x74,
	0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x2e, 0x54, 0x65, 0x72,
	0x6d, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x65,
	0x72, 0x6d, 0x12, 0x32, 0x0a, 0x03, 0x61, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x34, 0x2e, 0x41, 0x6e, 0x64, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x03, 0x61, 0x6e, 0x64, 0x12, 0x2f, 0x0a, 0x02, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x34, 0x2e, 0x4f, 0x72, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x02, 0x6f, 0x72, 0x12, 0x32, 0x0a, 0x03, 0x6e, 0x6f, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67,
	0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x2e, 0x4e, 0x6f, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x78, 0x54, 0x79, 0x70, 0x65, 0x52, 0x03, 0x6e, 0x6f, 0x74, 0x1a, 0x40, 0x0a, 0x0a, 0x53,
	0x69, 0x6d, 0x70, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x6e, 0x63,
	0x6c, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x6e, 0x63, 0x6c,
	0x75, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x42, 0xd1, 0x01,
	0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x42, 0x15, 0x51, 0x75, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x63, 0x68, 0x71,
	0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2d, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2d, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x67,
	0x69, 0x6e, 0x67, 0x2f, 0x76, 0x34, 0x3b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67,
	0x67, 0x69, 0x6e, 0x67, 0x76, 0x34, 0xa2, 0x02, 0x03, 0x45, 0x58, 0x58, 0xaa, 0x02, 0x0f, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x34, 0xca, 0x02,
	0x0f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x34,
	0xe2, 0x02, 0x1b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x5c,
	0x56, 0x34, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x10, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x56,
	0x34, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_event_logging_v4_query_complex_type_proto_rawDescOnce sync.Once
	file_event_logging_v4_query_complex_type_proto_rawDescData = file_event_logging_v4_query_complex_type_proto_rawDesc
)

func file_event_logging_v4_query_complex_type_proto_rawDescGZIP() []byte {
	file_event_logging_v4_query_complex_type_proto_rawDescOnce.Do(func() {
		file_event_logging_v4_query_complex_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_event_logging_v4_query_complex_type_proto_rawDescData)
	})
	return file_event_logging_v4_query_complex_type_proto_rawDescData
}

var file_event_logging_v4_query_complex_type_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_event_logging_v4_query_complex_type_proto_goTypes = []interface{}{
	(*QueryComplexType)(nil),                                        // 0: event_logging.v4.QueryComplexType
	(*QueryComplexType_AdvancedType)(nil),                           // 1: event_logging.v4.QueryComplexType.AdvancedType
	(*QueryComplexType_SimpleType)(nil),                             // 2: event_logging.v4.QueryComplexType.SimpleType
	(*QueryComplexType_AdvancedType_ChoiceWrapperAdvancedType)(nil), // 3: event_logging.v4.QueryComplexType.AdvancedType.ChoiceWrapperAdvancedType
	(*TermComplexType)(nil),                                         // 4: event_logging.v4.TermComplexType
	(*AndComplexType)(nil),                                          // 5: event_logging.v4.AndComplexType
	(*OrComplexType)(nil),                                           // 6: event_logging.v4.OrComplexType
	(*NotComplexType)(nil),                                          // 7: event_logging.v4.NotComplexType
}
var file_event_logging_v4_query_complex_type_proto_depIdxs = []int32{
	1, // 0: event_logging.v4.QueryComplexType.advanced:type_name -> event_logging.v4.QueryComplexType.AdvancedType
	2, // 1: event_logging.v4.QueryComplexType.simple:type_name -> event_logging.v4.QueryComplexType.SimpleType
	3, // 2: event_logging.v4.QueryComplexType.AdvancedType.choice_wrapper:type_name -> event_logging.v4.QueryComplexType.AdvancedType.ChoiceWrapperAdvancedType
	4, // 3: event_logging.v4.QueryComplexType.AdvancedType.ChoiceWrapperAdvancedType.term:type_name -> event_logging.v4.TermComplexType
	5, // 4: event_logging.v4.QueryComplexType.AdvancedType.ChoiceWrapperAdvancedType.and:type_name -> event_logging.v4.AndComplexType
	6, // 5: event_logging.v4.QueryComplexType.AdvancedType.ChoiceWrapperAdvancedType.or:type_name -> event_logging.v4.OrComplexType
	7, // 6: event_logging.v4.QueryComplexType.AdvancedType.ChoiceWrapperAdvancedType.not:type_name -> event_logging.v4.NotComplexType
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_event_logging_v4_query_complex_type_proto_init() }
func file_event_logging_v4_query_complex_type_proto_init() {
	if File_event_logging_v4_query_complex_type_proto != nil {
		return
	}
	file_event_logging_v4_logic_complex_type_proto_init()
	file_event_logging_v4_term_complex_type_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_event_logging_v4_query_complex_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryComplexType); i {
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
		file_event_logging_v4_query_complex_type_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryComplexType_AdvancedType); i {
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
		file_event_logging_v4_query_complex_type_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryComplexType_SimpleType); i {
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
		file_event_logging_v4_query_complex_type_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryComplexType_AdvancedType_ChoiceWrapperAdvancedType); i {
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
			RawDescriptor: file_event_logging_v4_query_complex_type_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_event_logging_v4_query_complex_type_proto_goTypes,
		DependencyIndexes: file_event_logging_v4_query_complex_type_proto_depIdxs,
		MessageInfos:      file_event_logging_v4_query_complex_type_proto_msgTypes,
	}.Build()
	File_event_logging_v4_query_complex_type_proto = out.File
	file_event_logging_v4_query_complex_type_proto_rawDesc = nil
	file_event_logging_v4_query_complex_type_proto_goTypes = nil
	file_event_logging_v4_query_complex_type_proto_depIdxs = nil
}
