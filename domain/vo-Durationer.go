// Created by interfacer; DO NOT EDIT

package domain

import (
	"github.com/golang/protobuf/proto"
)

// Durationer is an interface generated for "github.com/michilu/boilerplate/domain.Duration".
type Durationer interface {
	Descriptor() ([]byte, []int)
	GetMillis() int32
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
