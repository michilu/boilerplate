package profile

import (
	"context"
	"fmt"
	_ "net/http/pprof"
	"os"

	"cloud.google.com/go/profiler"
	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/meta"
	"github.com/michilu/boilerplate/service/slog"
	"github.com/spf13/viper"
	"go.opencensus.io/trace"
	"google.golang.org/grpc/codes"
)

// Profiling Go applications | Stackdriver Profiler | Google Cloud
// https://cloud.google.com/profiler/docs/profiling-go
func RunProfiler(ctx context.Context) error {
	const op = op + ".RunProfiler"
	if ctx == nil {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"}
		return err
	}
	ctx, s := trace.StartSpan(ctx, op)
	defer s.End()
	t := slog.Trace(ctx)

	v0 := meta.Get()
	v1 := profiler.Config{
		Service:        v0.GetName(),
		ServiceVersion: v0.GetSemver(),
		ProjectID:      viper.GetString("gcp.project.id"),
		DebugLogging:   viper.GetBool("service.profile.profiler.debug-logging"),
	}
	{
		v2 := fmt.Sprintf("%v", v1)
		s.AddAttributes(trace.StringAttribute("v2", v2))
		slog.Logger().Debug().Str("op", op).EmbedObject(t).Str("v2", v2).Msg("value")
	}
	{
		const v3 = "google.application.credentials"
		v4 := viper.GetString(v3)
		s.AddAttributes(trace.StringAttribute(v3, v4))
		slog.Logger().Debug().Str("op", op).EmbedObject(t).Str(v3, v4).Msg("value")
		os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", v4)
	}
	err := profiler.Start(v1)
	if err != nil {
		const op = op + ".profiler.Start"
		err := &errs.Error{Op: op, Code: codes.Internal, Err: err}
		s.SetStatus(trace.Status{Code: int32(codes.Internal), Message: err.Error()})
		slog.Logger().Error().Str("op", op).EmbedObject(t).Err(err).Msg(err.Error())
		return err
	}
	return nil
}
