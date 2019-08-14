package debug

import (
	"context"

	"github.com/michilu/boilerplate/service/debug"
	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/slog"
	"google.golang.org/grpc/codes"
)

// NewClientRepository returns a new ClientRepository
func NewClientRepository() debug.ClientRepository {
	return &clientRepository{}
}

type clientRepository struct{}

func (*clientRepository) Config(ctx context.Context) (debug.ClientWithCtxer, error) {
	const op = op + ".Config"
	v0, err := GenerateUUID()
	if err != nil {
		return nil, err
	}
	v1 := &debug.ClientWithCtx{
		Ctx:    ctx,
		Client: debug.Client{Id: v0},
	}
	return v1, nil
}

func (*clientRepository) Connect(m debug.ClientWithCtxer) error {
	const op = op + ".clientRepository.Connect"
	{
		slog.Logger().Debug().Str("op", op).Msg("start")
		defer slog.Logger().Debug().Str("op", op).Msg("end")
	}
	err := m.Validate()
	if err != nil {
		return err
	}
	ch := make(chan error)
	go func(ch chan<- error, m debug.ClientWithCtxer) {
		defer close(ch)
		ctx := m.GetCtx()
		select {
		case ch <- OpenDebugPort(m):
		case <-ctx.Done():
			ch <- ctx.Err()
		}
	}(ch, m)
	ctx := m.GetCtx()
	select {
	case <-ctx.Done():
		const op = op + ".ctx.Done"
		return &errs.Error{Op: op, Code: codes.Aborted, Err: ctx.Err()}
	case err := <-ch:
		if err != nil {
			return err
		}
		return nil
	}
}
