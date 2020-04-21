package errs

import (
	"github.com/rs/zerolog"
)

func (p Error) MarshalZerologObject(e *zerolog.Event) {
	e.
		Str("message", p.Error()).
		Str("op", p.Op)
	if p.Code != nil {
		e.Str("code", p.Code.String())
	}
}
