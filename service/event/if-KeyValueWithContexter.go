// Created by interfacer; DO NOT EDIT

package event

import (
	"context"
	"github.com/michilu/boilerplate/infra/keyvalue"
	"github.com/rs/zerolog"
)

// KeyValueWithContexter is an interface generated for "github.com/michilu/boilerplate/service/event.KeyValueWithContext".
type KeyValueWithContexter interface {
	GetContext() context.Context
	GetKeyValue() keyvalue.KeyValuer
	MarshalZerologObject(*zerolog.Event)
	String() string
	Validate() error
}
