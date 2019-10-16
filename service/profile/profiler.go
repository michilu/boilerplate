package profile

import (
	"context"
	"fmt"
	_ "net/http/pprof"

	"cloud.google.com/go/profiler"
	"github.com/michilu/boilerplate/service/config"
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

	v0, err := config.GCPProjectID(ctx)
	if err != nil {
		const op = op + ".config.GCPProjectID"
		err := &errs.Error{Op: op, Code: codes.Internal, Err: err}
		s.SetStatus(trace.Status{Code: int32(codes.Internal), Message: err.Error()})
		slog.Logger().Error().Str("op", op).EmbedObject(t).Err(err).Msg(err.Error())
		return err
	}
	v1 := meta.Get()
	v2 := profiler.Config{
		Service:        v1.GetName(),
		ServiceVersion: v1.GetSemver(),
		ProjectID:      string(v0),
		DebugLogging:   viper.GetBool("service.profile.profiler.debug-logging"),
	}
	{
		v3 := fmt.Sprintf("%v", v2)
		s.AddAttributes(trace.StringAttribute("v3", v3))
		slog.Logger().Debug().Str("op", op).EmbedObject(t).Str("v3", v3).Msg("value")
	}
	{
		err := profiler.Start(v2)
		if err != nil {
			const op = op + ".profiler.Start"
			err := &errs.Error{Op: op, Code: codes.Internal, Err: err}
			s.SetStatus(trace.Status{Code: int32(codes.Internal), Message: err.Error()})
			slog.Logger().Error().Str("op", op).EmbedObject(t).Err(err).Msg(err.Error())
			return err
		}
	}
	return nil
}
