// +build !windows

package slog

import (
	"io"
	"os"

	"github.com/rs/zerolog"
)

func newConsoleWriter() io.Writer {
	return &consoleWriter{
		zerolog.ConsoleWriter{
			Out: os.Stdout,
		}}
}
