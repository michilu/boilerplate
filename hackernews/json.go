// +build !gopherjs

package hackernews

import (
	"encoding/json"

	"github.com/michilu/boilerplate/v/errs"
)

type (
	feed struct {
		CommentsCount int64  `json:"comments_count"`
		Domain        string `json:"domain"`
		ID            int64  `json:"id"`
		Points        int64  `json:"points"`
		Time          int64  `json:"time"`
		TimeAgo       string `json:"time_ago"`
		Title         string `json:"title"`
		Type          string `json:"type"`
		URL           string `json:"url"`
		User          string `json:"user"`
	}
)

func jsonUnmarshalFeed(b []byte) ([]feed, error) {
	const op = "hackernews.jsonUnmarshalFeed!gopherjs"
	var (
		v []feed
	)
	err := json.Unmarshal(b, &v)
	if err != nil {
		return nil, &errs.Error{Op: op, Err: err}
	}
	return v, nil
}

func jsonUnmarshalItem(b []byte) (*feed, error) {
	const op = "hackernews.jsonUnmarshalItem!gopherjs"
	var (
		v feed
	)
	err := json.Unmarshal(b, &v)
	if err != nil {
		return nil, &errs.Error{Op: op, Err: err}
	}
	return &v, nil
}
