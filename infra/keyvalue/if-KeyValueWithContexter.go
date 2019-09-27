// Created by interfacer; DO NOT EDIT

package keyvalue

import (
	"context"
	"github.com/rs/zerolog"
)

// KeyValueWithContexter is an interface generated for "github.com/michilu/boilerplate/infra/keyvalue.KeyValueWithContext".
type KeyValueWithContexter interface {
	GetContext() context.Context
	GetKeyValue() KeyValuer
	MarshalZerologObject(*zerolog.Event)
	String() string
	Validate() error
}
