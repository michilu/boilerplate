package main

import (
	"net/http"

	"github.com/michilu/boilerplate/service/slog"
	"go.opencensus.io/trace"
	"google.golang.org/appengine"
)

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	const op = op + ".handlerIndex"
	ctx := appengine.NewContext(r)
	ctx, s := trace.StartSpan(ctx, op)
	defer s.End()
	t := slog.Trace(ctx, s)

	switch r.URL.Path {
	case "/":
		handlerWarmup(w, r)
	default:
		http.NotFound(w, r)
	}
	slog.Logger().Debug().Str("op", op).EmbedObject(t).Msg(op + ": return")
}
