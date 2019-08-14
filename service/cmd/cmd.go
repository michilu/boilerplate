package cmd

import (
	"os"
	"path/filepath"

	valid "github.com/asaskevich/govalidator"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc/codes"

	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/meta"
	"github.com/michilu/boilerplate/service/slog"
)

const (
	op = "service/cmd"
)

var (
	app *cobra.Command
)

// Init is initializer.
func Init(ns []func() (*cobra.Command, error)) {
	app = &cobra.Command{
		Use:   meta.Name(),
		Short: "A command-line tool that copies the Go files from the bazel-bin directory to anywhere.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return preRunE(cmd, args, flag)
		},
	}
	initFlag()
	cobra.OnInitialize(initialize)
	addCommand(ns)
}

func preRunE(cmd *cobra.Command, args []string, f *flags) error {
	const op = op + ".preRunE"
	ok, err := valid.ValidateStruct(&opt{f.config})
	if err != nil {
		return &errs.Error{Op: op, Err: err}
	}
	if !ok {
		return &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "invalid arguments"}
	}
	return nil
}

func initialize() {
	const op = op + ".initialize"
	f := flag

	if f.debug {
		err := slog.SetLevel("debug")
		if err != nil {
			slog.Logger().Fatal().
				Str("op", op).
				Err(&errs.Error{Op: op, Err: err}).
				Msg("error")
		}
	}

	switch f.config {
	case "":
		viper.AddConfigPath(filepath.Dir(os.Args[0]))
		viper.SetConfigName(meta.Name())
	default:
		viper.SetConfigFile(f.config)
	}

	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	switch err.(type) {
	case nil,
		viper.ConfigFileNotFoundError:
	default:
		slog.Logger().Fatal().
			Str("op", op).
			Err(&errs.Error{Op: op, Err: err}).
			Msg("error")
	}

	slog.Logger().Debug().
		Str("op", op).
		Str("config", viper.ConfigFileUsed()).
		Msg("using config file")

	debugFlag()
	setSem()
}

func addCommand(ns []func() (*cobra.Command, error)) {
	const op = op + ".addCommand"
	for _, n := range ns {
		c, err := n()
		if err != nil {
			slog.Logger().Fatal().
				Str("op", op).
				Err(&errs.Error{Op: op, Err: err}).
				Msg("error")
		}
		app.AddCommand(c)
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the app.
func Execute() {
	const op = op + ".Execute"
	err := app.Execute()
	if err != nil {
		slog.Logger().Fatal().
			Str("op", op).
			Err(err).
			Msg("error")
	}
}
