package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"sync"

	k "github.com/michilu/boilerplate/application/config"
	"github.com/michilu/boilerplate/service/config"
	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/slog"
	"github.com/spf13/viper"
	"go.opencensus.io/trace"
	"google.golang.org/appengine"
	"google.golang.org/grpc/codes"
)

var (
	_warmupOnce sync.Once
)

// Note that warmup requests are not guaranteed to be called.
// https://cloud.google.com/appengine/docs/standard/go/configuring-warmup-requests
func handlerWarmup(w http.ResponseWriter, r *http.Request) {
	const op = op + ".handlerWarmup"
	ctx := appengine.NewContext(r)
	ctx, s := trace.StartSpan(ctx, op)
	defer s.End()
	t := slog.Trace(ctx, s)

	var err0 error
	_warmupOnce.Do(func() {
		const op = op + "#func"
		if err := setConfig(ctx); err != nil {
			err := &errs.Error{Op: op, Code: codes.Internal, Err: err}
			s.SetStatus(trace.Status{Code: int32(codes.Internal), Message: err.Error()})
			slog.Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
			err0 = err
			return
		}
		for _, v := range []func(context.Context) error{} {
			v0 := v
			if err := v0(ctx); err != nil {
				err := &errs.Error{Op: op, Code: codes.Internal, Err: err}
				s.SetStatus(trace.Status{Code: int32(codes.Internal), Message: err.Error()})
				slog.Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
				err0 = err
				return
			}
		}
	})
	if err0 != nil {
		http.Error(w, err0.Error() /*error*/, http.StatusInternalServerError)
	} else {
		v0 := viper.GetString(k.GcpAppengineVersionId)
		slog.Logger().Debug().Str("op", op).EmbedObject(t).Str("v0", v0).Msg(op + ": value")
		fmt.Fprintln(w, v0)
	}
	slog.Logger().Debug().Str("op", op).EmbedObject(t).Msg(op + ": return")
}

func setConfig(ctx context.Context) error {
	const op = op + ".setConfig"
	for _, v := range []config.KV{
		{K: k.GcpAppengineVersionId, V: strings.SplitN(appengine.VersionID(ctx), ".", 2)[0]},
	} {
		v := v
		viper.Set(v.K, v.V)
	}
	return nil
}
