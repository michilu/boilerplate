package terminate

import (
	"context"

	"go.opencensus.io/trace"
)

//go:generate genny -in=../topic/topic.go -out=gen-topic-Context.go -pkg=$GOPACKAGE gen "ChanT=context.Context"
//go:generate genny -in=../pipe/pipe.go -out=gen-pipe-terminate.go -pkg=$GOPACKAGE gen "Name=terminate InT=context.Context OutT=context.Context"

const (
	op = "service/terminate"
)

// Terminate is terminator.
func Terminate(ctx context.Context) (context.Context, error) {
	const op = op + ".Terminate"
	if ctx == nil {
		return context.Background(), nil
	}
	s := trace.FromContext(ctx)
	if s != nil {
		s.End()
	}
	return ctx, nil
}
