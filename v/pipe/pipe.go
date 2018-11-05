package pipe

import (
	"context"

	"github.com/cheekybits/genny/generic"
	"google.golang.org/grpc/codes"

	"github.com/michilu/boilerplate/v/errs"
)

// InT is a placeholder for the genny.
type InT generic.Type

// OutT is a placeholder for the genny.
type OutT generic.Type

// GetPipeName returns new input(chan<- InT)/output(<-chan OutT) channels that embedded the given 'func(InT) OutT'.
func GetPipeName(
	ctx context.Context,
	fn func(InT) (OutT, error),
	fnErr func(error) bool,
) (
	chan<- InT,
	<-chan OutT,
) {
	const op = "pipe.getPipeName"

	if ctx == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"})
	}
	if fn == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'fn' is nil"})
	}
	if fnErr == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'fnErr' is nil"})
	}

	inCh := make(chan InT)
	outCh := make(chan OutT)

	go func() {
		const op = "pipe.getPipeName#go"
		defer close(outCh)
		for i := range inCh {
			o, err := fn(i)
			if err != nil {
				if fnErr(&errs.Error{Op: op, Err: err}) {
					return
				}
				continue
			}
			select {
			case <-ctx.Done():
				err := ctx.Err()
				if err != nil {
					fnErr(&errs.Error{Op: op, Err: err})
				}
				return
			case outCh <- o:
			default:
			}
		}
	}()

	return inCh, outCh
}
