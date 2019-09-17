package debug

import (
	"context"
	fmt "fmt"
	"os/exec"
	"strings"

	"go.opencensus.io/trace"
	"google.golang.org/grpc/codes"

	"github.com/michilu/boilerplate/service/debug"
	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/slog"
	"github.com/spf13/viper"
)

func OpenDebugPort(ctx context.Context, m debug.Clienter) error {
	const op = op + ".OpenDebugPort"
	if ctx == nil {
		return &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"}
	}
	ctx, s := trace.StartSpan(ctx, op)
	defer s.End()
	a := make([]trace.Attribute, 0)
	defer s.AddAttributes(a...)
	t := slog.Trace(ctx)

	{
		if m == nil {
			err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'm' is nil"}
			s.SetStatus(trace.Status{Code: int32(codes.InvalidArgument), Message: err.Error()})
			return err
		}
		slog.Logger().Debug().Str("op", op).EmbedObject(t).EmbedObject(m).Msg("arg")
	}
	{
		err := m.Validate()
		if err != nil {
			const op = op + ".Validate"
			err := &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
			s.SetStatus(trace.Status{Code: int32(codes.InvalidArgument), Message: err.Error()})
			return err
		}
	}
	const v0 = "application.debug.open-debug-port.command"
	v1 := viper.GetString(v0)
	a = append(a, trace.StringAttribute(v0, v1))
	if v1 == "" {
		const op = op + ".viper.GetString"
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: fmt.Sprintf("check '%s' in config.toml", v0)}
		s.SetStatus(trace.Status{Code: int32(codes.InvalidArgument), Message: err.Error()})
		return err
	}
	v2 := strings.Split(v1, " ")
	if len(v2) < 1 {
		const op = op + ".len"
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: fmt.Sprintf("check '%s' in config.toml", v0)}
		s.SetStatus(trace.Status{Code: int32(codes.InvalidArgument), Message: err.Error()})
		return err
	}
	file := v2[0]
	a = append(a, trace.StringAttribute("file", file))
	v3, err := exec.LookPath(file)
	if err != nil {
		const op = op + ".exec.LookPath"
		err := &errs.Error{Op: op, Code: codes.Internal, Err: err}
		s.SetStatus(trace.Status{Code: int32(codes.Internal), Message: err.Error()})
		return err
	}
	c := exec.Command(v3, v2[1:]...)
	v4, err := c.CombinedOutput()
	{
		const op = op + ".exec.Command.CombinedOutput"
		v5 := string(v4)
		a = append(a, trace.StringAttribute(op, v5))
		if err != nil {
			err := &errs.Error{Op: op, Code: codes.Unavailable, Err: err}
			s.SetStatus(trace.Status{Code: int32(codes.Unavailable), Message: err.Error()})
			return err
		}
	}
	return nil
}
