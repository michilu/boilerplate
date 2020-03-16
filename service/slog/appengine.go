package slog

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"cloud.google.com/go/logging"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	structpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/michilu/boilerplate/service/errs"
	"github.com/rs/zerolog"
	logtypepb "google.golang.org/genproto/googleapis/logging/type"
	logpb "google.golang.org/genproto/googleapis/logging/v2"
	"google.golang.org/grpc/codes"
)

var (
	// Now returns a time.Time.
	Now func() time.Time = time.Now
)

// NewAppengineLogging returns a new AppengineLoggingWriter.
func NewAppengineLogging(
	ctx context.Context,
	projectID string,
	logID string,
) (*AppengineLoggingWriter, error) {
	const op = op + ".NewAppengineLogging"
	if ctx == nil {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"}
		return nil, err
	}
	if projectID == "" {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'project' is ''"}
		return nil, err
	}
	if logID == "" {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'project' is ''"}
		return nil, err
	}
	v0 := &AppengineLoggingWriter{
		logger:    os.Stdout,
		projectID: projectID,
	}
	if ok := _logID.MatchString(logID); !ok {
		const op = op + ".Regexp.MatchString"
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: fmt.Sprintf("must be %v", _reLogID)}
		return nil, err
	}
	return v0, nil
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
	v1, err := w.toLogEntry(*v0)
	if err != nil {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: fmt.Sprintf("must be %v", _reLogID)}
		return 0, err
	}
	proto.MarshalText(w.logger, v1)
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

	v0 := NewEntry(p)
	v0.Severity = severity
	v1, err := w.toLogEntry(*v0)
	if err != nil {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: fmt.Sprintf("must be %v", _reLogID)}
		return 0, err
	}
	proto.MarshalText(w.logger, v1)
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

// https://github.com/googleapis/google-cloud-go/blob/57a019f/logging/logging.go#L887-L942
func (l *AppengineLoggingWriter) toLogEntry(e logging.Entry) (*logpb.LogEntry, error) {
	const op = op + ".AppengineLoggingWriter.toLogEntry"
	if e.LogName != "" {
		return nil, errors.New("logging: Entry.LogName should be not be set when writing")
	}
	t := e.Timestamp
	if t.IsZero() {
		t = Now()
	}
	ts, err := ptypes.TimestampProto(t)
	if err != nil {
		return nil, err
	}
	ent := &logpb.LogEntry{
		Timestamp:      ts,
		Severity:       logtypepb.LogSeverity(e.Severity),
		InsertId:       e.InsertID,
		Operation:      e.Operation,
		Labels:         e.Labels,
		Trace:          e.Trace,
		SpanId:         e.SpanID,
		Resource:       e.Resource,
		SourceLocation: e.SourceLocation,
		TraceSampled:   e.TraceSampled,
	}
	switch p := e.Payload.(type) {
	case string:
		ent.Payload = &logpb.LogEntry_TextPayload{TextPayload: p}
	default:
		s, err := toProtoStruct(p)
		if err != nil {
			return nil, err
		}
		ent.Payload = &logpb.LogEntry_JsonPayload{JsonPayload: s}
	}
	return ent, nil
}

// https://github.com/googleapis/google-cloud-go/blob/57a019f/logging/logging.go#L730-L788

// toProtoStruct converts v, which must marshal into a JSON object,
// into a Google Struct proto.
func toProtoStruct(v interface{}) (*structpb.Struct, error) {
	// Fast path: if v is already a *structpb.Struct, nothing to do.
	if s, ok := v.(*structpb.Struct); ok {
		return s, nil
	}
	// v is a Go value that supports JSON marshalling. We want a Struct
	// protobuf. Some day we may have a more direct way to get there, but right
	// now the only way is to marshal the Go value to JSON, unmarshal into a
	// map, and then build the Struct proto from the map.
	var jb []byte
	var err error
	if raw, ok := v.(json.RawMessage); ok { // needed for Go 1.7 and below
		jb = []byte(raw)
	} else {
		jb, err = json.Marshal(v)
		if err != nil {
			return nil, fmt.Errorf("logging: json.Marshal: %v", err)
		}
	}
	var m map[string]interface{}
	err = json.Unmarshal(jb, &m)
	if err != nil {
		return nil, fmt.Errorf("logging: json.Unmarshal: %v", err)
	}
	return jsonMapToProtoStruct(m), nil
}

func jsonMapToProtoStruct(m map[string]interface{}) *structpb.Struct {
	fields := map[string]*structpb.Value{}
	for k, v := range m {
		fields[k] = jsonValueToStructValue(v)
	}
	return &structpb.Struct{Fields: fields}
}

func jsonValueToStructValue(v interface{}) *structpb.Value {
	switch x := v.(type) {
	case bool:
		return &structpb.Value{Kind: &structpb.Value_BoolValue{BoolValue: x}}
	case float64:
		return &structpb.Value{Kind: &structpb.Value_NumberValue{NumberValue: x}}
	case string:
		return &structpb.Value{Kind: &structpb.Value_StringValue{StringValue: x}}
	case nil:
		return &structpb.Value{Kind: &structpb.Value_NullValue{}}
	case map[string]interface{}:
		return &structpb.Value{Kind: &structpb.Value_StructValue{StructValue: jsonMapToProtoStruct(x)}}
	case []interface{}:
		var vals []*structpb.Value
		for _, e := range x {
			vals = append(vals, jsonValueToStructValue(e))
		}
		return &structpb.Value{Kind: &structpb.Value_ListValue{ListValue: &structpb.ListValue{Values: vals}}}
	default:
		panic(fmt.Sprintf("bad type %T for JSON value", v))
	}
}
