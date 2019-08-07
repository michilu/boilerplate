package meta

import (
	"github.com/rs/zerolog"
)

// ZerologObject returns a MarshalZerologObject.
func ZerologObject() zerolog.LogObjectMarshaler {
	return &m
}

func (m *Meta) MarshalZerologObject(e *zerolog.Event) {
	e.RawJSON("version", m.JSON())
}
