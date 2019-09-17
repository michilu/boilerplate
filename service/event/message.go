package event

import (
	"context"
	fmt "fmt"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/michilu/boilerplate/infra/keyvalue"
	"github.com/michilu/boilerplate/service/errs"
	_ "github.com/mwitkow/go-proto-validators" // https://github.com/mwitkow/go-proto-validators/blob/master/validator.proto
	"github.com/rs/xid"
	"google.golang.org/grpc/codes"
)

// [DDD Reference - Domain Language](https://domainlanguage.com/ddd/reference/)
const (
	// Occurred is a tag for the time the event occurred.
	Occurred = "occurred"
	// Entered is a tag for the time the event entered.
	Entered = "entered"
)

//go:generate interfacer -for github.com/michilu/boilerplate/service/event.Event -as event.Eventer -o entity-Eventer.go

// NewEvent returns a timestamp for the time the event occurred.
func NewEvent(timeStamp *time.Time, origin string) (Eventer, error) {
	const op = op + ".NewEvent"
	if timeStamp.IsZero() {
		return nil, &errs.Error{Op: op, Code: codes.InvalidArgument, Message: fmt.Sprintf("must be given. timeStamp is zero(%v)", timeStamp)}
	}
	if origin == "" {
		return nil, &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. origin is ''"}
	}
	v0 := xid.New()
	var v1 Eventer = &Event{
		Id:        v0.Bytes(),
		Origin:    origin,
		TimePoint: make([]*TimePoint, 0),
	}
	v1, err := v1.AddTimePoint(Occurred, *timeStamp)
	if err != nil {
		err := &errs.Error{Op: op, Code: codes.Internal, Err: err}
		return nil, err
	}
	v1, err = v1.AddTimePoint(Entered, v0.Time())
	if err != nil {
		err := &errs.Error{Op: op, Code: codes.Internal, Err: err}
		return nil, err
	}
	return v1, nil
}

// AddTimePoint returns a new Eventer with given the TimePoint.
func (p *Event) AddTimePoint(tag string, timeStamp time.Time) (Eventer, error) {
	const op = op + ".Event.AddTimePoint"
	if tag == "" {
		return nil, &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. tag is ''"}
	}
	if timeStamp.IsZero() {
		return nil, &errs.Error{Op: op, Code: codes.InvalidArgument, Message: fmt.Sprintf("must be given. timeStamp is zero(%v)", timeStamp)}
	}
	return &Event{
		Id:     p.GetId(),
		Origin: p.GetOrigin(),
		TimePoint: append(p.GetTimePoint(),
			&TimePoint{
				Timestamp: &timestamp.Timestamp{
					Seconds: timeStamp.Unix(),
					Nanos:   int32(timeStamp.UnixNano()),
				},
				Tag: tag,
			},
		),
	}, nil
}

func (p *Event) GetKey() []byte {
	return []byte(fmt.Sprintf("%s+%s", p.GetOrigin(), p.GetId()))
}

//go:generate genny -in=../topic/topic.go -out=gen-topic-EventWithContexter.go -pkg=$GOPACKAGE gen "ChanT=EventWithContexter"
//go:generate interfacer -for github.com/michilu/boilerplate/service/event.EventWithContext -as event.EventWithContexter -o if-EventWithContexter.go

type EventWithContext struct {
	Context context.Context
	Event   Eventer
}

func (p *EventWithContext) GetContext() context.Context { return p.Context }
func (p *EventWithContext) GetEvent() Eventer           { return p.Event }

func (p *EventWithContext) Validate() error {
	const op = op + ".EventWithContext.Validate"
	if p.Context == nil {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. '*EventWithContext.Context' is nil"}
		return err
	}
	if p.Event == nil {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. '*EventWithContext.Event' is nil"}
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

//go:generate genny -in=../topic/topic.go -out=gen-topic-KeyValueWithContexter.go -pkg=$GOPACKAGE gen "ChanT=KeyValueWithContexter"
//go:generate interfacer -for github.com/michilu/boilerplate/service/event.KeyValueWithContext -as event.KeyValueWithContexter -o if-KeyValueWithContexter.go

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
