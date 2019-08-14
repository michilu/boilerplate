package debug

import (
	"github.com/google/uuid"
	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/slog"
	"google.golang.org/grpc/codes"
)

const (
	op = "service/debug"
)

func NewID() (string, error) {
	const op = op + ".NewID"
	{
		slog.Logger().Debug().Str("op", op).Msg("start")
		defer slog.Logger().Debug().Str("op", op).Msg("end")
	}
	v0, err := uuid.NewRandom()
	if err != nil {
		const op = op + ".uuid.NewRandom"
		return "", &errs.Error{Op: op, Code: codes.Unavailable, Err: err}
	}
	return v0.String(), nil
}
