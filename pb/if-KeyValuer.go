// Created by interfacer; DO NOT EDIT

package pb

import (
	"github.com/rs/zerolog"
	"google.golang.org/protobuf/runtime/protoiface"
)

// KeyValuer is an interface generated for "github.com/michilu/boilerplate/pb.KeyValue".
type KeyValuer interface {
	Descriptor() ([]byte, []int)
	GetKey() []byte
	GetValue() []byte
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
