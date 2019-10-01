// Created by interfacer; DO NOT EDIT

package keyvalue

import (
	"context"
	"github.com/rs/zerolog"
)

// KeyWithContexter is an interface generated for "github.com/michilu/boilerplate/infra/keyvalue.KeyWithContext".
type KeyWithContexter interface {
	GetContext() context.Context
	GetKey() Keyer
	MarshalZerologObject(*zerolog.Event)
	String() string
	Validate() error
}
