package echo

import (
	"context"

	valid "github.com/asaskevich/govalidator"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc/codes"

	"github.com/michilu/boilerplate/v/errs"
	"github.com/michilu/boilerplate/v/log"
	"github.com/michilu/boilerplate/v/pipe"
)

type (
	topic string

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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	inCh, outCh := GetPipeString(ctx, echo, pipe.DefaultErrorHandler)

	t := GetTopicString(topic("echo"))
	t.Subscribe(inCh)
	//t.Publish(ctx, outCh)

	t.Publisher(ctx) <- f.filepath
	<-outCh
}

//go:generate genny -in=../../v/pipe/pipe.go -out=gen-pipe-$GOFILE -pkg=$GOPACKAGE gen "Name=string InT=string OutT=string"
//go:generate genny -in=../../v/topic/topic.go -out=gen-topic-$GOFILE -pkg=$GOPACKAGE gen "ChanT=string"

func echo(s string) (string, error) {
	const op = "cmd.echo.echo"
	log.Logger().Info().
		Str("op", op).
		Msg(s)
	return s, nil
}
