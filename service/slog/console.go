package slog

import (
	"github.com/rs/zerolog"
)

type consoleWriter struct {
	zerolog.ConsoleWriter
}

func (c *consoleWriter) WriteLevel(l zerolog.Level, p []byte) (int, error) {
	switch l {
	case
		zerolog.FatalLevel,
		zerolog.PanicLevel:
		return c.Write(p)
	default:
	}
	if zerolog.DebugLevel < zerolog.GlobalLevel() {
		return len(p), nil
	}
	return c.Write(p)
}
