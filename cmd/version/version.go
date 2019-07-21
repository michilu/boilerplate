package version

import (
	"github.com/spf13/cobra"
	"google.golang.org/grpc/codes"

	"github.com/michilu/boilerplate/v/errs"
	"github.com/michilu/boilerplate/v/log"
	"github.com/michilu/boilerplate/v/meta"
)

const (
	op = "version"
)

// New returns a new command.
func New() (*cobra.Command, error) {
	return &cobra.Command{
		Use:   "version",
		Short: "print version",
		Run: func(cmd *cobra.Command, args []string) {
			const op = op + ".Run"
			s, err := meta.Yaml()
			if err != nil {
				log.Logger().Error().Str("op", op).Err(&errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}).Msg("error")
			}
			cmd.Print(s)
		},
	}, nil
}
