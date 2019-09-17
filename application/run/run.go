package run

import (
	"context"

	"github.com/michilu/boilerplate/service/pprof"
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
	a := make([]trace.Attribute, 0)
	defer s.AddAttributes(a...)

	{
		v0 := "service.trace.enable"
		v1 := viper.GetBool(v0)
		a = append(a, trace.BoolAttribute(v0, v1))
		if v1 {
			go exporter.Run()
		}
	}
	{
		v0 := "service.pprof.enable"
		v1 := viper.GetBool(v0)
		a = append(a, trace.BoolAttribute(v0, v1))
		if v1 {
			go pprof.Profile(ctx)
			go pprof.Run()
		}
	}
	{
		v0 := "application.event.enable"
		v1 := viper.GetBool(v0)
		a = append(a, trace.BoolAttribute(v0, v1))
		if v1 {
			go event.Dataflow(ctx)
		}
	}
	go debug.Dataflow(ctx)
	Dataflow(ctx)
}
