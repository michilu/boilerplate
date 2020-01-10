# topic
--
    import "github.com/michilu/boilerplate/service/topic"


## Usage

#### type ChanT

```go
type ChanT generic.Type
```

ChanT is a placeholder for the genny.

#### type Ier

```go
type Ier generic.Type
```

Ier is a placeholder for the genny.

#### type T

```go
type T generic.Type
```

T is a placeholder for the genny.

#### type TWithContext

```go
type TWithContext struct {
	Context context.Context
	T       Ier
}
```

TWithContext is T with context.Context.

#### func (*TWithContext) GetContext

```go
func (p *TWithContext) GetContext() context.Context
```
GetContext returns context.Context.

#### func (*TWithContext) GetT

```go
func (p *TWithContext) GetT() Ier
```
GetT returns Ier.

#### func (*TWithContext) MarshalZerologObject

```go
func (p *TWithContext) MarshalZerologObject(e *zerolog.Event)
```
MarshalZerologObject writes TWithContext to given zerolog.Event.

#### func (*TWithContext) String

```go
func (p *TWithContext) String() string
```
String returns TWithContext as string.

#### func (*TWithContext) Validate

```go
func (p *TWithContext) Validate() error
```
Validate returns error if failed validate.

#### type TopicChanT

```go
type TopicChanT interface {
	// Publish returns a '<-chan ChanT' that joins to the given topic.
	Publish(ctx context.Context, c <-chan ChanT)
	// Publisher returns a 'chan<- ChanT' that joins to the given topic.
	Publisher(ctx context.Context) chan<- ChanT
	// Subscribe returns a 'chan<- ChanT' that joins to the given topic.
	Subscribe(c chan<- ChanT)
}
```

TopicChanT is a topic.

#### func  GetTopicChanT

```go
func GetTopicChanT(topic interface{}) TopicChanT
```
GetTopicChanT returns a TopicChanT of the given topic.
