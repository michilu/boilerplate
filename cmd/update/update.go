package update

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/michilu/boilerplate/domain"
	"github.com/michilu/boilerplate/v/log"
	"github.com/michilu/boilerplate/v/pipe"
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
		log.Logger().Debug().Str("op", op).Msg("start")
		defer log.Logger().Debug().Str("op", op).Msg("end")
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	t := topic("update")
	tTerminate := domain.GetTopicStruct(t)
	tDurationer := domain.GetTopicDurationer(t)
	tTick := domain.GetTopicTimeTime(t)

	{
		iCh, oCh := domain.GetPipeRestart(ctx, domain.Restart, pipe.DefaultErrorHandler)
		tTerminate.Publish(ctx, oCh)
		tDurationer.Subscribe(iCh)
	}
	{
		iCh, oCh := domain.GetPipeUpdate(ctx, domain.Update, pipe.DefaultErrorHandler)
		tDurationer.Publish(ctx, oCh)
		tTick.Subscribe(iCh)
	}
	{
		iCh, oCh := domain.GetPipeStruct(ctx, domain.Terminate, pipe.DefaultErrorHandler)
		tTerminate.Subscribe(iCh)

		tTick.Publisher(ctx) <- domain.Now()
		<-oCh
	}
}
