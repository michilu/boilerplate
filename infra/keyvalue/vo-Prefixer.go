// Created by interfacer; DO NOT EDIT

package keyvalue

import (
	"github.com/rs/zerolog"
	"google.golang.org/protobuf/runtime/protoiface"
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
	XXX_Merge(protoiface.MessageV1)
	XXX_Size() int
	XXX_Unmarshal([]byte) error
}
