package slog

import (
	"context"
	"io"

	sentry "github.com/getsentry/sentry-go"
	k "github.com/michilu/boilerplate/application/config"
	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/meta"
	"github.com/spf13/viper"
	"go.opencensus.io/trace"
	"google.golang.org/grpc/codes"
)

func InitSentry(ctx context.Context) (io.Closer, error) {
	const op = op + ".InitSentry"
	if ctx == nil {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"}
		Logger().Err(err).Str("op", op).Msg(err.Error())
		return nil, err
	}
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	ctx, s := trace.StartSpan(ctx, op)
	defer s.End()
	t := Trace(ctx, s)

	v0 := viper.GetString(k.ServiceSlogSentryDsn)
	{
		s.AddAttributes(trace.StringAttribute("v0", v0))
		Logger().Debug().Str("op", op).EmbedObject(t).Str("v0", v0).Msg(op + ": value")
	}
	if v0 == "" {
		return nil, nil
	}
	v1 := viper.GetString(k.ServiceSlogSentryServerName)
	if v1 == "" {
		v2 := viper.GetString(k.ServiceSlogSentryServerNameAlias)
		if v2 != "" {
			v1 = viper.GetString(v2)
		}
	}
	{
		v3 := meta.Get()
		err := sentry.Init(sentry.ClientOptions{
			Dsn:         v0,
			ServerName:  v1,
			Release:     v3.Semver,
			Dist:        v3.Runtime.Os,
			Environment: v3.Channel,
		})
		if err != nil {
			err := &errs.Error{Op: op, Code: codes.Internal, Err: err}
			s.SetStatus(trace.Status{Code: int32(codes.Internal), Message: err.Error()})
			Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
		}
	}
	return nil, nil
}
