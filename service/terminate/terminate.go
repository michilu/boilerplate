package terminate

//go:generate genny -in=../topic/topic.go -out=gen-topic-struct.go -pkg=$GOPACKAGE gen "ChanT=struct{}"
//go:generate genny -in=../pipe/pipe.go -out=gen-pipe-struct.go -pkg=$GOPACKAGE gen "Name=struct InT=struct{} OutT=struct{}"

const (
	op = "terminate"
)

// Terminate is terminator.
func Terminate(struct{}) (struct{}, error) {
	return struct{}{}, nil
}
