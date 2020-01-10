# debug
--
    import "github.com/michilu/boilerplate/application/debug"


## Usage

#### func  Dataflow

```go
func Dataflow(ctx context.Context)
```

#### func  ErrorHandler

```go
func ErrorHandler(ctx context.Context, err error) bool
```
ErrorHandler ...

#### func  GenerateUUID

```go
func GenerateUUID(ctx context.Context) (string, error)
```

#### func  GetFanoutConfig

```go
func GetFanoutConfig(
	ctx context.Context,
	fn func(context.Context) ([]debug.ClientWithContexter, error),
	fnErr func(context.Context, error) bool,
) (
	chan<- context.Context,
	<-chan debug.ClientWithContexter,
)
```
GetFanoutConfig returns new input(chan<- ContextContext)/output(<-chan
DebugClientWithContexter) channels that embedded the given 'func(ContextContext)
DebugClientWithContexter'.

#### func  GetFanoutConnect

```go
func GetFanoutConnect(
	ctx context.Context,
	fn func(debug.ClientWithContexter) ([]context.Context, error),
	fnErr func(context.Context, error) bool,
) (
	chan<- debug.ClientWithContexter,
	<-chan context.Context,
)
```
GetFanoutConnect returns new input(chan<-
DebugClientWithContexter)/output(<-chan ContextContext) channels that embedded
the given 'func(DebugClientWithContexter) ContextContext'.

#### func  GetPipeConfig

```go
func GetPipeConfig(
	ctx context.Context,
	fn func(context.Context) (debug.ClientWithContexter, error),
	fnErr func(context.Context, error) bool,
) (
	chan<- context.Context,
	<-chan debug.ClientWithContexter,
)
```
GetPipeConfig returns new input(chan<- ContextContext)/output(<-chan
DebugClientWithContexter) channels that embedded the given 'func(ContextContext)
DebugClientWithContexter'.

#### func  GetPipeConnect

```go
func GetPipeConnect(
	ctx context.Context,
	fn func(debug.ClientWithContexter) (context.Context, error),
	fnErr func(context.Context, error) bool,
) (
	chan<- debug.ClientWithContexter,
	<-chan context.Context,
)
```
GetPipeConnect returns new input(chan<- DebugClientWithContexter)/output(<-chan
ContextContext) channels that embedded the given 'func(DebugClientWithContexter)
ContextContext'.

#### func  NewClientRepository

```go
func NewClientRepository() debug.ClientRepository
```
NewClientRepository returns a new ClientRepository

#### func  OpenDebugPort

```go
func OpenDebugPort(ctx context.Context, m debug.Clienter) error
```

#### type Config

```go
type Config struct {
}
```

Config ...

#### func (*Config) Config

```go
func (p *Config) Config(ctx context.Context) (debug.ClientWithContexter, error)
```

#### func (*Config) Connect

```go
func (p *Config) Connect(m debug.ClientWithContexter) (context.Context, error)
```

#### type ConfigGetContexter

```go
type ConfigGetContexter interface {
	GetContext() context.Context
}
```


#### type Configer

```go
type Configer interface {
	Config(context.Context) (debug.ClientWithContexter, error)
	Connect(debug.ClientWithContexter) (context.Context, error)
}
```

Configer is an interface generated for
"github.com/michilu/boilerplate/application/debug.Config".

#### func  NewConfiger

```go
func NewConfiger(ctx context.Context, clientRepo debug.ClientRepository) Configer
```
NewConfiger returns a new Configer

#### type ConnectGetContexter

```go
type ConnectGetContexter interface {
	GetContext() context.Context
}
```
