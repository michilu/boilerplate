// Created by interfacer; DO NOT EDIT

package event

import (
	"context"
	"github.com/rs/zerolog"
)

// EventWithContexter is an interface generated for "github.com/michilu/boilerplate/service/event.EventWithContext".
type EventWithContexter interface {
	GetContext() context.Context
	GetEvent() Eventer
	MarshalZerologObject(*zerolog.Event)
	String() string
	Validate() error
}
