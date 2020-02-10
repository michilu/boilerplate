# keyvalue
--
    import "github.com/michilu/boilerplate/infra/keyvalue"


## Usage

#### type Key

```go
type Key struct {
	Key                  []byte   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```

Key is Entity.

#### func (*Key) Descriptor

```go
func (*Key) Descriptor() ([]byte, []int)
```

#### func (*Key) GetKey

```go
func (m *Key) GetKey() []byte
```

#### func (*Key) MarshalZerologObject

```go
func (p *Key) MarshalZerologObject(e *zerolog.Event)
```

#### func (*Key) ProtoMessage

```go
func (*Key) ProtoMessage()
```

#### func (*Key) Reset

```go
func (m *Key) Reset()
```

#### func (*Key) String

```go
func (m *Key) String() string
```

#### func (*Key) Validate

```go
func (m *Key) Validate() error
```
Validate checks the field values on Key with the rules defined in the proto
definition for this message. If any rules are violated, an error is returned.

#### func (*Key) XXX_DiscardUnknown

```go
func (m *Key) XXX_DiscardUnknown()
```

#### func (*Key) XXX_Marshal

```go
func (m *Key) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*Key) XXX_Merge

```go
func (m *Key) XXX_Merge(src proto.Message)
```

#### func (*Key) XXX_Size

```go
func (m *Key) XXX_Size() int
```

#### func (*Key) XXX_Unmarshal

```go
func (m *Key) XXX_Unmarshal(b []byte) error
```

#### type KeyValidationError

```go
type KeyValidationError struct {
}
```

KeyValidationError is the validation error returned by Key.Validate if the
designated constraints aren't met.

#### func (KeyValidationError) Cause

```go
func (e KeyValidationError) Cause() error
```
Cause function returns cause value.

#### func (KeyValidationError) Error

```go
func (e KeyValidationError) Error() string
```
Error satisfies the builtin error interface

#### func (KeyValidationError) ErrorName

```go
func (e KeyValidationError) ErrorName() string
```
ErrorName returns error name.

#### func (KeyValidationError) Field

```go
func (e KeyValidationError) Field() string
```
Field function returns field value.

#### func (KeyValidationError) Key

```go
func (e KeyValidationError) Key() bool
```
Key function returns key value.

#### func (KeyValidationError) Reason

```go
func (e KeyValidationError) Reason() string
```
Reason function returns reason value.

#### type KeyValue

```go
type KeyValue struct {
	Key                  []byte   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```

KeyValue is Entity.

#### func (*KeyValue) Descriptor

```go
func (*KeyValue) Descriptor() ([]byte, []int)
```

#### func (*KeyValue) GetKey

```go
func (m *KeyValue) GetKey() []byte
```

#### func (*KeyValue) GetValue

```go
func (m *KeyValue) GetValue() []byte
```

#### func (*KeyValue) MarshalZerologObject

```go
func (p *KeyValue) MarshalZerologObject(e *zerolog.Event)
```

#### func (*KeyValue) ProtoMessage

```go
func (*KeyValue) ProtoMessage()
```

#### func (*KeyValue) Reset

```go
func (m *KeyValue) Reset()
```

#### func (*KeyValue) String

```go
func (m *KeyValue) String() string
```

#### func (*KeyValue) Validate

```go
func (m *KeyValue) Validate() error
```
Validate checks the field values on KeyValue with the rules defined in the proto
definition for this message. If any rules are violated, an error is returned.

#### func (*KeyValue) XXX_DiscardUnknown

```go
func (m *KeyValue) XXX_DiscardUnknown()
```

#### func (*KeyValue) XXX_Marshal

```go
func (m *KeyValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*KeyValue) XXX_Merge

```go
func (m *KeyValue) XXX_Merge(src proto.Message)
```

#### func (*KeyValue) XXX_Size

```go
func (m *KeyValue) XXX_Size() int
```

#### func (*KeyValue) XXX_Unmarshal

```go
func (m *KeyValue) XXX_Unmarshal(b []byte) error
```

#### type KeyValueCloser

```go
type KeyValueCloser interface {
	Close() error
	Delete(context.Context, Keyer) error
	Get(context.Context, Keyer) (KeyValuer, error)
	Put(context.Context, KeyValuer) error
}
```


#### type KeyValueValidationError

```go
type KeyValueValidationError struct {
}
```

KeyValueValidationError is the validation error returned by KeyValue.Validate if
the designated constraints aren't met.

#### func (KeyValueValidationError) Cause

```go
func (e KeyValueValidationError) Cause() error
```
Cause function returns cause value.

#### func (KeyValueValidationError) Error

```go
func (e KeyValueValidationError) Error() string
```
Error satisfies the builtin error interface

#### func (KeyValueValidationError) ErrorName

```go
func (e KeyValueValidationError) ErrorName() string
```
ErrorName returns error name.

#### func (KeyValueValidationError) Field

```go
func (e KeyValueValidationError) Field() string
```
Field function returns field value.

#### func (KeyValueValidationError) Key

```go
func (e KeyValueValidationError) Key() bool
```
Key function returns key value.

#### func (KeyValueValidationError) Reason

```go
func (e KeyValueValidationError) Reason() string
```
Reason function returns reason value.

#### type KeyValueWithContext

```go
type KeyValueWithContext struct {
	Context  context.Context
	KeyValue KeyValuer
}
```

KeyValueWithContext is KeyValue with context.Context.

#### func (*KeyValueWithContext) GetContext

```go
func (p *KeyValueWithContext) GetContext() context.Context
```
GetContext returns context.Context.

#### func (*KeyValueWithContext) GetKeyValue

```go
func (p *KeyValueWithContext) GetKeyValue() KeyValuer
```
GetKeyValue returns KeyValuer.

#### func (*KeyValueWithContext) MarshalZerologObject

```go
func (p *KeyValueWithContext) MarshalZerologObject(e *zerolog.Event)
```
MarshalZerologObject writes KeyValueWithContext to given zerolog.Event.

#### func (*KeyValueWithContext) String

```go
func (p *KeyValueWithContext) String() string
```
String returns KeyValueWithContext as string.

#### func (*KeyValueWithContext) Validate

```go
func (p *KeyValueWithContext) Validate() error
```
Validate returns error if failed validate.

#### type KeyValueWithContexter

```go
type KeyValueWithContexter interface {
	GetContext() context.Context
	GetKeyValue() KeyValuer
	MarshalZerologObject(*zerolog.Event)
	String() string
	Validate() error
}
```

KeyValueWithContexter is an interface generated for
"github.com/michilu/boilerplate/infra/keyvalue.KeyValueWithContext".

#### type KeyValuer

```go
type KeyValuer interface {
	Descriptor() ([]byte, []int)
	GetKey() []byte
	GetValue() []byte
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

KeyValuer is an interface generated for
"github.com/michilu/boilerplate/infra/keyvalue.KeyValue".

#### type KeyWithContext

```go
type KeyWithContext struct {
	Context context.Context
	Key     Keyer
}
```

KeyWithContext is Key with context.Context.

#### func (*KeyWithContext) GetContext

```go
func (p *KeyWithContext) GetContext() context.Context
```
GetContext returns context.Context.

#### func (*KeyWithContext) GetKey

```go
func (p *KeyWithContext) GetKey() Keyer
```
GetKey returns Keyer.

#### func (*KeyWithContext) MarshalZerologObject

```go
func (p *KeyWithContext) MarshalZerologObject(e *zerolog.Event)
```
MarshalZerologObject writes KeyWithContext to given zerolog.Event.

#### func (*KeyWithContext) String

```go
func (p *KeyWithContext) String() string
```
String returns KeyWithContext as string.

#### func (*KeyWithContext) Validate

```go
func (p *KeyWithContext) Validate() error
```
Validate returns error if failed validate.

#### type KeyWithContexter

```go
type KeyWithContexter interface {
	GetContext() context.Context
	GetKey() Keyer
	MarshalZerologObject(*zerolog.Event)
	String() string
	Validate() error
}
```

KeyWithContexter is an interface generated for
"github.com/michilu/boilerplate/infra/keyvalue.KeyWithContext".

#### type Keyer

```go
type Keyer interface {
	Descriptor() ([]byte, []int)
	GetKey() []byte
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

Keyer is an interface generated for
"github.com/michilu/boilerplate/infra/keyvalue.Key".

#### type LoadSaveCloser

```go
type LoadSaveCloser interface {
	Load(context.Context, Prefixer) (<-chan KeyValuer, error)
	Save(context.Context, KeyValuer) error
	//Delete(context.Context, Keyer) error
	Close() error
}
```


#### type Prefix

```go
type Prefix struct {
	Prefix               []byte   `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
```

Prefix is ValueObject.

#### func (*Prefix) Descriptor

```go
func (*Prefix) Descriptor() ([]byte, []int)
```

#### func (*Prefix) GetPrefix

```go
func (m *Prefix) GetPrefix() []byte
```

#### func (*Prefix) MarshalZerologObject

```go
func (p *Prefix) MarshalZerologObject(e *zerolog.Event)
```

#### func (*Prefix) ProtoMessage

```go
func (*Prefix) ProtoMessage()
```

#### func (*Prefix) Reset

```go
func (m *Prefix) Reset()
```

#### func (*Prefix) String

```go
func (m *Prefix) String() string
```

#### func (*Prefix) Validate

```go
func (m *Prefix) Validate() error
```
Validate checks the field values on Prefix with the rules defined in the proto
definition for this message. If any rules are violated, an error is returned.

#### func (*Prefix) XXX_DiscardUnknown

```go
func (m *Prefix) XXX_DiscardUnknown()
```

#### func (*Prefix) XXX_Marshal

```go
func (m *Prefix) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*Prefix) XXX_Merge

```go
func (m *Prefix) XXX_Merge(src proto.Message)
```

#### func (*Prefix) XXX_Size

```go
func (m *Prefix) XXX_Size() int
```

#### func (*Prefix) XXX_Unmarshal

```go
func (m *Prefix) XXX_Unmarshal(b []byte) error
```

#### type PrefixValidationError

```go
type PrefixValidationError struct {
}
```

PrefixValidationError is the validation error returned by Prefix.Validate if the
designated constraints aren't met.

#### func (PrefixValidationError) Cause

```go
func (e PrefixValidationError) Cause() error
```
Cause function returns cause value.

#### func (PrefixValidationError) Error

```go
func (e PrefixValidationError) Error() string
```
Error satisfies the builtin error interface

#### func (PrefixValidationError) ErrorName

```go
func (e PrefixValidationError) ErrorName() string
```
ErrorName returns error name.

#### func (PrefixValidationError) Field

```go
func (e PrefixValidationError) Field() string
```
Field function returns field value.

#### func (PrefixValidationError) Key

```go
func (e PrefixValidationError) Key() bool
```
Key function returns key value.

#### func (PrefixValidationError) Reason

```go
func (e PrefixValidationError) Reason() string
```
Reason function returns reason value.

#### type Prefixer

```go
type Prefixer interface {
	Descriptor() ([]byte, []int)
	GetPrefix() []byte
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

Prefixer is an interface generated for
"github.com/michilu/boilerplate/infra/keyvalue.Prefix".

#### type TopicKeyValueWithContexter

```go
type TopicKeyValueWithContexter interface {
	// Publish returns a '<-chan KeyValueWithContexter' that joins to the given topic.
	Publish(ctx context.Context, c <-chan KeyValueWithContexter)
	// Publisher returns a 'chan<- KeyValueWithContexter' that joins to the given topic.
	Publisher(ctx context.Context) chan<- KeyValueWithContexter
	// Subscribe returns a 'chan<- KeyValueWithContexter' that joins to the given topic.
	Subscribe(c chan<- KeyValueWithContexter)
}
```

TopicKeyValueWithContexter is a topic.

#### func  GetTopicKeyValueWithContexter

```go
func GetTopicKeyValueWithContexter(topic interface{}) TopicKeyValueWithContexter
```
GetTopicKeyValueWithContexter returns a TopicKeyValueWithContexter of the given
topic.

#### type TopicKeyWithContexter

```go
type TopicKeyWithContexter interface {
	// Publish returns a '<-chan KeyWithContexter' that joins to the given topic.
	Publish(ctx context.Context, c <-chan KeyWithContexter)
	// Publisher returns a 'chan<- KeyWithContexter' that joins to the given topic.
	Publisher(ctx context.Context) chan<- KeyWithContexter
	// Subscribe returns a 'chan<- KeyWithContexter' that joins to the given topic.
	Subscribe(c chan<- KeyWithContexter)
}
```

TopicKeyWithContexter is a topic.

#### func  GetTopicKeyWithContexter

```go
func GetTopicKeyWithContexter(topic interface{}) TopicKeyWithContexter
```
GetTopicKeyWithContexter returns a TopicKeyWithContexter of the given topic.
