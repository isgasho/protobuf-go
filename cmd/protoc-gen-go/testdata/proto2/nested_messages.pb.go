// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto2/nested_messages.proto

package proto2

import (
	protoreflect "github.com/golang/protobuf/v2/reflect/protoreflect"
	protoregistry "github.com/golang/protobuf/v2/reflect/protoregistry"
	protoiface "github.com/golang/protobuf/v2/runtime/protoiface"
	protoimpl "github.com/golang/protobuf/v2/runtime/protoimpl"
	sync "sync"
)

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

type Layer1 struct {
	L2                   *Layer1_Layer2        `protobuf:"bytes,1,opt,name=l2" json:"l2,omitempty"`
	L3                   *Layer1_Layer2_Layer3 `protobuf:"bytes,2,opt,name=l3" json:"l3,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (x *Layer1) Reset() {
	*x = Layer1{}
}

func (x *Layer1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Layer1) ProtoMessage() {}

func (x *Layer1) ProtoReflect() protoreflect.Message {
	return xxx_File_proto2_nested_messages_proto_messageTypes[0].MessageOf(x)
}

func (m *Layer1) XXX_Methods() *protoiface.Methods {
	return xxx_File_proto2_nested_messages_proto_messageTypes[0].Methods()
}

// Deprecated: Use Layer1.ProtoReflect.Type instead.
func (*Layer1) Descriptor() ([]byte, []int) {
	return xxx_File_proto2_nested_messages_proto_rawDescGZIP(), []int{0}
}

func (x *Layer1) GetL2() *Layer1_Layer2 {
	if x != nil {
		return x.L2
	}
	return nil
}

func (x *Layer1) GetL3() *Layer1_Layer2_Layer3 {
	if x != nil {
		return x.L3
	}
	return nil
}

type Layer1_Layer2 struct {
	L3                   *Layer1_Layer2_Layer3 `protobuf:"bytes,1,opt,name=l3" json:"l3,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (x *Layer1_Layer2) Reset() {
	*x = Layer1_Layer2{}
}

func (x *Layer1_Layer2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Layer1_Layer2) ProtoMessage() {}

func (x *Layer1_Layer2) ProtoReflect() protoreflect.Message {
	return xxx_File_proto2_nested_messages_proto_messageTypes[1].MessageOf(x)
}

func (m *Layer1_Layer2) XXX_Methods() *protoiface.Methods {
	return xxx_File_proto2_nested_messages_proto_messageTypes[1].Methods()
}

// Deprecated: Use Layer1_Layer2.ProtoReflect.Type instead.
func (*Layer1_Layer2) Descriptor() ([]byte, []int) {
	return xxx_File_proto2_nested_messages_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Layer1_Layer2) GetL3() *Layer1_Layer2_Layer3 {
	if x != nil {
		return x.L3
	}
	return nil
}

type Layer1_Layer2_Layer3 struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (x *Layer1_Layer2_Layer3) Reset() {
	*x = Layer1_Layer2_Layer3{}
}

func (x *Layer1_Layer2_Layer3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Layer1_Layer2_Layer3) ProtoMessage() {}

func (x *Layer1_Layer2_Layer3) ProtoReflect() protoreflect.Message {
	return xxx_File_proto2_nested_messages_proto_messageTypes[2].MessageOf(x)
}

func (m *Layer1_Layer2_Layer3) XXX_Methods() *protoiface.Methods {
	return xxx_File_proto2_nested_messages_proto_messageTypes[2].Methods()
}

// Deprecated: Use Layer1_Layer2_Layer3.ProtoReflect.Type instead.
func (*Layer1_Layer2_Layer3) Descriptor() ([]byte, []int) {
	return xxx_File_proto2_nested_messages_proto_rawDescGZIP(), []int{0, 0, 0}
}

var File_proto2_nested_messages_proto protoreflect.FileDescriptor

