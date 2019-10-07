// Created by interfacer; DO NOT EDIT

package keyvalue

import (
	"github.com/golang/protobuf/proto"
	"github.com/rs/zerolog"
)

// Keyer is an interface generated for "github.com/michilu/boilerplate/infra/keyvalue.Key".
type Keyer interface {
	Descriptor() ([]byte, []int)
	GetKey() []byte
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
