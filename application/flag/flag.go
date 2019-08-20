package flag

import (
	"runtime"

	"github.com/spf13/cobra"
)

const (
	defaultConfig = "config.toml"
)

var (
	flag *Flag
)

type (
	Flag struct {
		Config   string
		Parallel int
		Pprof    bool
		Verbose  bool
	}
)

func init() {
	flag = &Flag{}
}

func Get() Flag {
	return *flag
}

func GetP() *Flag {
	return flag
}

func InitCmdFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVar(&flag.Config, "config", defaultConfig, "config file")
	cmd.PersistentFlags().IntVarP(&flag.Parallel, "parallel", "p", runtime.NumCPU(), "parallel")
	cmd.PersistentFlags().BoolVar(&flag.Pprof, "pprof", false, "launch pprof")
	cmd.PersistentFlags().BoolVar(&flag.Verbose, "verbose", false, "verbose")
}
