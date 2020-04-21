package pb

import (
	"encoding/json"

	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/slog"
	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
)

func (p T) MarshalZerologObject(e *zerolog.Event) {
	const op = op + ".T"
	v, err := json.Marshal(p)
	if err != nil {
		const op = op + ".MarshalZerologObject.json.Marshal"
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
		slog.Logger().Err(err).Str("op", op).Msg(err.Error())
		return
	}
	e.RawJSON(op, v)
}
