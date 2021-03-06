// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/protobuf/type.proto

package known_proto

import (
	protoreflect "github.com/golang/protobuf/v2/reflect/protoreflect"
	protoregistry "github.com/golang/protobuf/v2/reflect/protoregistry"
	protoiface "github.com/golang/protobuf/v2/runtime/protoiface"
	protoimpl "github.com/golang/protobuf/v2/runtime/protoimpl"
	sync "sync"
)

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

// The syntax in which a protocol buffer element is defined.
type Syntax int32

const (
	// Syntax `proto2`.
	Syntax_SYNTAX_PROTO2 Syntax = 0
	// Syntax `proto3`.
	Syntax_SYNTAX_PROTO3 Syntax = 1
)

// Deprecated: Use Syntax.Type.Values instead.
var Syntax_name = map[int32]string{
	0: "SYNTAX_PROTO2",
	1: "SYNTAX_PROTO3",
}

// Deprecated: Use Syntax.Type.Values instead.
var Syntax_value = map[string]int32{
	"SYNTAX_PROTO2": 0,
	"SYNTAX_PROTO3": 1,
}

func (x Syntax) String() string {
	return protoimpl.X.EnumStringOf(x.Type(), protoreflect.EnumNumber(x))
}

func (Syntax) Type() protoreflect.EnumType {
	return xxx_File_google_protobuf_type_proto_enumTypes[0]
}

func (x Syntax) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Syntax.Type instead.
func (Syntax) EnumDescriptor() ([]byte, []int) {
	return xxx_File_google_protobuf_type_proto_rawDescGZIP(), []int{0}
}

// Basic field types.
type Field_Kind int32

const (
	// Field type unknown.
	Field_TYPE_UNKNOWN Field_Kind = 0
	// Field type double.
	Field_TYPE_DOUBLE Field_Kind = 1
	// Field type float.
	Field_TYPE_FLOAT Field_Kind = 2
	// Field type int64.
	Field_TYPE_INT64 Field_Kind = 3
	// Field type uint64.
	Field_TYPE_UINT64 Field_Kind = 4
	// Field type int32.
	Field_TYPE_INT32 Field_Kind = 5
	// Field type fixed64.
	Field_TYPE_FIXED64 Field_Kind = 6
	// Field type fixed32.
	Field_TYPE_FIXED32 Field_Kind = 7
	// Field type bool.
	Field_TYPE_BOOL Field_Kind = 8
	// Field type string.
	Field_TYPE_STRING Field_Kind = 9
	// Field type group. Proto2 syntax only, and deprecated.
	Field_TYPE_GROUP Field_Kind = 10
	// Field type message.
	Field_TYPE_MESSAGE Field_Kind = 11
	// Field type bytes.
	Field_TYPE_BYTES Field_Kind = 12
	// Field type uint32.
	Field_TYPE_UINT32 Field_Kind = 13
	// Field type enum.
	Field_TYPE_ENUM Field_Kind = 14
	// Field type sfixed32.
	Field_TYPE_SFIXED32 Field_Kind = 15
	// Field type sfixed64.
	Field_TYPE_SFIXED64 Field_Kind = 16
	// Field type sint32.
	Field_TYPE_SINT32 Field_Kind = 17
	// Field type sint64.
	Field_TYPE_SINT64 Field_Kind = 18
)

// Deprecated: Use Field_Kind.Type.Values instead.
var Field_Kind_name = map[int32]string{
	0:  "TYPE_UNKNOWN",
	1:  "TYPE_DOUBLE",
	2:  "TYPE_FLOAT",
	3:  "TYPE_INT64",
	4:  "TYPE_UINT64",
	5:  "TYPE_INT32",
	6:  "TYPE_FIXED64",
	7:  "TYPE_FIXED32",
	8:  "TYPE_BOOL",
	9:  "TYPE_STRING",
	10: "TYPE_GROUP",
	11: "TYPE_MESSAGE",
	12: "TYPE_BYTES",
	13: "TYPE_UINT32",
	14: "TYPE_ENUM",
	15: "TYPE_SFIXED32",
	16: "TYPE_SFIXED64",
	17: "TYPE_SINT32",
	18: "TYPE_SINT64",
}

