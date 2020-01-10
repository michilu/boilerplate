# pipe
--
    import "github.com/michilu/boilerplate/service/pipe"


## Usage

#### func  ErrorHandler

```go
func ErrorHandler(ctx context.Context, err error) (returns bool)
```
ErrorHandler is an error handler with error level.

#### func  FatalErrorHandler

```go
func FatalErrorHandler(ctx context.Context, err error) (returns bool)
```
FatalErrorHandler is an error handler with fatal level.

#### func  GetFanoutName

```go
func GetFanoutName(
	ctx context.Context,
	fn func(InT) ([]OutT, error),
	fnErr func(context.Context, error) bool,
) (
	chan<- InT,
	<-chan OutT,
)
```
GetFanoutName returns new input(chan<- InT)/output(<-chan OutT) channels that
embedded the given 'func(InT) OutT'.

#### func  GetPipeName

```go
func GetPipeName(
	ctx context.Context,
	fn func(InT) (OutT, error),
	fnErr func(context.Context, error) bool,
) (
	chan<- InT,
	<-chan OutT,
)
```
GetPipeName returns new input(chan<- InT)/output(<-chan OutT) channels that
embedded the given 'func(InT) OutT'.

#### func  Init

```go
func Init(ctx context.Context)
```

#### type InT

```go
type InT generic.Type
```

InT is a placeholder for the genny.

#### type NameGetContexter

```go
type NameGetContexter interface {
	GetContext() context.Context
}
```


#### type OutT

```go
type OutT generic.Type
```

OutT is a placeholder for the genny.
