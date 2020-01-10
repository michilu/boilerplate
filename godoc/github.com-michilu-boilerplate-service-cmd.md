# cmd
--
    import "github.com/michilu/boilerplate/service/cmd"


## Usage

#### func  NewCommand

```go
func NewCommand(
	logger func() ([]io.Writer, slog.Closer, error),
	defaults []config.KV,
	initCmdFlag func(*cobra.Command),
	subCmd []func() (*cobra.Command, error),
) (*cobra.Command, slog.Closer)
```
