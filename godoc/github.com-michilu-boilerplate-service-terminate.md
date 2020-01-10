# terminate
--
    import "github.com/michilu/boilerplate/service/terminate"


## Usage

#### func  GetFanoutTerminate

```go
func GetFanoutTerminate(
	ctx context.Context,
	fn func(context.Context) ([]context.Context, error),
	fnErr func(context.Context, error) bool,
) (
	chan<- context.Context,
	<-chan context.Context,
)
```
GetFanoutTerminate returns new input(chan<- ContextContext)/output(<-chan
ContextContext) channels that embedded the given 'func(ContextContext)
ContextContext'.

#### func  GetPipeTerminate

```go
func GetPipeTerminate(
	ctx context.Context,
	fn func(context.Context) (context.Context, error),
	fnErr func(context.Context, error) bool,
) (
	chan<- context.Context,
	<-chan context.Context,
)
```
GetPipeTerminate returns new input(chan<- ContextContext)/output(<-chan
ContextContext) channels that embedded the given 'func(ContextContext)
ContextContext'.

#### func  Terminate

```go
func Terminate(ctx context.Context) (context.Context, error)
```
Terminate is terminator.

#### type TerminateGetContexter

```go
type TerminateGetContexter interface {
	GetContext() context.Context
}
```


#### type TopicContextContext

```go
type TopicContextContext interface {
	// Publish returns a '<-chan ContextContext' that joins to the given topic.
	Publish(ctx context.Context, c <-chan context.Context)
	// Publisher returns a 'chan<- ContextContext' that joins to the given topic.
	Publisher(ctx context.Context) chan<- context.Context
	// Subscribe returns a 'chan<- ContextContext' that joins to the given topic.
	Subscribe(c chan<- context.Context)
}
```

TopicContextContext is a topic.

#### func  GetTopicContextContext

```go
func GetTopicContextContext(topic interface{}) TopicContextContext
```
GetTopicContextContext returns a TopicContextContext of the given topic.
