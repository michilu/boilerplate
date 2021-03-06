// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/type.proto

package pb

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Bool is ValueObject.
type Bool struct {
	Value                bool     `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bool) Reset()         { *m = Bool{} }
func (m *Bool) String() string { return proto.CompactTextString(m) }
func (*Bool) ProtoMessage()    {}
func (*Bool) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1b2ccf908edb774, []int{0}
}

func (m *Bool) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bool.Unmarshal(m, b)
}
func (m *Bool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bool.Marshal(b, m, deterministic)
}
func (m *Bool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bool.Merge(m, src)
}
func (m *Bool) XXX_Size() int {
	return xxx_messageInfo_Bool.Size(m)
}
func (m *Bool) XXX_DiscardUnknown() {
	xxx_messageInfo_Bool.DiscardUnknown(m)
}

var xxx_messageInfo_Bool proto.InternalMessageInfo

func (m *Bool) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

// Byte is ValueObject.
type Byte struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Byte) Reset()         { *m = Byte{} }
func (m *Byte) String() string { return proto.CompactTextString(m) }
func (*Byte) ProtoMessage()    {}
func (*Byte) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1b2ccf908edb774, []int{1}
}

func (m *Byte) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Byte.Unmarshal(m, b)
}
func (m *Byte) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Byte.Marshal(b, m, deterministic)
}
func (m *Byte) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Byte.Merge(m, src)
}
func (m *Byte) XXX_Size() int {
	return xxx_messageInfo_Byte.Size(m)
}
func (m *Byte) XXX_DiscardUnknown() {
	xxx_messageInfo_Byte.DiscardUnknown(m)
}

var xxx_messageInfo_Byte proto.InternalMessageInfo

func (m *Byte) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*Bool)(nil), "michilu.boilerplate.pb.Bool")
	proto.RegisterType((*Byte)(nil), "michilu.boilerplate.pb.Byte")
}

func init() { proto.RegisterFile("pb/type.proto", fileDescriptor_d1b2ccf908edb774) }

var fileDescriptor_d1b2ccf908edb774 = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x48, 0xd2, 0x2f,
	0xa9, 0x2c, 0x48, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xcb, 0xcd, 0x4c, 0xce, 0xc8,
	0xcc, 0x29, 0xd5, 0x4b, 0xca, 0xcf, 0xcc, 0x49, 0x2d, 0x2a, 0xc8, 0x49, 0x2c, 0x49, 0xd5, 0x2b,
	0x48, 0x92, 0x72, 0x4c, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcd,
	0x2b, 0xcb, 0xaf, 0x2c, 0x28, 0xca, 0xaf, 0xa8, 0xd4, 0x07, 0x6b, 0x4a, 0xd6, 0x4d, 0x4f, 0xcd,
	0xd3, 0x2d, 0x4b, 0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0x49, 0xd5, 0xc7, 0x60, 0x40, 0x8c, 0x56, 0x92,
	0xe1, 0x62, 0x71, 0xca, 0xcf, 0xcf, 0x11, 0x12, 0xe1, 0x62, 0x2d, 0x4b, 0xcc, 0x29, 0x4d, 0x95,
	0x60, 0x54, 0x60, 0xd4, 0xe0, 0x08, 0x82, 0x70, 0x94, 0x54, 0xb9, 0x58, 0x9c, 0x2a, 0x4b, 0x52,
	0x85, 0x64, 0x91, 0x65, 0x79, 0x9c, 0xd8, 0x7f, 0x39, 0xb1, 0x54, 0x31, 0x09, 0x30, 0x42, 0x95,
	0x39, 0xb1, 0x44, 0x31, 0x15, 0x24, 0x25, 0xb1, 0x81, 0x4d, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff,
	0xff, 0xf9, 0xb0, 0x2b, 0x36, 0xbd, 0x00, 0x00, 0x00,
}
