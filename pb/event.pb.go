// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/event.proto

package pb

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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
	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// origin is an identity of the operator who entered the event.
	Origin               string            `protobuf:"bytes,2,opt,name=origin,proto3" json:"origin,omitempty"`
	TimePoint            []*EventTimePoint `protobuf:"bytes,3,rep,name=time_point,json=timePoint,proto3" json:"time_point,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_37b0496b5d264a50, []int{0}
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

func (m *Event) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *Event) GetTimePoint() []*EventTimePoint {
	if m != nil {
		return m.TimePoint
	}
	return nil
}

// EventTimePoint is Entity.
type EventTimePoint struct {
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Tag                  string               `protobuf:"bytes,2,opt,name=tag,proto3" json:"tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *EventTimePoint) Reset()         { *m = EventTimePoint{} }
func (m *EventTimePoint) String() string { return proto.CompactTextString(m) }
func (*EventTimePoint) ProtoMessage()    {}
func (*EventTimePoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_37b0496b5d264a50, []int{1}
}

func (m *EventTimePoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventTimePoint.Unmarshal(m, b)
}
func (m *EventTimePoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventTimePoint.Marshal(b, m, deterministic)
}
func (m *EventTimePoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventTimePoint.Merge(m, src)
}
func (m *EventTimePoint) XXX_Size() int {
	return xxx_messageInfo_EventTimePoint.Size(m)
}
func (m *EventTimePoint) XXX_DiscardUnknown() {
	xxx_messageInfo_EventTimePoint.DiscardUnknown(m)
}

var xxx_messageInfo_EventTimePoint proto.InternalMessageInfo

func (m *EventTimePoint) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *EventTimePoint) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func init() {
	proto.RegisterType((*Event)(nil), "michilu.boilerplate.pb.Event")
	proto.RegisterType((*EventTimePoint)(nil), "michilu.boilerplate.pb.EventTimePoint")
}

func init() { proto.RegisterFile("pb/event.proto", fileDescriptor_37b0496b5d264a50) }

var fileDescriptor_37b0496b5d264a50 = []byte{
	// 283 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0x3f, 0x4f, 0xf3, 0x30,
	0x18, 0xc4, 0x65, 0xa7, 0xed, 0xdb, 0xb8, 0xaf, 0x2a, 0xe4, 0x01, 0x42, 0x96, 0x46, 0x1d, 0x50,
	0x96, 0x3a, 0x52, 0x59, 0x58, 0x89, 0xc4, 0x0c, 0x8a, 0x3a, 0xb1, 0xa0, 0x38, 0x79, 0x70, 0x1f,
	0x29, 0x89, 0xad, 0xe0, 0x44, 0x94, 0x8f, 0xd1, 0x8f, 0xdb, 0x09, 0x25, 0x21, 0x20, 0xfe, 0x6c,
	0xd6, 0x3d, 0xbf, 0x3b, 0xdf, 0xb1, 0xa5, 0x91, 0x11, 0xb4, 0x50, 0x59, 0x61, 0x6a, 0x6d, 0x35,
	0x3f, 0x2f, 0x31, 0xdb, 0x63, 0xd1, 0x08, 0xa9, 0xb1, 0x80, 0xda, 0x14, 0xa9, 0x05, 0x61, 0xa4,
	0x7f, 0xab, 0xd0, 0xee, 0x1b, 0x29, 0x32, 0x5d, 0x46, 0x50, 0xb5, 0xfa, 0x60, 0x6a, 0xfd, 0x7a,
	0x88, 0x7a, 0x53, 0xb6, 0x51, 0x50, 0x6d, 0xda, 0xb4, 0xc0, 0x3c, 0xb5, 0x10, 0xfd, 0x7a, 0x0c,
	0xd1, 0xfe, 0x4a, 0x69, 0xad, 0x0a, 0x18, 0x3c, 0xb2, 0x79, 0x8e, 0x2c, 0x96, 0xf0, 0x62, 0xd3,
	0xd2, 0x0c, 0xc0, 0xfa, 0x48, 0xd8, 0xf4, 0xae, 0xeb, 0xc2, 0x2f, 0x18, 0xc5, 0xdc, 0x23, 0x01,
	0x09, 0xff, 0xc7, 0xff, 0x4e, 0xf1, 0xe4, 0x8d, 0x9e, 0x91, 0x84, 0x62, 0xce, 0x57, 0x6c, 0xa6,
	0x6b, 0x54, 0x58, 0x79, 0x34, 0x20, 0xa1, 0xdb, 0x1f, 0xeb, 0xee, 0xf8, 0x21, 0xf3, 0x7b, 0xc6,
	0xba, 0xd8, 0x27, 0xa3, 0xb1, 0xb2, 0x9e, 0x13, 0x38, 0xe1, 0x62, 0x7b, 0x25, 0xfe, 0x1e, 0x25,
	0xfa, 0xcf, 0x76, 0x58, 0xc2, 0x43, 0x47, 0xc7, 0xf3, 0x53, 0x3c, 0x3d, 0x12, 0x3a, 0x27, 0x89,
	0x6b, 0x47, 0x71, 0x0d, 0x6c, 0xf9, 0x1d, 0xe3, 0x37, 0xcc, 0xfd, 0x6c, 0xde, 0x77, 0x5c, 0x6c,
	0x7d, 0x31, 0x6c, 0x13, 0xe3, 0x36, 0xb1, 0x1b, 0x89, 0xe4, 0x0b, 0xe6, 0x97, 0xcc, 0xb1, 0xa9,
	0xfa, 0x59, 0xbd, 0xd3, 0xe2, 0xc9, 0x23, 0x35, 0x52, 0xce, 0x7a, 0xff, 0xf5, 0x7b, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x5c, 0x6e, 0x38, 0x96, 0x96, 0x01, 0x00, 0x00,
}
