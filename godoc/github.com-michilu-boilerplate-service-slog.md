# slog
--
    import "github.com/michilu/boilerplate/service/slog"


## Usage

```go
var (
	Atrace = Trace
)
```

#### func  GetTraceIDTemplate

```go
func GetTraceIDTemplate() string
```

#### func  GetTraceURLTemplate

```go
func GetTraceURLTemplate() string
```

#### func  Logger

```go
func Logger() *zerolog.Logger
```
Logger returns the root logger.

#### func  NewEntry

```go
func NewEntry(p []byte) *logging.Entry
```

#### func  SetDefaultLogger

```go
func SetDefaultLogger(writer []io.Writer)
```
SetDefaultLogger sets up the zerolog.Logger

#### func  SetDefaultTracer

```go
func SetDefaultTracer(v Tracer)
```

#### func  SetTimeFieldFormat

```go
func SetTimeFieldFormat()
```
SetTimeFieldFormat sets up the zerolog.TimeFieldFormat

#### func  Trace

```go
func Trace(ctx context.Context, s *trace.Span) zerolog.LogObjectMarshaler
```

#### type Closer

```go
type Closer interface {
	Close() error
}
```


#### type HookMeta

```go
type HookMeta struct{}
```

HookMeta appends a meta information to zerolog.

#### func (HookMeta) Run

```go
func (HookMeta) Run(e *zerolog.Event, level zerolog.Level, msg string)
```
Run runs the hook with the event.

#### type StackdriverCloser

```go
type StackdriverCloser struct {
}
```


#### func (*StackdriverCloser) Close

```go
func (p *StackdriverCloser) Close() error
```

#### type StackdriverLoggingWriter

```go
type StackdriverLoggingWriter struct {
	Logger *logging.Logger
}
```

StackdriverLogging accepts pre-encoded JSON messages and writes them to Google
Stackdriver Logging and maps Zerolog levels to Stackdriver levels. The labels
argument is ignored if opts includes CommonLabels. The returned client should be
closed before the program exits.

#### func  NewStackdriverLogging

```go
func NewStackdriverLogging(
	ctx context.Context,
	projectID string,
	logID string,
	labels map[string]string,
	opts ...logging.LoggerOption,
) (*StackdriverLoggingWriter, *logging.Client, error)
```
NewStackdriverLogging returns a new StackdriverLoggingWriter.

#### func (*StackdriverLoggingWriter) Flush

```go
func (p *StackdriverLoggingWriter) Flush() error
```

#### func (*StackdriverLoggingWriter) GetParentProjects

```go
func (p *StackdriverLoggingWriter) GetParentProjects() string
```
GetParentProjects returns a string of parent projects.
https://godoc.org/cloud.google.com/go/logging#NewClient

#### func (*StackdriverLoggingWriter) GetTraceIDTemplate

```go
func (p *StackdriverLoggingWriter) GetTraceIDTemplate() string
```
GetTraceIDTemplate returns a template string of the stackdriver traceID.

#### func (*StackdriverLoggingWriter) GetTraceURLTemplate

```go
func (p *StackdriverLoggingWriter) GetTraceURLTemplate() string
```
GetTraceURLTemplate returns a template string of the stackdriver traces URL.

#### func (*StackdriverLoggingWriter) Write

```go
func (w *StackdriverLoggingWriter) Write(p []byte) (int, error)
```
Write always returns len(p), nil.

#### func (*StackdriverLoggingWriter) WriteLevel

```go
func (w *StackdriverLoggingWriter) WriteLevel(level zerolog.Level, p []byte) (int, error)
```
WriteLevel implements zerolog.LevelWriter. It always returns len(p), nil.

#### type StackdriverZerologWriter

```go
type StackdriverZerologWriter struct {
}
```


#### func  NewStackdriverZerologWriter

```go
func NewStackdriverZerologWriter(ctx context.Context) *StackdriverZerologWriter
```
NewStackdriverZerologWriter returns a new ZerologWriter.

#### func (*StackdriverZerologWriter) Gen

```go
func (p *StackdriverZerologWriter) Gen() ([]io.Writer, Closer, error)
```

#### func (*StackdriverZerologWriter) MarshalZerologObject

```go
func (p *StackdriverZerologWriter) MarshalZerologObject(e *zerolog.Event)
```

#### type TraceObject

```go
type TraceObject struct {
}
```

Trace is trace span handler for zerolog.

#### func (*TraceObject) MarshalZerologObject

```go
func (p *TraceObject) MarshalZerologObject(e *zerolog.Event)
```

#### type Tracer

```go
type Tracer interface {
	GetTraceIDTemplate() string
	GetTraceURLTemplate() string
}
```
