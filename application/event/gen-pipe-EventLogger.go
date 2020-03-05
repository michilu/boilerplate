// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package event

import (
	"context"

	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/event"
	"github.com/michilu/boilerplate/service/slog"
	"go.opencensus.io/trace"
	"google.golang.org/grpc/codes"
)

type EventLoggerGetContexter interface {
	GetContext() context.Context
}

// GetPipeEventLogger returns new input(chan<- EventEventWithContexter)/output(<-chan EventKeyValueWithContexter) channels that embedded the given 'func(EventEventWithContexter) EventKeyValueWithContexter'.
func GetPipeEventLogger(
	ctx context.Context,
	fn func(event.EventWithContexter) (event.KeyValueWithContexter, error),
	fnErr func(context.Context, error) bool,
) (
	chan<- event.EventWithContexter,
	<-chan event.KeyValueWithContexter,
) {
	const op = op + ".GetPipeEventLogger"

	if ctx == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"})
	}
	if fn == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'fn' is nil"})
	}
	if fnErr == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'fnErr' is nil"})
	}

	inCh := make(chan event.EventWithContexter)
	outCh := make(chan event.KeyValueWithContexter)

	go slog.Recover(ctx, func(ctx context.Context) error {
		const op = op + "#go"
		defer close(outCh)
		for i := range inCh {
			v0, err := fn(i)
			if err != nil {
				var vctx context.Context
				vctx, ok := i.(context.Context)
				if !ok {
					v1, ok := i.(EventLoggerGetContexter)
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
			case outCh <- v0:
			}
		}
		return nil
	})

	return inCh, outCh
}

// GetFanoutEventLogger returns new input(chan<- EventEventWithContexter)/output(<-chan EventKeyValueWithContexter) channels that embedded the given 'func(EventEventWithContexter) EventKeyValueWithContexter'.
func GetFanoutEventLogger(
	ctx context.Context,
	fn func(event.EventWithContexter) ([]event.KeyValueWithContexter, error),
	fnErr func(context.Context, error) bool,
) (
	chan<- event.EventWithContexter,
	<-chan event.KeyValueWithContexter,
) {
	const op = op + ".GetFanoutEventLogger"

	if ctx == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"})
	}
	if fn == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'fn' is nil"})
	}
	if fnErr == nil {
		panic(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'fnErr' is nil"})
	}

	inCh := make(chan event.EventWithContexter)
	outCh := make(chan event.KeyValueWithContexter)

	go slog.Recover(ctx, func(ctx context.Context) error {
		const op = op + "#go"
		defer close(outCh)
		for i := range inCh {
			v0, err := fn(i)
			if err != nil {
				var ctx0 context.Context
				ctx0, ok := i.(context.Context)
				if !ok {
					v1, ok := i.(EventLoggerGetContexter)
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
					if err := ctx.Err(); err != nil {
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
