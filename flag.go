package main

import (
	"github.com/michilu/boilerplate/service/slog"
)

var (
	flag *flags
)

type (
	flags struct {
		config   string
		ellipsis bool
		verbose  bool
	}
)

func initFlag() {
	const op = op + ".initFlag"
	if rootCmd == nil {
		slog.Logger().Fatal().Str("op", op).Msg("must be initialize `main.rootCmd` by initCmd")
	}
	flag = &flags{}
	f := flag
	defaultConfig := "config.toml"
	rootCmd.PersistentFlags().StringVar(&f.config, "config", defaultConfig, "config file")
	rootCmd.PersistentFlags().BoolVar(&f.verbose, "verbose", false, "verbose")
	rootCmd.PersistentFlags().BoolVar(&f.ellipsis, "ellipsis", false, "print ellipsis")
}
