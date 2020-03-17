package update

import (
	"context"

	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/pipe"
	"github.com/michilu/boilerplate/service/slog"
	"github.com/michilu/boilerplate/service/terminate"
	"github.com/michilu/boilerplate/service/topic"
	"github.com/michilu/boilerplate/service/update"
	"go.opencensus.io/trace"
	"google.golang.org/grpc/codes"
)

type (
	_topic string
)

func Dataflow(ctx context.Context) {
	const op = op + ".Dataflow"
	if ctx == nil {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"}
		slog.Logger().Err(err).Str("op", op).Msg(err.Error())
		return
	}
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	ctx, s := trace.StartSpan(ctx, op)
	defer s.End()

	t := _topic("update")
	tTick := topic.GetTopicContextContext(t)
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
		{
			s.End()
			<-oCh
		}
	}
}
