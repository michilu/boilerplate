package topic

import (
	"context"
	"fmt"
	"time"

	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/slog"
	"go.opencensus.io/trace"
	"google.golang.org/grpc/codes"
)

//go:generate genny -in=../topic/topic.go -out=gen-topic-context.go -pkg=$GOPACKAGE gen "ChanT=context.Context"

func ContextTicker(ctx context.Context, duration time.Duration) (<-chan context.Context, error) {
	const op = op + ".ContextTicker"
	if ctx == nil {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"}
		slog.Logger().Err(err).Str("op", op).Msg(err.Error())
		return nil, err
	}
	ctx, s := trace.StartSpan(ctx, op)
	defer s.End()
	t := slog.Trace(ctx, s)

	{
		s.AddAttributes(trace.StringAttribute("duration", fmt.Sprintf("%v", duration)))
		slog.Logger().Debug().Str("op", op).EmbedObject(t).Dur("duration", duration).Msg(op + ": arg")
	}
	oCh := make(chan context.Context)
	go slog.Recover(ctx, func(ctx context.Context) error {
		ticker(ctx, duration, oCh)
		return nil
	})
	{
		v0 := fmt.Sprintf("%v", oCh)
		s.AddAttributes(trace.StringAttribute("oCh", v0))
		slog.Logger().Debug().Str("op", op).EmbedObject(t).Str("oCh", v0).Msg(op + ": return")
	}
	return oCh, nil
}

func ticker(ctx context.Context, duration time.Duration, oCh chan<- context.Context) {
	const op = op + ".ticker"
	defer close(oCh)
	ticker := time.NewTicker(duration)
	defer ticker.Stop()
	for {
		const op = op + "#for"
		select {
		case <-ctx.Done():
			err := &errs.Error{Op: op, Code: codes.Canceled, Err: ctx.Err()}
			slog.Logger().Debug().Err(err).Str("op", op).Msg(err.Error())
			return
		case <-ticker.C:
		}
		m, _ := trace.StartSpan(context.Background(), op)
		select {
		case <-ctx.Done():
			err := &errs.Error{Op: op, Code: codes.Canceled, Err: ctx.Err()}
			slog.Logger().Debug().Err(err).Str("op", op).Msg(err.Error())
			return
		case oCh <- m:
		}
	}
}
