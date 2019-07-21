package main

import (
	"github.com/spf13/cobra"

	"github.com/michilu/boilerplate/v/cmd"

	"github.com/michilu/boilerplate/cmd/echo"
	"github.com/michilu/boilerplate/cmd/update"
	"github.com/michilu/boilerplate/cmd/version"
)

const (
	name   = "boilerplate"
	semver = "1.0.0-alpha"
)

var (
	ns = []func() (*cobra.Command, error){
		echo.New,
		update.New,
		version.New,
	}
)

func main() {
	cmd.Execute()
}
