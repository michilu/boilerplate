# debug
--
    import "github.com/michilu/boilerplate/github.com/michilu/boilerplate/service/debug"


## Usage

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
