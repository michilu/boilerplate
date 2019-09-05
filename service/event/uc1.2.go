package event

import (
	"github.com/michilu/boilerplate/service/errs"
	"google.golang.org/grpc/codes"
)

type Marshaler interface {
	XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
}

// Marshal returns a bytes from given Marshaler.
func Marshal(event Marshaler) ([]byte, error) {
	const op = op + ".Marshal"
	v0 := make([]byte, 0)
	v1, err := event.XXX_Marshal(v0, false)
	if err != nil {
		return nil, &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
	}
	return v1, nil
}
