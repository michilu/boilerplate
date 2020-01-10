# keystore
--
    import "github.com/michilu/boilerplate/infra/nutsdb/keystore"


## Usage

#### func  NewOptions

```go
func NewOptions() nutsdb.Options
```

#### func  NewRepository

```go
func NewRepository(ctx context.Context) (keyvalue.KeyValueCloser, func() error, error)
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

#### func (*Repository) Get

```go
func (p *Repository) Get(ctx context.Context, key keyvalue.Keyer) (keyvalue.KeyValuer, error)
```

#### func (*Repository) Save

```go
func (p *Repository) Save(ctx context.Context, keyvalue keyvalue.KeyValuer) error
```
