package event

import (
	"context"

	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/now"
	"github.com/michilu/boilerplate/service/slog"
	"go.opencensus.io/trace"
	"google.golang.org/grpc/codes"
)

// RestoreEvent returns an Eventer from given bytes.
func RestoreEvent(ctx context.Context, b []byte) (Eventer, error) {
	const op = op + ".RestoreEvent"
	if ctx == nil {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"}
		return nil, err
	}
	ctx, s := trace.StartSpan(ctx, op)
	defer s.End()
	a := make([]trace.Attribute, 0)
	defer s.AddAttributes(a...)
	t := slog.Trace(ctx)
	slog.Logger().Debug().Str("op", op).EmbedObject(t).Bytes("b", b).Msg("arg")

	v0 := &Event{}
	err := v0.XXX_Unmarshal(b)
	if err != nil {
		return nil, &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
	}
	v1 := make([]*TimePoint, 0, len(v0.TimePoint)+1)
	occurred := false
	for _, i := range v0.GetTimePoint() {
		switch i.GetTag() {
		case Entered:
			continue
		case Occurred:
			occurred = true
		}
		v1 = append(v1, i)
	}
	var v2 Eventer = &Event{
		Id:        v0.GetId(),
		TimePoint: v1,
	}
	v3 := now.Now()
	if !occurred {
		v2, err = v2.AddTimePoint(Occurred, v3)
		if err != nil {
			const op = op + ".AddTimePoint(Occurred)"
			return nil, &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
		}
	}
	v2, err = v2.AddTimePoint(Entered, v3)
	if err != nil {
		const op = op + ".AddTimePoint(Entered)"
		return nil, &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
	}
	slog.Logger().Debug().Str("op", op).EmbedObject(t).EmbedObject(v2).Msg("return")
	return v2, nil
}
