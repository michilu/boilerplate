package event

import (
	fmt "fmt"
	"time"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	"github.com/michilu/boilerplate/infra/keyvalue"
	"github.com/michilu/boilerplate/pb"
	"github.com/michilu/boilerplate/service/errs"
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

var (
	GetTopicKeyValueWithContexter = keyvalue.GetTopicKeyValueWithContexter
)

type (
	Event     = pb.Event
	Eventer   = pb.Eventer
	TimePoint = pb.EventTimePoint

	KeyValueWithContext   = keyvalue.KeyValueWithContext
	KeyValueWithContexter = keyvalue.KeyValueWithContexter
)

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

//go:generate genny -in=../topic/with-context.go -out=gen-EventWithContext.go -pkg=$GOPACKAGE gen "T=Event Ier=Eventer"
//go:generate genny -in=../topic/topic.go -out=gen-topic-EventWithContexter.go -pkg=$GOPACKAGE gen "ChanT=EventWithContexter"
//go:generate interfacer -for github.com/michilu/boilerplate/service/event.EventWithContext -as event.EventWithContexter -o if-EventWithContexter.go
