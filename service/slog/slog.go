package slog

import (
	"log"

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
func SetDefaultLogger() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	logger = zerolog.New(zerolog.MultiLevelWriter(
		newConsoleWriter(),
	)).
		Hook(&HookMeta{}).
		With().
		Caller().
		Timestamp().
		Logger()
	log.SetOutput(logger)
	log.SetFlags(0)
}

// SetLevel sets the global override for log level.
func SetLevel(s string) error {
	l, err := zerolog.ParseLevel(s)
	if err != nil {
		return err
	}
	zerolog.SetGlobalLevel(l)
	return nil
}

// Logger returns the root logger.
func Logger() *zerolog.Logger {
	return &logger
}
