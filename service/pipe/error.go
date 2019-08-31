package pipe

import (
	"context"
	"fmt"

	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/slog"
	"go.opencensus.io/trace"
	"google.golang.org/grpc/codes"
)

// ErrorHandler is an error handler with error level.
func ErrorHandler(ctx context.Context, err error) (returns bool) {
	const op = op + ".ErrorHandler"
	v0 := slog.Logger().Error().Str("op", op).Err(err)
	if ctx == nil {
		ctx = context.Background()
		slog.Logger().Error().Str("op", op).Err(&errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}).Msg("must be given. 'ctx' is nil")
	}
	{
		sp := trace.FromContext(ctx)
		if sp != nil {
			defer sp.End()
			v1 := sp.SpanContext()
			v0.
				Str("trace", fmt.Sprintf(slog.GetTraceIDTemplate(), v1.TraceID.String())).
				Str("spanID", v1.SpanID.String())
		}
	}
	v0.Msg("error")
	ctx, s := trace.StartSpan(ctx, op)
	defer s.End()
	s.SetStatus(trace.Status{Code: int32(codes.Unknown), Message: err.Error()})
	return
}

// FatalErrorHandler is an error handler with fatal level.
func FatalErrorHandler(ctx context.Context, err error) (returns bool) {
	const op = op + ".FatalErrorHandler"
	v0 := slog.Logger().Error().Str("op", op).Err(err)
	if ctx == nil {
		ctx = context.Background()
		slog.Logger().Error().Str("op", op).Err(&errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}).Msg("must be given. 'ctx' is nil")
	}
	{
		sp := trace.FromContext(ctx)
		if sp != nil {
			defer sp.End()
			v1 := sp.SpanContext()
			v0.
				Str("trace", fmt.Sprintf(slog.GetTraceIDTemplate(), v1.TraceID.String())).
				Str("spanID", v1.SpanID.String())
		}
	}
	v0.Msg("error")
	ctx, s := trace.StartSpan(ctx, op)
	defer s.End()
	s.SetStatus(trace.Status{Code: int32(codes.Unknown), Message: err.Error()})
	return true
}
