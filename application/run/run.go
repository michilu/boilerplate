package run

import (
	"context"

	k "github.com/michilu/boilerplate/application/config"
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
	t := slog.Trace(ctx)

	{
		v0 := k.ServiceProfileProfilerEnable
		v1 := viper.GetBool(v0)
		s.AddAttributes(trace.BoolAttribute(v0, v1))
		if v1 {
			err := profile.RunProfiler(ctx)
			if err != nil {
				const op = op + ".profile.RunProfiler"
				err := &errs.Error{Op: op, Err: err}
				slog.Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
			}
		}
	}
	{
		v0 := k.ServiceTraceEnable
		v1 := viper.GetBool(v0)
		s.AddAttributes(trace.BoolAttribute(v0, v1))
		if v1 {
			go exporter.Run()
		}
	}
	{
		v0 := k.ServiceProfilePprofEnable
		v1 := viper.GetBool(v0)
		s.AddAttributes(trace.BoolAttribute(v0, v1))
		if v1 {
			go profile.Profile(ctx)
			go profile.RunPprof()
		}
	}
	{
		v0 := k.ApplicationEventEnable
		v1 := viper.GetBool(v0)
		s.AddAttributes(trace.BoolAttribute(v0, v1))
		if v1 {
			go event.Dataflow(ctx)
		}
	}
	go debug.Dataflow(ctx)
	{
		s.End()
		Dataflow(ctx)
	}
}
