// +build !gopherjs

package hackernews

import (
	"io/ioutil"
	"net/http"

	"github.com/michilu/boilerplate/service/errs"
)

func httpGet(u string) ([]byte, error) {
	const op = op + ".httpGet!gopherjs"
	r, err := http.Get(u)
	if err != nil {
		return nil, &errs.Error{Op: op, Err: err}
	}
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, &errs.Error{Op: op, Err: err}
	}
	return b, nil
}
