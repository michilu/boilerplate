// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v/meta/message.proto

package meta

import (
	fmt "fmt"
	math "math"
	"time"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
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

// Meta is ValueObject of meta infomation
type Meta struct {
	Name                 string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Semver               string    `protobuf:"bytes,2,opt,name=semver,proto3" json:"semver,omitempty"`
	Runtime              *Runtime  `protobuf:"bytes,3,opt,name=runtime,proto3" json:"runtime,omitempty"`
	Build                time.Time `protobuf:"bytes,4,opt,name=build,proto3" json:"build,omitempty"`
	Serial               string    `protobuf:"bytes,5,opt,name=serial,proto3" json:"serial,omitempty"`
	Vcs                  *Vcs      `protobuf:"bytes,6,opt,name=vcs,proto3" json:"vcs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-" yaml:"-"`
	XXX_unrecognized     []byte    `json:"-" yaml:"-"`
	XXX_sizecache        int32     `json:"-" yaml:"-"`
}

func (m *Meta) Reset()         { *m = Meta{} }
func (m *Meta) String() string { return proto.CompactTextString(m) }
func (*Meta) ProtoMessage()    {}
func (*Meta) Descriptor() ([]byte, []int) {
	return fileDescriptor_30750152cff00771, []int{0}
}

func (m *Meta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Meta.Unmarshal(m, b)
}
func (m *Meta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Meta.Marshal(b, m, deterministic)
}
func (m *Meta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Meta.Merge(m, src)
}
func (m *Meta) XXX_Size() int {
	return xxx_messageInfo_Meta.Size(m)
}
func (m *Meta) XXX_DiscardUnknown() {
	xxx_messageInfo_Meta.DiscardUnknown(m)
}

var xxx_messageInfo_Meta proto.InternalMessageInfo

func (m *Meta) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Meta) GetSemver() string {
	if m != nil {
		return m.Semver
	}
	return ""
}

func (m *Meta) GetRuntime() *Runtime {
	if m != nil {
		return m.Runtime
	}
	return nil
}

func (m *Meta) GetBuild() time.Time {
	return m.Build
}

func (m *Meta) GetSerial() string {
	if m != nil {
		return m.Serial
	}
	return ""
}

func (m *Meta) GetVcs() *Vcs {
	if m != nil {
		return m.Vcs
	}
	return nil
}

// Runtime is ValueObject of Runtime
type Runtime struct {
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Arch                 string   `protobuf:"bytes,2,opt,name=arch,proto3" json:"arch,omitempty"`
	Os                   string   `protobuf:"bytes,3,opt,name=os,proto3" json:"os,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" yaml:"-"`
	XXX_unrecognized     []byte   `json:"-" yaml:"-"`
	XXX_sizecache        int32    `json:"-" yaml:"-"`
}

func (m *Runtime) Reset()         { *m = Runtime{} }
func (m *Runtime) String() string { return proto.CompactTextString(m) }
func (*Runtime) ProtoMessage()    {}
func (*Runtime) Descriptor() ([]byte, []int) {
	return fileDescriptor_30750152cff00771, []int{1}
}

func (m *Runtime) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Runtime.Unmarshal(m, b)
}
func (m *Runtime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Runtime.Marshal(b, m, deterministic)
}
func (m *Runtime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Runtime.Merge(m, src)
}
func (m *Runtime) XXX_Size() int {
	return xxx_messageInfo_Runtime.Size(m)
}
func (m *Runtime) XXX_DiscardUnknown() {
	xxx_messageInfo_Runtime.DiscardUnknown(m)
}

var xxx_messageInfo_Runtime proto.InternalMessageInfo

func (m *Runtime) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Runtime) GetArch() string {
	if m != nil {
		return m.Arch
	}
	return ""
}

func (m *Runtime) GetOs() string {
	if m != nil {
		return m.Os
	}
	return ""
}

// Vcs is ValueObject of VCS
type Vcs struct {
	Branch               string   `protobuf:"bytes,1,opt,name=branch,proto3" json:"branch,omitempty"`
	Hash                 string   `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" yaml:"-"`
	XXX_unrecognized     []byte   `json:"-" yaml:"-"`
	XXX_sizecache        int32    `json:"-" yaml:"-"`
}

func (m *Vcs) Reset()         { *m = Vcs{} }
func (m *Vcs) String() string { return proto.CompactTextString(m) }
func (*Vcs) ProtoMessage()    {}
func (*Vcs) Descriptor() ([]byte, []int) {
	return fileDescriptor_30750152cff00771, []int{2}
}

func (m *Vcs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vcs.Unmarshal(m, b)
}
func (m *Vcs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vcs.Marshal(b, m, deterministic)
}
func (m *Vcs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vcs.Merge(m, src)
}
func (m *Vcs) XXX_Size() int {
	return xxx_messageInfo_Vcs.Size(m)
}
func (m *Vcs) XXX_DiscardUnknown() {
	xxx_messageInfo_Vcs.DiscardUnknown(m)
}

var xxx_messageInfo_Vcs proto.InternalMessageInfo

