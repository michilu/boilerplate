package update

import (
	"github.com/spf13/cobra"

	"github.com/michilu/boilerplate/application/update"
)

// New returns a new command.
func New() (*cobra.Command, error) {
	c := &cobra.Command{
		Use:   "update",
		Short: "update",
		Run:   update.Run,
	}
	return c, nil
}
