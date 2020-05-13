package pipe

import (
	"context"

	"github.com/cheekybits/genny/generic"
	"go.opencensus.io/trace"
	"google.golang.org/grpc/codes"

	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/slog"
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

	go slog.Recover(ctx, func(ctx context.Context) error {
		const op = op + "#func"
		defer close(outCh)
		for v := range inCh {
			v0 := v
			v1, err := fn(v0)
			if err != nil {
				var vctx context.Context
				vctx, ok := v0.(context.Context)
				if !ok {
					v1, ok := v0.(NameGetContexter)
					if ok {
						vctx = v1.GetContext()
					} else {
						vctx = context.Background()
					}
				}
				if fnErr(vctx, &errs.Error{Op: op, Err: err}) {
					return nil
				}
				continue
			}
			select {
			case <-ctx.Done():
				if err := ctx.Err(); err != nil {
					fnErr(ctx, &errs.Error{Op: op, Err: err})
				}
				return nil
			case outCh <- v1:
			}
		}
		return nil
	})

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

	go slog.Recover(ctx, func(ctx context.Context) error {
		const op = op + "#func"
		defer close(outCh)
		for v := range inCh {
			v0 := v
			v1, err := fn(v0)
			if err != nil {
				var ctx0 context.Context
				ctx0, ok := v0.(context.Context)
				if !ok {
					v2, ok := v0.(NameGetContexter)
					if ok {
						ctx0 = v2.GetContext()
					} else {
						ctx0 = context.Background()
					}
				}
				v3 := fnErr(ctx0, &errs.Error{Op: op, Err: err})
				v4 := trace.FromContext(ctx0)
				if v4 != nil {
					v4.End()
				}
				if v3 {
					return nil
				}
				continue
			}
			for _, v := range v1 {
				v5 := v
				select {
				case <-ctx.Done():
					if err := ctx.Err(); err != nil {
						fnErr(ctx, &errs.Error{Op: op, Err: err})
					}
					return nil
				case outCh <- v5:
				}
			}
		}
		return nil
	})

	return inCh, outCh
}
