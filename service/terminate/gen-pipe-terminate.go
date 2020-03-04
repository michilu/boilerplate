// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package terminate

import (
	"context"

	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/slog"
	"go.opencensus.io/trace"
	"google.golang.org/grpc/codes"
)

type TerminateGetContexter interface {
	GetContext() context.Context
}

// GetPipeTerminate returns new input(chan<- ContextContext)/output(<-chan ContextContext) channels that embedded the given 'func(ContextContext) ContextContext'.
func GetPipeTerminate(
	ctx context.Context,
	fn func(context.Context) (context.Context, error),
	fnErr func(context.Context, error) bool,
) (
	chan<- context.Context,
	<-chan context.Context,
) {
	const op = op + ".GetPipeTerminate"

	if ctx == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"})
	}
	if fn == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'fn' is nil"})
	}
	if fnErr == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'fnErr' is nil"})
	}

	inCh := make(chan context.Context)
	outCh := make(chan context.Context)

	go slog.Recover(ctx, func(ctx context.Context) error {
		const op = op + "#go"
		defer close(outCh)
		for i := range inCh {
			v0, err := fn(i)
			if err != nil {
				var vctx context.Context
				vctx, ok := i.(context.Context)
				if !ok {
					v1, ok := i.(TerminateGetContexter)
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
				err := ctx.Err()
				if err != nil {
					fnErr(ctx, &errs.Error{Op: op, Err: err})
				}
				return nil
			case outCh <- v0:
			}
		}
		return nil
	})

	return inCh, outCh
}

// GetFanoutTerminate returns new input(chan<- ContextContext)/output(<-chan ContextContext) channels that embedded the given 'func(ContextContext) ContextContext'.
func GetFanoutTerminate(
	ctx context.Context,
	fn func(context.Context) ([]context.Context, error),
	fnErr func(context.Context, error) bool,
) (
	chan<- context.Context,
	<-chan context.Context,
) {
	const op = op + ".GetFanoutTerminate"

	if ctx == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"})
	}
	if fn == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'fn' is nil"})
	}
	if fnErr == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'fnErr' is nil"})
	}

	inCh := make(chan context.Context)
	outCh := make(chan context.Context)

	go slog.Recover(ctx, func(ctx context.Context) error {
		const op = op + "#go"
		defer close(outCh)
		for i := range inCh {
			v0, err := fn(i)
			if err != nil {
				var ctx0 context.Context
				ctx0, ok := i.(context.Context)
				if !ok {
					v1, ok := i.(TerminateGetContexter)
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
					return nil
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
					return nil
				case outCh <- v4:
				}
			}
		}
		return nil
	})

	return inCh, outCh
}
