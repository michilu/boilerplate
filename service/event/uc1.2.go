package event

import (
	"context"

	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/slog"
	"github.com/rs/zerolog"
	"go.opencensus.io/trace"
	"google.golang.org/grpc/codes"
)

type Marshaler interface {
	XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
	zerolog.LogObjectMarshaler
}

// StoreEvent returns a bytes from given Marshaler.
func StoreEvent(ctx context.Context, event Marshaler) ([]byte, error) {
	const op = op + ".StoreEvent"
	if ctx == nil {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"}
		return nil, err
	}
	ctx, s := trace.StartSpan(ctx, op)
	defer s.End()
	a := make([]trace.Attribute, 0)
	defer s.AddAttributes(a...)
	t := slog.Trace(ctx)
	slog.Logger().Debug().Str("op", op).EmbedObject(t).EmbedObject(event).Msg("arg")

	v0 := make([]byte, 0)
	v1, err := event.XXX_Marshal(v0, false)
	if err != nil {
		return nil, &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
	}
	slog.Logger().Debug().Str("op", op).EmbedObject(t).Bytes("v1", v1).Msg("return")
	return v1, nil
}