// Deprecated: Use Field_Kind.Type.Values instead.
var Field_Kind_value = map[string]int32{
	"TYPE_UNKNOWN":  0,
	"TYPE_DOUBLE":   1,
	"TYPE_FLOAT":    2,
	"TYPE_INT64":    3,
	"TYPE_UINT64":   4,
	"TYPE_INT32":    5,
	"TYPE_FIXED64":  6,
	"TYPE_FIXED32":  7,
	"TYPE_BOOL":     8,
	"TYPE_STRING":   9,
	"TYPE_GROUP":    10,
	"TYPE_MESSAGE":  11,
	"TYPE_BYTES":    12,
	"TYPE_UINT32":   13,
	"TYPE_ENUM":     14,
	"TYPE_SFIXED32": 15,
	"TYPE_SFIXED64": 16,
	"TYPE_SINT32":   17,
	"TYPE_SINT64":   18,
}

func (x Field_Kind) String() string {
	return protoimpl.X.EnumStringOf(x.Type(), protoreflect.EnumNumber(x))
}

func (Field_Kind) Type() protoreflect.EnumType {
	return xxx_File_google_protobuf_type_proto_enumTypes[1]
}

func (x Field_Kind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Field_Kind.Type instead.
func (Field_Kind) EnumDescriptor() ([]byte, []int) {
	return xxx_File_google_protobuf_type_proto_rawDescGZIP(), []int{1, 0}
}

// Whether a field is optional, required, or repeated.
type Field_Cardinality int32

const (
	// For fields with unknown cardinality.
	Field_CARDINALITY_UNKNOWN Field_Cardinality = 0
	// For optional fields.
	Field_CARDINALITY_OPTIONAL Field_Cardinality = 1
	// For required fields. Proto2 syntax only.
	Field_CARDINALITY_REQUIRED Field_Cardinality = 2
	// For repeated fields.
	Field_CARDINALITY_REPEATED Field_Cardinality = 3
)

// Deprecated: Use Field_Cardinality.Type.Values instead.
var Field_Cardinality_name = map[int32]string{
	0: "CARDINALITY_UNKNOWN",
	1: "CARDINALITY_OPTIONAL",
	2: "CARDINALITY_REQUIRED",
	3: "CARDINALITY_REPEATED",
}

// Deprecated: Use Field_Cardinality.Type.Values instead.
var Field_Cardinality_value = map[string]int32{
	"CARDINALITY_UNKNOWN":  0,
	"CARDINALITY_OPTIONAL": 1,
	"CARDINALITY_REQUIRED": 2,
	"CARDINALITY_REPEATED": 3,
}

func (x Field_Cardinality) String() string {
	return protoimpl.X.EnumStringOf(x.Type(), protoreflect.EnumNumber(x))
}

func (Field_Cardinality) Type() protoreflect.EnumType {
	return xxx_File_google_protobuf_type_proto_enumTypes[2]
}

func (x Field_Cardinality) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Field_Cardinality.Type instead.
func (Field_Cardinality) EnumDescriptor() ([]byte, []int) {
	return xxx_File_google_protobuf_type_proto_rawDescGZIP(), []int{1, 1}
}

// A protocol buffer message type.
type Type struct {
	// The fully qualified message name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The list of fields.
	Fields []*Field `protobuf:"bytes,2,rep,name=fields,proto3" json:"fields,omitempty"`
	// The list of types appearing in `oneof` definitions in this type.
	Oneofs []string `protobuf:"bytes,3,rep,name=oneofs,proto3" json:"oneofs,omitempty"`
	// The protocol buffer options.
	Options []*Option `protobuf:"bytes,4,rep,name=options,proto3" json:"options,omitempty"`
	// The source context.
	SourceContext *SourceContext `protobuf:"bytes,5,opt,name=source_context,json=sourceContext,proto3" json:"source_context,omitempty"`
	// The source syntax.
	Syntax               Syntax   `protobuf:"varint,6,opt,name=syntax,proto3,enum=google.protobuf.Syntax" json:"syntax,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (x *Type) Reset() {
	*x = Type{}
}

func (x *Type) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Type) ProtoMessage() {}

func (x *Type) ProtoReflect() protoreflect.Message {
	return xxx_File_google_protobuf_type_proto_messageTypes[0].MessageOf(x)
}

func (m *Type) XXX_Methods() *protoiface.Methods {
	return xxx_File_google_protobuf_type_proto_messageTypes[0].Methods()
}

// Deprecated: Use Type.ProtoReflect.Type instead.
func (*Type) Descriptor() ([]byte, []int) {
	return xxx_File_google_protobuf_type_proto_rawDescGZIP(), []int{0}
}

func (x *Type) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Type) GetFields() []*Field {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *Type) GetOneofs() []string {
	if x != nil {
		return x.Oneofs
	}
	return nil
}

func (x *Type) GetOptions() []*Option {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *Type) GetSourceContext() *SourceContext {
	if x != nil {
		return x.SourceContext
	}
	return nil
}

func (x *Type) GetSyntax() Syntax {
	if x != nil {
		return x.Syntax
	}
	return Syntax_SYNTAX_PROTO2
}

// A single field of a message type.
type Field struct {
	// The field type.
	Kind Field_Kind `protobuf:"varint,1,opt,name=kind,proto3,enum=google.protobuf.Field_Kind" json:"kind,omitempty"`
	// The field cardinality.
	Cardinality Field_Cardinality `protobuf:"varint,2,opt,name=cardinality,proto3,enum=google.protobuf.Field_Cardinality" json:"cardinality,omitempty"`
	// The field number.
	Number int32 `protobuf:"varint,3,opt,name=number,proto3" json:"number,omitempty"`
	// The field name.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// The field type URL, without the scheme, for message or enumeration
	// types. Example: `"type.googleapis.com/google.protobuf.Timestamp"`.
	TypeUrl string `protobuf:"bytes,6,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	// The index of the field type in `Type.oneofs`, for message or enumeration
	// types. The first type has index 1; zero means the type is not in the list.
	OneofIndex int32 `protobuf:"varint,7,opt,name=oneof_index,json=oneofIndex,proto3" json:"oneof_index,omitempty"`
	// Whether to use alternative packed wire representation.
	Packed bool `protobuf:"varint,8,opt,name=packed,proto3" json:"packed,omitempty"`
	// The protocol buffer options.
	Options []*Option `protobuf:"bytes,9,rep,name=options,proto3" json:"options,omitempty"`
	// The field JSON name.
	JsonName string `protobuf:"bytes,10,opt,name=json_name,json=jsonName,proto3" json:"json_name,omitempty"`
	// The string value of the default value of this field. Proto2 syntax only.
	DefaultValue         string   `protobuf:"bytes,11,opt,name=default_value,json=defaultValue,proto3" json:"default_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (x *Field) Reset() {
	*x = Field{}
}

func (x *Field) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Field) ProtoMessage() {}

func (x *Field) ProtoReflect() protoreflect.Message {
	return xxx_File_google_protobuf_type_proto_messageTypes[1].MessageOf(x)
}

func (m *Field) XXX_Methods() *protoiface.Methods {
	return xxx_File_google_protobuf_type_proto_messageTypes[1].Methods()
}

// Deprecated: Use Field.ProtoReflect.Type instead.
func (*Field) Descriptor() ([]byte, []int) {
	return xxx_File_google_protobuf_type_proto_rawDescGZIP(), []int{1}
}

func (x *Field) GetKind() Field_Kind {
	if x != nil {
		return x.Kind
	}
	return Field_TYPE_UNKNOWN
}

func (x *Field) GetCardinality() Field_Cardinality {
	if x != nil {
		return x.Cardinality
	}
	return Field_CARDINALITY_UNKNOWN
}

func (x *Field) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *Field) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Field) GetTypeUrl() string {
	if x != nil {
		return x.TypeUrl
	}
	return ""
}

func (x *Field) GetOneofIndex() int32 {
	if x != nil {
		return x.OneofIndex
	}
	return 0
}

func (x *Field) GetPacked() bool {
	if x != nil {
		return x.Packed
	}
	return false
}

func (x *Field) GetOptions() []*Option {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *Field) GetJsonName() string {
	if x != nil {
		return x.JsonName
	}
	return ""
}

func (x *Field) GetDefaultValue() string {
	if x != nil {
		return x.DefaultValue
	}
	return ""
}

// Enum type definition.
type Enum struct {
	// Enum type name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Enum value definitions.
	Enumvalue []*EnumValue `protobuf:"bytes,2,rep,name=enumvalue,proto3" json:"enumvalue,omitempty"`
	// Protocol buffer options.
	Options []*Option `protobuf:"bytes,3,rep,name=options,proto3" json:"options,omitempty"`
	// The source context.
	SourceContext *SourceContext `protobuf:"bytes,4,opt,name=source_context,json=sourceContext,proto3" json:"source_context,omitempty"`
	// The source syntax.
	Syntax               Syntax   `protobuf:"varint,5,opt,name=syntax,proto3,enum=google.protobuf.Syntax" json:"syntax,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (x *Enum) Reset() {
	*x = Enum{}
}

func (x *Enum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Enum) ProtoMessage() {}

func (x *Enum) ProtoReflect() protoreflect.Message {
	return xxx_File_google_protobuf_type_proto_messageTypes[2].MessageOf(x)
}

func (m *Enum) XXX_Methods() *protoiface.Methods {
	return xxx_File_google_protobuf_type_proto_messageTypes[2].Methods()
}

// Deprecated: Use Enum.ProtoReflect.Type instead.
func (*Enum) Descriptor() ([]byte, []int) {
	return xxx_File_google_protobuf_type_proto_rawDescGZIP(), []int{2}
}

func (x *Enum) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Enum) GetEnumvalue() []*EnumValue {
	if x != nil {
		return x.Enumvalue
	}
	return nil
}

