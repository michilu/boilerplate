package slog

import (
	"io"
	"log"
	"os"

	"github.com/rs/zerolog"
)

var (
	logger          zerolog.Logger
	timeFieldFormat string
)

// SetTimeFieldFormat sets up the zerolog.TimeFieldFormat
func SetTimeFieldFormat() {
	timeFieldFormat = zerolog.TimeFieldFormat
	zerolog.TimeFieldFormat = ""
}

// SetDefaultLogger sets up the zerolog.Logger
func SetDefaultLogger(verbose bool) {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	var w io.Writer = os.Stdout
	if !verbose {
		SetTimeFieldFormat()
	} else {
		w = newConsoleWriter()
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	logger = zerolog.New(zerolog.MultiLevelWriter(
		w,
	)).
		Hook(&HookMeta{}).
		With().
		Caller().
		Timestamp().
		Logger()
	log.SetOutput(logger)
	log.SetFlags(0)
}

// Logger returns the root logger.
func Logger() *zerolog.Logger {
	return &logger
}
