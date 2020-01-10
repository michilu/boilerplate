# hackernews
--
    import "github.com/michilu/boilerplate/hackernews"


## Usage

#### type HackerNews

```go
type HackerNews interface {
	// GetFeed returns a list of feeds given params.
	GetFeed(name string, page int) (interface{}, error)
	// GetItem returns a feed given params.
	GetItem(id string) (interface{}, error)
}
```

HackerNews is a interface of HackerNews.

#### func  NewHackerNewsAPI

```go
func NewHackerNewsAPI(baseUrl string) HackerNews
```
NewHackerNewsAPI returns a new HackerNews via API.
