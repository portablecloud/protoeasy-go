// Code generated by protoc-gen-go.
// source: google/protobuf/wrappers.proto
// DO NOT EDIT!

package google_protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Wrapper message for `double`.
//
// The JSON representation for `DoubleValue` is JSON number.
type DoubleValue struct {
	// The double value.
	Value float64 `protobuf:"fixed64,1,opt,name=value" json:"value,omitempty"`
}

func (m *DoubleValue) Reset()                    { *m = DoubleValue{} }
func (m *DoubleValue) String() string            { return proto.CompactTextString(m) }
func (*DoubleValue) ProtoMessage()               {}
func (*DoubleValue) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{0} }
func (*DoubleValue) XXX_WellKnownType() string   { return "DoubleValue" }

func (m *DoubleValue) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

// Wrapper message for `float`.
//
// The JSON representation for `FloatValue` is JSON number.
type FloatValue struct {
	// The float value.
	Value float32 `protobuf:"fixed32,1,opt,name=value" json:"value,omitempty"`
}

func (m *FloatValue) Reset()                    { *m = FloatValue{} }
func (m *FloatValue) String() string            { return proto.CompactTextString(m) }
func (*FloatValue) ProtoMessage()               {}
func (*FloatValue) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{1} }
func (*FloatValue) XXX_WellKnownType() string   { return "FloatValue" }

func (m *FloatValue) GetValue() float32 {
	if m != nil {
		return m.Value
	}
	return 0
}

// Wrapper message for `int64`.
//
// The JSON representation for `Int64Value` is JSON string.
type Int64Value struct {
	// The int64 value.
	Value int64 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *Int64Value) Reset()                    { *m = Int64Value{} }
func (m *Int64Value) String() string            { return proto.CompactTextString(m) }
func (*Int64Value) ProtoMessage()               {}
func (*Int64Value) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{2} }
func (*Int64Value) XXX_WellKnownType() string   { return "Int64Value" }

func (m *Int64Value) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

// Wrapper message for `uint64`.
//
// The JSON representation for `UInt64Value` is JSON string.
type UInt64Value struct {
	// The uint64 value.
	Value uint64 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *UInt64Value) Reset()                    { *m = UInt64Value{} }
func (m *UInt64Value) String() string            { return proto.CompactTextString(m) }
func (*UInt64Value) ProtoMessage()               {}
func (*UInt64Value) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{3} }
func (*UInt64Value) XXX_WellKnownType() string   { return "UInt64Value" }

func (m *UInt64Value) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

// Wrapper message for `int32`.
//
// The JSON representation for `Int32Value` is JSON number.
type Int32Value struct {
	// The int32 value.
	Value int32 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *Int32Value) Reset()                    { *m = Int32Value{} }
func (m *Int32Value) String() string            { return proto.CompactTextString(m) }
func (*Int32Value) ProtoMessage()               {}
func (*Int32Value) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{4} }
func (*Int32Value) XXX_WellKnownType() string   { return "Int32Value" }

func (m *Int32Value) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

// Wrapper message for `uint32`.
//
// The JSON representation for `UInt32Value` is JSON number.
type UInt32Value struct {
	// The uint32 value.
	Value uint32 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *UInt32Value) Reset()                    { *m = UInt32Value{} }
func (m *UInt32Value) String() string            { return proto.CompactTextString(m) }
func (*UInt32Value) ProtoMessage()               {}
func (*UInt32Value) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{5} }
func (*UInt32Value) XXX_WellKnownType() string   { return "UInt32Value" }

func (m *UInt32Value) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

