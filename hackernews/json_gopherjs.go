// +build gopherjs

package hackernews

import (
	"github.com/gopherjs/gopherjs/js"
)

var (
	jsonUnmarshalFeed = jsonUnmarshal
	jsonUnmarshalItem = jsonUnmarshal
)

func jsonUnmarshal(b []byte) (*js.Object, error) {
	return js.Global.Get("JSON").Call("parse", string(b)), nil
}
