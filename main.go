package main

import (
	"context"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	k "github.com/michilu/boilerplate/application/config"
	"github.com/michilu/boilerplate/service/cmd"
	"github.com/michilu/boilerplate/service/config"
	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/pipe"
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
		{K: k.ServiceProfilePprofAddr, V: ":8888"},
		{K: k.ServiceUpdateChannel, V: "release"},
		{K: k.ServiceUpdateUrl, V: "http://localhost:8000/"},
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
	Update   bool
	Verbose  bool
}

func initFlag(command *cobra.Command) {
	f := command.PersistentFlags()

	f.StringVar(&flag.Config, "config", "config.toml", "config file")
	viper.BindPFlag(k.ServiceConfigFile, f.Lookup("config"))

	f.BoolVar(&flag.Debug, "debug", false, "open the debug port")
	viper.BindPFlag(k.ServiceSlogDebug, f.Lookup("debug"))

	f.IntVar(&flag.Parallel, "parallel", runtime.NumCPU(), "specify the maximum number of concurrent")
	viper.BindPFlag(k.ServiceSemaphoreParallel, f.Lookup("parallel"))

	f.BoolVar(&flag.Pprof, "pprof", false, "launch the Go runtime/pprof")
	viper.BindPFlag(k.ServiceProfilePprofEnable, f.Lookup("pprof"))

	f.BoolVar(&flag.Profiler, "profiler", false, "enable the Google Stackdriver Profiler")
	viper.BindPFlag(k.ServiceProfileProfilerEnable, f.Lookup("profiler"))

	f.BoolVar(&flag.Trace, "trace", false, "enable the Google Stackdriver Trace")
	viper.BindPFlag(k.ServiceTraceEnable, f.Lookup("trace"))

	f.BoolVar(&flag.Update, "update", false, "enable the OTA update")
	viper.BindPFlag(k.ServiceUpdateEnable, f.Lookup("update"))

	f.BoolVar(&flag.Verbose, "verbose", false, "enable verbosely")
	viper.BindPFlag(k.ServiceSlogVerbose, f.Lookup("verbose"))
}

func main() {
	const op = op + ".main"
	ctx, s := trace.StartSpan(context.Background(), op)
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
			ctx, s := trace.StartSpan(ctx, op)
			defer s.End()
			t := slog.Trace(ctx)
			err := &errs.Error{Op: op, Code: codes.Unavailable, Err: err}
			s.SetStatus(trace.Status{Code: int32(codes.Unavailable), Message: err.Error()})
			slog.Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
			return
		}
	}()
	pipe.Init(ctx)
	ch := make(chan struct{})
	go func() {
		defer close(ch)
		err := c.Execute()
		if err != nil {
			const op = op + ".cmd.Execute"
			ctx, s := trace.StartSpan(ctx, op)
			defer s.End()
			t := slog.Trace(ctx)
			err := &errs.Error{Op: op, Code: codes.Unknown, Err: err}
			s.SetStatus(trace.Status{Code: int32(codes.Unknown), Message: err.Error()})
			slog.Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
			return
		}
		ch <- struct{}{}
	}()
	sCh := make(chan os.Signal)
	defer close(sCh)
	signal.Notify(sCh, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	s.End()
	select {
	case <-ch:
	case v := <-sCh:
		const op = op + ".signal.Notify"
		ctx, s := trace.StartSpan(ctx, op)
		defer s.End()
		t := slog.Trace(ctx)
		s.AddAttributes(trace.StringAttribute("signal", v.String()))
		slog.Logger().Info().Str("op", op).EmbedObject(t).Str("signal", v.String()).Msg(op + ": signal")
	}
}
