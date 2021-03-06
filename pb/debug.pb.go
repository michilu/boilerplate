// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/debug.proto

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

// DebugClient is Entity
type DebugClient struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DebugClient) Reset()         { *m = DebugClient{} }
func (m *DebugClient) String() string { return proto.CompactTextString(m) }
func (*DebugClient) ProtoMessage()    {}
func (*DebugClient) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dcfd92cd9abdfea, []int{0}
}

func (m *DebugClient) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DebugClient.Unmarshal(m, b)
}
func (m *DebugClient) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DebugClient.Marshal(b, m, deterministic)
}
func (m *DebugClient) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DebugClient.Merge(m, src)
}
func (m *DebugClient) XXX_Size() int {
	return xxx_messageInfo_DebugClient.Size(m)
}
func (m *DebugClient) XXX_DiscardUnknown() {
	xxx_messageInfo_DebugClient.DiscardUnknown(m)
}

var xxx_messageInfo_DebugClient proto.InternalMessageInfo

func (m *DebugClient) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*DebugClient)(nil), "michilu.boilerplate.pb.DebugClient")
}

func init() { proto.RegisterFile("pb/debug.proto", fileDescriptor_4dcfd92cd9abdfea) }

var fileDescriptor_4dcfd92cd9abdfea = []byte{
	// 158 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x48, 0xd2, 0x4f,
	0x49, 0x4d, 0x2a, 0x4d, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xcb, 0xcd, 0x4c, 0xce,
	0xc8, 0xcc, 0x29, 0xd5, 0x4b, 0xca, 0xcf, 0xcc, 0x49, 0x2d, 0x2a, 0xc8, 0x49, 0x2c, 0x49, 0xd5,
	0x2b, 0x48, 0x92, 0x72, 0x4c, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f,
	0xcd, 0x2b, 0xcb, 0xaf, 0x2c, 0x28, 0xca, 0xaf, 0xa8, 0xd4, 0x07, 0x6b, 0x4a, 0xd6, 0x4d, 0x4f,
	0xcd, 0xd3, 0x2d, 0x4b, 0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0x49, 0xd5, 0xc7, 0x60, 0x40, 0x8c, 0x56,
	0x52, 0xe3, 0xe2, 0x76, 0x01, 0xd9, 0xe4, 0x9c, 0x93, 0x99, 0x9a, 0x57, 0x22, 0x24, 0xce, 0xc5,
	0x94, 0x99, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0xe9, 0xc4, 0xfe, 0xcb, 0x89, 0xa5, 0x88, 0x49,
	0x80, 0x31, 0x88, 0x29, 0x33, 0xc5, 0x89, 0x25, 0x8a, 0xa9, 0x20, 0x29, 0x89, 0x0d, 0xac, 0xc9,
	0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xe0, 0x6c, 0xd5, 0x37, 0xa1, 0x00, 0x00, 0x00,
}
