package main

import (
	_ "net/http/pprof"

	"github.com/michilu/boilerplate/service/cmd"
	"github.com/michilu/boilerplate/service/config"
	"github.com/spf13/cobra"

	"github.com/michilu/boilerplate/application/flag"
	"github.com/michilu/boilerplate/presentation/cmd/run"
	"github.com/michilu/boilerplate/presentation/cmd/update"
	"github.com/michilu/boilerplate/presentation/cmd/version"
)

const (
	op = "main"
)

var (
	defaults = []config.KV{
		{"service.pprof.addr", ":8888"},
		{"service.update.channel", "release"},
		{"service.update.url", "http://localhost:8000/"},
	}
	subCmd = []func() (*cobra.Command, error){
		run.New,
		update.New,
		version.New,
	}
)

func main() {
	cmd.NewCommand(defaults, flag.InitCmdFlag, subCmd).Execute()
}
