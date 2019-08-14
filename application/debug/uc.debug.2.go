package debug

import (
	fmt "fmt"
	"os/exec"
	"strings"

	"google.golang.org/grpc/codes"

	"github.com/michilu/boilerplate/service/debug"
	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/slog"
	"github.com/spf13/viper"
)

func OpenDebugPort(v debug.Clienter) error {
	const op = op + ".Connect"
	{
		slog.Logger().Debug().Str("op", op).Msg("start")
		defer slog.Logger().Debug().Str("op", op).Msg("end")
	}
	err := v.Validate()
	if err != nil {
		const op = op + ".Validate"
		return &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
	}
	const s = "application.debug.open-debug-port.command"
	v0 := strings.Split(viper.GetString(s), " ")
	if len(v0) < 1 {
		const op = op + ".len"
		return &errs.Error{Op: op, Code: codes.InvalidArgument, Message: fmt.Sprintf("check '%s' in config.toml", s)}
	}
	file := v0[0]
	v1, err := exec.LookPath(file)
	if err != nil {
		const op = op + ".exec.LookPath"
		return &errs.Error{Op: op, Code: codes.Internal, Err: err}
	}
	c := exec.Command(v1, v0[1:]...)
	v2, err := c.CombinedOutput()
	{
		const op = op + ".exec.Command.CombinedOutput"
		if err != nil {
			slog.Logger().Error().Str("op", op).Str("output", string(v2)).Msg("error")
			return &errs.Error{Op: op, Code: codes.Unavailable, Err: err}
		} else {
			slog.Logger().Debug().Str("op", op).Str("output", string(v2)).Msg("debug")
		}
	}
	return nil
}
