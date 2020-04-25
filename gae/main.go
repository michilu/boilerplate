package main

import (
	"context"
	"io"
	"strings"

	k "github.com/michilu/boilerplate/application/config"
	"github.com/michilu/boilerplate/service/config"
	"github.com/michilu/boilerplate/service/slog"
	"github.com/spf13/viper"
	"go.opencensus.io/trace"
	"google.golang.org/grpc/codes"
)

var (
	defaults = []config.KV{
		{K: k.ServiceSlogVerbose, V: true},

		{K: k.GcpAppengineHostnameFmt, V: "https://%s-dot-%s.appspot.com%s"},
		// https://cloud.google.com/pubsub/docs/push#authenticating_standard_urls
		{K: k.GcpPubsubPushEndpointFmt, V: "/_ah/push-handlers/%s"},
		{K: k.GoogleProjectId, V: "boilerplate"},
		{K: k.GoogleServicesId, V: "default"},
	}
)

func main() {
	const op = op + ".main"
	ctx, s := trace.StartSpan(context.Background(), op)
	defer s.End()
	t := slog.Trace(ctx, s)

	initialize(ctx, defaults)
	if err := slog.Recover(ctx,
		httpServer,
	); err != nil {
		slog.Logger().Err(err).EmbedObject(t).Msg(err.Error())
	}
}

func initialize(ctx context.Context, v0 []config.KV) {
	const op = op + ".initialize"
	ctx, s := trace.StartSpan(ctx, op)
	defer s.End()
	t := slog.Trace(ctx, s)

	{
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(
			"-", "_",
			".", "_",
		))
		config.SetDefault(v0)
	}
	if v1, err := slog.NewAppengineLogging(ctx); err != nil {
		s.SetStatus(trace.Status{Code: int32(codes.InvalidArgument), Message: err.Error()})
		slog.Logger().Err(err).EmbedObject(t).Msg(err.Error())
	} else {
		slog.SetDefaultTracer(v1)
		slog.SetDefaultLogger([]io.Writer{v1})
	}
	slog.Logger().Debug().EmbedObject(t).Msg(op + ": return")
}