func (x *Enum) GetOptions() []*Option {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *Enum) GetSourceContext() *SourceContext {
	if x != nil {
		return x.SourceContext
	}
	return nil
}

func (x *Enum) GetSyntax() Syntax {
	if x != nil {
		return x.Syntax
	}
	return Syntax_SYNTAX_PROTO2
}

// Enum value definition.
type EnumValue struct {
	// Enum value name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Enum value number.
	Number int32 `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	// Protocol buffer options.
	Options              []*Option `protobuf:"bytes,3,rep,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (x *EnumValue) Reset() {
	*x = EnumValue{}
}

func (x *EnumValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnumValue) ProtoMessage() {}

func (x *EnumValue) ProtoReflect() protoreflect.Message {
	return xxx_File_google_protobuf_type_proto_messageTypes[3].MessageOf(x)
}

func (m *EnumValue) XXX_Methods() *protoiface.Methods {
	return xxx_File_google_protobuf_type_proto_messageTypes[3].Methods()
}

// Deprecated: Use EnumValue.ProtoReflect.Type instead.
func (*EnumValue) Descriptor() ([]byte, []int) {
	return xxx_File_google_protobuf_type_proto_rawDescGZIP(), []int{3}
}

func (x *EnumValue) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EnumValue) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *EnumValue) GetOptions() []*Option {
	if x != nil {
		return x.Options
	}
	return nil
}

