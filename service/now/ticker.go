package now

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

func ContextTicker(ctx context.Context, duration time.Duration) <-chan context.Context {
	const op = op + ".ContextTicker"
	if ctx == nil {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"}
		slog.Logger().Fatal().Str("op", op).Err(err).Msg(err.Error())
	}
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	ctx, s := trace.StartSpan(ctx, op)
	defer s.End()
	t := slog.Trace(ctx)

	{
		s.AddAttributes(trace.StringAttribute("duration", fmt.Sprintf("%v", duration)))
		slog.Logger().Debug().Str("op", op).EmbedObject(t).Dur("duration", duration).Msg("arg")
	}
	oCh := make(chan context.Context)
	go func(ctx context.Context, duration time.Duration, oCh chan<- context.Context) {
		const op = op + "#go"
		defer close(oCh)
		ticker := time.NewTicker(duration)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
			}
			m := context.Background()
			m, _ = trace.StartSpan(m, op)
			oCh <- m
		}
	}(ctx, duration, oCh)
	{
		v0 := fmt.Sprintf("%v", oCh)
		s.AddAttributes(trace.StringAttribute("oCh", v0))
		slog.Logger().Debug().Str("op", op).EmbedObject(t).Str("oCh", v0).Msg("return")
	}
	return oCh
}
