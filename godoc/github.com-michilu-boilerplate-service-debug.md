# debug
--
    import "github.com/michilu/boilerplate/service/debug"


## Usage

#### func  NewID

```go
func NewID() (string, error)
```

#### type Client

```go
type Client struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```

Client is Entity

#### func (*Client) Descriptor

```go
func (*Client) Descriptor() ([]byte, []int)
```

#### func (*Client) GetId

```go
func (m *Client) GetId() string
```

#### func (*Client) MarshalZerologObject

```go
func (p *Client) MarshalZerologObject(e *zerolog.Event)
```

#### func (*Client) ProtoMessage

```go
func (*Client) ProtoMessage()
```

#### func (*Client) Reset

```go
func (m *Client) Reset()
```

#### func (*Client) String

```go
func (m *Client) String() string
```

#### func (*Client) Validate

```go
func (m *Client) Validate() error
```
Validate checks the field values on Client with the rules defined in the proto
definition for this message. If any rules are violated, an error is returned.

#### func (*Client) XXX_DiscardUnknown

```go
func (m *Client) XXX_DiscardUnknown()
```

#### func (*Client) XXX_Marshal

```go
func (m *Client) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*Client) XXX_Merge

```go
func (m *Client) XXX_Merge(src proto.Message)
```

#### func (*Client) XXX_Size

```go
func (m *Client) XXX_Size() int
```

#### func (*Client) XXX_Unmarshal

```go
func (m *Client) XXX_Unmarshal(b []byte) error
```

#### type ClientRepository

```go
type ClientRepository interface {
	Config(context.Context) (ClientWithContexter, error)
	Connect(ClientWithContexter) error
}
```


#### type ClientValidationError

```go
type ClientValidationError struct {
}
```

ClientValidationError is the validation error returned by Client.Validate if the
designated constraints aren't met.

#### func (ClientValidationError) Cause

```go
func (e ClientValidationError) Cause() error
```
Cause function returns cause value.

#### func (ClientValidationError) Error

```go
func (e ClientValidationError) Error() string
```
Error satisfies the builtin error interface

#### func (ClientValidationError) ErrorName

```go
func (e ClientValidationError) ErrorName() string
```
ErrorName returns error name.

#### func (ClientValidationError) Field

```go
func (e ClientValidationError) Field() string
```
Field function returns field value.

#### func (ClientValidationError) Key

```go
func (e ClientValidationError) Key() bool
```
Key function returns key value.

#### func (ClientValidationError) Reason

```go
func (e ClientValidationError) Reason() string
```
Reason function returns reason value.

#### type ClientWithContext

```go
type ClientWithContext struct {
	Context context.Context
	Client  Clienter
}
```

ClientWithContext is Client with context.Context.

#### func (*ClientWithContext) GetClient

```go
func (p *ClientWithContext) GetClient() Clienter
```
GetClient returns Clienter.

#### func (*ClientWithContext) GetContext

```go
func (p *ClientWithContext) GetContext() context.Context
```
GetContext returns context.Context.

#### func (*ClientWithContext) MarshalZerologObject

```go
func (p *ClientWithContext) MarshalZerologObject(e *zerolog.Event)
```
MarshalZerologObject writes ClientWithContext to given zerolog.Event.

#### func (*ClientWithContext) String

```go
func (p *ClientWithContext) String() string
```
String returns ClientWithContext as string.

#### func (*ClientWithContext) Validate

```go
func (p *ClientWithContext) Validate() error
```
Validate returns error if failed validate.

#### type ClientWithContexter

```go
type ClientWithContexter interface {
	GetClient() Clienter
	GetContext() context.Context
	MarshalZerologObject(*zerolog.Event)
	String() string
	Validate() error
}
```

ClientWithContexter is an interface generated for
"github.com/michilu/boilerplate/service/debug.ClientWithContext".

#### type Clienter

```go
type Clienter interface {
	Descriptor() ([]byte, []int)
	GetId() string
	MarshalZerologObject(*zerolog.Event)
	ProtoMessage()
	Reset()
	String() string
	Validate() error
	XXX_DiscardUnknown()
	XXX_Marshal([]byte, bool) ([]byte, error)
	XXX_Merge(proto.Message)
	XXX_Size() int
	XXX_Unmarshal([]byte) error
}
```

Clienter is an interface generated for
"github.com/michilu/boilerplate/service/debug.Client".

#### type TopicClientWithContexter

```go
type TopicClientWithContexter interface {
	// Publish returns a '<-chan ClientWithContexter' that joins to the given topic.
	Publish(ctx context.Context, c <-chan ClientWithContexter)
	// Publisher returns a 'chan<- ClientWithContexter' that joins to the given topic.
	Publisher(ctx context.Context) chan<- ClientWithContexter
	// Subscribe returns a 'chan<- ClientWithContexter' that joins to the given topic.
	Subscribe(c chan<- ClientWithContexter)
}
```

TopicClientWithContexter is a topic.

#### func  GetTopicClientWithContexter

```go
func GetTopicClientWithContexter(topic interface{}) TopicClientWithContexter
```
GetTopicClientWithContexter returns a TopicClientWithContexter of the given
topic.

#### type TopicClienter

```go
type TopicClienter interface {
	// Publish returns a '<-chan Clienter' that joins to the given topic.
	Publish(ctx context.Context, c <-chan Clienter)
	// Publisher returns a 'chan<- Clienter' that joins to the given topic.
	Publisher(ctx context.Context) chan<- Clienter
	// Subscribe returns a 'chan<- Clienter' that joins to the given topic.
	Subscribe(c chan<- Clienter)
}
```

TopicClienter is a topic.

#### func  GetTopicClienter

```go
func GetTopicClienter(topic interface{}) TopicClienter
```
GetTopicClienter returns a TopicClienter of the given topic.
