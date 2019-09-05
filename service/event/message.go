package event

import (
	fmt "fmt"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/michilu/boilerplate/service/errs"
	_ "github.com/mwitkow/go-proto-validators" // https://github.com/mwitkow/go-proto-validators/blob/master/validator.proto
	"github.com/rs/xid"
	"google.golang.org/grpc/codes"
)

// [DDD Reference - Domain Language](https://domainlanguage.com/ddd/reference/)
const (
	// Occurred is a tag for the time the event occurred.
	Occurred = "occurred"
	// Entered is a tag for the time the event entered
	Entered = "entered"
)

//go:generate interfacer -for github.com/michilu/boilerplate/service/event.Event -as event.Eventer -o entity-Eventer.go

// NewEvent returns a timestamp for the time the event occurred.
func NewEvent(now *time.Time) Eventer {
	id := xid.New()
	return &Event{
		Id: id.Bytes(),
		TimePoint: []*TimePoint{
			&TimePoint{
				Tag: Occurred,
				Timestamp: &timestamp.Timestamp{
					Seconds: now.Unix(),
					Nanos:   int32(now.UnixNano()),
				},
			},
			&TimePoint{
				Tag: Entered,
				Timestamp: &timestamp.Timestamp{
					Seconds: id.Time().Unix(),
					Nanos:   int32(id.Time().UnixNano()),
				},
			},
		},
	}
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
		Id: p.GetId(),
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
