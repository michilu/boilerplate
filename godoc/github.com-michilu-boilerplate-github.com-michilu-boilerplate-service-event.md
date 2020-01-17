# event
--
    import "github.com/michilu/boilerplate/github.com/michilu/boilerplate/service/event"


## Usage

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

#### func (*Event) Descriptor

```go
func (*Event) Descriptor() ([]byte, []int)
```

#### func (*Event) GetId

```go
func (m *Event) GetId() []byte
```

#### func (*Event) GetOrigin

```go
func (m *Event) GetOrigin() string
```

#### func (*Event) GetTimePoint

```go
func (m *Event) GetTimePoint() []*TimePoint
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
