// Created by interfacer; DO NOT EDIT

package client

import (
	"github.com/golang/protobuf/proto"
)

// Debuger is an interface generated for "github.com/michilu/boilerplate/service/client.Debug".
type Debuger interface {
	Descriptor() ([]byte, []int)
	GetPort() uint32
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
