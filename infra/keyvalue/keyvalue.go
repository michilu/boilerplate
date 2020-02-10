package keyvalue

import (
	"context"
)

const (
	op = "infra/keyvalue"
)

type KeyValueCloser interface {
	Close() error
	Delete(context.Context, Keyer) error
	Get(context.Context, Keyer) (KeyValuer, error)
	Put(context.Context, KeyValuer) error
}
