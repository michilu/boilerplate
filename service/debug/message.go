package debug

import (
	"encoding/json"

	_ "github.com/mwitkow/go-proto-validators" // https://github.com/mwitkow/go-proto-validators/blob/master/validator.proto
	"github.com/rs/zerolog"
)

//go:generate genny -in=../../service/topic/topic.go -out=gen-topic-Clienter.go -pkg=$GOPACKAGE gen "ChanT=Clienter"

//go:generate interfacer -for github.com/michilu/boilerplate/service/debug.Client -as debug.Clienter -o entity-Clienter.go
//go:generate genny -in=../../service/topic/topic.go -out=gen-topic-Clienter.go -pkg=$GOPACKAGE gen "ChanT=Clienter"

func (p *Client) MarshalZerologObject(e *zerolog.Event) {
	if v, err := json.Marshal(&p); err == nil {
		e.RawJSON("Client", v)
	}
}

//go:generate genny -in=../topic/with-context.go -out=gen-ClientWithContext.go -pkg=$GOPACKAGE gen "T=Client Ier=Clienter"
//go:generate interfacer -for github.com/michilu/boilerplate/service/debug.ClientWithContext -as debug.ClientWithContexter -o if-ClientWithContexter.go
//go:generate genny -in=../../service/topic/topic.go -out=gen-topic-ClientWithContexter.go -pkg=$GOPACKAGE gen "ChanT=ClientWithContexter"
