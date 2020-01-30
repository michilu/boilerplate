package pipe

import (
	"context"

	"cloud.google.com/go/errorreporting"
	"github.com/michilu/boilerplate/service/config"
	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/meta"
	"github.com/michilu/boilerplate/service/slog"
	"go.opencensus.io/trace"
	"google.golang.org/grpc/codes"
)

var (
	_errorreportingClient *errorreporting.Client
)

func Init(ctx context.Context) {
	const op = op + ".Init"
	if ctx == nil {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"}
		slog.Logger().Err(err).Str("op", op).Msg(err.Error())
		return
	}
	ctx, s := trace.StartSpan(ctx, op)
	defer s.End()
	t := slog.Trace(ctx, s)

	v0, err := config.GCPProjectID(ctx)
	if err != nil {
		const op = op + ".config.GCPProjectID"
		err := &errs.Error{Op: op, Code: codes.Internal, Err: err}
		s.SetStatus(trace.Status{Code: int32(codes.Internal), Message: err.Error()})
		slog.Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
		return
	}
	v1 := meta.Get()
	v2, err := errorreporting.NewClient(ctx, string(v0), errorreporting.Config{
		ServiceName:    v1.Name,
		ServiceVersion: v1.Semver,
		OnError: func(err error) {
			const op = op + ".errorreporting.Client.OnError"
			{
				err := &errs.Error{Op: op, Code: codes.Unavailable, Err: err}
				slog.Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
			}
		},
	})
	if err != nil {
		const op = op + ".config.GCPCredentials"
		err := &errs.Error{Op: op, Code: codes.Internal, Err: err}
		s.SetStatus(trace.Status{Code: int32(codes.Internal), Message: err.Error()})
		slog.Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
		return
	}
	_errorreportingClient = v2
}

// ErrorHandler is an error handler with error level.
func ErrorHandler(ctx context.Context, err error) (returns bool) {
	const op = op + ".ErrorHandler"
	if ctx == nil {
		ctx = context.Background()
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"}
		slog.Logger().Err(err).Str("op", op).Msg(err.Error())
	}
	ctx, s := trace.StartSpan(ctx, op)
	defer s.End()
	t := slog.Trace(ctx, s)

	s.SetStatus(trace.Status{Code: int32(codes.Unknown), Message: err.Error()})
	slog.Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
	if _errorreportingClient != nil {
		_errorreportingClient.Report(errorreporting.Entry{Error: err})
	}
	return
}

// FatalErrorHandler is an error handler with fatal level.
func FatalErrorHandler(ctx context.Context, err error) (returns bool) {
	const op = op + ".FatalErrorHandler"
	if ctx == nil {
		ctx = context.Background()
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"}
		slog.Logger().Err(err).Str("op", op).Msg(err.Error())
	}
	ctx, s := trace.StartSpan(ctx, op)
	defer s.End()
	t := slog.Trace(ctx, s)

	s.SetStatus(trace.Status{Code: int32(codes.Unknown), Message: err.Error()})
	slog.Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
	if _errorreportingClient != nil {
		_errorreportingClient.Report(errorreporting.Entry{Error: err})
	}
	return true
}
