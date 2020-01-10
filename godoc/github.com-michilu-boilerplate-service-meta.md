# meta
--
    import "github.com/michilu/boilerplate/service/meta"


## Usage

#### func  Name

```go
func Name() string
```
Name returns a name.

#### func  Set

```go
func Set(v *Meta) error
```
Set sets a meta data.

#### func  Yaml

```go
func Yaml() (string, error)
```
Yaml returns Meta as YAML.

#### func  ZerologObject

```go
func ZerologObject() zerolog.LogObjectMarshaler
```
ZerologObject returns a MarshalZerologObject.

#### type Meta

```go
type Meta struct {
	Name                 string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Semver               string    `protobuf:"bytes,2,opt,name=semver,proto3" json:"semver,omitempty"`
	Channel              string    `protobuf:"bytes,3,opt,name=channel,proto3" json:"channel,omitempty"`
	Runtime              *Runtime  `protobuf:"bytes,4,opt,name=runtime,proto3" json:"runtime,omitempty"`
	Serial               string    `protobuf:"bytes,5,opt,name=serial,proto3" json:"serial,omitempty"`
	Build                time.Time `protobuf:"bytes,6,opt,name=build,proto3" json:"build,omitempty"`
	Vcs                  *Vcs      `protobuf:"bytes,7,opt,name=vcs,proto3" json:"vcs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-" yaml:"-"`
	XXX_unrecognized     []byte    `json:"-" yaml:"-"`
	XXX_sizecache        int32     `json:"-" yaml:"-"`
}
```

Meta is ValueObject of meta infomation

#### func  Get

```go
func Get() Meta
```
Get returns a Meta.

#### func (*Meta) Descriptor

```go
func (*Meta) Descriptor() ([]byte, []int)
```

#### func (*Meta) GetBuild

```go
func (m *Meta) GetBuild() time.Time
```

#### func (*Meta) GetChannel

```go
func (m *Meta) GetChannel() string
```

#### func (*Meta) GetName

```go
func (m *Meta) GetName() string
```

#### func (*Meta) GetRuntime

```go
func (m *Meta) GetRuntime() *Runtime
```

#### func (*Meta) GetSemver

```go
func (m *Meta) GetSemver() string
```

#### func (*Meta) GetSerial

```go
func (m *Meta) GetSerial() string
```

#### func (*Meta) GetVcs

```go
func (m *Meta) GetVcs() *Vcs
```

#### func (Meta) JSON

```go
func (m Meta) JSON() []byte
```

#### func (*Meta) MarshalZerologObject

```go
func (m *Meta) MarshalZerologObject(e *zerolog.Event)
```

#### func (*Meta) ProtoMessage

```go
func (*Meta) ProtoMessage()
```

#### func (*Meta) Reset

```go
func (m *Meta) Reset()
```

#### func (*Meta) String

```go
func (m *Meta) String() string
```

#### func (*Meta) Validate

```go
func (this *Meta) Validate() error
```

#### func (*Meta) XXX_DiscardUnknown

```go
func (m *Meta) XXX_DiscardUnknown()
```

#### func (*Meta) XXX_Marshal

```go
func (m *Meta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*Meta) XXX_Merge

```go
func (m *Meta) XXX_Merge(src proto.Message)
```

#### func (*Meta) XXX_Size

```go
func (m *Meta) XXX_Size() int
```

#### func (*Meta) XXX_Unmarshal

```go
func (m *Meta) XXX_Unmarshal(b []byte) error
```

#### type Metaer

```go
type Metaer interface {
	Descriptor() ([]byte, []int)
	GetBuild() time.Time
	GetChannel() string
	GetName() string
	GetRuntime() *Runtime
	GetSemver() string
	GetSerial() string
	GetVcs() *Vcs
	JSON() []byte
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

Metaer is an interface generated for
"github.com/michilu/boilerplate/service/meta.Meta".

#### type Runtime

```go
type Runtime struct {
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Arch                 string   `protobuf:"bytes,2,opt,name=arch,proto3" json:"arch,omitempty"`
	Os                   string   `protobuf:"bytes,3,opt,name=os,proto3" json:"os,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" yaml:"-"`
	XXX_unrecognized     []byte   `json:"-" yaml:"-"`
	XXX_sizecache        int32    `json:"-" yaml:"-"`
}
```

Runtime is ValueObject of Runtime

#### func (*Runtime) Descriptor

```go
func (*Runtime) Descriptor() ([]byte, []int)
```

#### func (*Runtime) GetArch

```go
func (m *Runtime) GetArch() string
```

#### func (*Runtime) GetOs

```go
func (m *Runtime) GetOs() string
```

#### func (*Runtime) GetVersion

```go
func (m *Runtime) GetVersion() string
```

#### func (*Runtime) ProtoMessage

```go
func (*Runtime) ProtoMessage()
```

#### func (*Runtime) Reset

```go
func (m *Runtime) Reset()
```

#### func (*Runtime) String

```go
func (m *Runtime) String() string
```

#### func (*Runtime) Validate

```go
func (this *Runtime) Validate() error
```

#### func (*Runtime) XXX_DiscardUnknown

```go
func (m *Runtime) XXX_DiscardUnknown()
```

#### func (*Runtime) XXX_Marshal

```go
func (m *Runtime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*Runtime) XXX_Merge

```go
func (m *Runtime) XXX_Merge(src proto.Message)
```

#### func (*Runtime) XXX_Size

```go
func (m *Runtime) XXX_Size() int
```

#### func (*Runtime) XXX_Unmarshal

```go
func (m *Runtime) XXX_Unmarshal(b []byte) error
```

#### type Vcs

```go
type Vcs struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Branch               string   `protobuf:"bytes,2,opt,name=branch,proto3" json:"branch,omitempty"`
	Tag                  string   `protobuf:"bytes,3,opt,name=tag,proto3" json:"tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" yaml:"-"`
	XXX_unrecognized     []byte   `json:"-" yaml:"-"`
	XXX_sizecache        int32    `json:"-" yaml:"-"`
}
```

Vcs is ValueObject of VCS

#### func (*Vcs) Descriptor

```go
func (*Vcs) Descriptor() ([]byte, []int)
```

#### func (*Vcs) GetBranch

```go
func (m *Vcs) GetBranch() string
```

#### func (*Vcs) GetHash

```go
func (m *Vcs) GetHash() string
```

#### func (*Vcs) GetTag

```go
func (m *Vcs) GetTag() string
```

#### func (*Vcs) ProtoMessage

```go
func (*Vcs) ProtoMessage()
```

#### func (*Vcs) Reset

```go
func (m *Vcs) Reset()
```

#### func (*Vcs) String

```go
func (m *Vcs) String() string
```

#### func (*Vcs) Validate

```go
func (this *Vcs) Validate() error
```

#### func (*Vcs) XXX_DiscardUnknown

```go
func (m *Vcs) XXX_DiscardUnknown()
```

#### func (*Vcs) XXX_Marshal

```go
func (m *Vcs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```

#### func (*Vcs) XXX_Merge

```go
func (m *Vcs) XXX_Merge(src proto.Message)
```

#### func (*Vcs) XXX_Size

```go
func (m *Vcs) XXX_Size() int
```

#### func (*Vcs) XXX_Unmarshal

```go
func (m *Vcs) XXX_Unmarshal(b []byte) error
```