// Wrapper message for `bool`.
//
// The JSON representation for `BoolValue` is JSON `true` and `false`.
type BoolValue struct {
	// The bool value.
	Value bool `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *BoolValue) Reset()                    { *m = BoolValue{} }
func (m *BoolValue) String() string            { return proto.CompactTextString(m) }
func (*BoolValue) ProtoMessage()               {}
func (*BoolValue) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{6} }
func (*BoolValue) XXX_WellKnownType() string   { return "BoolValue" }

func (m *BoolValue) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

// Wrapper message for `string`.
//
// The JSON representation for `StringValue` is JSON string.
type StringValue struct {
	// The string value.
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *StringValue) Reset()                    { *m = StringValue{} }
func (m *StringValue) String() string            { return proto.CompactTextString(m) }
func (*StringValue) ProtoMessage()               {}
func (*StringValue) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{7} }
func (*StringValue) XXX_WellKnownType() string   { return "StringValue" }

func (m *StringValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// Wrapper message for `bytes`.
//
// The JSON representation for `BytesValue` is JSON string.
type BytesValue struct {
	// The bytes value.
	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *BytesValue) Reset()                    { *m = BytesValue{} }
func (m *BytesValue) String() string            { return proto.CompactTextString(m) }
func (*BytesValue) ProtoMessage()               {}
func (*BytesValue) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{8} }
func (*BytesValue) XXX_WellKnownType() string   { return "BytesValue" }

func (m *BytesValue) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*DoubleValue)(nil), "google.protobuf.DoubleValue")
	proto.RegisterType((*FloatValue)(nil), "google.protobuf.FloatValue")
	proto.RegisterType((*Int64Value)(nil), "google.protobuf.Int64Value")
	proto.RegisterType((*UInt64Value)(nil), "google.protobuf.UInt64Value")
	proto.RegisterType((*Int32Value)(nil), "google.protobuf.Int32Value")
	proto.RegisterType((*UInt32Value)(nil), "google.protobuf.UInt32Value")
	proto.RegisterType((*BoolValue)(nil), "google.protobuf.BoolValue")
	proto.RegisterType((*StringValue)(nil), "google.protobuf.StringValue")
	proto.RegisterType((*BytesValue)(nil), "google.protobuf.BytesValue")
}

func init() { proto.RegisterFile("google/protobuf/wrappers.proto", fileDescriptor9) }

var fileDescriptor9 = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x2f, 0x2f, 0x4a, 0x2c,
	0x28, 0x48, 0x2d, 0x2a, 0xd6, 0x03, 0x8b, 0x08, 0xf1, 0x43, 0xe4, 0xf5, 0x60, 0xf2, 0x4a, 0xca,
	0x5c, 0xdc, 0x2e, 0xf9, 0xa5, 0x49, 0x39, 0xa9, 0x61, 0x89, 0x39, 0xa5, 0xa9, 0x42, 0x22, 0x5c,
	0xac, 0x65, 0x20, 0x86, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x63, 0x10, 0x84, 0xa3, 0xa4, 0xc4, 0xc5,
	0xe5, 0x96, 0x93, 0x9f, 0x58, 0x82, 0x45, 0x0d, 0x13, 0x92, 0x1a, 0xcf, 0xbc, 0x12, 0x33, 0x13,
	0x2c, 0x6a, 0x98, 0x61, 0x6a, 0x94, 0xb9, 0xb8, 0x43, 0x71, 0x29, 0x62, 0x41, 0x35, 0xc8, 0xd8,
	0x08, 0x8b, 0x1a, 0x56, 0x34, 0x83, 0xb0, 0x2a, 0xe2, 0x85, 0x29, 0x52, 0xe4, 0xe2, 0x74, 0xca,
	0xcf, 0xcf, 0xc1, 0xa2, 0x84, 0x03, 0xc9, 0x9c, 0xe0, 0x92, 0xa2, 0xcc, 0xbc, 0x74, 0x2c, 0x8a,
	0x38, 0x91, 0x1c, 0xe4, 0x54, 0x59, 0x92, 0x5a, 0x8c, 0x45, 0x0d, 0x0f, 0x54, 0x8d, 0x53, 0x00,
	0x97, 0x70, 0x72, 0x7e, 0xae, 0x1e, 0x5a, 0xe8, 0x3a, 0xf1, 0x86, 0x43, 0x83, 0x3f, 0x00, 0x24,
	0x12, 0xc0, 0xf8, 0x83, 0x91, 0x71, 0x11, 0x13, 0xb3, 0x7b, 0x80, 0xd3, 0x2a, 0x26, 0x39, 0x77,
	0x88, 0xda, 0x00, 0xa8, 0x5a, 0xbd, 0xf0, 0xd4, 0x9c, 0x1c, 0xef, 0xbc, 0xfc, 0xf2, 0xbc, 0x90,
	0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0xb0, 0x21, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf6,
	0x59, 0x0f, 0xd8, 0xd2, 0x01, 0x00, 0x00,
}
