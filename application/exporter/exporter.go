package exporter

import (
	"context"

	k "github.com/michilu/boilerplate/application/config"
	"github.com/michilu/boilerplate/service/config"
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
		const v0 = k.ServiceTraceEnable
		v1 := viper.GetBool(v0)
		s.AddAttributes(trace.BoolAttribute(v0, v1))
		if !v1 {
			return
		}
	}
	v2, err := config.GCPCredentials(ctx)
	if err != nil {
		const op = op + ".config.GCPCredentials"
		err := &errs.Error{Op: op, Code: codes.Internal, Err: err}
		s.SetStatus(trace.Status{Code: int32(codes.Internal), Message: err.Error()})
		slog.Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
		return
	}
	v3, err := config.GCPProjectID(ctx)
	if err != nil {
		const op = op + ".config.GCPProjectID"
		err := &errs.Error{Op: op, Code: codes.Internal, Err: err}
		s.SetStatus(trace.Status{Code: int32(codes.Internal), Message: err.Error()})
		slog.Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
		return
	}
	v4, _, err := sdserver.NewExporter(v3,
		gcp.CredentialsTokenSource(v2),
		GlobalMonitoredResource{projectID: string(v3)},
	)
	if err != nil {
		const op = op + ".sdserver.NewExporter"
		err := &errs.Error{Op: op, Code: codes.Internal, Err: err}
		s.SetStatus(trace.Status{Code: int32(codes.Internal), Message: err.Error()})
		slog.Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
		return
	}
	trace.RegisterExporter(v4)
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
}
