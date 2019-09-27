package keyvalue

import (
	"context"
	fmt "fmt"

	"github.com/michilu/boilerplate/service/errs"
	"google.golang.org/grpc/codes"
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

//go:generate genny -in=../../service/topic/topic.go -out=gen-topic-KeyValueWithContexter.go -pkg=$GOPACKAGE gen "ChanT=KeyValueWithContexter"
//go:generate interfacer -for github.com/michilu/boilerplate/infra/keyvalue.KeyValueWithContext -as keyvalue.KeyValueWithContexter -o if-KeyValueWithContexter.go

type KeyValueWithContext struct {
	Context  context.Context
	KeyValue KeyValuer
}

func (p *KeyValueWithContext) GetContext() context.Context { return p.Context }
func (p *KeyValueWithContext) GetKeyValue() KeyValuer      { return p.KeyValue }

func (p *KeyValueWithContext) Validate() error {
	const op = op + ".KeyValueWithContext.Validate"
	if p.Context == nil {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. '*KeyValueWithContext.Context' is nil"}
		return err
	}
	{
		if p.KeyValue == nil {
			err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. '*KeyValueWithContext.KeyValue' is nil"}
			return err
		}
		err := p.KeyValue.Validate()
		if err != nil {
			err := &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
			return err
		}
	}
	return nil
}

func (p *KeyValueWithContext) String() string {
	return fmt.Sprintf("KeyValueWithContext<Context: %v, KeyValue: %s>", p.GetContext(), p.GetKeyValue().String())
}
