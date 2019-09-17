package event

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/michilu/boilerplate/infra/keyvalue"
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

func NewRepository() (keyvalue.LoadSaveCloser, func() error, error) {
	const op = op + ".NewRepository"
	v0 := viper.GetString("infra.rocksdb.event.path")
	{
		v1 := filepath.Dir(v0)
		_, err := os.Stat(v1)
		if os.IsNotExist(err) {
			err := os.MkdirAll(v1, 0777)
			if err != nil {
				const op = op + ".os.MkdirAll"
				err := &errs.Error{Op: op, Code: codes.Unavailable, Err: err}
				return nil, nil, err
			}
		}
	}
	v1 := gorocksdb.NewDefaultOptions()
	v1.SetCreateIfMissing(true)
	v2, err := gorocksdb.OpenDb(v1, v0)
	if err != nil {
		const op = op + ".gorocksdb.OpenDb"
		err := &errs.Error{Op: op, Code: codes.Internal, Err: err}
		return nil, nil, err
	}
	v3, err := gorocksdb.OpenDbForReadOnly(gorocksdb.NewDefaultOptions(), v0, false)
	if err != nil {
		const op = op + ".gorocksdb.OpenDbForReadOnly"
		err := &errs.Error{Op: op, Code: codes.Internal, Err: err}
		return nil, nil, err
	}
	v4 := &Repository{db: v2, dbForReadOnly: v3}
	return v4, v4.Close, nil
}

func (p *Repository) Close() (err error) {
	const op = op + ".Repository.Close"
	defer func() {
		const op = op + ".recover"
		e := recover()
		if e != nil {
			err = &errs.Error{Op: op, Code: codes.Internal, Message: fmt.Sprintf("%v", e)}
		}
	}()
	if p.db != nil {
		defer p.db.Close()
	}
	if p.dbForReadOnly != nil {
		defer p.dbForReadOnly.Close()
	}
	return nil
}

func (p *Repository) Load(ctx context.Context, prefix keyvalue.Prefixer) (<-chan keyvalue.KeyValuer, error) {
	const op = op + ".Repository.Load"
	if ctx == nil {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"}
		return nil, err
	}
	ctx, s := trace.StartSpan(ctx, op)
	defer s.End()
	a := make([]trace.Attribute, 0)
	defer s.AddAttributes(a...)
	t := slog.Trace(ctx)

	{
		if prefix == nil {
			err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'prefix' is nil"}
			s.SetStatus(trace.Status{Code: int32(codes.InvalidArgument), Message: err.Error()})
			return nil, err
		}
		err := prefix.Validate()
		if err != nil {
			err := &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
			s.SetStatus(trace.Status{Code: int32(codes.InvalidArgument), Message: err.Error()})
			return nil, err
		}
		slog.Logger().Debug().Str("op", op).EmbedObject(t).EmbedObject(prefix).Msg("arg")
	}

	err := &errs.Error{Op: op, Code: codes.Unimplemented}
	s.SetStatus(trace.Status{Code: int32(codes.Unimplemented), Message: err.Error()})
	return nil, err
}

func (p *Repository) Save(ctx context.Context, keyvalue keyvalue.KeyValuer) error {
	const op = op + ".Repository.Save"
	if ctx == nil {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"}
		return err
	}
	ctx, s := trace.StartSpan(ctx, op)
	defer s.End()
	a := make([]trace.Attribute, 0)
	defer s.AddAttributes(a...)
	t := slog.Trace(ctx)

	{
		if keyvalue == nil {
			err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'keyvalue' is nil"}
			s.SetStatus(trace.Status{Code: int32(codes.InvalidArgument), Message: err.Error()})
			return err
		}
		err := keyvalue.Validate()
		if err != nil {
			err := &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
			s.SetStatus(trace.Status{Code: int32(codes.InvalidArgument), Message: err.Error()})
			return err
		}
		slog.Logger().Debug().Str("op", op).EmbedObject(t).EmbedObject(keyvalue).Msg("arg")
	}

	err := p.db.Put(gorocksdb.NewDefaultWriteOptions(), keyvalue.GetKey(), keyvalue.GetValue())
	if err != nil {
		const op = op + ".gorocksdb.DB.Put"
		err := &errs.Error{Op: op, Code: codes.Internal, Err: err}
		s.SetStatus(trace.Status{Code: int32(codes.Internal), Message: err.Error()})
		return err
	}
	return nil
}
