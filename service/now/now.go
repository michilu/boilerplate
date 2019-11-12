package now

import (
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/michilu/boilerplate/service/errs"
	"google.golang.org/grpc/codes"
)

const (
	op = "service/now"
)

var (
	// Now returns a time.Time.
	Now func() time.Time = time.Now
)

// TimestampFromTime returns a timestamp.Timestamp of the given time.Time.
func TimestampFromTime(v time.Time) *timestamp.Timestamp {
	v0 := v.Unix()
	return &timestamp.Timestamp{
		Seconds: v0,
		Nanos:   int32(v.UnixNano() - v0*1000000000),
	}
}

// TimeFromTimestamp returns a time.Time of the given timestamp.Timestamp.
func TimeFromTimestamp(v *timestamp.Timestamp) (time.Time, error) {
	const op = op + ".TimeFromTimestamp"
	if v == nil {
		err := &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "ust be given. v is 'nil'"}
		return time.Time{}, err
	}
	return time.Unix(v.Seconds, int64(v.Nanos)), nil
}
