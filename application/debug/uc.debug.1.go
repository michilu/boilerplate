package debug

import (
	"context"
	"fmt"
	"strings"

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
	t := slog.Trace(ctx)

	const (
		c0 = "application.debug.client.id"
		c1 = "application.debug.client.id.alias"
	)
	s.AddAttributes(trace.StringAttribute("c0", c0))
	v0 := viper.GetString(c0)
	if v0 == "" {
		v1 := viper.GetString(c1)
		if v1 != "" {
			v0 = strings.ReplaceAll(viper.GetString(v1), ":", "-")
		}
	}
	s.AddAttributes(trace.StringAttribute("v0", v0))
	if v0 == "" {
		slog.Logger().Debug().Str("op", op).EmbedObject(t).Str("v0", v0).Msg("return")
		return v0, nil
	}
	v2, err := debug.NewID()
	if err != nil {
		const op = op + ".debug.NewID"
		s.SetStatus(trace.Status{Code: int32(codes.Unknown), Message: err.Error()})
		slog.Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
		return "", err
	}
	s.AddAttributes(trace.StringAttribute("v2", v2))
	v3 := fmt.Sprintf("you can set to '%s' in config.toml", v0)
	s.AddAttributes(trace.StringAttribute("Warn", v3))
	slog.Logger().Warn().Str("op", op).EmbedObject(t).Str("value", v2).Msg(v3)
	slog.Logger().Debug().Str("op", op).EmbedObject(t).Str("v2", v2).Msg("return")
	return v2, nil
}
