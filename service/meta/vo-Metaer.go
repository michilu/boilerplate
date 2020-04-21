// Created by interfacer; DO NOT EDIT

package meta

import (
	"github.com/rs/zerolog"
	"google.golang.org/protobuf/runtime/protoiface"
	"time"
)

// Metaer is an interface generated for "github.com/michilu/boilerplate/service/meta.Meta".
type Metaer interface {
	Descriptor() ([]byte, []int)
	Flatten() map[string]interface{}
	GetBuild() time.Time
	GetChannel() string
	GetName() string
	GetRuntime() *Runtime
	GetSemver() string
	GetSerial() string
	GetVcs() *Vcs
	JSON() []byte
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
