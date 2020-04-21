// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service/slog/appengine.proto

package slog

import (
	fmt "fmt"
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

// AppengineEntry is ValueObject.
type AppengineEntry struct {
	Time                 string                   `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Trace                string                   `protobuf:"bytes,2,opt,name=trace,proto3" json:"trace,omitempty"`
	SourceLocation       *AppengineSourceLocation `protobuf:"bytes,3,opt,name=source_location,json=sourceLocation,proto3" json:"source_location,omitempty"`
	Severity             string                   `protobuf:"bytes,4,opt,name=severity,proto3" json:"severity,omitempty"`
	Message              string                   `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
	Data                 string                   `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *AppengineEntry) Reset()         { *m = AppengineEntry{} }
func (m *AppengineEntry) String() string { return proto.CompactTextString(m) }
func (*AppengineEntry) ProtoMessage()    {}
func (*AppengineEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_e181ce6b77c4050e, []int{0}
}

func (m *AppengineEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppengineEntry.Unmarshal(m, b)
}
func (m *AppengineEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppengineEntry.Marshal(b, m, deterministic)
}
func (m *AppengineEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppengineEntry.Merge(m, src)
}
func (m *AppengineEntry) XXX_Size() int {
	return xxx_messageInfo_AppengineEntry.Size(m)
}
func (m *AppengineEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_AppengineEntry.DiscardUnknown(m)
}

var xxx_messageInfo_AppengineEntry proto.InternalMessageInfo

func (m *AppengineEntry) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

func (m *AppengineEntry) GetTrace() string {
	if m != nil {
		return m.Trace
	}
	return ""
}

func (m *AppengineEntry) GetSourceLocation() *AppengineSourceLocation {
	if m != nil {
		return m.SourceLocation
	}
	return nil
}

func (m *AppengineEntry) GetSeverity() string {
	if m != nil {
		return m.Severity
	}
	return ""
}

func (m *AppengineEntry) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *AppengineEntry) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

// AppengineSourceLocation is ValueObject.
type AppengineSourceLocation struct {
	File                 string   `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
	Line                 string   `protobuf:"bytes,2,opt,name=line,proto3" json:"line,omitempty"`
	Function             string   `protobuf:"bytes,3,opt,name=function,proto3" json:"function,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppengineSourceLocation) Reset()         { *m = AppengineSourceLocation{} }
func (m *AppengineSourceLocation) String() string { return proto.CompactTextString(m) }
func (*AppengineSourceLocation) ProtoMessage()    {}
func (*AppengineSourceLocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_e181ce6b77c4050e, []int{1}
}

func (m *AppengineSourceLocation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppengineSourceLocation.Unmarshal(m, b)
}
func (m *AppengineSourceLocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppengineSourceLocation.Marshal(b, m, deterministic)
}
func (m *AppengineSourceLocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppengineSourceLocation.Merge(m, src)
}
func (m *AppengineSourceLocation) XXX_Size() int {
	return xxx_messageInfo_AppengineSourceLocation.Size(m)
}
func (m *AppengineSourceLocation) XXX_DiscardUnknown() {
	xxx_messageInfo_AppengineSourceLocation.DiscardUnknown(m)
}

var xxx_messageInfo_AppengineSourceLocation proto.InternalMessageInfo

func (m *AppengineSourceLocation) GetFile() string {
	if m != nil {
		return m.File
	}
	return ""
}

func (m *AppengineSourceLocation) GetLine() string {
	if m != nil {
		return m.Line
	}
	return ""
}

func (m *AppengineSourceLocation) GetFunction() string {
	if m != nil {
		return m.Function
	}
	return ""
}

func init() {
	proto.RegisterType((*AppengineEntry)(nil), "michilu.boilerplate.service.slog.AppengineEntry")
	proto.RegisterType((*AppengineSourceLocation)(nil), "michilu.boilerplate.service.slog.AppengineSourceLocation")
}

func init() { proto.RegisterFile("service/slog/appengine.proto", fileDescriptor_e181ce6b77c4050e) }

var fileDescriptor_e181ce6b77c4050e = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0xa9, 0x76, 0xab, 0x1b, 0x61, 0x85, 0x20, 0x18, 0xc4, 0x43, 0xd9, 0xd3, 0x9e, 0xb2,
	0xa0, 0x27, 0x8f, 0x0a, 0xde, 0x3c, 0xd5, 0x9b, 0x20, 0x92, 0xc6, 0xd9, 0x3a, 0x90, 0x26, 0x25,
	0x49, 0x17, 0xf6, 0x1f, 0xfb, 0x33, 0x64, 0xb2, 0x6d, 0xc0, 0x83, 0x78, 0x7b, 0x6f, 0x92, 0x37,
	0x2f, 0x5f, 0xd8, 0x6d, 0x00, 0xbf, 0x47, 0x0d, 0xdb, 0x60, 0x5c, 0xb7, 0x55, 0xc3, 0x00, 0xb6,
	0x43, 0x0b, 0x72, 0xf0, 0x2e, 0x3a, 0x5e, 0xf7, 0xa8, 0xbf, 0xd0, 0x8c, 0xb2, 0x75, 0x68, 0xc0,
	0x0f, 0x46, 0x45, 0x90, 0x53, 0x42, 0x52, 0x62, 0xfd, 0x5d, 0xb0, 0xd5, 0xe3, 0x9c, 0x7a, 0xb6,
	0xd1, 0x1f, 0x38, 0x67, 0x65, 0xc4, 0x1e, 0x44, 0x51, 0x17, 0x9b, 0x65, 0x93, 0x34, 0xbf, 0x62,
	0x8b, 0xe8, 0x95, 0x06, 0x71, 0x92, 0x86, 0x47, 0xc3, 0x5b, 0x76, 0x19, 0xdc, 0xe8, 0x35, 0x7c,
	0x18, 0xa7, 0x55, 0x44, 0x67, 0xc5, 0x69, 0x5d, 0x6c, 0x2e, 0xee, 0x1e, 0xe4, 0x7f, 0xc5, 0x32,
	0x97, 0xbe, 0xa6, 0x0d, 0x2f, 0xd3, 0x82, 0x66, 0x15, 0x7e, 0x79, 0x7e, 0xc3, 0xce, 0x03, 0xec,
	0xc1, 0x63, 0x3c, 0x88, 0x32, 0x95, 0x67, 0xcf, 0x05, 0x3b, 0xeb, 0x21, 0x04, 0xd5, 0x81, 0x58,
	0xa4, 0xa3, 0xd9, 0x12, 0xc3, 0xa7, 0x8a, 0x4a, 0x54, 0x47, 0x06, 0xd2, 0xeb, 0x77, 0x76, 0xfd,
	0x47, 0x29, 0x5d, 0xdf, 0xa1, 0xc9, 0xc8, 0xa4, 0x69, 0x66, 0xd0, 0xce, 0xc4, 0x49, 0xd3, 0x63,
	0x76, 0xa3, 0xd5, 0x99, 0x74, 0xd9, 0x64, 0xff, 0x54, 0xbd, 0x95, 0x04, 0xd6, 0x56, 0xe9, 0xeb,
	0xef, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xbf, 0x4e, 0x12, 0x14, 0x9a, 0x01, 0x00, 0x00,
}
