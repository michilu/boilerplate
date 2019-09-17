package event

import (
	"context"
	"fmt"

	"github.com/michilu/boilerplate/infra/keyvalue"
	"github.com/michilu/boilerplate/service/errs"
	"github.com/michilu/boilerplate/service/event"
	"google.golang.org/grpc/codes"
)

//go:generate interfacer -for github.com/michilu/boilerplate/application/event.EventWithContext -as event.EventWithContexter -o if-EventWithContexter.go
//go:generate genny -in=../../service/topic/topic.go -out=gen-topic-EventWithContexter.go -pkg=$GOPACKAGE gen "ChanT=EventWithContexter"

type EventWithContext struct {
	Context context.Context
	Event   event.Eventer
}

func (p *EventWithContext) GetContext() context.Context { return p.Context }
func (p *EventWithContext) GetEvent() event.Eventer     { return p.Event }

func (p *EventWithContext) Validate() error {
	const op = op + ".Event.Validate"
	if p.Context == nil {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. '*Event.Context' is nil"}
		return err
	}
	if p.Event == nil {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. '*Event.Event' is nil"}
		return err
	}
	err := p.Event.Validate()
	if err != nil {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
		return err
	}
	return nil
}

func (p *EventWithContext) String() string {
	return fmt.Sprintf("EventWithContext<Context: %v, Event: %s>", p.GetContext(), p.GetEvent().String())
}

//go:generate interfacer -for github.com/michilu/boilerplate/application/event.KeyValueWithContext -as event.KeyValueWithContexter -o if-KeyValueWithContexter.go
//go:generate genny -in=../../service/topic/topic.go -out=gen-topic-KeyValueWithContexter.go -pkg=$GOPACKAGE gen "ChanT=KeyValueWithContexter"

type KeyValueWithContext struct {
	Context  context.Context
	KeyValue keyvalue.KeyValuer
}

func (p *KeyValueWithContext) GetContext() context.Context     { return p.Context }
func (p *KeyValueWithContext) GetKeyValue() keyvalue.KeyValuer { return p.KeyValue }

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
