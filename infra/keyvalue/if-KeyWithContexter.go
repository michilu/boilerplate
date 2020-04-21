// Created by interfacer; DO NOT EDIT

package keyvalue

import (
	"context"
	"github.com/michilu/boilerplate/pb"
	"github.com/rs/zerolog"
)

// KeyWithContexter is an interface generated for "github.com/michilu/boilerplate/infra/keyvalue.KeyWithContext".
type KeyWithContexter interface {
	GetContext() context.Context
	GetKey() pb.Keyer
	MarshalZerologObject(*zerolog.Event)
	String() string
	Validate() error
}
