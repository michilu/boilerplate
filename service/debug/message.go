package debug

import (
	"context"
	"encoding/json"

	"github.com/michilu/boilerplate/service/errs"
	_ "github.com/mwitkow/go-proto-validators" // https://github.com/mwitkow/go-proto-validators/blob/master/validator.proto
	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
)

//go:generate genny -in=../../service/topic/topic.go -out=gen-topic-Clienter.go -pkg=$GOPACKAGE gen "ChanT=Clienter"

//go:generate interfacer -for github.com/michilu/boilerplate/service/debug.Client -as debug.Clienter -o entity-Clienter.go
//go:generate genny -in=../../service/topic/topic.go -out=gen-topic-Clienter.go -pkg=$GOPACKAGE gen "ChanT=Clienter"

func (p *Client) MarshalZerologObject(e *zerolog.Event) {
	if v, err := json.Marshal(&p); err == nil {
		e.RawJSON("Client", v)
	}
}

//go:generate interfacer -for github.com/michilu/boilerplate/service/debug.ClientWithContext -as debug.ClientWithContexter -o if-ClientWithContexter.go
//go:generate genny -in=../../service/topic/topic.go -out=gen-topic-ClientWithContexter.go -pkg=$GOPACKAGE gen "ChanT=ClientWithContexter"

// ClientWithContext is message
type ClientWithContext struct {
	Context context.Context
	Client
}

func (p *ClientWithContext) GetContext() context.Context { return p.Context }
func (p *ClientWithContext) Validate() error {
	const op = op + ".ClientWithContext.Validate"
	if p.Context == nil {
		return &errs.Error{Op: op, Code: codes.InvalidArgument, Message: "must be given. 'Context' is nil"}
	}
	err := p.Client.Validate()
	if err != nil {
		return &errs.Error{Op: op, Code: codes.InvalidArgument, Err: err}
	}
	return nil
}

func (p *ClientWithContext) MarshalZerologObject(e *zerolog.Event) {
	if v, err := json.Marshal(&p); err == nil {
		e.RawJSON("Client", v)
	}
}