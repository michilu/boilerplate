package keyvalue

import "context"

const (
	op = "infra/keyvalue"
)

//go:generate interfacer -for github.com/michilu/boilerplate/infra/keyvalue.KeyValue -as keyvalue.KeyValuer -o entity-KeyValuer.go
//go:generate interfacer -for github.com/michilu/boilerplate/infra/keyvalue.Prefix -as keyvalue.Prefixer -o vo-Prefixer.go

type LoadSaveCloser interface {
	Load(context.Context, Prefixer) (<-chan KeyValuer, error)
	Save(context.Context, KeyValuer) error
	Close() error
}
