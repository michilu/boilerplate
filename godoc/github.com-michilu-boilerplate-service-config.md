# config
--
    import "github.com/michilu/boilerplate/service/config"


## Usage

#### func  GCPCredentials

```go
func GCPCredentials(ctx context.Context) (*google.Credentials, error)
```

#### func  GCPProjectID

```go
func GCPProjectID(ctx context.Context) (gcp.ProjectID, error)
```

#### func  SetDefault

```go
func SetDefault(config ...[]KV)
```
SetDefault sets default values to config.

#### type KV

```go
type KV struct {
	K string
	V interface{}
}
```
