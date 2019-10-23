package event

import (
	"context"
	"encoding/json"
	"time"

	"github.com/michilu/boilerplate/infra/keyvalue"
	infra "github.com/michilu/boilerplate/infra/nutsdb/event"
	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/event"
	"github.com/michilu/boilerplate/service/now"
	"github.com/michilu/boilerplate/service/pipe"
	"github.com/michilu/boilerplate/service/slog"
	"github.com/michilu/boilerplate/service/terminate"
	_ "github.com/michilu/boilerplate/service/topic"
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
		slog.Logger().Fatal().Err(err).Str("op", op).Msg(err.Error())
	}
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	ctx, s := trace.StartSpan(ctx, op)
	defer s.End()
	t := slog.Trace(ctx)

	v0, v1, err := infra.NewRepository()
	if err != nil {
		err := &errs.Error{Op: op, Code: codes.Internal, Err: err}
		s.SetStatus(trace.Status{Code: int32(codes.Internal), Message: err.Error()})
		slog.Logger().Fatal().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
		return
	}
	defer func(ctx context.Context, v1 func() error) {
		const op = op + ".closer"
		err := v1()
		if err != nil {
			err := &errs.Error{Op: op, Code: codes.Internal, Err: err}
			s.SetStatus(trace.Status{Code: int32(codes.Internal), Message: err.Error()})
			slog.Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
			return
		}
	}(ctx, v1)

	tStart := terminate.GetTopicContextContext(topic("start"))
	tEvent := event.GetTopicEventWithContexter(topic("system"))
	tKeyValue := event.GetTopicKeyValueWithContexter(topic("event"))
	tTerminate := terminate.GetTopicContextContext(topic("terminate"))

	{
		iCh, oCh := GetPipeStart(ctx, Start, ErrorHandler)
		tStart.Subscribe(iCh)
		tEvent.Publish(ctx, oCh)
	}
	{
		iCh, oCh := GetPipeEventLogger(ctx, EventLogger, ErrorHandler)
		tEvent.Subscribe(iCh)
		tKeyValue.Publish(ctx, oCh)
	}
	{
		v1 := &Saver{Saver: v0}
		iCh, oCh := GetPipeSaver(ctx, v1.Save, ErrorHandler)
		tKeyValue.Subscribe(iCh)
		tTerminate.Publish(ctx, oCh)
	}

	{
		m, _ := trace.StartSpan(context.Background(), op)
		tStart.Publisher(ctx) <- m
	}
	s.End()
	<-ctx.Done()
}

// ErrorHandler ...
func ErrorHandler(ctx context.Context, err error) bool {
	defer time.Sleep(5 * time.Minute)
	return pipe.ErrorHandler(ctx, err)
}

//go:generate genny -in=../../service/pipe/pipe.go -out=gen-pipe-Start.go -pkg=$GOPACKAGE gen "Name=start InT=context.Context OutT=event.EventWithContexter"

func Start(ctx context.Context) (event.EventWithContexter, error) {
	const op = op + ".Start"
	if ctx == nil {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"}
		return nil, err
	}
	ctx, s := trace.StartSpan(ctx, op)
	defer s.End()
	t := slog.Trace(ctx)

	v0 := now.Now()
	v1, err := event.NewEvent(&v0, op)
	if err != nil {
		const op = op + ".event.NewEvent"
		err := &errs.Error{Op: op, Code: codes.Internal, Err: err}
		s.SetStatus(trace.Status{Code: int32(codes.Internal), Message: err.Error()})
		slog.Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
		return nil, err
	}
	v2 := &event.EventWithContext{
		Context: ctx,
		Event:   v1,
	}
	{
		err := v2.Validate()
		if err != nil {
			err := &errs.Error{Op: op, Code: codes.Internal, Err: err}
			s.SetStatus(trace.Status{Code: int32(codes.Internal), Message: err.Error()})
			slog.Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
			return nil, err
		}
	}
	s.AddAttributes(trace.StringAttribute("return", v2.String()))
	slog.Logger().Debug().Str("op", op).EmbedObject(t).EmbedObject(v2).Msg("return")
	return v2, nil
}

//go:generate genny -in=../../service/pipe/pipe.go -out=gen-pipe-EventLogger.go -pkg=$GOPACKAGE gen "Name=eventLogger InT=event.EventWithContexter OutT=event.KeyValueWithContexter"

func EventLogger(m event.EventWithContexter) (event.KeyValueWithContexter, error) {
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
	t := slog.Trace(ctx)

	s.AddAttributes(trace.StringAttribute("m", m.String()))
	{
		err := m.Validate()
		if err != nil {
			err := &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
			s.SetStatus(trace.Status{Code: int32(codes.InvalidArgument), Message: err.Error()})
			slog.Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
			return nil, err
		}
		slog.Logger().Debug().Str("op", op).EmbedObject(t).EmbedObject(m).Msg("arg")
	}

	v0, err := json.Marshal(m.GetEvent())
	if err != nil {
		const op = op + ".json.Marshal"
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
		s.SetStatus(trace.Status{Code: int32(codes.InvalidArgument), Message: err.Error()})
		slog.Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
		return nil, err
	}
	v1 := &event.KeyValueWithContext{
		Context: ctx,
		KeyValue: &keyvalue.KeyValue{
			Key:   m.GetEvent().GetKey(),
			Value: v0,
		},
	}
	{
		err := v1.Validate()
		if err != nil {
			err := &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
			s.SetStatus(trace.Status{Code: int32(codes.InvalidArgument), Message: err.Error()})
			slog.Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
			return nil, err
		}
	}
	s.AddAttributes(trace.StringAttribute("return", v1.String()))
	slog.Logger().Debug().Str("op", op).EmbedObject(t).EmbedObject(v1).Msg("return")
	return v1, nil
}

//go:generate genny -in=../../service/pipe/pipe.go -out=gen-pipe-Saver.go -pkg=$GOPACKAGE gen "Name=saver InT=event.KeyValueWithContexter OutT=context.Context"

type Saver struct {
	Saver event.Saver
}

func (p *Saver) Save(m event.KeyValueWithContexter) (context.Context, error) {
	const op = op + ".Saver.Save"
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
	t := slog.Trace(ctx)
	{
		err := m.Validate()
		if err != nil {
			err := &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
			s.SetStatus(trace.Status{Code: int32(codes.InvalidArgument), Message: err.Error()})
			slog.Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
			return nil, err
		}
		s.AddAttributes(trace.StringAttribute("m", m.String()))
		slog.Logger().Debug().Str("op", op).EmbedObject(t).EmbedObject(m).Msg("arg")
	}
	{
		err := event.SaveEventPayload(ctx, p.Saver, m.GetKeyValue())
		if err != nil {
			const op = op + ".SaveEventPayload"
			err := &errs.Error{Op: op, Code: codes.Internal, Err: err}
			s.SetStatus(trace.Status{Code: int32(codes.Internal), Message: err.Error()})
			slog.Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
			return nil, err
		}
	}
	return nil, nil
}
