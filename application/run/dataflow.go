package run

import (
	"context"
	"time"

	k "github.com/michilu/boilerplate/application/config"
	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/pipe"
	"github.com/michilu/boilerplate/service/slog"
	"github.com/michilu/boilerplate/service/terminate"
	"github.com/michilu/boilerplate/service/topic"
	"github.com/michilu/boilerplate/service/update"
	"github.com/spf13/viper"
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
	t := slog.Trace(ctx, s)

	{
		const c0 = k.ServiceUpdateEnable
		v0 := viper.GetString(c0)
		s.AddAttributes(trace.StringAttribute(c0, v0))
		v1 := viper.GetBool(c0)
		slog.Logger().Debug().Str("op", op).EmbedObject(t).Bool(c0, v1).Msg(op + ": value")
		if !v1 {
			<-ctx.Done()
			return
		}
	}

	tUpdate := topic.GetTopicContextContext(_topic("update"))
	tTerminate := terminate.GetTopicContextContext(_topic("terminate"))

	{
		iCh, oCh := update.GetPipeUpdate(ctx, update.Update, ErrorHandler)
		tTerminate.Publish(ctx, oCh)
		tUpdate.Subscribe(iCh)
	}
	{
		oCh, err := topic.ContextTicker(ctx, 3*time.Second)
		if err != nil {
			err := &errs.Error{Op: op, Code: codes.Internal, Err: err}
			s.SetStatus(trace.Status{Code: int32(codes.Internal), Message: err.Error()})
			slog.Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
			return
		}
		tUpdate.Publish(ctx, oCh)
	}
	{
		iCh, oCh := terminate.GetPipeTerminate(ctx, terminate.Terminate, pipe.FatalErrorHandler)
		tTerminate.Subscribe(iCh)
		{
			s.End()
			<-oCh
		}
	}
}

// ErrorHandler ...
func ErrorHandler(ctx context.Context, err error) bool {
	defer time.Sleep(5 * time.Minute)
	return pipe.ErrorHandler(ctx, err)
}
