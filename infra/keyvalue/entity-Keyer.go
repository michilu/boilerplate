// Created by interfacer; DO NOT EDIT

package keyvalue

import (
	"github.com/golang/protobuf/proto"
)

// Keyer is an interface generated for "github.com/michilu/boilerplate/infra/keyvalue.Key".
type Keyer interface {
	Descriptor() ([]byte, []int)
	GetKey() []byte
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
