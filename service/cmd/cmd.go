package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/michilu/boilerplate/service/config"
	"github.com/michilu/boilerplate/service/meta"
	"github.com/michilu/boilerplate/service/semaphore"
	"github.com/michilu/boilerplate/service/slog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	semaphore.Init(viper.GetInt("service.semaphore.parallel"))
	{
		s := viper.GetString("service.config.file")
		_, err := os.Stat(s)
		if err == nil {
			viper.SetConfigFile(s)
			err := viper.ReadInConfig()
			if err != nil {
				const op = op + ".viper.ReadInConfig"
				slog.Logger().Fatal().Str("op", op).Err(err).Msg("error")
			}
			slog.Logger().Debug().Str("op", op).Str("file", viper.ConfigFileUsed()).Msg("config")
		}
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(
			"-", "_",
			".", "_",
		))
		config.SetDefault(v)
	}
	slog.Logger().Debug().Str("op", op).Interface("viper", viper.AllSettings()).Msg("config")
}

func NewCommand(
	logger func() ([]io.Writer, func() error, error),
	defaults []config.KV,
	initCmdFlag func(*cobra.Command),
	subCmd []func() (*cobra.Command, error),
) (*cobra.Command, *func() error) {
	const op = op + ".NewCommand"
	var closer func() error
	initCmdFlag(rootCmd)
	cobra.OnInitialize(func() {
		const op = op + ".cobra.OnInitialize"
		initialize(defaults)
		var (
			w   []io.Writer
			err error
		)
		if logger != nil {
			w, closer, err = logger()
			if err != nil {
				const op = op + ".logger"
				os.Stderr.WriteString(fmt.Sprintf("op: %s: %s\n", op, err))
			}
		}
		slog.SetDefaultLogger(w)
	})
	for _, f := range subCmd {
		const op = op + ".subCmd"
		c, err := f()
		if err != nil {
			slog.Logger().Fatal().Str("op", op).Err(err).Msg("error")
		}
		rootCmd.AddCommand(c)
	}
	return rootCmd, &closer
}