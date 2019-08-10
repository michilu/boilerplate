package run

import (
	"github.com/spf13/cobra"

	"github.com/michilu/boilerplate/usecase/run"
)

// New returns a new command.
func New() (*cobra.Command, error) {
	c := &cobra.Command{
		Use:   "run",
		Short: "run",
		Run:   run.Run,
	}
	return c, nil
}
