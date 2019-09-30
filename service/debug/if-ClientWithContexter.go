// Created by interfacer; DO NOT EDIT

package debug

import (
	"context"
	"github.com/rs/zerolog"
)

// ClientWithContexter is an interface generated for "github.com/michilu/boilerplate/service/debug.ClientWithContext".
type ClientWithContexter interface {
	GetClient() Clienter
	GetContext() context.Context
	MarshalZerologObject(*zerolog.Event)
	String() string
	Validate() error
}
