package keystore

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	k "github.com/michilu/boilerplate/application/config"
	"github.com/michilu/boilerplate/infra/keyvalue"
	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/slog"
	"github.com/spf13/viper"
	"github.com/xujiajun/nutsdb"
	"go.opencensus.io/trace"
	"google.golang.org/grpc/codes"
)

const (
	op = "infra/nutsdb/keystore"
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

func NewRepository(ctx context.Context) (keyvalue.KeyValueCloser, func() error, error) {
	const op = op + ".NewRepository"
	if ctx == nil {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"}
		return nil, nil, err
	}
	ctx, s := trace.StartSpan(ctx, op)
	defer s.End()
	t := slog.Trace(ctx, s)

	const c0 = k.InfraNutsdbKeystorePath
	v0 := viper.GetString(c0)
	if v0 == "" {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: fmt.Sprintf("must be set '%v' in config.toml.", c0)}
		s.SetStatus(trace.Status{Code: int32(codes.InvalidArgument), Message: err.Error()})
		slog.Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
		return nil, nil, err
	}

	v1 := 0
MKDIR:
	{
		v1 := filepath.Dir(v0)
		_, err := os.Stat(v1)
		if os.IsNotExist(err) {
			err := os.MkdirAll(v1, 0777)
			if err != nil {
				const op = op + ".os.MkdirAll"
				err := &errs.Error{Op: op, Code: codes.Unavailable, Err: err}
				s.SetStatus(trace.Status{Code: int32(codes.Unavailable), Message: err.Error()})
				slog.Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
				return nil, nil, err
			}
		}
	}

	v2 := NewOptions()
	v2.Dir = v0
	v2.EntryIdxMode = nutsdb.HintKeyAndRAMIdxMode
	v3, err := nutsdb.Open(v2)
	if err != nil {
		const op = op + ".nutsdb.Open"
		{
			err := &errs.Error{Op: op, Code: codes.Internal, Err: err}
			slog.Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
			switch {
			case viper.GetBool(k.InfraNutsdbKeystoreAutoRecovery):
				// pass
			case 3 < v1:
				fallthrough
			default:
				s.SetStatus(trace.Status{Code: int32(codes.Internal), Message: err.Error()})
				return nil, nil, err
			}
		}
		{
			v1++
			err := os.RemoveAll(v0)
			if err != nil {
				const op = op + ".os.RemoveAll"
				err := &errs.Error{Op: op, Code: codes.Internal, Err: err}
				s.SetStatus(trace.Status{Code: int32(codes.Internal), Message: err.Error()})
				slog.Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
				return nil, nil, err
			}
		}
		goto MKDIR
	}
	v4 := &Repository{
		bucket: "keystore",
		db:     v3,
	}
	return v4, v4.Close, nil
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

func (p *Repository) Delete(ctx context.Context, key keyvalue.Keyer) error {
	const op = op + ".Repository.Delete"
	if ctx == nil {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"}
		return err
	}
	ctx, s := trace.StartSpan(ctx, op)
	defer s.End()
	t := slog.Trace(ctx, s)

	if key == nil {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'key' is nil"}
		s.SetStatus(trace.Status{Code: int32(codes.InvalidArgument), Message: err.Error()})
		slog.Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
		return err
	}
	{
		s.AddAttributes(trace.StringAttribute("key", key.String()))
		slog.Logger().Debug().Str("op", op).EmbedObject(t).EmbedObject(key).Msg(op + ": arg")
	}
	{
		err := key.Validate()
		if err != nil {
			err := &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
			s.SetStatus(trace.Status{Code: int32(codes.InvalidArgument), Message: err.Error()})
			slog.Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
			return err
		}
	}

	s.AddAttributes(trace.Int64Attribute("db.KeyCount", int64(p.db.KeyCount)))

	err := p.db.View(func(tx *nutsdb.Tx) error {
		err := tx.Delete(p.bucket, key.GetKey())
		if err != nil {
			const op = op + ".tx.Delete"
			err := &errs.Error{Op: op, Code: codes.Internal, Err: err}
			s.SetStatus(trace.Status{Code: int32(codes.Internal), Message: err.Error()})
			slog.Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
			return err
		}
		return nil
	})
	if err != nil {
		const op = op + ".db.View"
		err := &errs.Error{Op: op, Code: codes.Internal, Err: err}
		s.SetStatus(trace.Status{Code: int32(codes.Internal), Message: err.Error()})
		slog.Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
		return err
	}
	return nil
}

func (p *Repository) Get(ctx context.Context, key keyvalue.Keyer) (keyvalue.KeyValuer, error) {
	const op = op + ".Repository.Get"
	if ctx == nil {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"}
		return nil, err
	}
	ctx, s := trace.StartSpan(ctx, op)
	defer s.End()
	t := slog.Trace(ctx, s)

	if key == nil {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'key' is nil"}
		s.SetStatus(trace.Status{Code: int32(codes.InvalidArgument), Message: err.Error()})
		slog.Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
		return nil, err
	}
	{
		s.AddAttributes(trace.StringAttribute("key", key.String()))
		slog.Logger().Debug().Str("op", op).EmbedObject(t).EmbedObject(key).Msg(op + ": arg")
	}
	{
		err := key.Validate()
		if err != nil {
			err := &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
			s.SetStatus(trace.Status{Code: int32(codes.InvalidArgument), Message: err.Error()})
			slog.Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
			return nil, err
		}
	}

	s.AddAttributes(trace.Int64Attribute("db.KeyCount", int64(p.db.KeyCount)))

	var v0 *nutsdb.Entry
	err := p.db.View(func(tx *nutsdb.Tx) error {
		v1, err := tx.Get(p.bucket, key.GetKey())
		if err != nil {
			const op = op + ".tx.Get"
			err := &errs.Error{Op: op, Code: codes.Internal, Err: err}
			s.SetStatus(trace.Status{Code: int32(codes.Internal), Message: err.Error()})
			slog.Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
			return err
		}
		v0 = v1
		return nil
	})
	if err != nil {
		const op = op + ".db.View"
		err := &errs.Error{Op: op, Code: codes.Internal, Err: err}
		s.SetStatus(trace.Status{Code: int32(codes.Internal), Message: err.Error()})
		slog.Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
		return nil, err
	}
	s.AddAttributes(trace.StringAttribute("v0", fmt.Sprintf("%v", v0)))
	if v0 == nil {
		slog.Logger().Debug().Str("op", op).EmbedObject(t).Str("v0", "nil").Msg(op + ": return")
		return nil, nil
	}
	v2 := &keyvalue.KeyValue{
		Key:   v0.Key,
		Value: v0.Value,
	}
	slog.Logger().Debug().Str("op", op).EmbedObject(t).EmbedObject(v2).Msg(op + ": return")
	return v2, nil
}

func (p *Repository) Put(ctx context.Context, keyvalue keyvalue.KeyValuer) error {
	const op = op + ".Repository.Put"
	if ctx == nil {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'ctx' is nil"}
		return err
	}
	ctx, s := trace.StartSpan(ctx, op)
	defer s.End()
	t := slog.Trace(ctx, s)

	if keyvalue == nil {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'keyvalue' is nil"}
		s.SetStatus(trace.Status{Code: int32(codes.InvalidArgument), Message: err.Error()})
		slog.Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
		return err
	}
	{
		s.AddAttributes(trace.StringAttribute("keyvalue", keyvalue.String()))
		slog.Logger().Debug().Str("op", op).EmbedObject(t).EmbedObject(keyvalue).Msg(op + ": arg")
	}
	{
		err := keyvalue.Validate()
		if err != nil {
			err := &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
			s.SetStatus(trace.Status{Code: int32(codes.InvalidArgument), Message: err.Error()})
			slog.Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
			return err
		}
	}

	{
		s.AddAttributes(trace.Int64Attribute("db.KeyCount:before", int64(p.db.KeyCount)))
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
			const op = op + ".db.Update"
			err := &errs.Error{Op: op, Code: codes.Internal, Err: err}
			s.SetStatus(trace.Status{Code: int32(codes.Internal), Message: err.Error()})
			slog.Logger().Err(err).Str("op", op).EmbedObject(t).Msg(err.Error())
			return err
		}
		s.AddAttributes(trace.Int64Attribute("db.KeyCount:after", int64(p.db.KeyCount)))
	}
	return nil
}
