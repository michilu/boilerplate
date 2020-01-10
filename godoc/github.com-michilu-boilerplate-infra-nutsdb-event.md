# event
--
    import "github.com/michilu/boilerplate/infra/nutsdb/event"


## Usage

#### func  NewOptions

```go
func NewOptions() nutsdb.Options
```

#### func  NewRepository

```go
func NewRepository() (keyvalue.LoadSaveCloser, func() error, error)
```

#### type Repository

```go
type Repository struct {
}
```


#### func (*Repository) Close

```go
func (p *Repository) Close() error
```

#### func (*Repository) Load

```go
func (p *Repository) Load(ctx context.Context, prefix keyvalue.Prefixer) (<-chan keyvalue.KeyValuer, error)
```

#### func (*Repository) Save

```go
func (p *Repository) Save(ctx context.Context, keyvalue keyvalue.KeyValuer) error
```
