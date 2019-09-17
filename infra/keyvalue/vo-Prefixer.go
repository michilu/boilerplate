// Created by interfacer; DO NOT EDIT

package keyvalue

import (
	"github.com/golang/protobuf/proto"
	"github.com/rs/zerolog"
)

// Prefixer is an interface generated for "github.com/michilu/boilerplate/infra/keyvalue.Prefix".
type Prefixer interface {
	Descriptor() ([]byte, []int)
	GetPrefix() []byte
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
