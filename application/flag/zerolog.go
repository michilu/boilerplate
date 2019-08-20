package flag

import (
	"github.com/rs/zerolog"
)

// ZerologObject returns a MarshalZerologObject.
func ZerologObject() zerolog.LogObjectMarshaler {
	return flag
}

func (p Flag) MarshalZerologObject(e *zerolog.Event) {
	e.
		Str("Config", p.Config).
		Int("Parallel", p.Parallel).
		Bool("Pprof", p.Pprof).
		Bool("Verbose", p.Verbose)
}
