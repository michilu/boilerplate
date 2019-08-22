package slog

import (
	"io"
	"log"
	"os"

	"github.com/rs/zerolog"
	"github.com/spf13/viper"
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
func SetDefaultLogger(writer []io.Writer) {
	{
		level := zerolog.InfoLevel
		if viper.GetBool("service.slog.verbose") {
			level = zerolog.DebugLevel
		}
		zerolog.SetGlobalLevel(level)
	}
	{
		var w io.Writer = os.Stdout
		if viper.GetBool("service.slog.debug") {
			w = newConsoleWriter()
		} else {
			SetTimeFieldFormat()
		}
		if writer == nil {
			writer = []io.Writer{}
		}
		writer = append(writer, w)
		logger = zerolog.New(zerolog.MultiLevelWriter(writer...)).
			Hook(&HookMeta{}).With().Caller().Timestamp().Logger()
		log.SetOutput(logger)
	}
	log.SetFlags(0)
}

// Logger returns the root logger.
func Logger() *zerolog.Logger {
	return &logger
}
