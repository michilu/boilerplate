package update

import (
	"context"

	"github.com/michilu/boilerplate/service/slog"
	"github.com/spf13/cobra"
)

const (
	op = "application/update"
)

func Run(_ *cobra.Command, _ []string) {
	const op = op + ".Run"
	{
		slog.Logger().Debug().Str("op", op).Msg("start")
		defer slog.Logger().Debug().Str("op", op).Msg("end")
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	Dataflow(ctx)
}
