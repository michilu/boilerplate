package run

import (
	"context"
	"time"

	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/now"
	"github.com/michilu/boilerplate/service/pipe"
	"github.com/michilu/boilerplate/service/slog"
	"github.com/michilu/boilerplate/service/terminate"
	"github.com/michilu/boilerplate/service/update"
	"google.golang.org/grpc/codes"
)

type (
	topic string
)

func Dataflow(ctx context.Context) {
	const op = op + ".Dataflow"
	{
		slog.Logger().Debug().Str("op", op).Msg("start")
		defer slog.Logger().Debug().Str("op", op).Msg("end")
	}
	if ctx == nil {
		slog.Logger().Fatal().Str("op", op).
			Err(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"}).Msg("error")
	}
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	tTerminate := terminate.GetTopicStruct(topic("terminate"))
	tUpdate := now.GetTopicTimeTime(topic("update"))

	{
		iCh, oCh := update.GetPipeUpdate(ctx, update.Update, ErrorHandler)
		tTerminate.Publish(ctx, oCh)
		tUpdate.Subscribe(iCh)
	}
	{
		ticker := time.NewTicker(3 * time.Second)
		defer ticker.Stop()
		tUpdate.Publish(ctx, ticker.C)
	}
	{
		iCh, oCh := terminate.GetPipeStruct(ctx, terminate.Terminate, pipe.DefaultErrorHandler)
		tTerminate.Subscribe(iCh)
		<-oCh
	}
}

// ErrorHandler ...
func ErrorHandler(err error) bool {
	const op = op + ".ErrorHandler"
	slog.Logger().Error().Str("op", op).Err(err).Msg("error")
	return false
}
