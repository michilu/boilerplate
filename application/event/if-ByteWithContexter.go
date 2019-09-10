// Created by interfacer; DO NOT EDIT

package event

import (
	"context"
	"github.com/rs/zerolog"
)

// ByteWithContexter is an interface generated for "github.com/michilu/boilerplate/application/event.ByteWithContext".
type ByteWithContexter interface {
	GetByte() *Byte
	GetContext() context.Context
	MarshalZerologObject(*zerolog.Event)
	String() string
	Validate() error
}
