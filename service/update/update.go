package update

import (
	"context"
	"errors"

	k "github.com/michilu/boilerplate/application/config"
	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/meta"
	"github.com/michilu/boilerplate/service/slog"
	"github.com/michilu/go-selfupdate/selfupdate"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"go.opencensus.io/trace"
	"google.golang.org/grpc/codes"
)

const (
	op = "service/update"
)

// Update ...
//
//go:generate genny -in=../pipe/pipe.go -out=gen-pipe-Update.go -pkg=$GOPACKAGE gen "Name=update InT=context.Context OutT=context.Context"
func Update(ctx context.Context) (context.Context, error) {
	const op = op + ".Update"
	ctx, s := trace.StartSpan(ctx, op)
	defer s.End()
	t := slog.Trace(ctx, s)

	v0 := meta.Get()
	v1 := ""
	{
		const c0 = k.ServiceUpdateChannel
		v2 := viper.GetString(c0)
		{
			s.AddAttributes(trace.StringAttribute(c0, v2))
			slog.Logger().Debug().Str("op", op).EmbedObject(t).Str("v2", v2).Msg(op + ": value")
		}
		v1 = v2
	}
	if v1 == "" {
		v3 := v0.GetChannel()
		{
			s.AddAttributes(trace.StringAttribute("meta.channel", v3))
			slog.Logger().Debug().Str("op", op).EmbedObject(t).Str("v3", v3).Msg(op + ": value")
		}
		v1 = v3
	}
	{
		s.AddAttributes(trace.StringAttribute("v1", v1))
		slog.Logger().Debug().Str("op", op).EmbedObject(t).Str("v1", v1).Msg(op + ": value")
		if v1 == "" {
			const op = op + ".Meta.GetChannel"
			err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given Meta.Channel"}
			s.SetStatus(trace.Status{Code: int32(codes.InvalidArgument), Message: err.Error()})
			slog.Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
			return ctx, err
		}
	}
	v4 := viper.GetString(k.ServiceUpdateUrl) + v1 + "/"
	{
		s.AddAttributes(trace.StringAttribute("v4", v4))
		slog.Logger().Debug().Str("op", op).EmbedObject(t).Str("v4", v4).Msg(op + ": value")
	}
	v5 := &selfupdate.Updater{
		CurrentVersion: v0.GetSemver(),
		ApiURL:         v4,
		CmdName:        v0.GetName(),
		BinURL:         v4,
		DiffURL:        v4,
		Dir:            "assets/update/",
		ForceCheck:     true,
	}
	v6 := zerolog.Dict().
		Str("CurrentVersion", v5.CurrentVersion).
		Str("ApiURL", v5.ApiURL).
		Str("CmdName", v5.CmdName).
		Str("BinURL", v5.BinURL).
		Str("DiffURL", v5.DiffURL).
		Str("Dir", v5.Dir).
		Bool("ForceCheck", v5.ForceCheck).
		Dict("Info", zerolog.Dict().
			Str("Version", v5.Info.Version).
			Bytes("Sha256", v5.Info.Sha256))
	{
		s.AddAttributes(
			trace.StringAttribute("selfupdate.Updater.CurrentVersion", v5.CurrentVersion),
			trace.StringAttribute("selfupdate.Updater.ApiURL", v5.ApiURL),
			trace.StringAttribute("selfupdate.Updater.BinURL", v5.BinURL),
			trace.StringAttribute("selfupdate.Updater.DiffURL", v5.DiffURL),
			trace.StringAttribute("selfupdate.Updater.Dir", v5.Dir),
			trace.StringAttribute("selfupdate.Updater.CmdName", v5.CmdName),
			trace.BoolAttribute("selfupdate.Updater.ForceCheck", v5.ForceCheck),
		)
		slog.Logger().Debug().Str("op", op).EmbedObject(t).Dict("v6", v6).Msg(op + ": value")
	}
	{
		s.End()
		const op = op + ".Updater.ForegroundRun"
		slog.Logger().Debug().Str("op", op).EmbedObject(t).Msg(op + ": start")
		ok, err := v5.ForegroundRun(ctx)
		slog.Logger().Debug().Str("op", op).EmbedObject(t).Msg(op + ": end")
		ctx, s := trace.StartSpan(ctx, op)
		defer s.End()
		t := slog.Trace(ctx, s)
		if err != nil {
			if errors.Is(err, selfupdate.ErrNoAvailableUpdates) {
				err := &errs.Error{Op: op, Code: codes.NotFound, Err: err}
				slog.Logger().Info().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
				return ctx, nil
			}
			err := &errs.Error{Op: op, Code: codes.Unavailable, Err: err}
			s.SetStatus(trace.Status{Code: int32(codes.Unavailable), Message: err.Error()})
			slog.Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
			return ctx, err
		}
		if !ok {
			err := &errs.Error{Op: op, Code: codes.Unavailable, Message: "no updates. must be try after"}
			s.SetStatus(trace.Status{Code: int32(codes.Unavailable), Message: err.Error()})
			slog.Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
			return ctx, err
		}
	}
	slog.Logger().Info().Str("op", op).EmbedObject(t).Msg(op + ": updated")
	return ctx, nil
}
