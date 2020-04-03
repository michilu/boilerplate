package slog

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"

	"cloud.google.com/go/logging"
	k "github.com/michilu/boilerplate/application/config"
	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/now"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"github.com/valyala/fastjson"
	"google.golang.org/grpc/codes"
)

var (
	_stderrMu sync.Mutex
	_stdoutMu sync.Mutex
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
		stdout:    os.Stdout,
		stderr:    os.Stderr,
		projectID: v0,
	}
	return v1, nil
}

// AppengineLogging accepts pre-encoded JSON messages and writes them to Google Appengine Logging
// and maps Zerolog levels to Appengine levels.
// The labels argument is ignored if opts includes CommonLabels.
// The returned client should be closed before the program exits.
type AppengineLoggingWriter struct {
	stdout          *os.File
	stderr          *os.File
	parentProjects  string
	projectID       string
	traceIDTemplate string
}

// Write always returns len(p), nil.
func (w *AppengineLoggingWriter) Write(p []byte) (int, error) {
	const op = op + ".AppengineLoggingWriter.Write"
	v0 := NewAppengineEntry(p)
	v1, err := json.Marshal(v0)
	if err != nil {
		const op = op + ".json.Marshal"
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
		return 0, err
	}
	{
		_stdoutMu.Lock()
		defer _stdoutMu.Unlock()
	}
	if _, err := fmt.Fprintf(w.stdout, "%s\n", v1); err != nil {
		const op = op + ".io.File.Write"
		err := &errs.Error{Op: op, Code: codes.Internal, Err: err}
		return 0, err
	}
	return len(p), nil
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

	v0 := NewAppengineEntry(p)
	v0.Severity = severity.String()
	v1, err := json.Marshal(v0)
	if err != nil {
		const op = op + ".json.Marshal"
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
		return 0, err
	}
	v2 := w.stdout
	if logging.Warning < severity {
		v2 = w.stderr
		_stderrMu.Lock()
		defer _stderrMu.Unlock()
	} else {
		_stdoutMu.Lock()
		defer _stdoutMu.Unlock()
	}
	if _, err := fmt.Fprintf(v2, "%s\n", v1); err != nil {
		const op = op + ".io.File.Write"
		err := &errs.Error{Op: op, Code: codes.Internal, Err: err}
		return 0, err
	}
	return len(p), nil
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

// https://github.com/yfuruyama/stackdriver-request-context-log/blob/9427d3313129102b89c613ebf0c60be10f52a5ee/stackdriver.go#L263-L298

//go:generate interfacer -for github.com/michilu/boilerplate/service/slog.AppengineEntry -as AppengineEntryer -o vo-AppengineEntryer.go

func NewAppengineEntry(p []byte) *AppengineEntry {
	// get source location
	var location SourceLocation
	if pc, file, line, ok := runtime.Caller(6); ok {
		if function := runtime.FuncForPC(pc); function != nil {
			location.Function = function.Name()
		}
		location.Line = fmt.Sprintf("%d", line)
		parts := strings.Split(file, "/")
		location.File = parts[len(parts)-1] // use short file name
	}

	v0 := &AppengineEntry{
		Time:           now.Now().Format(time.RFC3339Nano),
		SourceLocation: &location,
	}
	v1 := rawJSON(p)
	v2, err := fastjson.ParseBytes(v1)
	if err == nil {
		v0.Trace = string(v2.GetStringBytes("trace"))
		v0.Message = string(v2.GetStringBytes("message"))
		v0.Data = string(p)
	}
	return v0
}
