package slog

import (
	"context"
	"fmt"

	"github.com/rs/zerolog"
	"go.opencensus.io/trace"
)

var (
	Atrace = Trace
)

func Trace(ctx context.Context, s *trace.Span) zerolog.LogObjectMarshaler {
	v0 := s.SpanContext()
	s.AddAttributes(
		trace.StringAttribute("trace", v0.TraceID.String()),
		trace.StringAttribute("span", v0.SpanID.String()),
	)
	return &TraceObject{ctx}
}

// Trace is trace span handler for zerolog.
type TraceObject struct {
	ctx context.Context
}

func (p *TraceObject) MarshalZerologObject(e *zerolog.Event) {
	if p.ctx == nil {
		return
	}
	v0 := trace.FromContext(p.ctx)
	if v0 == nil {
		return
	}
	v1 := v0.SpanContext()
	e.
		Str("trace", fmt.Sprintf(GetTraceIDTemplate(), v1.TraceID.String())).
		Str("spanID", v1.SpanID.String())
	return
}
