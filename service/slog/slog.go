package slog

import (
	"io"
	"log"
	"os"

	k "github.com/michilu/boilerplate/application/config"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

const (
	op = "service/slog"
)

var (
	logger          zerolog.Logger
	timeFieldFormat string
	tracer          Tracer
)

type Tracer interface {
	GetTraceIDTemplate() string
	GetTraceURLTemplate() string
}

// SetTimeFieldFormat sets up the zerolog.TimeFieldFormat
func SetTimeFieldFormat() {
	timeFieldFormat = zerolog.TimeFieldFormat
	zerolog.TimeFieldFormat = ""
}

// SetDefaultLogger sets up the zerolog.Logger
func SetDefaultLogger(writer []io.Writer) {
	{
		level := zerolog.InfoLevel
		if viper.GetBool(k.ServiceSlogVerbose) {
			level = zerolog.DebugLevel
		}
		zerolog.SetGlobalLevel(level)
	}
	{
		var w io.Writer = os.Stdout
		if viper.GetBool(k.ServiceSlogDebug) {
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

func SetDefaultTracer(v Tracer) {
	tracer = v
}

func GetTraceIDTemplate() string {
	if tracer == nil {
		return ""
	}
	return tracer.GetTraceIDTemplate()
}

func GetTraceURLTemplate() string {
	if tracer == nil {
		return ""
	}
	return tracer.GetTraceURLTemplate()
}
