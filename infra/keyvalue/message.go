package keyvalue

import (
	"context"

	"github.com/michilu/boilerplate/pb"
)

type Keyer = pb.Keyer
type KeyValue = pb.KeyValue
type KeyValuer = pb.KeyValuer
type KeyPrefixer = pb.KeyPrefixer

type LoadSaveCloser interface {
	Load(context.Context, KeyPrefixer) (<-chan KeyValuer, error)
	Save(context.Context, KeyValuer) error
	//Delete(context.Context, Keyer) error
	Close() error
}

//go:generate genny -in=../../service/topic/with-context.go -out=gen-KeyWithContext.go -pkg=$GOPACKAGE gen "T=Key Ier=Keyer"
//go:generate genny -in=../../service/topic/with-context.go -out=gen-KeyValueWithContext.go -pkg=$GOPACKAGE gen "T=KeyValue Ier=KeyValuer"

//go:generate interfacer -for github.com/michilu/boilerplate/infra/keyvalue.KeyWithContext -as keyvalue.KeyWithContexter -o if-KeyWithContexter.go
//go:generate interfacer -for github.com/michilu/boilerplate/infra/keyvalue.KeyValueWithContext -as keyvalue.KeyValueWithContexter -o if-KeyValueWithContexter.go

//go:generate genny -in=../../service/topic/topic.go -out=gen-topic-KeyWithContexter.go -pkg=$GOPACKAGE gen "ChanT=KeyWithContexter"
//go:generate genny -in=../../service/topic/topic.go -out=gen-topic-KeyValueWithContexter.go -pkg=$GOPACKAGE gen "ChanT=KeyValueWithContexter"
