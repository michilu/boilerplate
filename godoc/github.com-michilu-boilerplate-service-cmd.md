# cmd
--
    import "github.com/michilu/boilerplate/service/cmd"


## Usage

#### func  NewCommand

```go
func NewCommand(
	resource *Resource,
	defaults []config.KV,
	initCmdFlag func(*cobra.Command),
	subCmd []func() (*cobra.Command, error),
) *cobra.Command
```

#### type Resource

```go
type Resource struct {
	Context  context.Context
	Resource []func(context.Context) (io.Closer, error)
}
```


#### func (Resource) Close

```go
func (p Resource) Close() (err error)
```

#### func (*Resource) Init

```go
func (p *Resource) Init() error
```
