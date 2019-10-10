package event

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/slog"
	"github.com/rs/zerolog"
	"go.opencensus.io/trace"
	"google.golang.org/grpc/codes"
)

type Message interface {
	proto.Message
	zerolog.LogObjectMarshaler
}

// StoreEvent returns a bytes from given Marshaler.
func StoreEvent(ctx context.Context, message Message) ([]byte, error) {
	const op = op + ".StoreEvent"
	if ctx == nil {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"}
		return nil, err
	}
	ctx, s := trace.StartSpan(ctx, op)
	defer s.End()
	t := slog.Trace(ctx)
	slog.Logger().Debug().Str("op", op).EmbedObject(t).EmbedObject(message).Msg("arg")

	v1, err := proto.Marshal(message)
	if err != nil {
		const op = op + ".proto.Marshal"
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
		s.SetStatus(trace.Status{Code: int32(codes.InvalidArgument), Message: err.Error()})
		slog.Logger().Error().Str("op", op).EmbedObject(t).Err(err).Msg("error")
		return v1, err
	}
	slog.Logger().Debug().Str("op", op).EmbedObject(t).Bytes("v1", v1).Msg("return")
	return v1, nil
}
