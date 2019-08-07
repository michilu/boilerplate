package now

//go:generate genny -in=../topic/topic.go -out=gen-topic-time.go -pkg=$GOPACKAGE gen "ChanT=time.Time"
//go:generate genny -in=../pipe/pipe.go -out=gen-pipe-time.go -pkg=$GOPACKAGE gen "Name=time InT=time.Time OutT=struct{}"

import (
	"time"
)

const (
	op = "now"
)

var (
	// Now returns time.Time
	Now func() time.Time = time.Now
)
