// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package update

import (
	"context"

	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/slog"
	"go.opencensus.io/trace"
	"google.golang.org/grpc/codes"
)

type UpdateGetContexter interface {
	GetContext() context.Context
}

// GetPipeUpdate returns new input(chan<- ContextContext)/output(<-chan ContextContext) channels that embedded the given 'func(ContextContext) ContextContext'.
func GetPipeUpdate(
	ctx context.Context,
	fn func(context.Context) (context.Context, error),
	fnErr func(context.Context, error) bool,
) (
	chan<- context.Context,
	<-chan context.Context,
) {
	const op = op + ".GetPipeUpdate"

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
		const op = op + "#func"
		defer close(outCh)
		for v := range inCh {
			v0 := v
			v1, err := fn(v0)
			if err != nil {
				var vctx context.Context
				vctx, ok := v0.(context.Context)
				if !ok {
					v1, ok := v0.(UpdateGetContexter)
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

// GetFanoutUpdate returns new input(chan<- ContextContext)/output(<-chan ContextContext) channels that embedded the given 'func(ContextContext) ContextContext'.
func GetFanoutUpdate(
	ctx context.Context,
	fn func(context.Context) ([]context.Context, error),
	fnErr func(context.Context, error) bool,
) (
	chan<- context.Context,
	<-chan context.Context,
) {
	const op = op + ".GetFanoutUpdate"

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
		const op = op + "#func"
		defer close(outCh)
		for v := range inCh {
			v0 := v
			v1, err := fn(v0)
			if err != nil {
				var ctx0 context.Context
				ctx0, ok := v0.(context.Context)
				if !ok {
					v2, ok := v0.(UpdateGetContexter)
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
