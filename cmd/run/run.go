package run

import (
	"context"
	"time"

	"github.com/michilu/boilerplate/service/now"
	"github.com/michilu/boilerplate/service/pipe"
	"github.com/michilu/boilerplate/service/slog"
	"github.com/michilu/boilerplate/service/terminate"
	"github.com/michilu/boilerplate/service/update"
	"github.com/spf13/cobra"
)

const (
	op = "cmd.run"
)

type (
	topic string
)

// New returns a new command.
func New() (*cobra.Command, error) {
	c := &cobra.Command{
		Use:   "run",
		Short: "run",
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

	tTerminate := terminate.GetTopicStruct(topic("terminate"))
	tUpdate := now.GetTopicTimeTime(topic("update"))

	{
		iCh, oCh := update.GetPipeUpdate(ctx, update.Update, ErrorHandler)
		tTerminate.Publish(ctx, oCh)
		tUpdate.Subscribe(iCh)
	}
	{
		ticker := time.NewTicker(3 * time.Second)
		defer ticker.Stop()
		tUpdate.Publish(ctx, ticker.C)
	}
	{
		iCh, oCh := terminate.GetPipeStruct(ctx, terminate.Terminate, pipe.DefaultErrorHandler)
		tTerminate.Subscribe(iCh)
		<-oCh
	}
}

// ErrorHandler ...
func ErrorHandler(err error) bool {
	const op = op + ".ErrorHandler"
	slog.Logger().Error().Str("op", op).Err(err).Msg("error")
	return false
}
