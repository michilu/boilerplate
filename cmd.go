package main

import (
	"github.com/spf13/cobra"

	"github.com/michilu/boilerplate/service/meta"
)

var (
	rootCmd *cobra.Command
)

func initCmd() {
	rootCmd = &cobra.Command{
		Use: meta.Name(),
	}
}
