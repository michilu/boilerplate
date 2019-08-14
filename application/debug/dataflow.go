package debug

import (
	"context"

	"github.com/michilu/boilerplate/service/debug"
	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/pipe"
	"github.com/michilu/boilerplate/service/slog"
	"github.com/michilu/boilerplate/service/terminate"
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
	tStruct := terminate.GetTopicStruct(tDebug)
	tClient := debug.GetTopicClientWithCtxer(tDebug)

	config := NewConfiger(ctx, NewClientRepository())
	{
		iCh, oCh := GetPipeConfig(ctx, config.Config, pipe.DefaultErrorHandler)
		tStruct.Subscribe(iCh)
		tClient.Publish(ctx, oCh)
	}
	{
		iCh, oCh := GetPipeConnect(ctx, config.Connect, pipe.DefaultErrorHandler)
		tClient.Subscribe(iCh)
		tStruct.Publish(ctx, oCh)
	}

	tStruct.Publisher(ctx) <- struct{}{}
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

//go:generate genny -in=../../service/pipe/pipe.go -out=gen-pipe-Config.go -pkg=$GOPACKAGE gen "Name=config InT=struct{} OutT=debug.ClientWithCtxer"

func (p *Config) Config(_ struct{}) (debug.ClientWithCtxer, error) {
	const op = op + ".Config.Call"
	{
		slog.Logger().Debug().Str("op", op).Msg("start")
		defer slog.Logger().Debug().Str("op", op).Msg("end")
	}
	v0, err := p.clientRepo.Config(p.ctx)
	if err != nil {
		return nil, err
	}
	return v0, nil
}

//go:generate genny -in=../../service/pipe/pipe.go -out=gen-pipe-Connect.go -pkg=$GOPACKAGE gen "Name=connect InT=debug.ClientWithCtxer OutT=struct{}"

func (p *Config) Connect(m debug.ClientWithCtxer) (struct{}, error) {
	const op = op + ".Connect"
	{
		slog.Logger().Debug().Str("op", op).Msg("start")
		defer slog.Logger().Debug().Str("op", op).Msg("end")
	}
	err := m.Validate()
	if err != nil {
		return struct{}{}, &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
	}
	err = p.clientRepo.Connect(m)
	if err != nil {
		return struct{}{}, err
	}
	return struct{}{}, nil
}
