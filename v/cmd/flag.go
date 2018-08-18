package cmd

import (
	"fmt"
	"runtime"

	"github.com/michilu/boilerplate/v/errs"
	"github.com/michilu/boilerplate/v/log"
	"github.com/michilu/boilerplate/v/meta"
	"github.com/michilu/boilerplate/v/semaphore"
)

var (
	flag *flags
)

type (
	flags struct {
		config   string
		debug    bool
		parallel int
	}

	opt struct {
		C string `valid:"fileexists"`
	}
)

func initFlag() {
	flag = &flags{}
	f := flag
	app.PersistentFlags().StringVar(&f.config, "config", "", fmt.Sprintf("config file (default is %s.yaml)", meta.Name()))
	app.PersistentFlags().BoolVar(&f.debug, "debug", false, "debug mode")
	app.PersistentFlags().IntVarP(&f.parallel, "parallel", "p", runtime.NumCPU(), "parallel")
}

func debugFlag() {
	const op = "cmd.debugFlag"

	e := log.Logger().Debug()
	if !e.Enabled() {
		return
	}

	f := flag
	e.
		Str("op", op).
		Str("config", f.config).
		Bool("debug", f.debug).
		Int("parallel", f.parallel).
		Msg("flag")
}

func setSem() {
	const op = "cmd.setSem"

	err := semaphore.SetParallel(flag.parallel)
	if err != nil {
		log.Logger().Fatal().
			Str("op", op).
			Int("flag.parallel", flag.parallel).
			Err(&errs.Error{Op: op, Err: err}).
			Msg("error")
	}
}
