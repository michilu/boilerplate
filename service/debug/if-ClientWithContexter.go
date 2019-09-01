// Created by interfacer; DO NOT EDIT

package debug

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/rs/zerolog"
)

// ClientWithContexter is an interface generated for "github.com/michilu/boilerplate/service/debug.ClientWithContext".
type ClientWithContexter interface {
	Descriptor() ([]byte, []int)
	GetContext() context.Context
	GetId() string
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
