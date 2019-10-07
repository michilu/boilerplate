package debug

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/michilu/boilerplate/service/debug"
	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/slog"
	"github.com/spf13/viper"
	"go.opencensus.io/trace"
	"google.golang.org/grpc/codes"
)

func GenerateUUID(ctx context.Context) (string, error) {
	const op = op + ".GenerateUUID"
	if ctx == nil {
		return "", &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"}
	}
	ctx, s := trace.StartSpan(ctx, op)
	defer s.End()

	const v0 = "application.debug.client.id"
	s.AddAttributes(trace.StringAttribute("v0", v0))
	v1 := viper.GetString(v0)
	s.AddAttributes(trace.StringAttribute("v1", v1))
	v2, err := uuid.Parse(v1)
	if err == nil {
		v3 := v2.String()
		s.AddAttributes(trace.StringAttribute("v3", v3))
		slog.Logger().Debug().Str("op", op).Str("v3", v3).Msg("return")
		return v3, nil
	} else {
		const op = op + ".uuid.Parse"
		v4 := fmt.Sprintf("check '%s' in config.toml", v0)
		s.AddAttributes(trace.StringAttribute("Warn", v4))
		slog.Logger().Warn().Str("op", op).Err(err).Str("value", v1).Msg(v4)
	}
	v5, err := debug.NewID()
	if err != nil {
		s.SetStatus(trace.Status{Code: int32(codes.Unknown), Message: err.Error()})
		return "", err
	}
	s.AddAttributes(trace.StringAttribute("v5", v5))
	v6 := fmt.Sprintf("you can set to '%s' in config.toml", v0)
	s.AddAttributes(trace.StringAttribute("Warn", v6))
	slog.Logger().Warn().Str("op", op).Str("value", v5).Msg(v6)
	slog.Logger().Debug().Str("op", op).Str("v5", v5).Msg("return")
	return v5, nil
}
