// Created by interfacer; DO NOT EDIT

package debug

import (
	"context"
	"github.com/golang/protobuf/proto"
)

// ClientWithCtxer is an interface generated for "github.com/michilu/boilerplate/service/debug.ClientWithCtx".
type ClientWithCtxer interface {
	Descriptor() ([]byte, []int)
	GetCtx() context.Context
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
