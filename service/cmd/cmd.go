package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"

	k "github.com/michilu/boilerplate/application/config"
	"github.com/michilu/boilerplate/service/config"
	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/meta"
	"github.com/michilu/boilerplate/service/semaphore"
	"github.com/michilu/boilerplate/service/slog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc/codes"
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
	semaphore.Init(viper.GetInt(k.ServiceSemaphoreParallel))
	{
		s := viper.GetString(k.ServiceConfigFile)
		_, err := os.Stat(s)
		if err == nil {
			viper.SetConfigFile(s)
			err := viper.ReadInConfig()
			if err != nil {
				const op = op + ".viper.ReadInConfig"
				err := &errs.Error{Op: op, Code: codes.Internal, Err: err, Message: fmt.Sprintf("check the format of '%s'", s)}
				os.Stderr.WriteString(fmt.Sprintf("error: op: %s: %s\n", op, err))
				os.Exit(1)
			}
		}
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(
			"-", "_",
			".", "_",
		))
		config.SetDefault(v)
	}
}

func NewCommand(
	logger func() ([]io.Writer, slog.Closer, error),
	defaults []config.KV,
	initCmdFlag func(*cobra.Command),
	subCmd []func() (*cobra.Command, error),
) (*cobra.Command, slog.Closer) {
	const op = op + ".NewCommand"
	var closer slog.Closer
	initCmdFlag(rootCmd)
	cobra.OnInitialize(func() {
		const op = op + ".cobra.OnInitialize"
		initialize(defaults)
		var (
			w   []io.Writer
			err error
		)
		if logger != nil {
			slog.Logger().Debug().Str("op", op).Strs("os.Args", os.Args).Msg(op + ": value")
			w, closer, err = logger()
			if err != nil {
				const op = op + ".logger"
				if os.Args[1] == "run" {
					os.Stderr.WriteString(fmt.Sprintf("op: %s: %s\n", op, err))
					os.Exit(1)
				}
			}
		}
		{
			slog.SetDefaultLogger(w)
			slog.Logger().Debug().Str("op", op).Str("file", viper.ConfigFileUsed()).Msg(op + ": config")
			slog.Logger().Debug().Str("op", op).Interface("viper", viper.AllSettings()).Msg(op + ": config")
		}
	})
	for _, f := range subCmd {
		const op = op + ".subCmd"
		c, err := f()
		if err != nil {
			slog.Logger().Fatal().Err(err).Str("op", op).Msg(err.Error())
		}
		rootCmd.AddCommand(c)
	}
	return rootCmd, closer
}
