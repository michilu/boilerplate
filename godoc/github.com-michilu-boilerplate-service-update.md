# update
--
    import "github.com/michilu/boilerplate/service/update"


## Usage

#### func  GetFanoutUpdate

```go
func GetFanoutUpdate(
	ctx context.Context,
	fn func(context.Context) ([]context.Context, error),
	fnErr func(context.Context, error) bool,
) (
	chan<- context.Context,
	<-chan context.Context,
)
```
GetFanoutUpdate returns new input(chan<- ContextContext)/output(<-chan
ContextContext) channels that embedded the given 'func(ContextContext)
ContextContext'.

#### func  GetPipeUpdate

```go
func GetPipeUpdate(
	ctx context.Context,
	fn func(context.Context) (context.Context, error),
	fnErr func(context.Context, error) bool,
) (
	chan<- context.Context,
	<-chan context.Context,
)
```
GetPipeUpdate returns new input(chan<- ContextContext)/output(<-chan
ContextContext) channels that embedded the given 'func(ContextContext)
ContextContext'.

#### func  Update

```go
func Update(ctx context.Context) (context.Context, error)
```
Update ...

go:generate genny -in=../pipe/pipe.go -out=gen-pipe-Update.go -pkg=$GOPACKAGE
gen "Name=update InT=context.Context OutT=context.Context"

#### type UpdateGetContexter

```go
type UpdateGetContexter interface {
	GetContext() context.Context
}
```
