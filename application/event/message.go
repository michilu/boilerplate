package event

import (
	"context"
	"fmt"

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

//go:generate interfacer -for github.com/michilu/boilerplate/application/event.ByteWithContext -as event.ByteWithContexter -o if-ByteWithContexter.go
//go:generate genny -in=../../service/topic/topic.go -out=gen-topic-ByteWithContexter.go -pkg=$GOPACKAGE gen "ChanT=ByteWithContexter"

type ByteWithContext struct {
	Context context.Context
	Byte    *Byte
}

func (p *ByteWithContext) GetContext() context.Context { return p.Context }
func (p *ByteWithContext) GetByte() *Byte              { return p.Byte }

func (p *ByteWithContext) Validate() error {
	const op = op + ".Byte.Validate"
	if p.Context == nil {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. '*Byte.Context' is nil"}
		return err
	}
	err := p.Byte.Validate()
	if err != nil {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
		return err
	}
	return nil
}

func (p *ByteWithContext) String() string {
	return fmt.Sprintf("ByteWithContext<Context: %v, Byte: %s>", p.GetContext(), p.GetByte().String())
}
