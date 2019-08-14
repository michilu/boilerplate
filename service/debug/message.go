package debug

import (
	// https://github.com/mwitkow/go-proto-validators/blob/master/validator.proto
	_ "github.com/mwitkow/go-proto-validators"

	"context"

	"github.com/michilu/boilerplate/service/errs"
	"google.golang.org/grpc/codes"
)

//go:generate genny -in=../../service/topic/topic.go -out=gen-topic-Clienter.go -pkg=$GOPACKAGE gen "ChanT=Clienter"

//go:generate interfacer -for github.com/michilu/boilerplate/service/debug.Client -as debug.Clienter -o entity-Clienter.go
//go:generate genny -in=../../service/topic/topic.go -out=gen-topic-Clienter.go -pkg=$GOPACKAGE gen "ChanT=Clienter"

//go:generate interfacer -for github.com/michilu/boilerplate/service/debug.ClientWithCtx -as debug.ClientWithCtxer -o if-ClientWithCtxer.go
//go:generate genny -in=../../service/topic/topic.go -out=gen-topic-ClientWithCtxer.go -pkg=$GOPACKAGE gen "ChanT=ClientWithCtxer"

// ClientWithCtx is message
type ClientWithCtx struct {
	Ctx context.Context
	Client
}

func (p *ClientWithCtx) GetCtx() context.Context { return p.Ctx }
func (p *ClientWithCtx) Validate() error {
	const op = op + ".ClientWithCtx.Validate"
	if p.Ctx == nil {
		return &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'Ctx' is nil"}
	}
	err := p.Client.Validate()
	if err != nil {
		return &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
	}
	return nil
}
