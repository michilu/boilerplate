package hackernews

import (
	"strconv"
	"strings"

	"github.com/michilu/boilerplate/v/errs"
)

const (
	defaultBaseUrl = "https://api.hnpwa.com/v0"
)

var (
	cacheFeedKey    string
	cacheFeedResult interface{}
)

type (
	hnAPI struct {
		b string
	}
)

// NewHackerNewsAPI returns a new HackerNews via API.
func NewHackerNewsAPI(baseUrl string) HackerNews {
	if baseUrl == "" {
		baseUrl = defaultBaseUrl
	}
	return &hnAPI{b: baseUrl}
}

func (h *hnAPI) GetFeed(n string, p int) (interface{}, error) {
	const op = "hackernews.hnAPI.GetFeed"
	var (
		url = strings.Join([]string{h.b, "/", n, "/", strconv.FormatInt(int64(p), 10), ".json"}, "")
	)
	if cacheFeedKey == url {
		return cacheFeedResult, nil
	}
	b, err := httpGet(url)
	if err != nil {
		return nil, &errs.Error{Op: op, Err: err}
	}
	v, err := jsonUnmarshalFeed(b)
	if err != nil {
		return nil, &errs.Error{Op: op, Err: err}
	}
	cacheFeedKey = url
	cacheFeedResult = v
	return v, nil
}

func (h *hnAPI) GetItem(i string) (interface{}, error) {
	const op = "hackernews.hnAPI.GetItem"
	var (
		url = strings.Join([]string{h.b, "/item/", i, ".json"}, "")
	)
	b, err := httpGet(url)
	if err != nil {
		return nil, &errs.Error{Op: op, Err: err}
	}
	v, err := jsonUnmarshalItem(b)
	if err != nil {
		return nil, &errs.Error{Op: op, Err: err}
	}
	return *v, nil
}
