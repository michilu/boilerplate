package version

import (
	"github.com/spf13/cobra"

	"github.com/michilu/boilerplate/application/version"
)

// New returns a new command.
func New() (*cobra.Command, error) {
	return &cobra.Command{
		Use:   "version",
		Short: "print version",
		Run:   version.Run,
	}, nil
}
