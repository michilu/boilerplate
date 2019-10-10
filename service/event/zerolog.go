package event

import (
	"encoding/json"

	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/slog"
	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
)

func (p *Event) MarshalZerologObject(e *zerolog.Event) {
	const op = op + ".Event.MarshalZerologObject"
	v, err := json.Marshal(&p)
	if err != nil {
		const op = op + ".json.Marshal"
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
		slog.Logger().Error().Str("op", op).Err(err).Msg(err.Error())
		return
	}
	e.RawJSON("event", v)
}
