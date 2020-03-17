// Created by interfacer; DO NOT EDIT

package slog

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/any"
)

// AppengineEntryer is an interface generated for "github.com/michilu/boilerplate/service/slog.AppengineEntry".
type AppengineEntryer interface {
	Descriptor() ([]byte, []int)
	GetData() map[string]*any.Any
	GetMessage() string
	GetSeverity() string
	GetSourceLocation() *SourceLocation
	GetTime() string
	GetTrace() string
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
