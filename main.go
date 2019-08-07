package main

import (
	"github.com/spf13/cobra"

	"github.com/michilu/boilerplate/cmd/run"
	"github.com/michilu/boilerplate/cmd/update"
	"github.com/michilu/boilerplate/cmd/version"
	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/slog"
)

const (
	op = "main"
)

var (
	subCmd = []func() (*cobra.Command, error){
		run.New,
		update.New,
		version.New,
	}
)

func init() {
	initLog()
	initMeta()
	initCmd()
	initFlag()
}

func initLog() {
	slog.SetTimeFieldFormat()
	slog.SetDefaultLogger()
}

func main() {
	const op = op + ".main"

	cobra.OnInitialize(config)

	for _, f := range subCmd {
		c, err := f()
		if err != nil {
			slog.Logger().Fatal().
				Str("op", op).
				Err(&errs.Error{Op: op, Err: err}).
				Msg("error")
		}
		rootCmd.AddCommand(c)
	}

	rootCmd.Execute()
}
