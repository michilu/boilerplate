package main

import (
	_ "net/http/pprof"
	"runtime"

	"github.com/michilu/boilerplate/service/cmd"
	"github.com/michilu/boilerplate/service/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/michilu/boilerplate/presentation/cmd/run"
	"github.com/michilu/boilerplate/presentation/cmd/update"
	"github.com/michilu/boilerplate/presentation/cmd/version"
)

const (
	op = "main"
)

var (
	defaults = []config.KV{
		{K: "service.pprof.addr", V: ":8888"},
		{K: "service.update.channel", V: "release"},
		{K: "service.update.url", V: "http://localhost:8000/"},
	}
	subCmd = []func() (*cobra.Command, error){
		run.New,
		update.New,
		version.New,
	}
	flag = &Flag{}
)

type Flag struct {
	Config   string
	Debug    bool
	Parallel int
	Pprof    bool
	Verbose  bool
}

func initFlag(command *cobra.Command) {
	f := command.PersistentFlags()

	f.StringVar(&flag.Config, "config", "config.toml", "config file")
	viper.BindPFlag("service.config.file", f.Lookup("config"))

	f.BoolVar(&flag.Debug, "debug", false, "debug")
	viper.BindPFlag("service.slog.debug", f.Lookup("debug"))

	f.IntVarP(&flag.Parallel, "parallel", "p", runtime.NumCPU(), "parallel")
	viper.BindPFlag("service.semaphore.parallel", f.Lookup("parallel"))

	f.BoolVar(&flag.Pprof, "pprof", false, "launch pprof")
	viper.BindPFlag("service.pprof.enable", f.Lookup("pprof"))

	f.BoolVar(&flag.Verbose, "verbose", false, "verbose")
	viper.BindPFlag("service.slog.verbose", f.Lookup("verbose"))
}

func main() {
	cmd.NewCommand(defaults, initFlag, subCmd).Execute()
}
