package run

import (
	"context"

	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/profile"
	"github.com/michilu/boilerplate/service/slog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.opencensus.io/trace"

	"github.com/michilu/boilerplate/application/debug"
	"github.com/michilu/boilerplate/application/event"
	"github.com/michilu/boilerplate/application/exporter"
)

const (
	op = "application/run"
)

func Run(_ *cobra.Command, _ []string) {
	const op = op + ".Run"
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ctx, s := trace.StartSpan(ctx, op)
	defer s.End()

	{
		v0 := "service.profile.profiler.enable"
		v1 := viper.GetBool(v0)
		s.AddAttributes(trace.BoolAttribute(v0, v1))
		if v1 {
			err := profile.RunProfiler(ctx)
			if err != nil {
				const op = op + ".profile.RunProfiler"
				err := &errs.Error{Op: op, Err: err}
				slog.Logger().Error().Str("op", op).Err(err).Msg(err.Error())
			}
		}
	}
	{
		v0 := "service.trace.enable"
		v1 := viper.GetBool(v0)
		s.AddAttributes(trace.BoolAttribute(v0, v1))
		if v1 {
			go exporter.Run()
		}
	}
	{
		v0 := "service.profile.pprof.enable"
		v1 := viper.GetBool(v0)
		s.AddAttributes(trace.BoolAttribute(v0, v1))
		if v1 {
			go profile.Profile(ctx)
			go profile.RunPprof()
		}
	}
	{
		v0 := "application.event.enable"
		v1 := viper.GetBool(v0)
		s.AddAttributes(trace.BoolAttribute(v0, v1))
		if v1 {
			go event.Dataflow(ctx)
		}
	}
	go debug.Dataflow(ctx)
	Dataflow(ctx)
}
