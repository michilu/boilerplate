package event

import (
	"context"

	"github.com/golang/protobuf/proto"
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
	t := slog.Trace(ctx, s)
	slog.Logger().Debug().Str("op", op).EmbedObject(t).Bytes("b", b).Msg(op + ": arg")

	v0 := &Event{}
	if err := proto.Unmarshal(b, v0); err != nil {
		const op = op + ".proto.Unmarshal"
		return nil, &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
	}
	v1 := make([]*TimePoint, 0, len(v0.TimePoint)+1)
	occurred := false
	for _, v := range v0.GetTimePoint() {
		v2 := v
		switch v2.GetTag() {
		case Entered:
			continue
		case Occurred:
			occurred = true
		}
		v1 = append(v1, v2)
	}
	var v3 Eventer = &Event{
		Id:        v0.GetId(),
		TimePoint: v1,
	}
	v4 := now.Now()
	var err error
	if !occurred {
		v3, err = v3.AddTimePoint(Occurred, v4)
		if err != nil {
			const op = op + ".AddTimePoint(Occurred)"
			err := &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
			s.SetStatus(trace.Status{Code: int32(codes.InvalidArgument), Message: err.Error()})
			slog.Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
			return nil, err
		}
	}
	v3, err = v3.AddTimePoint(Entered, v4)
	if err != nil {
		const op = op + ".AddTimePoint(Entered)"
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
		s.SetStatus(trace.Status{Code: int32(codes.InvalidArgument), Message: err.Error()})
		slog.Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
		return nil, err
	}
	slog.Logger().Debug().Str("op", op).EmbedObject(t).EmbedObject(v3).Msg(op + ": return")
	return v3, nil
}
