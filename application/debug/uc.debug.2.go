package debug

import (
	"context"
	"fmt"
	"os/exec"
	"strings"

	k "github.com/michilu/boilerplate/application/config"
	"github.com/michilu/boilerplate/service/debug"
	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/slog"
	"github.com/spf13/viper"
	"go.opencensus.io/trace"
	"google.golang.org/grpc/codes"
)

func OpenDebugPort(ctx context.Context, m debug.Clienter) error {
	const op = op + ".OpenDebugPort"
	if ctx == nil {
		return &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"}
	}
	ctx, s := trace.StartSpan(ctx, op)
	defer s.End()
	t := slog.Trace(ctx, s)

	if m == nil {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'm' is nil"}
		s.SetStatus(trace.Status{Code: int32(codes.InvalidArgument), Message: err.Error()})
		slog.Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
		return err
	}
	{
		s.AddAttributes(trace.StringAttribute("m", m.String()))
		slog.Logger().Debug().Str("op", op).EmbedObject(t).EmbedObject(m).Msg(op + ": arg")
	}
	{
		err := m.Validate()
		if err != nil {
			const op = op + ".Validate"
			err := &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
			s.SetStatus(trace.Status{Code: int32(codes.InvalidArgument), Message: err.Error()})
			slog.Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
			return err
		}
	}
	const v0 = k.ApplicationDebugOpenDebugPortCommand
	v1 := viper.GetString(v0)
	s.AddAttributes(trace.StringAttribute(v0, v1))
	if v1 == "" {
		const op = op + ".viper.GetString"
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: fmt.Sprintf("check '%s' in config.toml", v0)}
		s.SetStatus(trace.Status{Code: int32(codes.InvalidArgument), Message: err.Error()})
		slog.Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
		return err
	}
	v2 := strings.Split(v1, " ")
	if len(v2) < 1 {
		const op = op + ".len"
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: fmt.Sprintf("check '%s' in config.toml", v0)}
		s.SetStatus(trace.Status{Code: int32(codes.InvalidArgument), Message: err.Error()})
		slog.Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
		return err
	}
	file := v2[0]
	s.AddAttributes(trace.StringAttribute("file", file))
	v3, err := exec.LookPath(file)
	if err != nil {
		const op = op + ".exec.LookPath"
		err := &errs.Error{Op: op, Code: codes.Internal, Err: err}
		s.SetStatus(trace.Status{Code: int32(codes.Internal), Message: err.Error()})
		slog.Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
		return err
	}
	c := exec.Command(v3, v2[1:]...)
	v4, err := c.CombinedOutput()
	{
		const op = op + ".exec.Command.CombinedOutput"
		v5 := string(v4)
		s.AddAttributes(trace.StringAttribute(op, v5))
		if err != nil {
			err := &errs.Error{Op: op, Code: codes.Unavailable, Err: err}
			s.SetStatus(trace.Status{Code: int32(codes.Unavailable), Message: err.Error()})
			slog.Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
			return err
		}
	}
	return nil
}
