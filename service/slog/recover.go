package slog

import (
	"context"

	"github.com/StudioSol/async"
)

func Recover(ctx context.Context, task ...async.Task) {
	const op = op + ".Recover"
	if err := async.Run(ctx, task...); err != nil {
		Logger().Err(err).Str("op", op).Msg(err.Error())
	}
}
