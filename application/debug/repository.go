package debug

import (
	"context"

	"github.com/michilu/boilerplate/service/debug"
	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/slog"
	"go.opencensus.io/trace"
	"google.golang.org/grpc/codes"
)

// NewClientRepository returns a new ClientRepository
func NewClientRepository() debug.ClientRepository {
	return &clientRepository{}
}

type clientRepository struct{}

func (*clientRepository) Config(ctx context.Context) (debug.ClientWithContexter, error) {
	const op = op + ".Config"
	if ctx == nil {
		return nil, &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'm' is nil"}
	}
	ctx, s := trace.StartSpan(ctx, op)
	defer s.End()
	a := make([]trace.Attribute, 0)
	defer s.AddAttributes(a...)

	v0, err := GenerateUUID(ctx)
	if err != nil {
		s.SetStatus(trace.Status{Code: int32(codes.Internal), Message: err.Error()})
		return nil, err
	}
	a = append(a, trace.StringAttribute("GenerateUUID", v0))
	v1 := &debug.ClientWithContext{
		Context: ctx,
		Client:  debug.Client{Id: v0},
	}
	a = append(a, trace.StringAttribute("debug.ClientWithContext", v1.String()))
	return v1, nil
}

func (*clientRepository) Connect(m debug.ClientWithContexter) error {
	const op = op + ".clientRepository.Connect"
	if m == nil {
		return &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'm' is nil"}
	}
	ctx := m.GetContext()
	if ctx == nil {
		return &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"}
	}
	ctx, s := trace.StartSpan(ctx, op)
	defer s.End()
	a := make([]trace.Attribute, 0)
	defer s.AddAttributes(a...)
	t := slog.Trace(ctx)

	{
		if m == nil {
			err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'm' is nil"}
			s.SetStatus(trace.Status{Code: int32(codes.InvalidArgument), Message: err.Error()})
			return err
		}
		slog.Logger().Debug().Str("op", op).EmbedObject(t).EmbedObject(m).Msg("arg")
		a = append(a, trace.StringAttribute("debug.ClientWithContexter", m.String()))
	}
	{
		err := m.Validate()
		if err != nil {
			s.SetStatus(trace.Status{Code: int32(codes.InvalidArgument), Message: err.Error()})
			return err
		}
	}
	ch := make(chan error)
	go func(ctx context.Context, ch chan<- error, m debug.ClientWithContexter) {
		defer close(ch)
		ctx, s := trace.StartSpan(ctx, op)
		defer s.End()
		a := make([]trace.Attribute, 0)
		defer s.AddAttributes(a...)
		vctx := m.GetContext()
		select {
		case ch <- OpenDebugPort(ctx, m):
		case <-vctx.Done():
			err := vctx.Err()
			if err != nil {
				s.SetStatus(trace.Status{Code: int32(codes.Unknown), Message: err.Error()})
			}
			ch <- vctx.Err()
		}
	}(ctx, ch, m)
	vctx := m.GetContext()
	select {
	case <-vctx.Done():
		const op = op + ".ctx.Done"
		err := &errs.Error{Op: op, Code: codes.Aborted, Err: vctx.Err()}
		s.SetStatus(trace.Status{Code: int32(codes.Aborted), Message: err.Error()})
		return err
	case err := <-ch:
		if err != nil {
			s.SetStatus(trace.Status{Code: int32(codes.Unknown), Message: err.Error()})
			return err
		}
		return nil
	}
}
