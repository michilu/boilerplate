package echo

import (
	valid "github.com/asaskevich/govalidator"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc/codes"

	"github.com/michilu/boilerplate/v/bus"
	"github.com/michilu/boilerplate/v/errs"
	"github.com/michilu/boilerplate/v/log"
)

const (
	topic = "echo"
)

type (
	flag struct {
		filepath string
	}

	opt struct {
		F string `valid:"fileexists"`
	}
)

// New returns a new command.
func New() (*cobra.Command, error) {
	const op = "cmd.echo.new"
	f := &flag{}
	c := &cobra.Command{
		Use:   "echo",
		Short: "echo",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return preRunE(cmd, args, f)
		},
		Run: func(cmd *cobra.Command, args []string) {
			run(cmd, args, f)
		},
	}
	c.Flags().StringVarP(&f.filepath, "file", "f", "", "path to an exists file")
	err := viper.BindPFlag("file", c.Flags().Lookup("file"))
	if err != nil {
		return nil, &errs.Error{Op: op, Err: err}
	}
	return c, nil
}

func preRunE(cmd *cobra.Command, args []string, f *flag) error {
	const op = "cmd.echo.preRunE"
	ok, err := valid.ValidateStruct(&opt{f.filepath})
	if err != nil {
		return &errs.Error{Op: op, Err: err}
	}
	if !ok {
		return &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "invalid arguments"}
	}
	return nil
}

func run(cmd *cobra.Command, args []string, f *flag) {
	const op = "cmd.echo.run"

	err := bus.Subscribe(topic, echo)
	if err != nil {
		log.Logger().Fatal().
			Str("op", op).
			Err(&errs.Error{Op: op, Err: err}).
			Msg("error")
	}
	bus.Publish(topic, f.filepath)
}

func echo(s string) {
	const op = "cmd.echo.echo"

	defer func() {
		const op = "cmd.echo.echo#defer"
		err := bus.Unsubscribe(topic, echo)
		if err != nil {
			log.Logger().Fatal().
				Str("op", op).
				Err(&errs.Error{Op: op, Err: err}).
				Msg("error")
		}
	}()

	log.Logger().Debug().
		Str("op", op).
		Str("s", s).
		Msg("echo a file")

	log.Logger().Info().
		Str("op", op).
		Msg(s)
}
