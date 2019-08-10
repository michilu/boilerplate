package main

import (
	"context"
	_ "net/http/pprof"

	"github.com/spf13/cobra"

	"github.com/michilu/boilerplate/presentation/cmd/run"
	"github.com/michilu/boilerplate/presentation/cmd/update"
	"github.com/michilu/boilerplate/presentation/cmd/version"
	"github.com/michilu/boilerplate/service/pprof"
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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go pprof.Profile(ctx)
	go pprof.Run()

	cobra.OnInitialize(config)
	for _, f := range subCmd {
		const op = op + ".subCmd"
		c, err := f()
		if err != nil {
			slog.Logger().Fatal().Str("op", op).Err(err).Msg("error")
		}
		rootCmd.AddCommand(c)
	}
	rootCmd.Execute()
}
