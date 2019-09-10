package event

import (
	"context"

	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/slog"
	"github.com/spf13/viper"
	"github.com/tecbot/gorocksdb"
	"go.opencensus.io/trace"
	"google.golang.org/grpc/codes"
)

const (
	op = "infra/rocksdb/event"
)

type Repository struct {
	db            *gorocksdb.DB
	dbForReadOnly *gorocksdb.DB
}

func NewRepository() (*Repository, func() error, error) {
	const op = op + ".NewRepository"
	v0 := viper.GetString("infra.rocksdb.event.path")
	v1, err := gorocksdb.OpenDb(gorocksdb.NewDefaultOptions(), v0)
	if err != nil {
		const op = op + ".gorocksdb.OpenDb"
		err := &errs.Error{Op: op, Code: codes.Internal, Err: err}
		return nil, nil, err
	}
	v2, err := gorocksdb.OpenDbForReadOnly(gorocksdb.NewDefaultOptions(), v0, false)
	if err != nil {
		const op = op + ".gorocksdb.OpenDbForReadOnly"
		err := &errs.Error{Op: op, Code: codes.Internal, Err: err}
		return nil, nil, err
	}
	v3 := &Repository{db: v1, dbForReadOnly: v2}
	return v3, v3.Close, nil
}

func (p *Repository) Close() error {
	return nil
}

func (p *Repository) Load(ctx context.Context, prefix string) (<-chan []byte, error) {
	const op = op + ".Repository.Load"
	if ctx == nil {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"}
		return nil, err
	}
	ctx, s := trace.StartSpan(ctx, op)
	defer s.End()
	a := make([]trace.Attribute, 0)
	defer s.AddAttributes(a...)
	slog.Logger().Debug().Str("op", op).EmbedObject(slog.Trace(ctx)).Str("prefix", prefix).Msg("arg")

	if prefix == "" {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'prefix' is ''"}
		s.SetStatus(trace.Status{Code: int32(codes.InvalidArgument), Message: err.Error()})
		return nil, err
	}

	err := &errs.Error{Op: op, Code: codes.Unimplemented}
	s.SetStatus(trace.Status{Code: int32(codes.Unimplemented), Message: err.Error()})
	return nil, err
}

func (p *Repository) Save(ctx context.Context, key string, payload []byte) error {
	const op = op + ".Repository.Save"
	if ctx == nil {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"}
		return err
	}
	ctx, s := trace.StartSpan(ctx, op)
	defer s.End()
	a := make([]trace.Attribute, 0)
	defer s.AddAttributes(a...)
	slog.Logger().Debug().Str("op", op).EmbedObject(slog.Trace(ctx)).Str("key", key).Bytes("payload", payload).Msg("arg")

	if key == "" {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'key' is ''"}
		s.SetStatus(trace.Status{Code: int32(codes.InvalidArgument), Message: err.Error()})
		return err
	}
	if len(payload) == 0 {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'payload' is empty"}
		s.SetStatus(trace.Status{Code: int32(codes.InvalidArgument), Message: err.Error()})
		return err
	}

	err := p.db.Put(gorocksdb.NewDefaultWriteOptions(), []byte(key), payload)
	if err != nil {
		const op = op + ".gorocksdb.DB.Put"
		err := &errs.Error{Op: op, Code: codes.Internal, Err: err}
		s.SetStatus(trace.Status{Code: int32(codes.Internal), Message: err.Error()})
		return err
	}
	return nil
}
