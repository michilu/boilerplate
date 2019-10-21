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
	if ctx == nil {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"}
		slog.Logger().Fatal().Str("op", op).Err(err).Msg(err.Error())
	}
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	ctx, s := trace.StartSpan(ctx, op)
	defer s.End()

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
		{
			m, _ := trace.StartSpan(context.Background(), op)
			tTick.Publisher(ctx) <- m
		}
		s.End()
		<-oCh
	}
}
