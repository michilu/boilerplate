package main

import (
	"context"
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
		{K: "service.profile.pprof.addr", V: ":8888"},
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
	Profiler bool
	Trace    bool
	Verbose  bool
}

func initFlag(command *cobra.Command) {
	f := command.PersistentFlags()

	f.StringVar(&flag.Config, "config", "config.toml", "config file")
	viper.BindPFlag("service.config.file", f.Lookup("config"))

	f.BoolVar(&flag.Debug, "debug", false, "open the debug port")
	viper.BindPFlag("service.slog.debug", f.Lookup("debug"))

	f.IntVar(&flag.Parallel, "parallel", runtime.NumCPU(), "specify the maximum number of concurrent")
	viper.BindPFlag("service.semaphore.parallel", f.Lookup("parallel"))

	f.BoolVar(&flag.Pprof, "pprof", false, "launch the Go runtime/pprof")
	viper.BindPFlag("service.profile.pprof.enable", f.Lookup("pprof"))

	f.BoolVar(&flag.Profiler, "profiler", false, "enable the Google Stackdriver Profiler")
	viper.BindPFlag("service.profile.profiler.enable", f.Lookup("profiler"))

	f.BoolVar(&flag.Trace, "trace", false, "enable the Google Stackdriver Trace")
	viper.BindPFlag("service.trace.enable", f.Lookup("trace"))

	f.BoolVar(&flag.Verbose, "verbose", false, "enable verbosely")
	viper.BindPFlag("service.slog.verbose", f.Lookup("verbose"))
}

func main() {
	const op = op + ".main"
	ctx := context.Background()
	ctx, s := trace.StartSpan(ctx, op)
	defer s.End()

	v0 := slog.NewStackdriverZerologWriter(ctx)
	c, closer := cmd.NewCommand(v0.Gen, defaults, initFlag, subCmd)
	defer func() {
		if closer == nil {
			return
		}
		err := closer.Close()
		if err != nil {
			const op = op + ".closer"
			err := &errs.Error{Op: op, Code: codes.Unavailable, Err: err}
			s.SetStatus(trace.Status{Code: int32(codes.Unavailable), Message: err.Error()})
			slog.Logger().Fatal().Str("op", op).Err(err).Msg(err.Error())
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
			slog.Logger().Fatal().Str("op", op).Err(err).Msg(err.Error())
		}
		ch <- struct{}{}
	}()
	sCh := make(chan os.Signal)
	defer close(sCh)
	signal.Notify(sCh, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	select {
	case <-ch:
	case v := <-sCh:
		const op = op + ".signal.Notify"
		s.AddAttributes(trace.StringAttribute("signal", v.String()))
		slog.Logger().Info().Str("op", op).Str("signal", v.String()).Msg("signal")
	}
}
