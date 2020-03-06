package slog

import (
	"context"

	"github.com/StudioSol/async"
	"github.com/michilu/boilerplate/service/errs"
	"google.golang.org/grpc/codes"
)

func Recover(ctx context.Context, task ...async.Task) error {
	const op = op + ".Recover"
	if err := async.Run(ctx, task...); err != nil {
		err := &errs.Error{Op: op, Code: codes.Unknown, Err: err}
		Logger().Err(err).Str("op", op).Msg(err.Error())
		return err
	}
	return nil
}
