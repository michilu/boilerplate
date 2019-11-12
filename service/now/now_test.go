package now

import (
	"reflect"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
)

func TestTimestampFromTime(t *testing.T) {
	type args struct {
		v time.Time
	}
	tests := []struct {
		name string
		args args
		want *timestamp.Timestamp
	}{
		{"zero", args{}, &timestamp.Timestamp{Seconds: -62135596800}},
		{"unixtime", args{time.Unix(0, 0)}, &timestamp.Timestamp{}},
		{"ok", args{time.Unix(0, 1234567890123456789)}, &timestamp.Timestamp{Seconds: 1234567890, Nanos: 123456789}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimestampFromTime(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimestampFromTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeFromTimestamp(t *testing.T) {
	type args struct {
		v *timestamp.Timestamp
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{"nil", args{}, time.Time{}, true},
		{"zero", args{&timestamp.Timestamp{}}, time.Unix(0, 0), false},
		{"ok", args{&timestamp.Timestamp{Seconds: 1234567890, Nanos: 123456789}}, time.Unix(1234567890, 123456789), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TimeFromTimestamp(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("TimeFromTimestamp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeFromTimestamp() = %v, want %v", got, tt.want)
			}
		})
	}
}
