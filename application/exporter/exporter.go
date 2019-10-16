package exporter

import (
	"context"
	"os"

	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/slog"
	"github.com/spf13/viper"
	"go.opencensus.io/trace"
	"gocloud.dev/gcp"
	"gocloud.dev/server/sdserver"
	"google.golang.org/grpc/codes"
)

const (
	op = "application/exporter"
)

type GlobalMonitoredResource struct {
	projectID string
}

func (g GlobalMonitoredResource) MonitoredResource() (string, map[string]string) {
	return "global", map[string]string{"project_id": g.projectID}
}

func Run() {
	const op = op + ".Run"
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ctx, s := trace.StartSpan(ctx, op)
	defer s.End()
	t := slog.Trace(ctx)

	{
		const v0 = "google.application.credentials"
		v1 := viper.GetString(v0)
		s.AddAttributes(trace.StringAttribute(v0, v1))
		slog.Logger().Debug().Str("op", op).EmbedObject(t).Str(v0, v1).Msg("value")
		os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", v1)
	}
	{
		const v0 = "service.trace.enable"
		v1 := viper.GetBool(v0)
		s.AddAttributes(trace.BoolAttribute(v0, v1))
		if !v1 {
			return
		}
	}
	credentials, err := gcp.DefaultCredentials(ctx)
	if err != nil {
		const op = op + ".gcp.DefaultCredentials"
		err := &errs.Error{Op: op, Code: codes.Internal, Err: err}
		s.SetStatus(trace.Status{Code: int32(codes.Internal), Message: err.Error()})
		slog.Logger().Fatal().Str("op", op).EmbedObject(t).Err(err).Msg(err.Error())
	}
	projectID, err := gcp.DefaultProjectID(credentials)
	if err != nil {
		const op = op + ".gcp.DefaultProjectID"
		err := &errs.Error{Op: op, Code: codes.Internal, Err: err}
		s.SetStatus(trace.Status{Code: int32(codes.Internal), Message: err.Error()})
		slog.Logger().Fatal().Str("op", op).EmbedObject(t).Err(err).Msg(err.Error())
	}
	exporter, _, err := sdserver.NewExporter(projectID,
		gcp.CredentialsTokenSource(credentials),
		GlobalMonitoredResource{projectID: string(projectID)},
	)
	if err != nil {
		const op = op + ".sdserver.NewExporter"
		err := &errs.Error{Op: op, Code: codes.Internal, Err: err}
		s.SetStatus(trace.Status{Code: int32(codes.Internal), Message: err.Error()})
		slog.Logger().Fatal().Str("op", op).EmbedObject(t).Err(err).Msg(err.Error())
	}
	trace.RegisterExporter(exporter)
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
}
