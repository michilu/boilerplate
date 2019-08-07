package slog

import (
	"io"
	"os"

	"github.com/rs/zerolog"
)

func newConsoleWriter() io.Writer {
	return zerolog.ConsoleWriter{
		Out:     os.Stdout,
		NoColor: true,
	}
}
