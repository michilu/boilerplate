package cmd

import (
	"fmt"

	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/semaphore"
	"github.com/michilu/boilerplate/service/slog"
	"github.com/spf13/viper"

	"github.com/michilu/boilerplate/application/flag"
)

type opt struct {
	C string `valid:"fileexists"`
}

func debugOn() {
	const op = op + ".debugOn"

	f := flag.Get()

	e := slog.Logger().Debug()
	if !e.Enabled() {
		return
	}

	e.
		Str("op", op).
		Object("flag", f).
		Str("config", fmt.Sprintf("%v", viper.AllSettings())).
		Msg("config")
}

func setSem() {
	const op = op + ".setSem"
	f := flag.Get()
	err := semaphore.SetParallel(f.Parallel)
	if err != nil {
		slog.Logger().Fatal().Str("op", op).Int("flag.parallel", f.Parallel).Err(&errs.Error{Op: op, Err: err}).Msg("error")
	}
}
