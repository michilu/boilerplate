package profile

import (
	"context"
	_ "net/http/pprof"
	"os"
	"runtime"
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/pkg/profile"
	"github.com/spf13/viper"
	"go.opencensus.io/trace"
	"google.golang.org/grpc/codes"

	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/now"
	"github.com/michilu/boilerplate/service/slog"
)

func RunPprof() error {
	const op = op + ".RunPprof"
	runtime.SetBlockProfileRate(1)
	e := gin.Default()
	pprof.Register(e)
	err := e.Run(viper.GetString("service.profile.pprof.addr"))
	if err != nil {
		const op = op + ".gin.Run"
		return &errs.Error{Op: op, Err: err}
	}
	return nil
}

func Profile(ctx context.Context) {
	const (
		op = op + ".Pprof"
		s  = "assets/pprof"
	)
	before := s + "/cpu.pprof"
	d := viper.GetDuration("service.profile.pprof.duration")
	if d == 0 {
		slog.Logger().Warn().Str("op", op).
			Err(&errs.Error{Op: op, Code: codes.InvalidArgument, Message: "service.pprof.duration"}).
			Dur("value", d).Msg("warnning")
		d = 5 * time.Minute
	}
	t := time.NewTicker(d)
	defer t.Stop()
	for {
		p := profile.Start(profile.ProfilePath(s))
		select {
		case <-ctx.Done():
			p.Stop()
			return
		case <-t.C:
			p.Stop()
		}
		ctx, s := trace.StartSpan(ctx, op)
		t := slog.Trace(ctx)
		after := before + "." + now.Now().UTC().Format(time.RFC3339)
		err := os.Rename(before, after)
		if err != nil {
			const op = op + ".os.Rename"
			slog.Logger().Err(err).Str("op", op).EmbedObject(t).Str("before", before).Str("after", after).Msg(err.Error())
		} else {
			slog.Logger().Info().Str("op", op).EmbedObject(t).Str("before", before).Str("after", after).Msg("rotated")
		}
		s.End()
	}
}
