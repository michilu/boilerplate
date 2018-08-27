package main

import (
	"github.com/augustoroman/promise"
	"github.com/gopherjs/gopherjs/js"

	"github.com/michilu/boilerplate/hackernews"
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
	go func() {
		v, err := h.h.GetFeed(n, p)
		if err != nil {
			j.Reject(err)
			return
		}
		j.Resolve(v)
	}()
	return j.Js()
}

func (h *hnGopherJS) GetItem(i string) *js.Object {
	var (
		j promise.Promise
	)
	go func() {
		v, err := h.h.GetItem(i)
		if err != nil {
			j.Reject(err)
			return
		}
		j.Resolve(v)
	}()
	return j.Js()
}
