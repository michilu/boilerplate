package slog

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"cloud.google.com/go/logging"
	k "github.com/michilu/boilerplate/application/config"
	"github.com/michilu/boilerplate/service/errs"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"google.golang.org/grpc/codes"
)

var (
	// Now returns a time.Time.
	Now func() time.Time = time.Now
)

// NewAppengineLogging returns a new AppengineLoggingWriter.
func NewAppengineLogging(ctx context.Context) (*AppengineLoggingWriter, error) {
	const op = op + ".NewAppengineLogging"
	if ctx == nil {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"}
		return nil, err
	}
	v0 := viper.GetString(k.GoogleProjectId)
	v1 := &AppengineLoggingWriter{
		logger:    os.Stdout,
		projectID: v0,
	}
	return v1, nil
}

// AppengineLogging accepts pre-encoded JSON messages and writes them to Google Appengine Logging
// and maps Zerolog levels to Appengine levels.
// The labels argument is ignored if opts includes CommonLabels.
// The returned client should be closed before the program exits.
type AppengineLoggingWriter struct {
	logger          *os.File
	parentProjects  string
	projectID       string
	traceIDTemplate string
}

// Write always returns len(p), nil.
func (w *AppengineLoggingWriter) Write(p []byte) (int, error) {
	const op = op + ".AppengineLoggingWriter.Write"
	v0 := NewEntry(p)
	v1, err := json.Marshal(v0)
	if err != nil {
		const op = op + ".json.Marshal"
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
		return 0, err
	}
	if v2, err := fmt.Fprintf(w.logger, "%s\n", v1); err != nil {
		const op = op + ".io.File.Write"
		err := &errs.Error{Op: op, Code: codes.Internal, Err: err}
		return 0, err
	} else {
		return v2, nil
	}
}

// WriteLevel implements zerolog.LevelWriter. It always returns len(p), nil.
func (w *AppengineLoggingWriter) WriteLevel(level zerolog.Level, p []byte) (int, error) {
	const op = op + ".AppengineLoggingWriter.WriteLevel"
	severity := logging.Default

	// https://godoc.org/github.com/rs/zerolog#Level
	// https://godoc.org/cloud.google.com/go/logging#pkg-constants
	switch level {
	case zerolog.NoLevel:
		severity = logging.Default
	case zerolog.DebugLevel:
		severity = logging.Debug
	case zerolog.InfoLevel:
		severity = logging.Info
	case zerolog.WarnLevel:
		severity = logging.Warning
	case zerolog.ErrorLevel:
		severity = logging.Error
	case zerolog.FatalLevel:
		severity = logging.Critical
	case zerolog.PanicLevel:
		severity = logging.Alert
	}

	v0 := NewEntry(p)
	v0.Severity = severity
	v1, err := json.Marshal(v0)
	if err != nil {
		const op = op + ".json.Marshal"
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
		return 0, err
	}
	if v2, err := fmt.Fprintf(w.logger, "%s\n", v1); err != nil {
		const op = op + ".io.File.Write"
		err := &errs.Error{Op: op, Code: codes.Internal, Err: err}
		return 0, err
	} else {
		return v2, nil
	}
}

// GetTraceIDTemplate returns a template string of the stackdriver traceID.
func (p *AppengineLoggingWriter) GetTraceIDTemplate() string {
	return p.GetParentProjects() + "/traces/%s"
}

// GetTraceURLTemplate returns a template string of the stackdriver traces URL.
func (p *AppengineLoggingWriter) GetTraceURLTemplate() string {
	return "https://console.cloud.google.com/traces/traces?tid=%s"
}

// GetParentProjects returns a string of parent projects.
// https://godoc.org/cloud.google.com/go/logging#NewClient
func (p *AppengineLoggingWriter) GetParentProjects() string {
	return "projects/" + p.projectID
}
