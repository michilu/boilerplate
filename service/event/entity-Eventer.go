// Created by interfacer; DO NOT EDIT

package event

import (
	"github.com/rs/zerolog"
	"google.golang.org/protobuf/runtime/protoiface"
	"time"
)

// Eventer is an interface generated for "github.com/michilu/boilerplate/service/event.Event".
type Eventer interface {
	AddTimePoint(string, time.Time) (Eventer, error)
	Descriptor() ([]byte, []int)
	GetId() []byte
	GetKey() []byte
	GetOrigin() string
	GetTimePoint() []*TimePoint
	MarshalZerologObject(*zerolog.Event)
	ProtoMessage()
	Reset()
	String() string
	Validate() error
	XXX_DiscardUnknown()
	XXX_Marshal([]byte, bool) ([]byte, error)
	XXX_Merge(protoiface.MessageV1)
	XXX_Size() int
	XXX_Unmarshal([]byte) error
}
