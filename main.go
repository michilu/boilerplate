package main

import (
	"context"
	_ "net/http/pprof"

	"github.com/michilu/boilerplate/cmd/run"
	"github.com/michilu/boilerplate/cmd/update"
	"github.com/michilu/boilerplate/cmd/version"
	"github.com/michilu/boilerplate/service/pprof"
	"github.com/michilu/boilerplate/service/slog"
	"github.com/spf13/cobra"
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
