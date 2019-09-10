// Created by interfacer; DO NOT EDIT

package event

import (
	"context"
	"github.com/michilu/boilerplate/service/event"
	"github.com/rs/zerolog"
)

// EventWithContexter is an interface generated for "github.com/michilu/boilerplate/application/event.EventWithContext".
type EventWithContexter interface {
	GetContext() context.Context
	GetEvent() event.Eventer
	MarshalZerologObject(*zerolog.Event)
	String() string
	Validate() error
}
