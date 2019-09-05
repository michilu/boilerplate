// Created by interfacer; DO NOT EDIT

package event

import (
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/rs/zerolog"
)

// Eventer is an interface generated for "github.com/michilu/boilerplate/service/event.Event".
type Eventer interface {
	AddTimePoint(string, time.Time) (Eventer, error)
	Descriptor() ([]byte, []int)
	GetId() []byte
	GetTimePoint() []*TimePoint
	MarshalZerologObject(*zerolog.Event)
	ProtoMessage()
	Reset()
	String() string
	Validate() error
	XXX_DiscardUnknown()
	XXX_Marshal([]byte, bool) ([]byte, error)
	XXX_Merge(proto.Message)
	XXX_Size() int
	XXX_Unmarshal([]byte) error
}
