package cmd

import (
	"context"
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

type Resource struct {
	Context  context.Context
	Resource []func(context.Context) (io.Closer, error)
	closer   []io.Closer
}

func (p *Resource) Init() error {
	v0 := make([]io.Closer, 0, len(p.Resource))
	for _, v := range p.Resource {
		v1 := v
		v2, err := v1(p.Context)
		if err != nil {
			return err
		}
		if v2 == nil {
			continue
		}
		v0 = append(v0, v2)
	}
	p.closer = v0
	return nil
}

func (p Resource) Close() (err error) {
	for _, v := range p.closer {
		v0 := v
		err = v0.Close()
	}
	return
}

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
		if _, err := os.Stat(s); err == nil {
			viper.SetConfigFile(s)
			if err := viper.ReadInConfig(); err != nil {
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
	resource *Resource,
	defaults []config.KV,
	initCmdFlag func(*cobra.Command),
	subCmd []func() (*cobra.Command, error),
) *cobra.Command {
	const op = op + ".NewCommand"
	initCmdFlag(rootCmd)
	cobra.OnInitialize(func() {
		const op = op + ".cobra.OnInitialize"
		initialize(defaults)
		if resource != nil {
			slog.Logger().Debug().Str("op", op).Strs("os.Args", os.Args).Msg(op + ": value")
			if err := resource.Init(); err != nil {
				const op = op + ".init"
				if os.Args[1] == "run" {
					os.Stderr.WriteString(fmt.Sprintf("op: %s: %s\n", op, err))
					os.Exit(1)
				}
			}
		}
	})
	for _, v := range subCmd {
		const op = op + ".subCmd"
		v0 := v
		c, err := v0()
		if err != nil {
			slog.Logger().Err(err).Str("op", op).Msg(err.Error())
			continue
		}
		rootCmd.AddCommand(c)
	}
	return rootCmd
}
