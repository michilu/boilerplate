package pb

import (
	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/michilu/boilerplate/service/errs"
	"google.golang.org/grpc/codes"
)

func (p *Event) GetKey() []byte {
	return []byte(fmt.Sprintf("%s+%s", p.GetOrigin(), p.GetId()))
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
			&EventTimePoint{
				Timestamp: &timestamp.Timestamp{
					Seconds: timeStamp.Unix(),
					Nanos:   int32(timeStamp.UnixNano()),
				},
				Tag: tag,
			},
		),
	}, nil
}
