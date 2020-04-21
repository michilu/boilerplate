package event

import (
	"context"

	"github.com/michilu/boilerplate/infra/keyvalue"
)

type Loader interface {
	Load(context.Context, keyvalue.KeyPrefixer) (<-chan keyvalue.KeyValuer, error)
}

type Saver interface {
	Save(context.Context, keyvalue.KeyValuer) error
}

type Closer interface {
	Close() error
}
