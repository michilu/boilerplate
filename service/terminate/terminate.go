package terminate

import (
	"context"

	"github.com/michilu/boilerplate/service/errs"
	"go.opencensus.io/trace"
	"google.golang.org/grpc/codes"
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
		ctx := context.Background()
		ctx, _ = trace.StartSpan(ctx, op)
		return ctx, &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"}
	}
	ctx, s := trace.StartSpan(ctx, op)
	defer s.End()
	return ctx, nil
}
