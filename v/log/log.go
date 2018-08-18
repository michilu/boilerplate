package log

import (
	"log"
	"os"

	"github.com/rs/zerolog"
)

var (
	logger          zerolog.Logger
	timeFieldFormat string
)

func init() {
	setTimeFieldFormat()
	setDefaultLogger()
}

func setTimeFieldFormat() {
	timeFieldFormat = zerolog.TimeFieldFormat
	zerolog.TimeFieldFormat = ""
}

func setDefaultLogger() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	logger = zerolog.New(os.Stdout).With().
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
	switch l {
	case zerolog.DebugLevel:
		zerolog.TimeFieldFormat = timeFieldFormat
		logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().
			Timestamp().
			Logger()
		log.SetOutput(logger)
	}
	return nil
}

// Logger returns the root logger.
func Logger() *zerolog.Logger {
	return &logger
}
