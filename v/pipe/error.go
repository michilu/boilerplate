package pipe

import (
	"github.com/michilu/boilerplate/v/errs"
	"github.com/michilu/boilerplate/v/log"
)

// DefaultErrorHandler is a default error handler.
func DefaultErrorHandler(err error) bool {
	const op = "v.pipe.DefaultErrorHandler"
	log.Logger().Fatal().
		Str("op", op).
		Err(&errs.Error{Op: op, Err: err}).
		Msg("error")
	return true
}
