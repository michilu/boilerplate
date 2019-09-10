package event

import (
	"context"

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
func SaveEventPayload(ctx context.Context, repository Repository, key Keyer, payload []byte) error {
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
	slog.Logger().Debug().Str("op", op).EmbedObject(t).EmbedObject(key).Bytes("payload", payload).Msg("arg")

	if repository == nil {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. repository is nil"}
		s.SetStatus(trace.Status{Code: int32(codes.InvalidArgument), Message: err.Error()})
		return err
	}
	if key == nil {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. key is nil"}
		s.SetStatus(trace.Status{Code: int32(codes.InvalidArgument), Message: err.Error()})
		return err
	}
	if len(payload) == 0 {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. payload is empty"}
		s.SetStatus(trace.Status{Code: int32(codes.InvalidArgument), Message: err.Error()})
		return err
	}

	v0 := key.GetKey()
	if v0 == "" {
		const op = op + ".Keyer.GetKey"
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. key.GetKey() is ''"}
		s.SetStatus(trace.Status{Code: int32(codes.InvalidArgument), Message: err.Error()})
		return err
	}

	err := repository.Save(ctx, v0, payload)
	if err != nil {
		const op = op + ".Repository.Save"
		err := &errs.Error{Op: op, Code: codes.Internal, Err: err}
		s.SetStatus(trace.Status{Code: int32(codes.Internal), Message: err.Error()})
		return err
	}
	slog.Logger().Debug().Str("op", op).EmbedObject(t).Str("key", v0).Msg("arg")
	return nil
}
