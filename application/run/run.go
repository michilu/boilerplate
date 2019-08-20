package run

import (
	"context"

	"github.com/michilu/boilerplate/service/pprof"
	"github.com/michilu/boilerplate/service/slog"
	"github.com/spf13/cobra"

	"github.com/michilu/boilerplate/application/debug"
	"github.com/michilu/boilerplate/application/flag"
)

const (
	op = "application/run"
)

func Run(_ *cobra.Command, _ []string) {
	const op = op + ".Run"
	{
		slog.Logger().Debug().Str("op", op).Msg("start")
		defer slog.Logger().Debug().Str("op", op).Msg("end")
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	f := flag.Get()
	if f.Pprof {
		go pprof.Profile(ctx)
		go pprof.Run()
	}
	go debug.Dataflow(ctx)
	Dataflow(ctx)
}
