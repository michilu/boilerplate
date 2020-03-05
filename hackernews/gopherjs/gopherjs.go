package main

import (
	"context"

	"github.com/augustoroman/promise"
	"github.com/gopherjs/gopherjs/js"
	"github.com/michilu/boilerplate/hackernews"
	"github.com/michilu/boilerplate/service/slog"
)

type (
	hnGopherJS struct {
		h hackernews.HackerNews
	}
)

// NewHackerNewsGopherJS returns a new js.Object of HackerNews.
func NewHackerNewsGopherJS(baseUrl string) *js.Object {
	return js.MakeWrapper(&hnGopherJS{h: hackernews.NewHackerNewsAPI(baseUrl)})
}

func (h *hnGopherJS) GetFeed(n string, p int) *js.Object {
	var (
		j promise.Promise
	)
	go slog.Recover(context.Background(), func(ctx context.Context) error {
		v, err := h.h.GetFeed(n, p)
		if err != nil {
			j.Reject(err)
			return nil
		}
		j.Resolve(v)
		return nil
	})
	return j.Js()
}

func (h *hnGopherJS) GetItem(i string) *js.Object {
	var (
		j promise.Promise
	)
	go slog.Recover(context.Background(), func(ctx context.Context) error {
		v, err := h.h.GetItem(i)
		if err != nil {
			j.Reject(err)
			return nil
		}
		j.Resolve(v)
		return nil
	})
	return j.Js()
}
