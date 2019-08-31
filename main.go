package main

import (
	"context"
	"io"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/michilu/boilerplate/service/cmd"
	"github.com/michilu/boilerplate/service/config"
	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/slog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.opencensus.io/trace"
	"google.golang.org/grpc/codes"

	"github.com/michilu/boilerplate/presentation/cmd/run"
	"github.com/michilu/boilerplate/presentation/cmd/update"
	"github.com/michilu/boilerplate/presentation/cmd/version"
)

const (
	op = "main"
)

var (
	defaults = []config.KV{
		{K: "service.pprof.addr", V: ":8888"},
		{K: "service.update.channel", V: "release"},
		{K: "service.update.url", V: "http://localhost:8000/"},
	}
	subCmd = []func() (*cobra.Command, error){
		run.New,
		update.New,
		version.New,
	}
	flag = &Flag{}
)

type Flag struct {
	Config   string
	Debug    bool
	Parallel int
	Pprof    bool
	Trace    bool
	Verbose  bool
}

func initFlag(command *cobra.Command) {
	f := command.PersistentFlags()

	f.StringVar(&flag.Config, "config", "config.toml", "config file")
	viper.BindPFlag("service.config.file", f.Lookup("config"))

	f.BoolVar(&flag.Debug, "debug", false, "debug")
	viper.BindPFlag("service.slog.debug", f.Lookup("debug"))

	f.IntVarP(&flag.Parallel, "parallel", "p", runtime.NumCPU(), "parallel")
	viper.BindPFlag("service.semaphore.parallel", f.Lookup("parallel"))

	f.BoolVar(&flag.Pprof, "pprof", false, "launch pprof")
	viper.BindPFlag("service.pprof.enable", f.Lookup("pprof"))

	f.BoolVar(&flag.Trace, "trace", false, "trace")
	viper.BindPFlag("service.trace.enable", f.Lookup("trace"))

	f.BoolVar(&flag.Verbose, "verbose", false, "verbose")
	viper.BindPFlag("service.slog.verbose", f.Lookup("verbose"))
}

func logger(ctx context.Context) (func() ([]io.Writer, func() error, error), error) {
	const op = op + ".logger"
	if ctx == nil {
		return nil, &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"}
	}

	return func() ([]io.Writer, func() error, error) {
		const op = op + "#func"
		ctx, s := trace.StartSpan(ctx, op)
		defer s.End()
		a := make([]trace.Attribute, 0)
		defer s.AddAttributes(a...)

		writer, client, err := slog.NewStackdriverLogging(
			ctx,
			viper.GetString("gcp.project.id"),
			viper.GetString("gcp.logging.id"),
			nil,
		)
		if err != nil {
			const op = op + ".slog.NewStackdriverLogging"
			err := &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
			s.SetStatus(trace.Status{Code: int32(codes.InvalidArgument), Message: err.Error()})
			return nil, nil, err
		}
		slog.SetDefaultTracer(writer)
		closer := func() error {
			const op = op + ".closer"
			slog.Logger().Info().Str("op", op).Msg("start clean up")
			err := client.Close()
			if err != nil {
				const op = op + ".client.Close"
				return &errs.Error{Op: op, Code: codes.Unavailable, Err: err}
			}
			slog.Logger().Info().Str("op", op).Msg("cleaned up")
			return nil
		}
		return []io.Writer{writer}, closer, nil
	}, nil
}

func main() {
	const op = op + ".main"
	ctx := context.Background()
	ctx, s := trace.StartSpan(ctx, op)
	defer s.End()
	a := make([]trace.Attribute, 0)
	defer s.AddAttributes(a...)

	f, err := logger(ctx)
	if err != nil {
		const op = op + ".logger"
		err := &errs.Error{Op: op, Code: codes.Internal, Err: err}
		s.SetStatus(trace.Status{Code: int32(codes.Internal), Message: err.Error()})
		slog.Logger().Fatal().Str("op", op).Err(err).Msg("error")
	}
	c, closer := cmd.NewCommand(f, defaults, initFlag, subCmd)
	defer func() {
		if closer == nil {
			return
		}
		if *closer == nil {
			return
		}
		err := (*closer)()
		if err != nil {
			const op = op + ".closer"
			err := &errs.Error{Op: op, Code: codes.Unavailable, Err: err}
			s.SetStatus(trace.Status{Code: int32(codes.Unavailable), Message: err.Error()})
			slog.Logger().Fatal().Str("op", op).Err(err).Msg("error")
		}
	}()
	ch := make(chan struct{})
	go func() {
		defer close(ch)
		err := c.Execute()
		if err != nil {
			const op = op + ".cmd.Execute"
			err := &errs.Error{Op: op, Code: codes.Unknown, Err: err}
			s.SetStatus(trace.Status{Code: int32(codes.Unknown), Message: err.Error()})
			slog.Logger().Fatal().Str("op", op).Err(err).Msg("error")
		}
		ch <- struct{}{}
	}()
	sCh := make(chan os.Signal)
	signal.Notify(sCh, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	select {
	case <-ch:
	case v := <-sCh:
		const op = op + ".signal.Notify"
		a = append(a, trace.StringAttribute("signal", v.String()))
		slog.Logger().Info().Str("op", op).Str("signal", v.String()).Msg("signal")
	}
}
