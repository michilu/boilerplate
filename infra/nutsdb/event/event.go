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
	"github.com/xujiajun/nutsdb"
	"go.opencensus.io/trace"
	"google.golang.org/grpc/codes"
)

const (
	op = "infra/nutsdb/event"
)

type Repository struct {
	bucket string
	db     *nutsdb.DB
}

func NewOptions() nutsdb.Options {
	v0 := nutsdb.DefaultOptions
	return nutsdb.Options{
		EntryIdxMode:         v0.EntryIdxMode,
		SegmentSize:          v0.SegmentSize,
		NodeNum:              v0.NodeNum,
		RWMode:               v0.RWMode,
		SyncEnable:           v0.SyncEnable,
		StartFileLoadingMode: v0.StartFileLoadingMode,
	}
}

func NewRepository() (keyvalue.LoadSaveCloser, func() error, error) {
	const op = op + ".NewRepository"
	v0 := viper.GetString("infra.nutsdb.event.path")
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

	v1 := NewOptions()
	v1.Dir = v0
	v1.EntryIdxMode = nutsdb.HintKeyAndRAMIdxMode
	v2, err := nutsdb.Open(v1)
	if err != nil {
		const op = op + ".nutsdb.Open"
		err := &errs.Error{Op: op, Code: codes.Internal, Err: err}
		return nil, nil, err
	}
	v3 := &Repository{
		bucket: "event",
		db:     v2,
	}
	return v3, v3.Close, nil
}

func (p *Repository) Close() error {
	const op = op + ".Repository.Close"
	if p.db == nil {
		return nil
	}
	err := p.db.Close()
	if err != nil {
		const op = op + ".nutsdb.DB.Close"
		err := &errs.Error{Op: op, Code: codes.Internal, Err: err}
		return err
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
	t := slog.Trace(ctx)

	{
		if prefix == nil {
			err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'prefix' is nil"}
			s.SetStatus(trace.Status{Code: int32(codes.InvalidArgument), Message: err.Error()})
			slog.Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
			return nil, err
		}
		err := prefix.Validate()
		if err != nil {
			err := &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
			s.SetStatus(trace.Status{Code: int32(codes.InvalidArgument), Message: err.Error()})
			slog.Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
			return nil, err
		}
		slog.Logger().Debug().Str("op", op).EmbedObject(t).EmbedObject(prefix).Msg("arg")
	}

	err := &errs.Error{Op: op, Code: codes.Unimplemented}
	s.SetStatus(trace.Status{Code: int32(codes.Unimplemented), Message: err.Error()})
	slog.Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
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
	t := slog.Trace(ctx)

	{
		if keyvalue == nil {
			err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'keyvalue' is nil"}
			s.SetStatus(trace.Status{Code: int32(codes.InvalidArgument), Message: err.Error()})
			slog.Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
			return err
		}
		err := keyvalue.Validate()
		if err != nil {
			err := &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
			s.SetStatus(trace.Status{Code: int32(codes.InvalidArgument), Message: err.Error()})
			slog.Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
			return err
		}
		s.AddAttributes(trace.StringAttribute("arg", keyvalue.String()))
		slog.Logger().Debug().Str("op", op).EmbedObject(t).EmbedObject(keyvalue).Msg("arg")
	}

	{
		err := p.db.Update(func(tx *nutsdb.Tx) error {
			v0 := keyvalue.GetKey()
			v1 := keyvalue.GetValue()
			s.AddAttributes(
				trace.StringAttribute("v0", fmt.Sprintf("%v", v0)),
				trace.StringAttribute("v1", fmt.Sprintf("%v", v1)),
				trace.Int64Attribute("v0 size", int64(len(v0))),
				trace.Int64Attribute("v1 size", int64(len(v1))),
			)
			err := tx.Put(p.bucket, v0, v1, 0)
			if err != nil {
				const op = op + ".nutsdb.Tx.Put"
				err := &errs.Error{Op: op, Code: codes.Internal, Err: err}
				s.SetStatus(trace.Status{Code: int32(codes.Internal), Message: err.Error()})
				slog.Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}
	return nil
}
