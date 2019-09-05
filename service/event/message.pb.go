// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service/event/message.proto

package event

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/mwitkow/go-proto-validators"
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

// Event is Entity.
type Event struct {
	Id                   []byte       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TimePoint            []*TimePoint `protobuf:"bytes,2,rep,name=time_point,json=timePoint,proto3" json:"time_point,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_a73432f895dbcfa7, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Event) GetTimePoint() []*TimePoint {
	if m != nil {
		return m.TimePoint
	}
	return nil
}

// TimePoint is Entity.
type TimePoint struct {
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Tag                  string               `protobuf:"bytes,2,opt,name=tag,proto3" json:"tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TimePoint) Reset()         { *m = TimePoint{} }
func (m *TimePoint) String() string { return proto.CompactTextString(m) }
func (*TimePoint) ProtoMessage()    {}
func (*TimePoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_a73432f895dbcfa7, []int{1}
}

func (m *TimePoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimePoint.Unmarshal(m, b)
}
func (m *TimePoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimePoint.Marshal(b, m, deterministic)
}
func (m *TimePoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimePoint.Merge(m, src)
}
func (m *TimePoint) XXX_Size() int {
	return xxx_messageInfo_TimePoint.Size(m)
}
func (m *TimePoint) XXX_DiscardUnknown() {
	xxx_messageInfo_TimePoint.DiscardUnknown(m)
}

var xxx_messageInfo_TimePoint proto.InternalMessageInfo

func (m *TimePoint) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *TimePoint) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func init() {
	proto.RegisterType((*Event)(nil), "michilu.boilerplate.service.event.Event")
	proto.RegisterType((*TimePoint)(nil), "michilu.boilerplate.service.event.TimePoint")
}

func init() { proto.RegisterFile("service/event/message.proto", fileDescriptor_a73432f895dbcfa7) }

var fileDescriptor_a73432f895dbcfa7 = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8e, 0x41, 0x4f, 0xb3, 0x40,
	0x10, 0x86, 0x3f, 0x68, 0xda, 0x2f, 0x6c, 0x3d, 0x71, 0x30, 0x04, 0x0f, 0xc5, 0x9e, 0x38, 0xd8,
	0xdd, 0xa4, 0x26, 0xc6, 0x73, 0x13, 0xef, 0x06, 0x3d, 0x18, 0x2f, 0x75, 0x81, 0x71, 0x3b, 0x91,
	0xed, 0x10, 0x18, 0xe8, 0xcf, 0x35, 0xf1, 0x97, 0x18, 0x16, 0x8b, 0x47, 0x6f, 0xbb, 0x33, 0xef,
	0xf3, 0xce, 0x23, 0xae, 0x5a, 0x68, 0x7a, 0x2c, 0x40, 0x41, 0x0f, 0x47, 0x56, 0x16, 0xda, 0x56,
	0x1b, 0x90, 0x75, 0x43, 0x4c, 0xe1, 0xb5, 0xc5, 0xe2, 0x80, 0x55, 0x27, 0x73, 0xc2, 0x0a, 0x9a,
	0xba, 0xd2, 0x0c, 0xf2, 0x07, 0x90, 0x0e, 0x88, 0xef, 0x0c, 0xf2, 0xa1, 0xcb, 0x65, 0x41, 0x56,
	0xd9, 0x13, 0xf2, 0x07, 0x9d, 0x94, 0xa1, 0x8d, 0xe3, 0x37, 0xbd, 0xae, 0xb0, 0xd4, 0x4c, 0x4d,
	0xab, 0xa6, 0xe7, 0x58, 0x1d, 0xaf, 0x0c, 0x91, 0xa9, 0x40, 0xb9, 0x5f, 0xde, 0xbd, 0x2b, 0x46,
	0x0b, 0x2d, 0x6b, 0x5b, 0x8f, 0x81, 0x35, 0x8b, 0xf9, 0xc3, 0x70, 0x21, 0xbc, 0x14, 0x3e, 0x96,
	0x91, 0x97, 0x78, 0xe9, 0xc5, 0x6e, 0xf1, 0xf5, 0xb9, 0xf2, 0xeb, 0x7f, 0x99, 0x8f, 0x65, 0xf8,
	0x24, 0xc4, 0xc0, 0xec, 0x6b, 0xc2, 0x23, 0x47, 0x7e, 0x32, 0x4b, 0x97, 0xdb, 0x1b, 0xf9, 0xa7,
	0xb1, 0x7c, 0x46, 0x0b, 0x8f, 0x03, 0x33, 0xb6, 0xbd, 0x79, 0x59, 0xc0, 0xe7, 0xd1, 0x7a, 0x2f,
	0x82, 0x69, 0x1f, 0xde, 0x8b, 0x60, 0xb2, 0x72, 0x02, 0xcb, 0x6d, 0x2c, 0x47, 0x6f, 0x79, 0xf6,
	0x76, 0x75, 0x2e, 0x91, 0xfd, 0x86, 0xc3, 0x48, 0xcc, 0x58, 0x9b, 0xc8, 0x4f, 0xbc, 0x34, 0x18,
	0xcf, 0xbc, 0x78, 0xd9, 0x30, 0xda, 0xfd, 0x7f, 0x9d, 0x3b, 0x8d, 0x7c, 0xe1, 0x1a, 0x6e, 0xbf,
	0x03, 0x00, 0x00, 0xff, 0xff, 0xc7, 0xe5, 0x24, 0x4f, 0x81, 0x01, 0x00, 0x00,
}
