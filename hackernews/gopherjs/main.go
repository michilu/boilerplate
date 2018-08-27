package main

import (
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	js.Global.Set("hackernews", map[string]interface{}{
		"HackerNews": NewHackerNewsGopherJS,
	})
}
