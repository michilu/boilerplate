# keyvalue
--
    import "github.com/michilu/boilerplate/github.com/michilu/boilerplate/infra/keyvalue"


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
