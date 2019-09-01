package errs

import (
	"github.com/rs/zerolog"
)

func (p *Error) MarshalZerologObject(e *zerolog.Event) {
	e.
		Str("message", p.Message).
		Str("op", p.Op)
	if p.Err != nil {
		e.Err(p.Err)
	}
	if p.Code != nil {
		e.Str("code", p.Code.String())
	}
}
