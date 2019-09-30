package keyvalue

import (
	"context"
)

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

//go:generate genny -in=../../service/topic/with-context.go -out=gen-KeyValueWithContext.go -pkg=$GOPACKAGE gen "T=KeyValue Ier=KeyValuer"
//go:generate genny -in=../../service/topic/topic.go -out=gen-topic-KeyValueWithContexter.go -pkg=$GOPACKAGE gen "ChanT=KeyValueWithContexter"
//go:generate interfacer -for github.com/michilu/boilerplate/infra/keyvalue.KeyValueWithContext -as keyvalue.KeyValueWithContexter -o if-KeyValueWithContexter.go
