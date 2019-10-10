package version

import (
	"fmt"

	"github.com/spf13/cobra"
	"google.golang.org/grpc/codes"

	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/meta"
	"github.com/michilu/boilerplate/service/slog"
)

const (
	op = "application/version"
)

func Run(cmd *cobra.Command, _ []string) {
	const op = op + ".Run"
	s, err := meta.Yaml()
	if err != nil {
		const op = op + ".meta.Meta"
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
		slog.Logger().Error().Str("op", op).Err(err).Msg(err.Error())
		return
	}
	fmt.Printf("%s", s)
}
