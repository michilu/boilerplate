package event

import (
	"context"

	"github.com/michilu/boilerplate/infra/keyvalue"
	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/slog"
	"github.com/rs/zerolog"
	"go.opencensus.io/trace"
	"google.golang.org/grpc/codes"
)

type Keyer interface {
	GetKey() string
	zerolog.LogObjectMarshaler
}

// SaveEventPayload saves an event payload.
func SaveEventPayload(ctx context.Context, repository Saver, keyvalue keyvalue.KeyValuer) error {
	const op = op + ".SaveEventPayload"
	if ctx == nil {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"}
		return err
	}
	ctx, s := trace.StartSpan(ctx, op)
	defer s.End()
	a := make([]trace.Attribute, 0)
	defer s.AddAttributes(a...)
	t := slog.Trace(ctx)

	if repository == nil {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. repository is nil"}
		s.SetStatus(trace.Status{Code: int32(codes.InvalidArgument), Message: err.Error()})
		return err
	}
	if keyvalue == nil {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. keyvalue is nil"}
		s.SetStatus(trace.Status{Code: int32(codes.InvalidArgument), Message: err.Error()})
		return err
	}
	{
		err := keyvalue.Validate()
		if err != nil {
			err := &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
			s.SetStatus(trace.Status{Code: int32(codes.InvalidArgument), Message: err.Error()})
			return err
		}
		slog.Logger().Debug().Str("op", op).EmbedObject(t).EmbedObject(keyvalue).Msg("arg")
	}
	{
		err := repository.Save(ctx, keyvalue)
		if err != nil {
			const op = op + ".Repository.Save"
			err := &errs.Error{Op: op, Code: codes.Internal, Err: err}
			s.SetStatus(trace.Status{Code: int32(codes.Internal), Message: err.Error()})
			return err
		}
	}
	return nil
}