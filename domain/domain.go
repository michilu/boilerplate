package domain

//go:generate genny -in=../v/topic/topic.go -out=gen-topic-struct.go -pkg=$GOPACKAGE gen "ChanT=struct{}"
//go:generate genny -in=../v/pipe/pipe.go -out=gen-pipe-struct.go -pkg=$GOPACKAGE gen "Name=struct InT=struct{} OutT=struct{}"

//go:generate genny -in=../v/topic/topic.go -out=gen-topic-time.go -pkg=$GOPACKAGE gen "ChanT=time.Time"
//go:generate genny -in=../v/pipe/pipe.go -out=gen-pipe-time.go -pkg=$GOPACKAGE gen "Name=time InT=time.Time OutT=struct{}"

import (
	"time"
)

const (
	op = "domain"
)

var (
	// Now returns time.Time
	Now func() time.Time = time.Now
)

// Terminate is terminator.
func Terminate(struct{}) (struct{}, error) {
	return struct{}{}, nil
}
