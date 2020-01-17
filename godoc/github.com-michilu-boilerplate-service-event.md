# event
--
    import "github.com/michilu/boilerplate/service/event"


## Usage

```go
const (
	// Occurred is a tag for the time the event occurred.
	Occurred = "occurred"
	// Entered is a tag for the time the event entered.
	Entered = "entered"
)
```
[DDD Reference - Domain Language](https://domainlanguage.com/ddd/reference/)

```go
var (
	GetTopicKeyValueWithContexter = keyvalue.GetTopicKeyValueWithContexter
)
```

#### func  SaveEventPayload

```go
func SaveEventPayload(ctx context.Context, repository Saver, keyvalue keyvalue.KeyValuer) error
```
SaveEventPayload saves an event payload.

#### func  StoreEvent

```go
func StoreEvent(ctx context.Context, message Message) ([]byte, error)
```
StoreEvent returns a bytes from given Marshaler.

#### type Closer

```go
type Closer interface {
	Close() error
}
```


#### type Event

```go
type Event struct {
	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// origin is an identity of the operator who entered the event.
	Origin               string       `protobuf:"bytes,2,opt,name=origin,proto3" json:"origin,omitempty"`
	TimePoint            []*TimePoint `protobuf:"bytes,3,rep,name=time_point,json=timePoint,proto3" json:"time_point,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}
```

Event is Entity.

#### func (*Event) AddTimePoint

```go
func (p *Event) AddTimePoint(tag string, timeStamp time.Time) (Eventer, error)
```
AddTimePoint returns a new Eventer with given the TimePoint.

#### func (*Event) Descriptor

```go
func (*Event) Descriptor() ([]byte, []int)
```

#### func (*Event) GetId

```go
func (m *Event) GetId() []byte
```

#### func (*Event) GetKey

```go
func (p *Event) GetKey() []byte
```

#### func (*Event) GetOrigin

```go
func (m *Event) GetOrigin() string
```

#### func (*Event) GetTimePoint

```go
func (m *Event) GetTimePoint() []*TimePoint
```

#### func (*Event) MarshalZerologObject

```go
func (p *Event) MarshalZerologObject(e *zerolog.Event)
```

#### func (*Event) ProtoMessage

```go
func (*Event) ProtoMessage()
```

#### func (*Event) Reset

```go
func (m *Event) Reset()
```

#### func (*Event) String

```go
func (m *Event) String() string
```

#### func (*Event) Validate

```go
func (m *Event) Validate() error
```
Validate checks the field values on Event with the rules defined in the proto
definition for this message. If any rules are violated, an error is returned.

#### func (*Event) XXX_DiscardUnknown

```go
func (m *Event) XXX_DiscardUnknown()
```

#### func (*Event) XXX_Marshal

```go
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*Event) XXX_Merge

```go
func (m *Event) XXX_Merge(src proto.Message)
```

#### func (*Event) XXX_Size

```go
func (m *Event) XXX_Size() int
```

#### func (*Event) XXX_Unmarshal

```go
func (m *Event) XXX_Unmarshal(b []byte) error
```

#### type EventValidationError

```go
type EventValidationError struct {
}
```

EventValidationError is the validation error returned by Event.Validate if the
designated constraints aren't met.

#### func (EventValidationError) Cause

```go
func (e EventValidationError) Cause() error
```
Cause function returns cause value.

#### func (EventValidationError) Error

```go
func (e EventValidationError) Error() string
```
Error satisfies the builtin error interface

#### func (EventValidationError) ErrorName

```go
func (e EventValidationError) ErrorName() string
```
ErrorName returns error name.

#### func (EventValidationError) Field

```go
func (e EventValidationError) Field() string
```
Field function returns field value.

#### func (EventValidationError) Key

```go
func (e EventValidationError) Key() bool
```
Key function returns key value.

#### func (EventValidationError) Reason

```go
func (e EventValidationError) Reason() string
```
Reason function returns reason value.

#### type EventWithContext

```go
type EventWithContext struct {
	Context context.Context
	Event   Eventer
}
```

EventWithContext is Event with context.Context.

#### func (*EventWithContext) GetContext

```go
func (p *EventWithContext) GetContext() context.Context
```
GetContext returns context.Context.

#### func (*EventWithContext) GetEvent

```go
func (p *EventWithContext) GetEvent() Eventer
```
GetEvent returns Eventer.

#### func (*EventWithContext) MarshalZerologObject

```go
func (p *EventWithContext) MarshalZerologObject(e *zerolog.Event)
```
MarshalZerologObject writes EventWithContext to given zerolog.Event.

#### func (*EventWithContext) String

```go
func (p *EventWithContext) String() string
```
String returns EventWithContext as string.

#### func (*EventWithContext) Validate

```go
func (p *EventWithContext) Validate() error
```
Validate returns error if failed validate.

#### type EventWithContexter

```go
type EventWithContexter interface {
	GetContext() context.Context
	GetEvent() Eventer
	MarshalZerologObject(*zerolog.Event)
	String() string
	Validate() error
}
```

