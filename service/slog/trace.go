package slog

import (
	"context"
	"fmt"

	"github.com/michilu/boilerplate/service/meta"
	"github.com/rs/zerolog"
	"go.opencensus.io/trace"
)

var (
	Atrace = Trace
)

func Trace(ctx context.Context, s *trace.Span) zerolog.LogObjectMarshaler {
	v0 := meta.Get().Flatten()
	v1 := make([]trace.Attribute, 0, len(v0)+2)
	for k, v := range v0 {
		v2 := k
		v3 := v
		v1 = append(v1, trace.StringAttribute(
			fmt.Sprintf("meta/%s", v2),
			fmt.Sprintf("%s", v3),
		))
	}
	v4 := s.SpanContext()
	v1 = append(v1,
		trace.StringAttribute("trace", v4.TraceID.String()),
		trace.StringAttribute("span", v4.SpanID.String()),
	)
	s.AddAttributes(v1...)
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
	v2 := v1.TraceID.String()
	e.
		Str("traceURL", fmt.Sprintf(GetTraceURLTemplate(), v2)).
		Str("trace", fmt.Sprintf(GetTraceIDTemplate(), v2)).
		Str("spanID", v1.SpanID.String())
	return
}
