# event
--
    import "github.com/michilu/boilerplate/application/event"


## Usage

#### func  Dataflow

```go
func Dataflow(ctx context.Context) error
```

#### func  ErrorHandler

```go
func ErrorHandler(ctx context.Context, err error) bool
```
ErrorHandler ...

#### func  EventLogger

```go
func EventLogger(m event.EventWithContexter) (event.KeyValueWithContexter, error)
```

#### func  GetFanoutEventLogger

```go
func GetFanoutEventLogger(
	ctx context.Context,
	fn func(event.EventWithContexter) ([]event.KeyValueWithContexter, error),
	fnErr func(context.Context, error) bool,
) (
	chan<- event.EventWithContexter,
	<-chan event.KeyValueWithContexter,
)
```
GetFanoutEventLogger returns new input(chan<-
EventEventWithContexter)/output(<-chan EventKeyValueWithContexter) channels that
embedded the given 'func(EventEventWithContexter) EventKeyValueWithContexter'.

#### func  GetFanoutSaver

```go
func GetFanoutSaver(
	ctx context.Context,
	fn func(event.KeyValueWithContexter) ([]context.Context, error),
	fnErr func(context.Context, error) bool,
) (
	chan<- event.KeyValueWithContexter,
	<-chan context.Context,
)
```
GetFanoutSaver returns new input(chan<-
EventKeyValueWithContexter)/output(<-chan ContextContext) channels that embedded
the given 'func(EventKeyValueWithContexter) ContextContext'.

#### func  GetFanoutStart

```go
func GetFanoutStart(
	ctx context.Context,
	fn func(context.Context) ([]event.EventWithContexter, error),
	fnErr func(context.Context, error) bool,
) (
	chan<- context.Context,
	<-chan event.EventWithContexter,
)
```
GetFanoutStart returns new input(chan<- ContextContext)/output(<-chan
EventEventWithContexter) channels that embedded the given 'func(ContextContext)
EventEventWithContexter'.

#### func  GetPipeEventLogger

```go
func GetPipeEventLogger(
	ctx context.Context,
	fn func(event.EventWithContexter) (event.KeyValueWithContexter, error),
	fnErr func(context.Context, error) bool,
) (
	chan<- event.EventWithContexter,
	<-chan event.KeyValueWithContexter,
)
```
GetPipeEventLogger returns new input(chan<-
EventEventWithContexter)/output(<-chan EventKeyValueWithContexter) channels that
embedded the given 'func(EventEventWithContexter) EventKeyValueWithContexter'.

#### func  GetPipeSaver

```go
func GetPipeSaver(
	ctx context.Context,
	fn func(event.KeyValueWithContexter) (context.Context, error),
	fnErr func(context.Context, error) bool,
) (
	chan<- event.KeyValueWithContexter,
	<-chan context.Context,
)
```
GetPipeSaver returns new input(chan<- EventKeyValueWithContexter)/output(<-chan
ContextContext) channels that embedded the given
'func(EventKeyValueWithContexter) ContextContext'.

#### func  GetPipeStart

```go
func GetPipeStart(
	ctx context.Context,
	fn func(context.Context) (event.EventWithContexter, error),
	fnErr func(context.Context, error) bool,
) (
	chan<- context.Context,
	<-chan event.EventWithContexter,
)
```
GetPipeStart returns new input(chan<- ContextContext)/output(<-chan
EventEventWithContexter) channels that embedded the given 'func(ContextContext)
EventEventWithContexter'.

#### func  Start

```go
func Start(ctx context.Context) (event.EventWithContexter, error)
```

#### type EventLoggerGetContexter

```go
type EventLoggerGetContexter interface {
	GetContext() context.Context
}
```


#### type Saver

```go
type Saver struct {
	Saver event.Saver
}
```


#### func (*Saver) Save

```go
func (p *Saver) Save(m event.KeyValueWithContexter) (context.Context, error)
```

#### type SaverGetContexter

```go
type SaverGetContexter interface {
	GetContext() context.Context
}
```


#### type StartGetContexter

```go
type StartGetContexter interface {
	GetContext() context.Context
}
```
