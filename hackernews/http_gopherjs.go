// +build gopherjs

package hackernews

import (
	"github.com/michilu/boilerplate/service/errs"
	"honnef.co/go/js/xhr"
)

func httpGet(u string) ([]byte, error) {
	const op = op + ".httpGet+gopherjs"
	b, err := xhr.Send("GET", u, nil)
	if err != nil {
		return nil, &errs.Error{Op: op, Err: err}
	}
	return b, nil
}
