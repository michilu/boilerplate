package update

import (
	"time"

	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/meta"
	"github.com/michilu/boilerplate/service/slog"
	"github.com/michilu/go-selfupdate/selfupdate"
	"github.com/spf13/viper"
	"google.golang.org/grpc/codes"
)

const (
	op = "update"
)

// Update ...
//
//go:generate genny -in=../pipe/pipe.go -out=gen-pipe-Update.go -pkg=$GOPACKAGE gen "Name=update InT=time.Time OutT=struct{}"
func Update(m time.Time) (struct{}, error) {
	const op = op + ".Update"
	{
		slog.Logger().Debug().Str("op", op).Msg("start")
		defer slog.Logger().Debug().Str("op", op).Msg("end")
	}
	v := meta.Get()
	channel := v.GetChannel()
	if channel == "" {
		channel = viper.GetString("service.update.channel")
	}
	if channel == "" {
		const op = op + ".Meta.GetChannel"
		return struct{}{}, &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given Meta.Channel"}
	}
	url := viper.GetString("service.update.url") + channel + "/"
	updater := &selfupdate.Updater{
		CurrentVersion: v.GetSemver(),
		ApiURL:         url,
		BinURL:         url,
		DiffURL:        url,
		Dir:            "assets/update/",
		CmdName:        v.GetName(),
		ForceCheck:     true,
	}
	err := updater.BackgroundRun()
	if err != nil {
		const op = op + ".updater.BackgroundRun"
		return struct{}{}, &errs.Error{Op: op, Code: codes.Unavailable, Err: err}
	}
	return struct{}{}, nil
}