var xxx_File_proto2_nested_messages_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x2f, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15,
	0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0x22, 0xcc, 0x01, 0x0a, 0x06, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x31,
	0x12, 0x34, 0x0a, 0x02, 0x6c, 0x32, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67,
	0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x32, 0x2e, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x31, 0x2e, 0x4c, 0x61, 0x79, 0x65,
	0x72, 0x32, 0x52, 0x02, 0x6c, 0x32, 0x12, 0x3b, 0x0a, 0x02, 0x6c, 0x33, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x2e, 0x4c, 0x61, 0x79, 0x65, 0x72,
	0x31, 0x2e, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x32, 0x2e, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x33, 0x52,
	0x02, 0x6c, 0x33, 0x1a, 0x4f, 0x0a, 0x06, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x32, 0x12, 0x3b, 0x0a,
	0x02, 0x6c, 0x33, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67, 0x6f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0x2e, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x31, 0x2e, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x32, 0x2e,
	0x4c, 0x61, 0x79, 0x65, 0x72, 0x33, 0x52, 0x02, 0x6c, 0x33, 0x1a, 0x08, 0x0a, 0x06, 0x4c, 0x61,
	0x79, 0x65, 0x72, 0x33, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
}

var (
	xxx_File_proto2_nested_messages_proto_rawDesc_once sync.Once
	xxx_File_proto2_nested_messages_proto_rawDesc_data = xxx_File_proto2_nested_messages_proto_rawDesc
)

func xxx_File_proto2_nested_messages_proto_rawDescGZIP() []byte {
	xxx_File_proto2_nested_messages_proto_rawDesc_once.Do(func() {
		xxx_File_proto2_nested_messages_proto_rawDesc_data = protoimpl.X.CompressGZIP(xxx_File_proto2_nested_messages_proto_rawDesc_data)
	})
	return xxx_File_proto2_nested_messages_proto_rawDesc_data
}

var xxx_File_proto2_nested_messages_proto_messageTypes = make([]protoimpl.MessageType, 3)
var xxx_File_proto2_nested_messages_proto_goTypes = []interface{}{
	(*Layer1)(nil),               // 0: goproto.protoc.proto2.Layer1
	(*Layer1_Layer2)(nil),        // 1: goproto.protoc.proto2.Layer1.Layer2
	(*Layer1_Layer2_Layer3)(nil), // 2: goproto.protoc.proto2.Layer1.Layer2.Layer3
}
var xxx_File_proto2_nested_messages_proto_depIdxs = []int32{
	1, // goproto.protoc.proto2.Layer1.l2:type_name -> goproto.protoc.proto2.Layer1.Layer2
	2, // goproto.protoc.proto2.Layer1.l3:type_name -> goproto.protoc.proto2.Layer1.Layer2.Layer3
	2, // goproto.protoc.proto2.Layer1.Layer2.l3:type_name -> goproto.protoc.proto2.Layer1.Layer2.Layer3
}

func init() { xxx_File_proto2_nested_messages_proto_init() }
func xxx_File_proto2_nested_messages_proto_init() {
	if File_proto2_nested_messages_proto != nil {
		return
	}
	File_proto2_nested_messages_proto = protoimpl.FileBuilder{
		RawDescriptor:      xxx_File_proto2_nested_messages_proto_rawDesc,
		GoTypes:            xxx_File_proto2_nested_messages_proto_goTypes,
		DependencyIndexes:  xxx_File_proto2_nested_messages_proto_depIdxs,
		MessageOutputTypes: xxx_File_proto2_nested_messages_proto_messageTypes,
		FilesRegistry:      protoregistry.GlobalFiles,
		TypesRegistry:      protoregistry.GlobalTypes,
	}.Init()
	xxx_File_proto2_nested_messages_proto_rawDesc = nil
	xxx_File_proto2_nested_messages_proto_goTypes = nil
	xxx_File_proto2_nested_messages_proto_depIdxs = nil
}
