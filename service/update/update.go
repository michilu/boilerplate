package update

import (
	"context"

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
	t := slog.Trace(ctx)

	v := meta.Get()
	channel := v.GetChannel()
	if channel == "" {
		const v = "service.update.channel"
		channel = viper.GetString(v)
		s.AddAttributes(trace.StringAttribute(v, channel))
	}
	s.AddAttributes(trace.StringAttribute("channel", channel))
	if channel == "" {
		const op = op + ".Meta.GetChannel"
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given Meta.Channel"}
		s.SetStatus(trace.Status{Code: int32(codes.InvalidArgument), Message: err.Error()})
		slog.Logger().Error().Str("op", op).EmbedObject(t).Err(err).Msg("error")
		return ctx, err
	}
	url := viper.GetString("service.update.url") + channel + "/"
	s.AddAttributes(trace.StringAttribute("url", url))
	updater := &selfupdate.Updater{
		CurrentVersion: v.GetSemver(),
		ApiURL:         url,
		BinURL:         url,
		DiffURL:        url,
		Dir:            "assets/update/",
		CmdName:        v.GetName(),
		ForceCheck:     true,
	}
	s.AddAttributes(
		trace.StringAttribute("selfupdate.Updater.CurrentVersion", updater.CurrentVersion),
		trace.StringAttribute("selfupdate.Updater.ApiURL", updater.ApiURL),
		trace.StringAttribute("selfupdate.Updater.BinURL", updater.BinURL),
		trace.StringAttribute("selfupdate.Updater.DiffURL", updater.DiffURL),
		trace.StringAttribute("selfupdate.Updater.Dir", updater.Dir),
		trace.StringAttribute("selfupdate.Updater.CmdName", updater.CmdName),
		trace.BoolAttribute("selfupdate.Updater.ForceCheck", updater.ForceCheck),
	)
	v0 := zerolog.Dict().
		Str("CurrentVersion", updater.CurrentVersion).
		Str("ApiURL", updater.ApiURL).
		Str("CmdName", updater.CmdName).
		Str("BinURL", updater.BinURL).
		Str("DiffURL", updater.DiffURL).
		Str("Dir", updater.Dir).
		Bool("ForceCheck", updater.ForceCheck).
		Dict("Info", zerolog.Dict().
			Str("Version", updater.Info.Version).
			Bytes("Sha256", updater.Info.Sha256))
	slog.Logger().Debug().Str("op", op).EmbedObject(t).Dict("updater", v0).Msg("arg")
	{
		ok, err := updater.ForegroundRun()
		const op = op + ".updater.BackgroundRun"
		if err != nil {
			err := &errs.Error{Op: op, Code: codes.Unavailable, Err: err}
			s.SetStatus(trace.Status{Code: int32(codes.Unavailable), Message: err.Error()})
			slog.Logger().Error().Str("op", op).EmbedObject(t).Err(err).Msg("error")
			return ctx, err
		}
		if !ok {
			err := &errs.Error{Op: op, Code: codes.Unavailable, Message: "no updates. must be try after"}
			s.SetStatus(trace.Status{Code: int32(codes.Unavailable), Message: err.Error()})
			slog.Logger().Error().Str("op", op).EmbedObject(t).Err(err).Msg("error")
			return ctx, err
		}
	}
	slog.Logger().Info().Str("op", op).EmbedObject(t).Dict("updater", v0).Msg("updated")
	return ctx, nil
}
