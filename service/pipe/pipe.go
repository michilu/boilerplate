package pipe

import (
	"context"

	"github.com/cheekybits/genny/generic"
	"go.opencensus.io/trace"
	"google.golang.org/grpc/codes"

	"github.com/michilu/boilerplate/service/errs"
)

// InT is a placeholder for the genny.
type InT generic.Type

// OutT is a placeholder for the genny.
type OutT generic.Type

type NameGetContexter interface {
	GetContext() context.Context
}

// GetPipeName returns new input(chan<- InT)/output(<-chan OutT) channels that embedded the given 'func(InT) OutT'.
func GetPipeName(
	ctx context.Context,
	fn func(InT) (OutT, error),
	fnErr func(context.Context, error) bool,
) (
	chan<- InT,
	<-chan OutT,
) {
	const op = op + ".GetPipeName"

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
		const op = op + "#go"
		defer close(outCh)
		for i := range inCh {
			v0, err := fn(i)
			if err != nil {
				var vctx context.Context
				vctx, ok := i.(context.Context)
				if !ok {
					v1, ok := i.(NameGetContexter)
					if ok {
						vctx = v1.GetContext()
					} else {
						vctx = context.Background()
					}
				}
				if fnErr(vctx, &errs.Error{Op: op, Err: err}) {
					return
				}
				continue
			}
			select {
			case <-ctx.Done():
				err := ctx.Err()
				if err != nil {
					fnErr(ctx, &errs.Error{Op: op, Err: err})
				}
				return
			case outCh <- v0:
			}
		}
	}()

	return inCh, outCh
}

// GetFanoutName returns new input(chan<- InT)/output(<-chan OutT) channels that embedded the given 'func(InT) OutT'.
func GetFanoutName(
	ctx context.Context,
	fn func(InT) ([]OutT, error),
	fnErr func(context.Context, error) bool,
) (
	chan<- InT,
	<-chan OutT,
) {
	const op = op + ".GetFanoutName"

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
		const op = op + "#go"
		defer close(outCh)
		for i := range inCh {
			v0, err := fn(i)
			if err != nil {
				var ctx0 context.Context
				ctx0, ok := i.(context.Context)
				if !ok {
					v1, ok := i.(NameGetContexter)
					if ok {
						ctx0 = v1.GetContext()
					} else {
						ctx0 = context.Background()
					}
				}
				v2 := fnErr(ctx0, &errs.Error{Op: op, Err: err})
				v3 := trace.FromContext(ctx0)
				if v3 != nil {
					v3.End()
				}
				if v2 {
					return
				}
				continue
			}
			for _, v4 := range v0 {
				select {
				case <-ctx.Done():
					err := ctx.Err()
					if err != nil {
						fnErr(ctx, &errs.Error{Op: op, Err: err})
					}
					return
				case outCh <- v4:
				}
			}
		}
	}()

	return inCh, outCh
}