// A protocol buffer option, which can be attached to a message, field,
// enumeration, etc.
type Option struct {
	// The option's name. For protobuf built-in options (options defined in
	// descriptor.proto), this is the short name. For example, `"map_entry"`.
	// For custom options, it should be the fully-qualified name. For example,
	// `"google.api.http"`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The option's value packed in an Any message. If the value is a primitive,
	// the corresponding wrapper type defined in google/protobuf/wrappers.proto
	// should be used. If the value is an enum, it should be stored as an int32
	// value using the google.protobuf.Int32Value type.
	Value                *Any     `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (x *Option) Reset() {
	*x = Option{}
}

func (x *Option) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Option) ProtoMessage() {}

func (x *Option) ProtoReflect() protoreflect.Message {
	return xxx_File_google_protobuf_type_proto_messageTypes[4].MessageOf(x)
}

func (m *Option) XXX_Methods() *protoiface.Methods {
	return xxx_File_google_protobuf_type_proto_messageTypes[4].Methods()
}

// Deprecated: Use Option.ProtoReflect.Type instead.
func (*Option) Descriptor() ([]byte, []int) {
	return xxx_File_google_protobuf_type_proto_rawDescGZIP(), []int{4}
}

func (x *Option) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Option) GetValue() *Any {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_google_protobuf_type_proto protoreflect.FileDescriptor

var xxx_File_google_protobuf_type_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61,
	0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d,
	0x02, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x6e, 0x65, 0x6f, 0x66, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x6e, 0x65,
	0x6f, 0x66, 0x73, 0x12, 0x31, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x45, 0x0a, 0x0e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x0d,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x2f, 0x0a,
	0x06, 0x73, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x52, 0x06, 0x73, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x22, 0xb4,
	0x06, 0x0a, 0x05, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x2f, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x4b,
	0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x44, 0x0a, 0x0b, 0x63, 0x61, 0x72,
	0x64, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c, 0x69,
	0x74, 0x79, 0x52, 0x0b, 0x63, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12,
	0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x74,
	0x79, 0x70, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74,
	0x79, 0x70, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6f, 0x6e, 0x65,
	0x6f, 0x66, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x63, 0x6b, 0x65,
	0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x12,
	0x31, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6a, 0x73, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0xc8, 0x02, 0x0a, 0x04, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x10, 0x0a,
	0x0c, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12,
	0x0f, 0x0a, 0x0b, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x4f, 0x55, 0x42, 0x4c, 0x45, 0x10, 0x01,
	0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x10, 0x02,
	0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x36, 0x34, 0x10, 0x03,
	0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x49, 0x4e, 0x54, 0x36, 0x34, 0x10,
	0x04, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x33, 0x32, 0x10,
	0x05, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x49, 0x58, 0x45, 0x44, 0x36,
	0x34, 0x10, 0x06, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x49, 0x58, 0x45,
	0x44, 0x33, 0x32, 0x10, 0x07, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x4f,
	0x4f, 0x4c, 0x10, 0x08, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x52,
	0x49, 0x4e, 0x47, 0x10, 0x09, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x47, 0x52,
	0x4f, 0x55, 0x50, 0x10, 0x0a, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x45,
	0x53, 0x53, 0x41, 0x47, 0x45, 0x10, 0x0b, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x42, 0x59, 0x54, 0x45, 0x53, 0x10, 0x0c, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x55, 0x49, 0x4e, 0x54, 0x33, 0x32, 0x10, 0x0d, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x10, 0x0e, 0x12, 0x11, 0x0a, 0x0d, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x53, 0x46, 0x49, 0x58, 0x45, 0x44, 0x33, 0x32, 0x10, 0x0f, 0x12, 0x11, 0x0a, 0x0d, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x53, 0x46, 0x49, 0x58, 0x45, 0x44, 0x36, 0x34, 0x10, 0x10, 0x12, 0x0f, 0x0a,
	0x0b, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x49, 0x4e, 0x54, 0x33, 0x32, 0x10, 0x11, 0x12, 0x0f,
	0x0a, 0x0b, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x49, 0x4e, 0x54, 0x36, 0x34, 0x10, 0x12, 0x22,
	0x74, 0x0a, 0x0b, 0x43, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x17,
	0x0a, 0x13, 0x43, 0x41, 0x52, 0x44, 0x49, 0x4e, 0x41, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x41, 0x52, 0x44, 0x49,
	0x4e, 0x41, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x10,
	0x01, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x41, 0x52, 0x44, 0x49, 0x4e, 0x41, 0x4c, 0x49, 0x54, 0x59,
	0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x43,
	0x41, 0x52, 0x44, 0x49, 0x4e, 0x41, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x52, 0x45, 0x50, 0x45, 0x41,
	0x54, 0x45, 0x44, 0x10, 0x03, 0x22, 0xff, 0x01, 0x0a, 0x04, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x65, 0x6e, 0x75, 0x6d, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x09, 0x65, 0x6e, 0x75, 0x6d, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x31, 0x0a, 0x07,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x45, 0x0a, 0x0e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x0d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x2f, 0x0a, 0x06, 0x73, 0x79, 0x6e, 0x74, 0x61, 0x78,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x52,
	0x06, 0x73, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x22, 0x6a, 0x0a, 0x09, 0x45, 0x6e, 0x75, 0x6d, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x31, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x22, 0x48, 0x0a, 0x06, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2a, 0x2e, 0x0a,
	0x06, 0x53, 0x79, 0x6e, 0x74, 0x61, 0x78, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x59, 0x4e, 0x54, 0x41,
	0x58, 0x5f, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x32, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x59,
	0x4e, 0x54, 0x41, 0x58, 0x5f, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x33, 0x10, 0x01, 0x42, 0x83, 0x01,
	0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x42, 0x09, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x76,
	0x32, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x3b, 0x6b, 0x6e,
	0x6f, 0x77, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x03, 0x47,
	0x50, 0x42, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x57, 0x65, 0x6c, 0x6c, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	xxx_File_google_protobuf_type_proto_rawDesc_once sync.Once
	xxx_File_google_protobuf_type_proto_rawDesc_data = xxx_File_google_protobuf_type_proto_rawDesc
)

func xxx_File_google_protobuf_type_proto_rawDescGZIP() []byte {
	xxx_File_google_protobuf_type_proto_rawDesc_once.Do(func() {
		xxx_File_google_protobuf_type_proto_rawDesc_data = protoimpl.X.CompressGZIP(xxx_File_google_protobuf_type_proto_rawDesc_data)
	})
	return xxx_File_google_protobuf_type_proto_rawDesc_data
}

var xxx_File_google_protobuf_type_proto_enumTypes = make([]protoreflect.EnumType, 3)
var xxx_File_google_protobuf_type_proto_messageTypes = make([]protoimpl.MessageType, 5)
var xxx_File_google_protobuf_type_proto_goTypes = []interface{}{
	(Syntax)(0),            // 0: google.protobuf.Syntax
	(Field_Kind)(0),        // 1: google.protobuf.Field.Kind
	(Field_Cardinality)(0), // 2: google.protobuf.Field.Cardinality
	(*Type)(nil),           // 3: google.protobuf.Type
	(*Field)(nil),          // 4: google.protobuf.Field
	(*Enum)(nil),           // 5: google.protobuf.Enum
	(*EnumValue)(nil),      // 6: google.protobuf.EnumValue
	(*Option)(nil),         // 7: google.protobuf.Option
	(*SourceContext)(nil),  // 8: google.protobuf.SourceContext
	(*Any)(nil),            // 9: google.protobuf.Any
}
var xxx_File_google_protobuf_type_proto_depIdxs = []int32{
	4, // google.protobuf.Type.fields:type_name -> google.protobuf.Field
	7, // google.protobuf.Type.options:type_name -> google.protobuf.Option
	8, // google.protobuf.Type.source_context:type_name -> google.protobuf.SourceContext
	0, // google.protobuf.Type.syntax:type_name -> google.protobuf.Syntax
	1, // google.protobuf.Field.kind:type_name -> google.protobuf.Field.Kind
	2, // google.protobuf.Field.cardinality:type_name -> google.protobuf.Field.Cardinality
	7, // google.protobuf.Field.options:type_name -> google.protobuf.Option
	6, // google.protobuf.Enum.enumvalue:type_name -> google.protobuf.EnumValue
	7, // google.protobuf.Enum.options:type_name -> google.protobuf.Option
	8, // google.protobuf.Enum.source_context:type_name -> google.protobuf.SourceContext
	0, // google.protobuf.Enum.syntax:type_name -> google.protobuf.Syntax
	7, // google.protobuf.EnumValue.options:type_name -> google.protobuf.Option
	9, // google.protobuf.Option.value:type_name -> google.protobuf.Any
}

func init() { xxx_File_google_protobuf_type_proto_init() }
func xxx_File_google_protobuf_type_proto_init() {
	if File_google_protobuf_type_proto != nil {
		return
	}
	xxx_File_google_protobuf_any_proto_init()
	xxx_File_google_protobuf_source_context_proto_init()
	File_google_protobuf_type_proto = protoimpl.FileBuilder{
		RawDescriptor:      xxx_File_google_protobuf_type_proto_rawDesc,
		GoTypes:            xxx_File_google_protobuf_type_proto_goTypes,
		DependencyIndexes:  xxx_File_google_protobuf_type_proto_depIdxs,
		EnumOutputTypes:    xxx_File_google_protobuf_type_proto_enumTypes,
		MessageOutputTypes: xxx_File_google_protobuf_type_proto_messageTypes,
		FilesRegistry:      protoregistry.GlobalFiles,
		TypesRegistry:      protoregistry.GlobalTypes,
	}.Init()
	xxx_File_google_protobuf_type_proto_rawDesc = nil
	xxx_File_google_protobuf_type_proto_goTypes = nil
	xxx_File_google_protobuf_type_proto_depIdxs = nil
}
