// +build gopherjs

package hackernews

import (
	"github.com/michilu/boilerplate/v/errs"
	"honnef.co/go/js/xhr"
)

func httpGet(u string) ([]byte, error) {
	const op = "hackernews.httpGet+gopherjs"
	b, err := xhr.Send("GET", u, nil)
	if err != nil {
		return nil, &errs.Error{Op: op, Err: err}
	}
	return b, nil
}
