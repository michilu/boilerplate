package config

import (
	"context"
	"os"

	"github.com/michilu/boilerplate/service/errs"
	"github.com/spf13/viper"
	"go.opencensus.io/trace"
	"gocloud.dev/gcp"
	"golang.org/x/oauth2/google"
	"google.golang.org/grpc/codes"
)

var (
	gcpCredentials *google.Credentials
	gcpProjectID   gcp.ProjectID
)

func GCPCredentials(ctx context.Context) (*google.Credentials, error) {
	const op = op + ".GCPCredentials"
	ctx, s := trace.StartSpan(ctx, op)
	defer s.End()

	if gcpCredentials != nil {
		return gcpCredentials, nil
	}
	{
		const c0 = "google.application.credentials"
		v0 := viper.GetString(c0)
		s.AddAttributes(trace.StringAttribute(c0, v0))
		os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", v0)
	}
	v0, err := gcp.DefaultCredentials(ctx)
	if err != nil {
		const op = op + ".gcp.DefaultCredentials"
		err := &errs.Error{Op: op, Code: codes.Internal, Err: err}
		s.SetStatus(trace.Status{Code: int32(codes.Internal), Message: err.Error()})
		return nil, err
	}
	gcpCredentials = v0
	return v0, nil
}

func GCPProjectID(ctx context.Context) (gcp.ProjectID, error) {
	const op = op + ".GCPProjectID"
	ctx, s := trace.StartSpan(ctx, op)
	defer s.End()

	if gcpProjectID != "" {
		return gcpProjectID, nil
	}
	v0, err := GCPCredentials(ctx)
	if err != nil {
		const op = op + ".GCPCredentials"
		err := &errs.Error{Op: op, Code: codes.Internal, Err: err}
		s.SetStatus(trace.Status{Code: int32(codes.Internal), Message: err.Error()})
	}
	v1, err := gcp.DefaultProjectID(v0)
	if err != nil {
		const op = op + ".gcp.DefaultProjectID"
		err := &errs.Error{Op: op, Code: codes.Internal, Err: err}
		s.SetStatus(trace.Status{Code: int32(codes.Internal), Message: err.Error()})
	}
	gcpProjectID = v1
	return v1, nil
}
