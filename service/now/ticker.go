package now

import (
	"context"
	"fmt"
	"time"

	"github.com/michilu/boilerplate/service/slog"
	"go.opencensus.io/trace"
)

//go:generate genny -in=../topic/topic.go -out=gen-topic-context.go -pkg=$GOPACKAGE gen "ChanT=context.Context"

func ContextTicker(ctx context.Context, duration time.Duration) <-chan context.Context {
	const op = op + ".ContextTicker"
	oCh := make(chan context.Context)
	go func(ctx context.Context, duration time.Duration, oCh chan<- context.Context) {
		const op = op + "#go"
		defer close(oCh)
		{
			sp := trace.FromContext(ctx)
			if sp != nil {
				v0 := sp.SpanContext()
				slog.Logger().Debug().Str("op", op).
					Str("trace", fmt.Sprintf(slog.GetTraceIDTemplate(), v0.TraceID.String())).
					Str("spanID", v0.SpanID.String())
				sp.End()
			}
		}
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
	return oCh
}