EventWithContexter is an interface generated for
"github.com/michilu/boilerplate/service/event.EventWithContext".

#### type Eventer

```go
type Eventer interface {
	AddTimePoint(string, time.Time) (Eventer, error)
	Descriptor() ([]byte, []int)
	GetId() []byte
	GetKey() []byte
	GetOrigin() string
	GetTimePoint() []*TimePoint
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

Eventer is an interface generated for
"github.com/michilu/boilerplate/service/event.Event".

#### func  NewEvent

```go
func NewEvent(timeStamp *time.Time, origin string) (Eventer, error)
```
NewEvent returns a timestamp for the time the event occurred.

#### func  RestoreEvent

```go
func RestoreEvent(ctx context.Context, b []byte) (Eventer, error)
```
RestoreEvent returns an Eventer from given bytes.

#### type KeyValueWithContext

```go
type KeyValueWithContext = keyvalue.KeyValueWithContext
```


#### type KeyValueWithContexter

```go
type KeyValueWithContexter = keyvalue.KeyValueWithContexter
```


#### type Keyer

```go
type Keyer interface {
	GetKey() string
	zerolog.LogObjectMarshaler
}
```


#### type Loader

```go
type Loader interface {
	Load(context.Context, keyvalue.Prefixer) (<-chan keyvalue.KeyValuer, error)
}
```


#### type Message

```go
type Message interface {
	proto.Message
	zerolog.LogObjectMarshaler
}
```


#### type Saver

```go
type Saver interface {
	Save(context.Context, keyvalue.KeyValuer) error
}
```


#### type TimePoint

```go
type TimePoint struct {
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Tag                  string               `protobuf:"bytes,2,opt,name=tag,proto3" json:"tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}
```

TimePoint is Entity.

#### func (*TimePoint) Descriptor

```go
func (*TimePoint) Descriptor() ([]byte, []int)
```

#### func (*TimePoint) GetTag

```go
func (m *TimePoint) GetTag() string
```

#### func (*TimePoint) GetTimestamp

```go
func (m *TimePoint) GetTimestamp() *timestamp.Timestamp
```

#### func (*TimePoint) ProtoMessage

```go
func (*TimePoint) ProtoMessage()
```

#### func (*TimePoint) Reset

```go
func (m *TimePoint) Reset()
```

#### func (*TimePoint) String

```go
func (m *TimePoint) String() string
```

#### func (*TimePoint) Validate

```go
func (m *TimePoint) Validate() error
```
Validate checks the field values on TimePoint with the rules defined in the
proto definition for this message. If any rules are violated, an error is
returned.

#### func (*TimePoint) XXX_DiscardUnknown

```go
func (m *TimePoint) XXX_DiscardUnknown()
```

#### func (*TimePoint) XXX_Marshal

```go
func (m *TimePoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*TimePoint) XXX_Merge

```go
func (m *TimePoint) XXX_Merge(src proto.Message)
```

#### func (*TimePoint) XXX_Size

```go
func (m *TimePoint) XXX_Size() int
```

#### func (*TimePoint) XXX_Unmarshal

```go
func (m *TimePoint) XXX_Unmarshal(b []byte) error
```

#### type TimePointValidationError

```go
type TimePointValidationError struct {
}
```

TimePointValidationError is the validation error returned by TimePoint.Validate
if the designated constraints aren't met.

#### func (TimePointValidationError) Cause

```go
func (e TimePointValidationError) Cause() error
```
Cause function returns cause value.

#### func (TimePointValidationError) Error

```go
func (e TimePointValidationError) Error() string
```
Error satisfies the builtin error interface

#### func (TimePointValidationError) ErrorName

```go
func (e TimePointValidationError) ErrorName() string
```
ErrorName returns error name.

#### func (TimePointValidationError) Field

```go
func (e TimePointValidationError) Field() string
```
Field function returns field value.

#### func (TimePointValidationError) Key

```go
func (e TimePointValidationError) Key() bool
```
Key function returns key value.

#### func (TimePointValidationError) Reason

```go
func (e TimePointValidationError) Reason() string
```
Reason function returns reason value.

#### type TopicEventWithContexter

```go
type TopicEventWithContexter interface {
	// Publish returns a '<-chan EventWithContexter' that joins to the given topic.
	Publish(ctx context.Context, c <-chan EventWithContexter)
	// Publisher returns a 'chan<- EventWithContexter' that joins to the given topic.
	Publisher(ctx context.Context) chan<- EventWithContexter
	// Subscribe returns a 'chan<- EventWithContexter' that joins to the given topic.
	Subscribe(c chan<- EventWithContexter)
}
```

TopicEventWithContexter is a topic.

#### func  GetTopicEventWithContexter

```go
func GetTopicEventWithContexter(topic interface{}) TopicEventWithContexter
```
GetTopicEventWithContexter returns a TopicEventWithContexter of the given topic.
