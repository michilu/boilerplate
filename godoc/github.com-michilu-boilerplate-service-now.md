# now
--
    import "github.com/michilu/boilerplate/service/now"


## Usage

```go
var (
	// Now returns a time.Time.
	Now func() time.Time = time.Now
)
```

#### func  ContextTicker

```go
func ContextTicker(ctx context.Context, duration time.Duration) (<-chan context.Context, error)
```

#### func  TimeFromTimestamp

```go
func TimeFromTimestamp(v *timestamp.Timestamp) (time.Time, error)
```
TimeFromTimestamp returns a time.Time of the given timestamp.Timestamp.

#### func  TimestampFromTime

```go
func TimestampFromTime(v time.Time) *timestamp.Timestamp
```
TimestampFromTime returns a timestamp.Timestamp of the given time.Time.

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
