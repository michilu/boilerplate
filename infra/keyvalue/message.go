package keyvalue

import (
	"context"
)

const (
	op = "infra/keyvalue"
)

type LoadSaveCloser interface {
	Load(context.Context, Prefixer) (<-chan KeyValuer, error)
	Save(context.Context, KeyValuer) error
	//Delete(context.Context, Keyer) error
	Close() error
}

//go:generate interfacer -for github.com/michilu/boilerplate/infra/keyvalue.Key -as keyvalue.Keyer -o entity-Keyer.go
//go:generate interfacer -for github.com/michilu/boilerplate/infra/keyvalue.KeyValue -as keyvalue.KeyValuer -o entity-KeyValuer.go
//go:generate interfacer -for github.com/michilu/boilerplate/infra/keyvalue.Prefix -as keyvalue.Prefixer -o vo-Prefixer.go

//go:generate genny -in=../../service/topic/with-context.go -out=gen-KeyWithContext.go -pkg=$GOPACKAGE gen "T=Key Ier=Keyer"
//go:generate genny -in=../../service/topic/with-context.go -out=gen-KeyValueWithContext.go -pkg=$GOPACKAGE gen "T=KeyValue Ier=KeyValuer"

//go:generate interfacer -for github.com/michilu/boilerplate/infra/keyvalue.KeyWithContext -as keyvalue.KeyWithContexter -o if-KeyWithContexter.go
//go:generate interfacer -for github.com/michilu/boilerplate/infra/keyvalue.KeyValueWithContext -as keyvalue.KeyValueWithContexter -o if-KeyValueWithContexter.go

//go:generate genny -in=../../service/topic/topic.go -out=gen-topic-KeyWithContexter.go -pkg=$GOPACKAGE gen "ChanT=KeyWithContexter"
//go:generate genny -in=../../service/topic/topic.go -out=gen-topic-KeyValueWithContexter.go -pkg=$GOPACKAGE gen "ChanT=KeyValueWithContexter"
