// Created by interfacer; DO NOT EDIT

package debug

import (
	"github.com/golang/protobuf/proto"
)

// Clienter is an interface generated for "github.com/michilu/boilerplate/service/debug.Client".
type Clienter interface {
	Descriptor() ([]byte, []int)
	GetId() string
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
