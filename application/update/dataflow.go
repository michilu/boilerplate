package update

import (
	"context"

	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/now"
	"github.com/michilu/boilerplate/service/pipe"
	"github.com/michilu/boilerplate/service/slog"
	"github.com/michilu/boilerplate/service/terminate"
	"github.com/michilu/boilerplate/service/update"
	"go.opencensus.io/trace"
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

	t := topic("update")
	tTick := now.GetTopicContextContext(t)
	tTerminate := terminate.GetTopicContextContext(t)

	{
		iCh, oCh := update.GetPipeUpdate(ctx, update.Update, pipe.FatalErrorHandler)
		tTerminate.Publish(ctx, oCh)
		tTick.Subscribe(iCh)
	}
	{
		iCh, oCh := terminate.GetPipeTerminate(ctx, terminate.Terminate, pipe.FatalErrorHandler)
		tTerminate.Subscribe(iCh)

		m := context.Background()
		m, _ = trace.StartSpan(m, op)
		tTick.Publisher(ctx) <- m
		<-oCh
	}
}
