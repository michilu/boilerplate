package keyvalue

import (
	"context"

	"github.com/michilu/boilerplate/pb"
)

const (
	op = "infra/keyvalue"
)

type Key = pb.Key

type KeyValueCloser interface {
	Close() error
	Delete(context.Context, Keyer) error
	Get(context.Context, Keyer) (KeyValuer, error)
	Put(context.Context, KeyValuer) error
}