func (m *Vcs) GetBranch() string {
	if m != nil {
		return m.Branch
	}
	return ""
}

func (m *Vcs) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*Meta)(nil), "michilu.boilerplate.v.meta.Meta")
	proto.RegisterType((*Runtime)(nil), "michilu.boilerplate.v.meta.Runtime")
	proto.RegisterType((*Vcs)(nil), "michilu.boilerplate.v.meta.Vcs")
}

func init() { proto.RegisterFile("v/meta/message.proto", fileDescriptor_30750152cff00771) }

var fileDescriptor_30750152cff00771 = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xc1, 0x6a, 0xe3, 0x30,
	0x10, 0x45, 0x8e, 0xe3, 0xb0, 0xb3, 0x7b, 0x12, 0x21, 0x98, 0x1c, 0x12, 0x93, 0xbd, 0xe4, 0x12,
	0x99, 0xec, 0x42, 0x3e, 0x20, 0xd0, 0x63, 0x20, 0xe8, 0x10, 0x4a, 0x2f, 0x45, 0x56, 0x44, 0x2c,
	0x6a, 0x59, 0x41, 0x92, 0x9d, 0xcf, 0x2d, 0xf4, 0xd8, 0xaf, 0x28, 0x92, 0x9d, 0x96, 0x06, 0x9a,
	0xdb, 0xcc, 0x9b, 0x99, 0xf7, 0xe6, 0x69, 0x04, 0xe3, 0x36, 0x57, 0xc2, 0xb1, 0x5c, 0x09, 0x6b,
	0xd9, 0x49, 0x90, 0xb3, 0xd1, 0x4e, 0xe3, 0xa9, 0x92, 0xbc, 0x94, 0x55, 0x43, 0x0a, 0x2d, 0x2b,
	0x61, 0xce, 0x15, 0x73, 0x82, 0xb4, 0xc4, 0x77, 0x4e, 0x37, 0x27, 0xe9, 0xca, 0xa6, 0x20, 0x5c,
	0xab, 0x5c, 0x5d, 0xa4, 0x7b, 0xd1, 0x97, 0xfc, 0xa4, 0x57, 0x61, 0x70, 0xd5, 0xb2, 0x4a, 0x1e,
	0x99, 0xd3, 0xc6, 0xe6, 0x9f, 0x61, 0xc7, 0xb9, 0x78, 0x47, 0x10, 0xef, 0x84, 0x63, 0x78, 0x0a,
	0x71, 0xcd, 0x94, 0x48, 0x51, 0x86, 0x96, 0xbf, 0xb6, 0xc9, 0xdb, 0xeb, 0x3c, 0x7a, 0x44, 0x34,
	0x60, 0x78, 0x06, 0x89, 0x15, 0xaa, 0x15, 0x26, 0x8d, 0xbe, 0x55, 0x7b, 0x14, 0x3f, 0xc0, 0xc8,
	0x34, 0xb5, 0x93, 0x4a, 0xa4, 0x83, 0x0c, 0x2d, 0x7f, 0xff, 0xfb, 0x4b, 0x7e, 0x5e, 0x95, 0xd0,
	0xae, 0xb5, 0x63, 0xc9, 0x10, 0xbd, 0xce, 0xe2, 0x31, 0x0c, 0x8b, 0x46, 0x56, 0xc7, 0x34, 0xf6,
	0x2a, 0xb4, 0x4b, 0xf0, 0xc4, 0x8b, 0x1b, 0xc9, 0xaa, 0x74, 0x18, 0xe0, 0x3e, 0xc3, 0x6b, 0x18,
	0xb4, 0xdc, 0xa6, 0x49, 0x10, 0x9c, 0xdf, 0x13, 0x3c, 0x70, 0x4b, 0x7d, 0xef, 0xe2, 0x19, 0x46,
	0xbd, 0x38, 0xce, 0x60, 0xd4, 0x0a, 0x63, 0xa5, 0xae, 0x6f, 0x1c, 0x5f, 0x61, 0xff, 0x20, 0xcc,
	0xf0, 0xf2, 0xc6, 0x72, 0xc0, 0xf0, 0x04, 0x22, 0x6d, 0x83, 0xd7, 0xaf, 0x4a, 0xa4, 0xed, 0x62,
	0x0d, 0x83, 0x03, 0xb7, 0x7e, 0xe5, 0xc2, 0xb0, 0x9a, 0x97, 0x1d, 0x37, 0xed, 0x33, 0x8c, 0x21,
	0x2e, 0x99, 0xed, 0x29, 0x69, 0x88, 0xb7, 0x1b, 0x98, 0x71, 0xad, 0xee, 0xac, 0xbf, 0xfd, 0xb3,
	0xeb, 0x7e, 0xc1, 0xde, 0x1f, 0x6c, 0x8f, 0x9e, 0x62, 0x8f, 0x16, 0x49, 0xb8, 0xdf, 0xff, 0x8f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x62, 0x95, 0x28, 0x12, 0x2b, 0x02, 0x00, 0x00,
}
