package slog

import (
	"github.com/rs/zerolog"

	"github.com/michilu/boilerplate/service/meta"
)

type (
	// HookMeta appends a meta information to zerolog.
	HookMeta struct{}
)

// Run runs the hook with the event.
func (HookMeta) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	e.Object("meta", meta.ZerologObject())
}
