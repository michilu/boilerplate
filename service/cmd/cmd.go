package cmd

import (
	"os"
	"strings"

	"github.com/michilu/boilerplate/service/config"
	"github.com/michilu/boilerplate/service/meta"
	"github.com/michilu/boilerplate/service/slog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/michilu/boilerplate/application/flag"
)

const (
	op = "service/cmd"
)

var (
	rootCmd *cobra.Command
)

func init() {
	rootCmd = &cobra.Command{
		Use: meta.Name(),
	}
}

func initialize(v []config.KV) {
	const op = op + ".initialize"

	slog.Init()

	{
		f := flag.Get()
		_, err := os.Stat(f.Config)
		if err == nil {
			viper.SetConfigFile(f.Config)
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
		config.SetDefault(v)
	}

	slog.Logger().Debug().Str("op", op).Str("config", viper.ConfigFileUsed()).Msg("using config file")

	debugOn()
	setSem()
}

func NewCommand(
	defaults []config.KV,
	initCmdFlag func(*cobra.Command),
	subCmd []func() (*cobra.Command, error),
) *cobra.Command {
	const op = op + ".NewCommand"
	initCmdFlag(rootCmd)
	cobra.OnInitialize(func() { initialize(defaults) })
	for _, f := range subCmd {
		const op = op + ".subCmd"
		c, err := f()
		if err != nil {
			slog.Logger().Fatal().Str("op", op).Err(err).Msg("error")
		}
		rootCmd.AddCommand(c)
	}
	return rootCmd
}
