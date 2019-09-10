package event

import (
	"context"
	"time"

	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/event"
	"github.com/michilu/boilerplate/service/now"
	"github.com/michilu/boilerplate/service/pipe"
	"github.com/michilu/boilerplate/service/slog"
	"github.com/michilu/boilerplate/service/terminate"
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
		slog.Logger().Fatal().Str("op", op).Err(err).Msg("error")
	}
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	ctx, s := trace.StartSpan(ctx, op)
	defer s.End()
	a := make([]trace.Attribute, 0)
	defer s.AddAttributes(a...)

	tStart := terminate.GetTopicContextContext(topic("start"))
	tEvent := GetTopicEventWithContexter(topic("system"))
	tByte := GetTopicByteWithContexter(topic("event"))
	tTerminate := terminate.GetTopicContextContext(topic("terminate"))

	{
		iCh, oCh := GetPipeStart(ctx, Start, ErrorHandler)
		tStart.Subscribe(iCh)
		tEvent.Publish(ctx, oCh)
	}
	{
		iCh, oCh := GetPipeEventLogger(ctx, EventLogger, ErrorHandler)
		tEvent.Subscribe(iCh)
		tByte.Publish(ctx, oCh)
	}
	{
		iCh, oCh := GetPipeSaver(ctx, Saver, ErrorHandler)
		tByte.Subscribe(iCh)
		tTerminate.Publish(ctx, oCh)
	}

	m := context.Background()
	m, _ = trace.StartSpan(m, op)
	tStart.Publisher(ctx) <- m
	<-ctx.Done()
}

// ErrorHandler ...
func ErrorHandler(ctx context.Context, err error) bool {
	defer time.Sleep(5 * time.Minute)
	return pipe.ErrorHandler(ctx, err)
}

//go:generate genny -in=../../service/pipe/pipe.go -out=gen-pipe-Start.go -pkg=$GOPACKAGE gen "Name=start InT=context.Context OutT=EventWithContexter"

func Start(ctx context.Context) (EventWithContexter, error) {
	const op = op + ".Start"
	if ctx == nil {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"}
		return nil, err
	}
	ctx, s := trace.StartSpan(ctx, op)
	defer s.End()
	a := make([]trace.Attribute, 0)
	defer s.AddAttributes(a...)

	v0 := now.Now()
	v1, err := event.NewEvent(&v0, op)
	if err != nil {
		err := &errs.Error{Op: op, Code: codes.Internal, Err: err}
		s.SetStatus(trace.Status{Code: int32(codes.Internal), Message: err.Error()})
		return nil, err
	}
	v2 := &EventWithContext{
		Context: ctx,
		Event:   v1,
	}
	err = v2.Validate()
	if err != nil {
		err := &errs.Error{Op: op, Code: codes.Internal, Err: err}
		s.SetStatus(trace.Status{Code: int32(codes.Internal), Message: err.Error()})
		return nil, err
	}
	a = append(a, trace.StringAttribute("return", v2.String()))
	slog.Logger().Debug().Str("op", op).EmbedObject(slog.Trace(ctx)).EmbedObject(v2).Msg("return")
	return v2, nil
}

//go:generate genny -in=../../service/pipe/pipe.go -out=gen-pipe-EventLogger.go -pkg=$GOPACKAGE gen "Name=eventLogger InT=EventWithContexter OutT=ByteWithContexter"

func EventLogger(m EventWithContexter) (ByteWithContexter, error) {
	const op = op + ".EventLogger"
	if m == nil {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'm' is nil"}
		return nil, err
	}
	ctx := m.GetContext()
	if ctx == nil {
		const op = op + ".Messager.GetContext"
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"}
		return nil, err
	}
	ctx, s := trace.StartSpan(ctx, op)
	defer s.End()
	a := make([]trace.Attribute, 0)
	defer s.AddAttributes(a...)
	slog.Logger().Debug().Str("op", op).EmbedObject(slog.Trace(ctx)).EmbedObject(m).Msg("arg")

	a = append(a, trace.StringAttribute("m", m.String()))
	err := m.Validate()
	if err != nil {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
		s.SetStatus(trace.Status{Code: int32(codes.InvalidArgument), Message: err.Error()})
		return nil, err
	}

	v0, err := event.Marshal(m.GetEvent())
	if err != nil {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
		s.SetStatus(trace.Status{Code: int32(codes.InvalidArgument), Message: err.Error()})
		return nil, err
	}
	v1 := &ByteWithContext{
		Context: ctx,
		Byte: &Byte{
			Byte: v0,
		},
	}
	err = v1.Validate()
	if err != nil {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
		s.SetStatus(trace.Status{Code: int32(codes.InvalidArgument), Message: err.Error()})
		return nil, err
	}
	a = append(a, trace.StringAttribute("return", v1.String()))
	slog.Logger().Debug().Str("op", op).EmbedObject(slog.Trace(ctx)).EmbedObject(v1).Msg("return")
	return v1, nil
}

//go:generate genny -in=../../service/pipe/pipe.go -out=gen-pipe-Saver.go -pkg=$GOPACKAGE gen "Name=saver InT=ByteWithContexter OutT=context.Context"

func Saver(m ByteWithContexter) (context.Context, error) {
	const op = op + ".Saver"
	if m == nil {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'm' is nil"}
		return nil, err
	}
	ctx := m.GetContext()
	if ctx == nil {
		const op = op + ".Messager.GetContext"
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"}
		return nil, err
	}
	ctx, s := trace.StartSpan(ctx, op)
	defer s.End()
	a := make([]trace.Attribute, 0)
	defer s.AddAttributes(a...)
	slog.Logger().Debug().Str("op", op).EmbedObject(slog.Trace(ctx)).EmbedObject(m).Msg("arg")

	a = append(a, trace.StringAttribute("m", m.String()))
	err := m.Validate()
	if err != nil {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
		s.SetStatus(trace.Status{Code: int32(codes.InvalidArgument), Message: err.Error()})
		return nil, err
	}
	return nil, nil
}
