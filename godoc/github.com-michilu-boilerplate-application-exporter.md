# exporter
--
    import "github.com/michilu/boilerplate/application/exporter"


## Usage

#### func  Run

```go
func Run(ctx context.Context) error
```

#### type GlobalMonitoredResource

```go
type GlobalMonitoredResource struct {
}
```


#### func (GlobalMonitoredResource) MonitoredResource

```go
func (g GlobalMonitoredResource) MonitoredResource() (string, map[string]string)
```
