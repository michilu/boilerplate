package version

import (
	"github.com/spf13/cobra"

	"github.com/michilu/boilerplate/v/meta"
)

// New returns a new command.
func New() (*cobra.Command, error) {
	return &cobra.Command{
		Use:   "version",
		Short: "print version",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Print(meta.Get())
		},
	}, nil
}
