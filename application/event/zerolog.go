package event

import (
	"encoding/json"

	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/slog"
	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
)

func (p *EventWithContext) MarshalZerologObject(e *zerolog.Event) {
	const op = op + ".EventWithContext.MarshalZerologObject"
	b, err := json.Marshal(p.GetEvent())
	if err != nil {
		const op = op + ".json.Marshal"
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
		slog.Logger().Error().Str("op", op).Err(err).Msg("error")
		return
	}
	e.RawJSON("eventWithContext", b)
}

func (p *ByteWithContext) MarshalZerologObject(e *zerolog.Event) {
	const op = op + ".ByteWithContext.MarshalZerologObject"
	b, err := json.Marshal(p.GetByte())
	if err != nil {
		const op = op + ".json.Marshal"
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
		slog.Logger().Error().Str("op", op).Err(err).Msg("error")
		return
	}
	e.RawJSON("byteWithContext", b)
}
