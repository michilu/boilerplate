package hackernews

type (
	// HackerNews is a interface of HackerNews.
	HackerNews interface {
		// GetFeed returns a list of feeds given params.
		GetFeed(name string, page int) (interface{}, error)
		// GetItem returns a feed given params.
		GetItem(id string) (interface{}, error)
	}
)
