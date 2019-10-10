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

	tUpdate := now.GetTopicContextContext(topic("update"))
	tTerminate := terminate.GetTopicContextContext(topic("terminate"))

	{
		iCh, oCh := update.GetPipeUpdate(ctx, update.Update, ErrorHandler)
		tTerminate.Publish(ctx, oCh)
		tUpdate.Subscribe(iCh)
	}
	{
		oCh := now.ContextTicker(ctx, 3*time.Second)
		tUpdate.Publish(ctx, oCh)
	}
	{
		iCh, oCh := terminate.GetPipeTerminate(ctx, terminate.Terminate, pipe.FatalErrorHandler)
		tTerminate.Subscribe(iCh)
		<-oCh
	}
}

// ErrorHandler ...
func ErrorHandler(ctx context.Context, err error) bool {
	defer time.Sleep(5 * time.Minute)
	return pipe.ErrorHandler(ctx, err)
}
