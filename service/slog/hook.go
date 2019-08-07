package slog

import (
	"fmt"

	"github.com/rs/zerolog"
	"github.com/spf13/viper"

	"github.com/michilu/boilerplate/service/meta"
)

type (
	// HookMeta appends a meta information to zerolog.
	HookMeta struct{}
)

// Run runs the hook with the event.
func (HookMeta) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	if level != zerolog.DebugLevel {
		if viper.GetBool("ellipsis") {
			fmt.Print(".")
		}
	}
	e.Object("meta", meta.ZerologObject())
}
