package errs

import (
	"github.com/rs/zerolog"
)

func (p *Error) MarshalZerologObject(e *zerolog.Event) {
	e.
		Str("Message", p.Message).
		Str("Op", p.Op)
	if p.Err != nil {
		e.Err(p.Err)
	}
	if p.Code != nil {
		e.Str("Code", p.Code.String())
	}
}
