package update

import (
	"context"

	"github.com/michilu/boilerplate/service/now"
	"github.com/michilu/boilerplate/service/pipe"
	"github.com/michilu/boilerplate/service/slog"
	"github.com/michilu/boilerplate/service/terminate"
	"github.com/michilu/boilerplate/service/update"
	"github.com/spf13/cobra"
)

const (
	op = "cmd.update"
)

type (
	topic string
)

// New returns a new command.
func New() (*cobra.Command, error) {
	c := &cobra.Command{
		Use:   "update",
		Short: "update",
		Run:   run,
	}
	return c, nil
}

func run(cmd *cobra.Command, args []string) {
	const op = op + ".run"
	{
		slog.Logger().Debug().Str("op", op).Msg("start")
		defer slog.Logger().Debug().Str("op", op).Msg("end")
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	t := topic("update")
	tTerminate := terminate.GetTopicStruct(t)
	tTick := now.GetTopicTimeTime(t)

	{
		iCh, oCh := update.GetPipeUpdate(ctx, update.Update, pipe.DefaultErrorHandler)
		tTerminate.Publish(ctx, oCh)
		tTick.Subscribe(iCh)
	}
	{
		iCh, oCh := terminate.GetPipeStruct(ctx, terminate.Terminate, pipe.DefaultErrorHandler)
		tTerminate.Subscribe(iCh)

		tTick.Publisher(ctx) <- now.Now()
		<-oCh
	}
}
