package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"

	cfg "github.com/michilu/boilerplate/service/config"
	"github.com/michilu/boilerplate/service/slog"
)

func config() {
	const op = op + ".config"
	f := flag

	if f.debug {
		err := slog.SetLevel("debug")
		if err != nil {
			const op = op + ".slog.SetLevel"
			slog.Logger().Fatal().Str("op", op).Err(err).Msg("error")
		}
	}

	{
		_, err := os.Stat(f.config)
		if err == nil {
			viper.SetConfigFile(f.config)
			err := viper.ReadInConfig()
			if err != nil {
				const op = op + ".viper.ReadInConfig"
				slog.Logger().Fatal().Str("op", op).Err(err).Msg("error")
			}
		}
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(
			"-", "_",
			".", "_",
		))
		cfg.SetDefault()
	}

	slog.Logger().Debug().Str("op", op).Str("config", viper.ConfigFileUsed()).Msg("using config file")
	debugOn()
}

func debugOn() {
	const op = op + ".debugOn"

	e := slog.Logger().Debug()
	if !e.Enabled() {
		return
	}

	f := flag
	e.
		Str("op", op).
		Bool("debug", f.debug).
		Str("config", fmt.Sprintf("%v", viper.AllSettings())).
		Msg("flag")
}
