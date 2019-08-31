// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package debug

import (
	"context"

	"github.com/michilu/boilerplate/service/debug"
	"github.com/michilu/boilerplate/service/errs"
	"google.golang.org/grpc/codes"
)

type ConfigGetContexter interface {
	GetContext() context.Context
}

// GetPipeConfig returns new input(chan<- ContextContext)/output(<-chan DebugClientWithContexter) channels that embedded the given 'func(ContextContext) DebugClientWithContexter'.
func GetPipeConfig(
	ctx context.Context,
	fn func(context.Context) (debug.ClientWithContexter, error),
	fnErr func(context.Context, error) bool,
) (
	chan<- context.Context,
	<-chan debug.ClientWithContexter,
) {
	const op = op + ".GetPipeConfig"

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
	outCh := make(chan debug.ClientWithContexter)

	go func() {
		const op = op + "#go"
		defer close(outCh)
		for i := range inCh {
			o, err := fn(i)
			if err != nil {
				var vctx context.Context
				vctx, ok := i.(context.Context)
				if !ok {
					v, ok := i.(ConfigGetContexter)
					if ok {
						vctx = v.GetContext()
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
			case outCh <- o:
			default:
			}
		}
	}()

	return inCh, outCh
}

// GetFanoutConfig returns new input(chan<- ContextContext)/output(<-chan DebugClientWithContexter) channels that embedded the given 'func(ContextContext) DebugClientWithContexter'.
func GetFanoutConfig(
	ctx context.Context,
	fn func(context.Context) ([]debug.ClientWithContexter, error),
	fnErr func(context.Context, error) bool,
) (
	chan<- context.Context,
	<-chan debug.ClientWithContexter,
) {
	const op = op + ".GetFanoutConfig"

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
	outCh := make(chan debug.ClientWithContexter)

	go func() {
		const op = op + "#go"
		defer close(outCh)
		for i := range inCh {
			o, err := fn(i)
			if err != nil {
				var vctx context.Context
				vctx, ok := i.(context.Context)
				if !ok {
					v, ok := i.(ConfigGetContexter)
					if ok {
						vctx = v.GetContext()
					} else {
						vctx = context.Background()
					}
				}
				if fnErr(vctx, &errs.Error{Op: op, Err: err}) {
					return
				}
				continue
			}
			for _, v := range o {
				select {
				case <-ctx.Done():
					err := ctx.Err()
					if err != nil {
						fnErr(ctx, &errs.Error{Op: op, Err: err})
					}
					return
				case outCh <- v:
				default:
				}
			}
		}
	}()

	return inCh, outCh
}
