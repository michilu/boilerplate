package keyvalue

import (
	"encoding/json"

	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/slog"
	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
)

func (p *KeyValue) MarshalZerologObject(e *zerolog.Event) {
	const op = op + ".KeyValue.MarshalZerologObject"
	v, err := json.Marshal(&p)
	if err != nil {
		const op = op + ".json.Marshal"
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
		slog.Logger().Error().Str("op", op).Err(err).Msg("error")
		return
	}
	e.RawJSON("keyValue", v)
}

func (p *Prefix) MarshalZerologObject(e *zerolog.Event) {
	const op = op + ".Prefix.MarshalZerologObject"
	v, err := json.Marshal(&p)
	if err != nil {
		const op = op + ".json.Marshal"
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
		slog.Logger().Error().Str("op", op).Err(err).Msg("error")
		return
	}
	e.RawJSON("prefix", v)
}
