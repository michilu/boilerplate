package debug

import (
	"context"
	"time"

	"github.com/michilu/boilerplate/service/debug"
	"github.com/michilu/boilerplate/service/errs"
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

	tDebug := topic("debug")
	tCtx := terminate.GetTopicContextContext(tDebug)
	tClient := debug.GetTopicClientWithContexter(tDebug)

	config := NewConfiger(ctx, NewClientRepository())
	{
		iCh, oCh := GetPipeConfig(ctx, config.Config, pipe.FatalErrorHandler)
		tCtx.Subscribe(iCh)
		tClient.Publish(ctx, oCh)
	}
	{
		iCh, oCh := GetPipeConnect(ctx, config.Connect, ErrorHandler)
		tClient.Subscribe(iCh)
		tCtx.Publish(ctx, oCh)
	}

	m := context.Background()
	m, _ = trace.StartSpan(m, op)
	tCtx.Publisher(ctx) <- m
	<-ctx.Done()
}

// NewConfiger returns a new Configer
func NewConfiger(ctx context.Context, clientRepo debug.ClientRepository) Configer {
	return &Config{
		ctx:        ctx,
		clientRepo: clientRepo,
	}
}

//go:generate interfacer -for github.com/michilu/boilerplate/application/debug.Config -as debug.Configer -o if-Configer.go

// Config ...
type Config struct {
	ctx        context.Context
	clientRepo debug.ClientRepository
}

//go:generate genny -in=../../service/pipe/pipe.go -out=gen-pipe-Config.go -pkg=$GOPACKAGE gen "Name=config InT=context.Context OutT=debug.ClientWithContexter"

func (p *Config) Config(ctx context.Context) (debug.ClientWithContexter, error) {
	const op = op + ".Config.Call"
	if ctx == nil {
		return nil, &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"}
	}
	ctx, s := trace.StartSpan(ctx, op)
	defer s.End()
	a := make([]trace.Attribute, 0)
	defer s.AddAttributes(a...)

	v0, err := p.clientRepo.Config(ctx)
	if err != nil {
		s.SetStatus(trace.Status{Code: int32(codes.Unknown), Message: err.Error()})
		return nil, err
	}
	a = append(a, trace.StringAttribute("v0", v0.String()))
	return v0, nil
}

//go:generate genny -in=../../service/pipe/pipe.go -out=gen-pipe-Connect.go -pkg=$GOPACKAGE gen "Name=connect InT=debug.ClientWithContexter OutT=context.Context"

func (p *Config) Connect(m debug.ClientWithContexter) (context.Context, error) {
	const op = op + "Config.Connect"
	if m == nil {
		return nil, &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'm' is nil"}
	}
	ctx := m.GetContext()
	if ctx == nil {
		return nil, &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"}
	}
	ctx, s := trace.StartSpan(ctx, op)
	defer s.End()
	a := make([]trace.Attribute, 0)
	defer s.AddAttributes(a...)

	a = append(a, trace.StringAttribute("m", m.String()))
	err := m.Validate()
	if err != nil {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
		s.SetStatus(trace.Status{Code: int32(codes.InvalidArgument), Message: err.Error()})
		return nil, err
	}
	err = p.clientRepo.Connect(m)
	ctx = context.Background()
	if err != nil {
		s.SetStatus(trace.Status{Code: int32(codes.Unknown), Message: err.Error()})
		return ctx, err
	}
	ctx, _ = trace.StartSpan(ctx, op)
	return ctx, nil
}

// ErrorHandler ...
func ErrorHandler(ctx context.Context, err error) bool {
	defer time.Sleep(5 * time.Minute)
	return pipe.ErrorHandler(ctx, err)
}
